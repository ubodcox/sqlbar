package users

import (
	"encoding/json"
	"net/http"
	"sqlbar/server/src/core"
	"sqlbar/server/src/logs"

	//"sqlbar/server/src/messengers"
	"strconv"
)

type (
	// API struct
	API struct {
	}
	// Response struct
	/* Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} */

	// ResponseAuth struct
	/* ResponseAuth struct {
		Code   int    `json:"code"`
		Url    string `json:"url"`
		Cookie string `json:"cookie"`
	} */

	// DataRole struct
	DataRole struct {
		ID      int64  `json:"id"`
		Caption string `json:"caption"`
	}

	// Data struct
	Data struct {
		ID         int64    `json:"id"`
		FirstName  string   `json:"firstname"`
		LastName   string   `json:"lastname"`
		Post       string   `json:"post"`
		Phone      string   `json:"phone"`
		Email      string   `json:"email"`
		Registered string   `json:"registered"`
		Role       DataRole `json:"role"`
		RoleInt    int      `json:"roleint"`
		IsDeleted  bool     `json:"isdeleted"`
	}

	// response struct
	response struct {
		core.DefaultResponse
		/* Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Menus []core.Menu `json:"menus"`
		User  core.User   `json:"user"` */
		Datas []Data `json:"datas"`
	}

	// EmptyResponse struct
	/* EmptyResponse struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
	} */
)

// response struct
/* response struct {
		core.DefaultResponse
		Steps []Step `json:"steps"`
	}
) */

var api API

// makeResponse func
func makeResponse(code int, msg string, s *core.Session, d []Data) *response {
	return &response{
		*core.MakeDefaultResponse(code, msg, s),
		d,
	}
}

//var templateStr string
//var productid int64

/* func core.MakeResponse(w http.ResponseWriter, r *http.Request, code int, msg string) {
	/* {if templateStr != "" {
		http.Redirect(w, r, "../../../products/info.html?id="+strconv.FormatInt(productid, 10), 303)
	} else { * /
	json.NewEncoder(w).Encode(
		Response{
			Code: code,
			Msg:  msg,
		})
	//}
} */

func (api *API) noauth(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "api", "noauth")
	defer logs.Log.PopFuncName()

	core.EnableCors(&w)

	json.NewEncoder(w).Encode(
		core.MakeEmptyResponse(1, ""),
	)

	logs.Log.Debug("code: 1")
}

func (api *API) auth(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "api", "auth")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("r.ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "Error parse form"))
		//json.NewEncoder(w).Encode(core.MakeResponse(-1, "error parse form"))
		return
	}
	//logs.Log.Debug("test", 1)
	core.EnableCors(&w)

	user := User{}
	if _, ok := r.Form["email"]; ok {
		user.Email = r.Form["email"][0]
	}
	if _, ok := r.Form["password"]; ok {
		user.Password = r.Form["password"][0]
	}
	//logs.Log.Debug("test", 2)
	err = sql.exists(&user)
	if err != nil {
		logs.Log.Error("sql.exists", err, "user:", user)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "User not exists"))
		//core.MakeResponse(w, r, nil, -2, "User not exists")
		//http.Redirect(w, r, "../auth.html", 303)
		return
	}

	//logs.Log.Debug("test", 3)
	users, err := sql.data(user.ID, true)
	//err = data(&user)
	if err != nil {
		logs.Log.Error("userSelect", err, "user:", user)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-3, "User not selected"))
		//log.Println("users.ExecAuth.Handler", err.Error())
		//http.Redirect(w, r, "../auth.html", 303)
		return
	}

	if len(users.List) != 1 {
		logs.Log.Error("users != 1")
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-3, "users != 1"))
		//log.Println("users.ExecAuth.Handler", err.Error())
		//http.Redirect(w, r, "../auth.html", 303)
		return
	}

	user = users.List[0]

	//var cookie string
	WriteCookie(w, r, user)

	var url string = "../../../other/title.html"
	/*logs.Log.Debug("user", user, user.Role.HasOrderRights(), user.Role.HasProductRights(), user.Role.HasUserRights())
	if user.Role.HasOrderRights() {
		url = "../../../orders/data.html"
	} else if user.Role.HasProductRights() {
		url = "../../../products/data.html"
	} else if user.Role.HasUserRights() {
		url = "../../../users/data.html"
	} */

	//core.MakeResponse(w, r, nil, 1, url)
	json.NewEncoder(w).Encode(core.MakeEmptyResponse(1, url))
	//cookie, _ := r.Cookie("token")

	/* json.NewEncoder(w).Encode(
	ResponseAuth{
		Code:   1,
		Url:    url,
		Cookie: cookie.,
	}) */

	/*
		templateStr = ""
		productid = -1
		err := r.ParseForm()
		if err != nil {
			logs.Log.Error("ParseForm", err)
			core.MakeResponse(w, r, -1, "error parse form")
			//json.NewEncoder(w).Encode(core.MakeResponse(-1, "error parse form"))
			return
		}
		//log.Println("orders.api.add 1")
		core.EnableCors(&w)
		var firstname, lastname, phone, email, message string
		//log.Println("orders.api.add 2")
		if _, ok := r.Form["firstName"]; ok {
			firstname = r.Form["firstName"][0]
			//logs.Log.Debug("2 no err", firstname)
		} else {
			//logs.Log.Debug("2 err")
		}
		if _, ok := r.Form["lastName"]; ok {
			lastname = r.Form["lastName"][0]
		}
		if _, ok := r.Form["phone"]; ok {
			phone = r.Form["phone"][0]
		}
		if _, ok := r.Form["email"]; ok {
			email = r.Form["email"][0]
		}
		if _, ok := r.Form["message"]; ok {
			message = r.Form["message"][0]
		}
		if _, ok := r.Form["productid"]; ok {
			productid, _ = strconv.ParseInt(r.Form["productid"][0], 10, 64)
		}
		if _, ok := r.Form["template"]; ok {
			templateStr = r.Form["template"][0]
		}
		//log.Println("orders.api.add 3")
		//log.Println("orders.api.add 3.0.1", firstname, lastname)
		if (firstname == "") && (lastname == "") {
			logs.Log.Error("check name", "no firstname or lastname")
			core.MakeResponse(w, r, -2, "no firstname or lastname")
			return
		}
		//log.Println("orders.api.add 3.1")
		if email == "" {
			logs.Log.Error("t productsSelect.email", "no email")
			core.MakeResponse(w, r, -2, "no email")
			return
		}
		//log.Println("orders.api.add 3.2")
		if !messengers.IsEmailValid(email) {
			logs.Log.Error("messengers.IsEmailValid", "email not valid")
			core.MakeResponse(w, r, -3, "email not valid")
			return
		}
		//log.Println("orders.api.add 4")
		err = sql.Order(firstname, lastname, phone, email, message, productid)
		if err != nil {
			logs.Log.Error("sql.Order", err.Error())
			core.MakeResponse(w, r, -4, "error write to db")
			return
		}
		err = messengers.SendMail(email, firstname, phone)
		if err != nil {
			logs.Log.Error("messengers.SendMail", err.Error())
			core.MakeResponse(w, r, -5, "error send mail")
			return
		}
		//log.Println("orders.api.add 5")
		core.MakeResponse(w, r, 1, "")
		//log.Println("orders.api.add 6")
		//http.Redirect(w, r, "https://github.com/WTSSystem/WTS/raw/master/WTS.exe#", 301)
		//counter++

		/*else {
			log.Println("Name or Email not valid:", firstname, email)
		}*/
	/*
			//log.Println("indexHandler", "enter")
			t, err := template.ParseFiles("templates/index.html")
			if err != nil {
				log.Println("indexHandler", err.Error())
				return
			}

			now := time.Now()
			usa := now.Add(time.Duration(5) * time.Hour)
			//log.Println("indexHandler", "exit")
			data := old.LandingData{
				LocalTime: now.Format(time.Kitchen),
				UsaTime:   usa.Format(time.Kitchen),
				Count:     strconv.Itoa(old.SubsCount())}

			//log.Println(w)
			//log.Println(t)
			t.ExecuteTemplate(w, "index", data)

		}

		log.Println("users.ExecAuth.Handler")
		r.ParseForm()

		/*token := core.Create(user.Email, user.Password)
			log.Println("session core.Create: ", token, core.Sessions)

			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := &http.Cookie{Name: "token", Value: token, Expires: expiration,
				Path: "/", MaxAge: 86400}
			http.SetCookie(w, cookie)

			session := core.Sessions.Cache[token]
			session.Email = user.Email
			session.Password = user.Password
			if user.Role.HasOrderRights() {
				http.Redirect(w, r, "../orders/data.html", 303)
			} else if user.Role.HasProductRights() {
				http.Redirect(w, r, "../products/data.html", 303)
			} else if user.Role.HasUserRights() {
				http.Redirect(w, r, "../users/data.html", 303)
			}
			/*t, err := template.ParseFiles("templates/pages/orders/data.html")
			if err != nil {
				log.Println("usersA uthHandler.ParseFiles", err.Error())
				return
			}

			err = t.ExecuteTemplate(w, "pages/users/auth.html", nil)
			if err != nil {
				log.Println("usersA uthHandler.ExecuteTemplate", err.Error())
				return
			}
		} */
}

func (api *API) get(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "api", "get")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	session := r.Context().Value("Session").(*core.Session)

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeDefaultResponse(-1, "error parse form", session))
		return
	}

	/*var id int64 = -1 //list := List{}
	if _, ok := r.Form["id"]; ok {
		id, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	core.EnableCors(&w)
	lists, err := sql.data(id, session.Role.HasUserRights())
	if err != nil {
		logs.Log.Error("sql.Data", err.Error())
		json.NewEncoder(w).Encode(core.MakeDefaultResponse(-4, "error read data from db", session))
		return
	}

	logs.Log.Debug("lists", lists)

	json.NewEncoder(w).Encode(
		makeResponse(1, "", session, lists.toData()),
	)*/

	/* data, err := sql.data(-1, session.Role.HasUserRights())
	if err != nil {
		logs.Log.Error("userList", err.Error())
		return
	} */

	logs.Log.Debug("code: 1")
}

func (api *API) insert(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "api", "insert")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	//session := r.Context().Value("Session").(*core.Session)

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "error parse form"))
		//json.NewEncoder(w).Encode(core.MakeResponse(-1, "error parse form"))
		return
	}

	core.EnableCors(&w)
	user := User{}
	//var password, firstname, lastname, post, phone, email string
	if _, ok := r.Form["password"]; ok {
		user.Password = r.Form["password"][0]
	}
	if _, ok := r.Form["firstName"]; ok {
		user.FirstName = r.Form["firstName"][0]
	}
	if _, ok := r.Form["lastName"]; ok {
		user.LastName = r.Form["lastName"][0]
	}
	if _, ok := r.Form["post"]; ok {
		user.Post = r.Form["post"][0]
	}
	if _, ok := r.Form["phone"]; ok {
		user.Phone = r.Form["phone"][0]
	}
	if _, ok := r.Form["email"]; ok {
		user.Email = r.Form["email"][0]
	}
	if _, ok := r.Form["role"]; ok {
		val, _ := strconv.Atoi(r.Form["role"][0])
		user.Role = core.GetRole(core.UserRoleID(val))
	}

	if (user.Password == "") || (user.FirstName == "") {
		logs.Log.Error("wrong first name or password")
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "no first name or password"))
		return
	}

	if user.Email == "" {
		logs.Log.Error("wrong first name or password")
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "no email"))
		return
	}

	/*if !messengers.IsEmailValid(user.Email) {
		logs.Log.Error("messengers.IsEmailValid", user.Email)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "wrong email"))
		return
	}*/

	err = sql.insert(user)
	if err != nil {
		logs.Log.Error("sql.insert", err.Error())
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-4, "error write to db"))
		return
	}

	url := "../../../users/data.html"
	json.NewEncoder(w).Encode(core.MakeEmptyResponse(1, url))
	logs.Log.Debug("code: 1")
}

func (api *API) set(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "api", "set")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	//session := r.Context().Value("Session").(*core.Session)

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "error parse form"))
		//json.NewEncoder(w).Encode(core.MakeResponse(-1, "error parse form"))
		return
	}

	core.EnableCors(&w)
	user := User{}
	//var password, firstname, lastname, post, phone, email string
	if _, ok := r.Form["id"]; ok {
		user.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}
	if _, ok := r.Form["password"]; ok {
		user.Password = r.Form["password"][0]
	}
	if _, ok := r.Form["firstName"]; ok {
		user.FirstName = r.Form["firstName"][0]
	}
	if _, ok := r.Form["lastName"]; ok {
		user.LastName = r.Form["lastName"][0]
	}
	if _, ok := r.Form["post"]; ok {
		user.Post = r.Form["post"][0]
	}
	if _, ok := r.Form["phone"]; ok {
		user.Phone = r.Form["phone"][0]
	}
	if _, ok := r.Form["email"]; ok {
		user.Email = r.Form["email"][0]
	}
	if _, ok := r.Form["role"]; ok {
		val, _ := strconv.Atoi(r.Form["role"][0])
		user.Role = core.GetRole(core.UserRoleID(val))
	}

	if /* (user.Password == "") || */ user.FirstName == "" {
		logs.Log.Error("no first name")
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "no first name"))
		return
	}

	if user.Email == "" {
		logs.Log.Error("no email")
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "no email"))
		return
	}

	/*if !messengers.IsEmailValid(user.Email) {
		logs.Log.Error("messengers.IsEmailValid", user.Email)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "wrong email"))
		return
	}*/

	err = sql.update(user)
	if err != nil {
		logs.Log.Error("sql.update", err.Error())
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-4, "error write to db"))
		return
	}

	url := "../../../users/data.html"
	json.NewEncoder(w).Encode(core.MakeEmptyResponse(1, url))
	logs.Log.Debug("code: 1")
}

func (api *API) delete(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "api", "delete")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	//session := r.Context().Value("Session").(*core.Session)

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "error parse form"))
		//json.NewEncoder(w).Encode(core.MakeResponse(-1, "error parse form"))
		return
	}

	core.EnableCors(&w)
	user := User{}
	if _, ok := r.Form["id"]; ok {
		user.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	if user.ID < 0 {
		logs.Log.Error("check id", "no id")
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "no id"))
		return
	}

	/* order := Order{}
	if _, ok := r.Form["id"]; ok {
		order.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	} */

	err = sql.delete(user)
	if err != nil {
		logs.Log.Error("sql.delete", err.Error())
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-4, "error write to db"))
		return
	}

	url := "../../../users/data.html"
	json.NewEncoder(w).Encode(core.MakeEmptyResponse(1, url))
	logs.Log.Debug("code: 1")
}
