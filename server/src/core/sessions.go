package core

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"sqlbar/server/src/logs"
	"strconv"
	"sync"
	"time"
)

//http://www.gorillatoolkit.org/pkg/sessions

type (
	//Session struct
	Session struct {
		Expired    time.Time
		ID         int64
		Password   string
		FirstName  string
		LastName   string
		Post       string
		Phone      string
		Email      string
		Role       UserRoles
		Registered time.Time
	}

	//SessionManager struct
	SessionManager struct {
		Cache map[string]Session
		mutex sync.RWMutex
	}
)

var (
	secInMin  int64 = 60
	tokenSalt       = "sALTaLTf4" //TODO: вынести в конфиг
	// Sessions manager
	Sessions = SessionManager{
		Cache: make(map[string]Session),
		mutex: sync.RWMutex{},
	}
)

func expiredRoutine() {
	for {
		time.Sleep(time.Second * 30)
		CheckExpired()
	}
}

func init() {
	logs.Log.PushFuncName("core", "sessions", "init")
	defer logs.Log.PopFuncName()

	logs.Log.Info("IMPORTED")
	//TODO: go expiredRoutine()
}

// ReadCookie func
func ReadCookie(w http.ResponseWriter, r *http.Request) (s *Session, err error) {
	logs.Log.PushFuncName("core", "sessions", "ReadCookie")
	defer logs.Log.PopFuncName()

	//logs.Log.Debug("BEGIN")

	var session *Session
	cookie, err := r.Cookie("token")
	if err != nil {
		logs.Log.Error("Cookie('token')", err)
		http.Redirect(w, r, "../../auth.html", 303)
		err = errors.New("no token")
		return
	}

	session, err = CheckAuth(cookie.Value)
	if err != nil {
		logs.Log.Warning("CheckAuth(cookie.Value)", err, Sessions.Cache, "token:", cookie.Value)
		http.Redirect(w, r, "../../auth.html", 303)
		err = errors.New("no auth")
		return
	}

	logs.Log.Debug("END", Sessions.Cache, session, cookie)
	s = session
	return
}

// IsAdmin func
func (s *Session) IsAdmin() bool {
	return (s.Role.ID == UroSuperAdmin) || (s.Role.ID == UroAdmin)
}

//CheckExpired func
func CheckExpired() {
	logs.Log.PushFuncName("core", "sessions", "CheckExpired")
	defer logs.Log.PopFuncName()

	Sessions.mutex.Lock()
	defer Sessions.mutex.Unlock()

	for token, v := range Sessions.Cache {
		t := time.Now()
		if v.Expired.Unix() > t.Unix() {
			delete(Sessions.Cache, token)
			logs.Log.Info("session deleted:", token)
		}
	}
}

//Create func
func Create(email, pass string) string {
	logs.Log.PushFuncName("core", "sessions", "Create")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN", email, pass)
	t := time.Now().Unix()
	unsignedToken := email + pass + strconv.Itoa(int(t)) + tokenSalt
	token := encryptString(unsignedToken)
	Upsert(token)
	logs.Log.Debug("END") //, "sessions:", Sessions.Cache)
	return token
}

//Upsert func
func Upsert(token string) (s *Session) {
	logs.Log.PushFuncName("core", "sessions", "Upsert")
	defer logs.Log.PopFuncName()

	Sessions.mutex.Lock()
	defer Sessions.mutex.Unlock()

	t := time.Now().Unix()
	Sessions.Cache[token] = Session{}
	logs.Log.Debug("Created sessions:", Sessions.Cache)
	session, ok := Sessions.Cache[token]
	if !ok {
		logs.Log.Error("Sessions.Cache[token] error: no session with token:" + token)
		return
	}

	session.Expired = time.Unix(0, t+60*secInMin) // 60 minute
	s = &session
	logs.Log.Debug("session:", s) //, "sessions:", Sessions.Cache)
	return
}

//CheckAuth func
func CheckAuth(token string) (s *Session, err error) {
	logs.Log.PushFuncName("core", "sessions", "CheckAuth")
	defer logs.Log.PopFuncName()

	Sessions.mutex.RLock()
	defer Sessions.mutex.RUnlock()

	logs.Log.Debug("BEGIN", token)
	session, ok := Sessions.Cache[token]
	if ok != true {
		err = errors.New("Sessions.Cache[token]: no session with token: " + token)
		return
	}

	s = &session
	logs.Log.Debug("END", "session:", s) //, "sessions:", Sessions.Cache)
	return
}

// UpdateSession func
func UpdateSession(token string, s *Session) {
	logs.Log.PushFuncName("core", "sessions", "UpdateSession")
	defer logs.Log.PopFuncName()

	Sessions.mutex.Lock()
	defer Sessions.mutex.Unlock()

	Sessions.Cache[token] = *s
	logs.Log.Info("updated OK") //, "sessions:", Sessions.Cache)
}

func encryptString(str string) string {
	h := sha256.Sum256([]byte(str))
	return base64.StdEncoding.EncodeToString(h[:])
}

/*func main() {
	token := tokenCreate("login", "pass")
	fmt.Println("token", token)
	fmt.Println("cache", sessionCache.Cache)
	if authCheck(token) {
		fmt.Println("token exists")
	}
}*/

//package main

/*
import (
	"fmt"
	"sync"
)

var globalSessions *session.Manager
var provides = make(map[string]Provider)

type Manager struct {
	cookieName  string     //private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

// Register makes a session provider available by the provided name.
// If a Register is called twice with the same name or if the driver is nil,
// it pa nics.
func Register(name string, provider Provider) {
	if provider == nil {
		p anic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		p anic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}*/

//globalSessions = NewManager("memory", "gosessionid", 3600)
