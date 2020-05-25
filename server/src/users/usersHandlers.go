package users

/* import (
	"core"
	"html/template"
	"logs"
	"messengers"
	"net/http"
	"strconv"
) */

/* func usersAuthHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "usersHandlers", "usersAuthHandler")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	r.ParseForm()

	t, err := template.ParseFiles("templates/pages/users/auth.html")
	if err != nil {
		logs.Log.Error("template.ParseFiles", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, "users/auth.html", nil)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}
} */

/* func usersExecAuthHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.Error("users.ExecAuth.Handler")
	r.ParseForm()

	user := User{}
	if _, ok := r.Form["email"]; ok {
		user.Email = r.Form["email"][0]
	}
	if _, ok := r.Form["password"]; ok {
		user.Password = r.Form["password"][0]
	}

	err := userExists(&user)
	if err != nil {
		logs.Log.Error("users.ExecAuth.Handler", err.Error())
		http.Redirect(w, r, "../auth.html", 303)
		return
	}

	err = userSelect(&user)
	if err != nil {
		logs.Log.Error("users.ExecAuth.Handler", err.Error())
		http.Redirect(w, r, "../auth.html", 303)
		return
	}

	//session :=
	WriteCookie(w, r, user)
	/*token := core.Create(user.Email, user.Password)
	logs.Log.Error("session core.Create: ", token, core.Sessions)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := &http.Cookie{Name: "token", Value: token, Expires: expiration,
		Path: "/", MaxAge: 86400}
	http.SetCookie(w, cookie)

	session := core.Sessions.Cache[token]
	session.Email = user.Email
	session.Password = user.Password*
	if user.Role.HasOrderRights() {
		http.Redirect(w, r, "../orders/data.html", 303)
	} else if user.Role.HasProductRights() {
		http.Redirect(w, r, "../products/data.html", 303)
	} else if user.Role.HasUserRights() {
		http.Redirect(w, r, "../users/data.html", 303)
	}
	/*t, err := template.ParseFiles("templates/pages/orders/data.html")
	if err != nil {
		logs.Log.Error("usersAuth Handler.ParseFiles", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, "pages/users/auth.html", nil)
	if err != nil {
		logs.Log.Error("usersAut hHandler.ExecuteTemplate", err.Error())
		return
	}*
} */

/* func usersListHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "usersHandlers", "usersListHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("ReadCookie", err.Error())
		return
	}

	logs.Log.Debug("DEBUG:", session)
	logs.Log.Debug("DEBUG:", session.Role)
	logs.Log.Debug("DEBUG:", session.Role.HasUserRights())
	if !session.Role.HasUserRights() {
		logs.Log.Error("HasUserRights", "no rights")
		return
	}

	t, err := template.ParseFiles("templates/pages/users/data.html")
	if err != nil {
		logs.Log.Error("template.ParseFiles", err.Error())
		return
	}

	data, err := userList(session.Role.HasUserRights())
	if err != nil {
		logs.Log.Error("userList", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, "users/data.html", data)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}
} */

/* func usersInsertHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "usersHandlers", "usersInsertHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("core.ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriUserAdd) {
		logs.Log.Error("HasRight UriUserAdd", err.Error())
		return
	}

	/*sub := Subscribe{}
	//var firstname, lastname, phone, email string
	if _, ok := r.Form["firstName"]; ok {
		sub.FirstName = r.Form["firstName"][0]
	}
	if _, ok := r.Form["lastName"]; ok {
		sub.LastName = r.Form["lastName"][0]
	}
	if _, ok := r.Form["phone"]; ok {
		sub.Phone = r.Form["phone"][0]
	}
	if _, ok := r.Form["email"]; ok {
		sub.Email = r.Form["email"][0]
	}

	sub.Registered = time.Now()

	if sub.FirstName != "" && sub.Email != "" {
		if isEmailValid(sub.Email) {

		} else {
			logs.Log.Error("Email not valid:", sub.Email)
		}
	} /*else {
		logs.Log.Error("Name or Email not valid:", firstname, email)
	}* /

	//logs.Log.Error("indexHandler", "enter")* /
	t, err := template.ParseFiles("templates/pages/users/insert.html")
	if err != nil {
		logs.Log.Error("template.ParseFiles", err.Error())
		return
	}

	//data := subscribesAdd(sub)

	t.ExecuteTemplate(w, "users/insert.html", nil)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}
} */

/* func usersUpdateHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "usersHandlers", "usersUpdateHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("core.ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriUserEdit) {
		logs.Log.Error("HasRight UriUserEdit", err.Error())
		return
	}

	user := User{}

	if _, ok := r.Form["id"]; ok {
		user.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}
	/*if _, ok := r.Form["firstName"]; ok {
		sub.FirstName = r.Form["firstName"][0]
	}
	if _, ok := r.Form["lastName"]; ok {
		sub.LastName = r.Form["lastName"][0]
	}
	if _, ok := r.Form["phone"]; ok {
		sub.Phone = r.Form["phone"][0]
	}
	if _, ok := r.Form["email"]; ok {
		sub.Email = r.Form["email"][0]
	}
	if _, ok := r.Form["registered"]; ok {
		layout := "2006-01-02T15:04:05.000Z"
		sub.Registered, _ = time.Parse(layout, r.Form["registered"][0])
	}* /

	err = userSelect(&user)
	if err != nil {
		logs.Log.Error("usersSelect", err.Error())
		return
	}

	t, err := template.ParseFiles("templates/pages/users/update.html")
	if err != nil {
		logs.Log.Error("template.ParseFiles", err.Error())
		return
	}

	logs.Log.Error(user)
	t.ExecuteTemplate(w, "users/update.html", user)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}
} */

/*func subscribesDeleteHandler(w http.ResponseWriter, r *http.Request) {
logs.Log.Error("subscribesDeleteHandler.begin")
r.ParseForm()
sub := Subscribe{}

if _, ok := r.Form["id"]; ok {
	sub.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
}

if sub.FirstName != "" && sub.Email != "" {
	if isEmailValid(sub.Email) {

	} else {
		logs.Log.Error("Email not valid:", sub.Email)
	}
} /*else {
	logs.Log.Error("Name or Email not valid:", firstname, email)
}*/

//logs.Log.Error("indexHandler", "enter")
/*	t, err := template.ParseFiles("templates/subscribesDelete.html")
	if err != nil {
		logs.Log.Error("subscribesDelete", err.Error())
		return
	}

	data := subscribesDelete(sub)

	t.ExecuteTemplate(w, "subscribesDelete", data)
}*/

/* func usersExecInsertHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "usersHandlers", "usersExecInsertHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("core.ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriUserAdd) {
		logs.Log.Error("HasRight UriUserAdd", err.Error())
		return
	}

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

	if user.Password != "" && user.FirstName != "" && user.Email != "" {
		if messengers.IsEmailValid(user.Email) {
			err := userAdd(user)
			if err == nil {
				//err := sendMail(email, firstname, phone)
				//logs.Log.Error("err: ", err)
				//if err == nil {
				/*http.Redirect(w, r, "https://github.com/WTSSystem/WTS/raw/master/WTS.exe#", 301)
					//counter++
				} else {
					logs.Log.Error("sendMail error:", err)
				}* /
			} else {
				logs.Log.Error("user error:", err)
			}
		} else {
			logs.Log.Error("Email not valid:", user.Email)
		}
	} /*else {
		logs.Log.Error("Name or Email not valid:", firstname, email)
	}* /

	//logs.Log.Error("indexHandler", "enter")

	t, err := template.ParseFiles("templates/pages/users/data.html")
	if err != nil {
		logs.Log.Error("template.ParseFiles", err.Error())
		return
	}

	data, err := userList(session.Role.HasUserRights())
	if err != nil {
		logs.Log.Error("userList", err.Error())
		return
	}

	//logs.Log.Error(data)
	err = t.ExecuteTemplate(w, "users/data.html", data)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}
} */

/* func usersExecUpdateHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "usersHandlers", "usersExecUpdateHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("core.ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriUserEdit) {
		logs.Log.Error("HasRight UriUserEdit", err.Error())
		return
	}

	user := User{}
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
		logs.Log.Error("ROLE:", user.Role)
	} else {
		logs.Log.Error("ROLE:", "NO ROLE", r.Form)
	}

	if user.FirstName != "" && user.Email != "" {
		if messengers.IsEmailValid(user.Email) {
			err := usersUpdate(user)
			if err == nil {
				//err := sendMail(email, firstname, phone)
				//logs.Log.Error("err: ", err)
				//if err == nil {
				/*http.Redirect(w, r, "https://github.com/WTSSystem/WTS/raw/master/WTS.exe#", 301)
					//counter++
				} else {
					logs.Log.Error("sendMail error:", err)
				}* /
			} else {
				logs.Log.Error("user error:", err)
			}
		} else {
			logs.Log.Error("Email not valid:", user.Email)
		}
	} /*else {
		logs.Log.Error("Name or Email not valid:", firstname, email)
	}* /

	//logs.Log.Error("indexHandler", "enter")

	t, err := template.ParseFiles("templates/pages/users/data.html")
	if err != nil {
		logs.Log.Error("template.ParseFiles", err.Error())
		return
	}

	data, err := userList(session.Role.HasUserRights())
	if err != nil {
		logs.Log.Error("usersList", err.Error())
		return
	}

	//logs.Log.Error(data)
	err = t.ExecuteTemplate(w, "users/data.html", data)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}
} */

/* func usersExecDeleteHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("users", "usersHandlers", "usersExecDeleteHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("core.ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriUserDel) {
		logs.Log.Error("HasRight UriUserDel", err.Error())
		return
	}

	user := User{}
	if _, ok := r.Form["id"]; ok {
		user.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	err = userDelete(user)
	if err != nil {
		logs.Log.Error("userDelete", err.Error())
	}

	t, err := template.ParseFiles("templates/pages/users/data.html")
	if err != nil {
		logs.Log.Error("template.ParseFiles", err.Error())
		return
	}

	data, err := userList(session.Role.HasUserRights())
	if err != nil {
		logs.Log.Error("usersList", err.Error())
		return
	}

	//logs.Log.Error(data)
	err = t.ExecuteTemplate(w, "users/data.html", data)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}

	//logs.Log.Error("indexHandler", "enter")
	/*	err = subscribesDelete(sub)
		if err != nil {
			t, err := template.ParseFiles("templates/subscribesDelete.html")
			if err != nil {
				logs.Log.Error("subscribesDelete", err.Error())
				return
			}
		}

		t.ExecuteTemplate(w, "subscribesDelete", data)* /
} */

/*t, err := template.ParseFiles("templates/pages/subscribes/data.html")
if err != nil {
	logs.Log.Error("subscribesListHandler.ParseFiles", err.Error())
	return
}

data, err := subscribesList()
if err != nil {
	logs.Log.Error("subscribesListHandler.subscribesList", err.Error())
	return
}

//logs.Log.Error(data)
err = t.ExecuteTemplate(w, "pages/subscribes/data.html", data)
if err != nil {
	logs.Log.Error("subscribesListHandler.ExecuteTemplate", err.Error())
	return
}*/
