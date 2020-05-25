package core

import (
	"context"
	"html/template"
	"net/http"
	"sqlbar/server/src/config"
	"sqlbar/server/src/logs"
	"sqlbar/server/src/postgresql"
	"strconv"

	pq "github.com/lib/pq"
)

type (
	// Submenu struct
	Submenu struct {
		Name string `json:"name"`
		Link string `json:"link"`
		Icon string `json:"icon"`
	}

	// Menu struct
	Menu struct {
		Name     string    `json:"name"`
		Submenus []Submenu `json:"submenus"`
	}

	// User struct
	User struct {
		ID       int64  `json:"id"`
		FullName string `json:"name"`
		Avatar   string `json:"avatar"`
	}

	// Server struct
	Server struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	// List struct
	List struct {
		ID    int64
		Name  string
		Date  pq.NullTime //time.Time
		Count int64
	}
	// Lists struct
	Lists struct {
		List []List `json:"lists"`
	}

	// EmptyResponse struct for non auth responses
	EmptyResponse struct {
		Code   int    `json:"code"`
		Msg    string `json:"msg"`
		Server Server `json:"server"`
	}

	// DefaultResponse struct for auth responses
	DefaultResponse struct {
		EmptyResponse
		Menus []Menu `json:"menus"`
		User  User   `json:"user"`
	}

	// IDData struct
	IDData struct {
		ID   int64 `json:"id"`
		Port int   `json:"port"`
	}
)

// MakeEmptyResponse func
func MakeEmptyResponse(code int, msg string) *EmptyResponse {
	return &EmptyResponse{
		code,
		msg,
		MakeServer(),
	}
}

// MakeDefaultResponse func
func MakeDefaultResponse(code int, msg string, s *Session) *DefaultResponse {
	return &DefaultResponse{
		*MakeEmptyResponse(code, msg),
		MakeMenus(s),
		MakeUser(s),
	}
}

func data() (lists Lists, err error) {
	logs.Log.PushFuncName("core", "front", "data")
	defer logs.Log.PopFuncName()

	query := "select id, name, date, count from fLists() order by id"
	rows, err := postgresql.Db.Query(query)
	if err != nil {
		logs.Log.Error("postgresql.Db.Query", err)
		return
	}
	defer rows.Close()

	lists = Lists{List: []List{}}
	for rows.Next() {
		list := List{}
		err := rows.Scan(&list.ID, &list.Name, &list.Date, &list.Count)
		if err != nil {
			logs.Log.Error("rows.Scan", err)
			continue
		}
		lists.List = append(lists.List, list)
	}

	return
}

// MakeMenus func
func MakeMenus(s *Session) (ms []Menu) {
	menu := Menu{
		Name:     "MAIN NAVIGATION",
		Submenus: make([]Submenu, 0)}

	if (s.Role.ID == UroAdmin) || (s.Role.ID == UroSuperAdmin) {
		submenu := Submenu{Name: "Lists", Link: "/lists/data.html", Icon: "fa-clone"}
		menu.Submenus = append(menu.Submenus, submenu)
	}
	if (s.Role.ID == UroAdmin) || (s.Role.ID == UroSuperAdmin) || (s.Role.ID == UroModerator) {
		submenu := Submenu{Name: "Products", Link: "/products/data.html", Icon: "fa-cubes"}
		menu.Submenus = append(menu.Submenus, submenu)
	}
	if (s.Role.ID == UroAdmin) || (s.Role.ID == UroSuperAdmin) || (s.Role.ID == UroModerator) {
		submenu := Submenu{Name: "Contacts", Link: "/contacts/data.html", Icon: "fa-newspaper-o"}
		menu.Submenus = append(menu.Submenus, submenu)
	}
	if (s.Role.ID == UroSysAdmin) || (s.Role.ID == UroSuperAdmin) {
		submenu := Submenu{Name: "Users", Link: "/users/data.html", Icon: "fa-users"}
		menu.Submenus = append(menu.Submenus, submenu)
	}
	if (s.Role.ID == UroSysAdmin) || (s.Role.ID == UroSuperAdmin) {
		submenu := Submenu{Name: "Settings", Link: "/other/settings.html", Icon: "fa-gear"}
		menu.Submenus = append(menu.Submenus, submenu)
	}
	ms = append(ms, menu)

	if (s.Role.ID == UroAdmin) || (s.Role.ID == UroModerator) || (s.Role.ID == UroSuperAdmin) {
		menu := Menu{
			Name:     "LEADS",
			Submenus: make([]Submenu, 0)}
		lists, err := data()
		if err != nil {
			logs.Log.Error("sql.Data", err.Error())
			/* json.NewEncoder(w).Encode(
				MakeDefaultResponse(-4, "error read data from db", w, r, s),
			) */
			return
		}
		for _, list := range lists.List {
			submenu := Submenu{
				Name: list.Name,
				Link: "/orders/data.html?id=" + strconv.FormatInt(list.ID, 10),
				Icon: "fa-table"}
			menu.Submenus = append(menu.Submenus, submenu)
		}
		ms = append(ms, menu)
	}
	return
}

//MakeUser func
func MakeUser(s *Session) (u User) {
	u.FullName = s.FirstName + " " + s.LastName
	u.Avatar = "/templates/dist/img/user2-160x160.jpg"
	return
}

// MakeServer func
func MakeServer() (s Server) {
	s.Version = config.ServerVersion
	s.Name = config.ServerName
	return
}

//MakeResponse func
/* func MakeResponse(w http.ResponseWriter, r *http.Request, s *Session, code int, msg string) {
	json.NewEncoder(w).Encode(
		DefaultResponse{
			Code:  code,
			Msg:   msg,
			Menus: MakeMenus(w, r, s),
			User:  MakeUser(w, r, s),
		})
} */

//MiddleWare type
type MiddleWare func(next http.HandlerFunc) http.HandlerFunc

//ChainMiddleWare func
func ChainMiddleWare(mw ...MiddleWare) MiddleWare {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}

//WithLogging func
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logs.Log.PushFuncName("core", "front", "WithLogging")
		defer logs.Log.PopFuncName()

		logs.Log.Debug("BEGIN", r.RemoteAddr, r.RequestURI)
		defer logs.Log.Debug("END")
		next.ServeHTTP(w, r)
	}
}

//WithAuth func
func WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logs.Log.PushFuncName("core", "front", "WithAuth")
		defer logs.Log.PopFuncName()

		r.ParseForm()

		session, err := ReadCookie(w, r)
		if err != nil {
			logs.Log.Error("ReadCookie", err.Error())
			http.Error(w, "wrong pin", http.StatusForbidden) //Redirect404(w, r)
			return
		}

		/* if !session.Role.HasOrderRights() {
			logs.Log.Error("HasOrderRights", "no order rights")
			http.Error(w, "wrong pin", http.StatusForbidden) //Redirect404(w, r)
			return
		} */

		logs.Log.Info("SESSION OK", session)
		ctx := context.WithValue(r.Context(), "Session", session)

		logs.Log.SessionName = session.Email

		next.ServeHTTP(w, r.WithContext(ctx))

		logs.Log.SessionName = ""
	}
}

//DefaultHandler func
func DefaultHandler(w http.ResponseWriter, r *http.Request, folder string, module string) {
	logs.Log.PushFuncName("core", "front", "DefaultHandler")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN", folder, module)
	defer logs.Log.Debug("END")

	t, err := template.ParseFiles("templates/pages/" + folder + "/" + module + ".html")
	if err != nil {
		logs.Log.Error("ParseFiles", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, folder+"/"+module+".html", nil)
	if err != nil {
		logs.Log.Error("ExecuteTemplate", err.Error())
		return
	}
}

// IDHandler func
/*func IDHandler(w http.ResponseWriter, r *http.Request, folder string, module string) {
	logs.Log.PushFuncName("core", "front", "IDHandler")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN", folder, module)
	defer logs.Log.Debug("END")

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(MakeDefaultResponse(-1, "error parse form", nil))
		return
	}

	var id int64 = -1
	if _, ok := r.Form["id"]; ok {
		id, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	EnableCors(&w)

	t, err := template.ParseFiles("templates/pages/" + folder + "/" + module + ".html")
	if err != nil {
		logs.Log.Error("ParseFiles", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, folder+"/"+module+".html", IDData{ID: id, Port: config.ServerPort})
	if err != nil {
		logs.Log.Error("ExecuteTemplate", err.Error())
		return
	}
}

	t, err := template.ParseFiles("templates/pages/" + folder + "/" + module + ".html")
	if err != nil {
		logs.Log.Error("ParseFiles", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, folder+"/"+module+".html", IDData{ID: id, Port: config.ServerPort})
	if err != nil {
		logs.Log.Error("ExecuteTemplate", err.Error())
		return
	}
}*/

// ParseStr func
func ParseStr(r *http.Request, param string, def string) string {
	if _, ok := r.Form[param]; ok {
		return r.Form[param][0]
	}
	return def
}

// ParseInt func
func ParseInt(r *http.Request, param string, def int) int {
	if _, ok := r.Form[param]; ok {
		value, _ := strconv.Atoi(r.Form[param][0])
		return value
	}
	return def
}

// ParseInt64 func
func ParseInt64(r *http.Request, param string, def int64) int64 {
	if _, ok := r.Form[param]; ok {
		value, _ := strconv.ParseInt(r.Form[param][0], 10, 64)
		return value
	}
	return def
}
