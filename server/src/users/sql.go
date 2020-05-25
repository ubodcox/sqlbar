package users

import (
	"errors"
	"sqlbar/server/src/core"
	"sqlbar/server/src/logs"
	"sqlbar/server/src/postgresql"
	"strconv"
	"time"
)

type (
	// User struct
	User struct {
		ID            int64
		Password      string
		FirstName     string
		LastName      string
		Post          string
		Phone         string
		Email         string
		Registered    time.Time
		RegisteredStr string
		Role          core.UserRoles
		RoleInt       int
		IsDeleted     bool
	}
	// Users struct
	Users struct {
		List []User
	}
)

type (
	// SQL struct
	SQL struct{}
)

var sql SQL

//TODO: add encrypt for password

func (user *User) toData() (data Data) {
	data.ID = user.ID
	data.FirstName = user.FirstName
	data.LastName = user.LastName
	data.Post = user.Post
	data.Phone = user.Phone
	data.Email = user.Email
	data.Registered = user.RegisteredStr
	data.Role = DataRole{int64(user.Role.ID), user.Role.Caption}
	data.RoleInt = user.RoleInt
	data.IsDeleted = user.IsDeleted
	return
}

func (users *Users) toData() (datas []Data) {
	datas = make([]Data, 0)
	for _, user := range users.List {
		data := user.toData()
		datas = append(datas, data)
	}
	return
}

// data func
func (apiSql *SQL) data(id int64, isadmin bool) (users Users, err error) {
	logs.Log.PushFuncName("users", "usersSql", "userList")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	//logs.Log.Error("user.List")
	sql := "select id, firstname, lastname, post, phone, email, registered, roleid, isdeleted from tusers"
	var where string = ""
	if id >= 0 {
		where += "(id = " + strconv.FormatInt(id, 10) + ")"
	}
	if !isadmin {
		if where != "" {
			where += " AND "
		}
		where += "(isdeleted = false)"
	}
	if where != "" {
		sql += " where " + where
	}
	sql += " order by id"
	logs.Log.Warning(sql)
	rows, err := postgresql.Db.Query(sql)
	if err != nil {
		logs.Log.Error("postgresql.Db.Query", users, err)
		return
	}
	defer rows.Close()

	users = Users{List: []User{}}
	for rows.Next() {
		user := User{}
		var roleid int
		err := rows.Scan(&user.ID, &user.FirstName,
			&user.LastName, &user.Post, &user.Phone, &user.Email, &user.Registered, &roleid, &user.IsDeleted)
		if err != nil {
			logs.Log.Error("rows.Scan", err)
			continue
		}
		user.Role = core.GetRole(core.UserRoleID(roleid))
		user.RoleInt = roleid
		user.RegisteredStr = core.TimeToA(user.Registered)

		users.List = append(users.List, user)
	}

	return
}

// insert func
func (apiSql *SQL) insert(user User) (err error) {
	logs.Log.PushFuncName("users", "sql", "insert")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	logs.Log.Debug("user:", user)
	user.Registered = time.Now()

	sql := "insert into tusers (password, firstname, lastname, post, phone, email, registered, roleid) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8) returning id"
	row := postgresql.Db.QueryRow(sql, user.Password, user.FirstName,
		user.LastName, user.Post, user.Phone, user.Email, user.Registered, user.Role.ID)
	/*if err != nil {
		pan ic(err)
	}*/

	err = row.Scan(&user.ID)
	//user.ID, err = row.LastInsertId()
	if err != nil {
		logs.Log.Error("row.Scan", err)
		return
	}

	if user.ID <= 0 {
		err = errors.New("no id")
	}

	return
}

// update func
func (apiSql *SQL) update(user User) (err error) {
	logs.Log.PushFuncName("users", "sql", "update")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	logs.Log.Debug("user:", user)
	if user.Password != "" {
		sql := "update tusers " +
			"set password = $1, firstname = $2, lastname = $3, post = $4, phone = $5, " +
			"email = $6, roleid = $7 " +
			"where id = $8"
		_, err = postgresql.Db.Exec(sql, user.Password, user.FirstName,
			user.LastName, user.Post, user.Phone, user.Email, user.Role.ID, user.ID)
		if err != nil {
			logs.Log.Error("postgresql.Db.Exec", err)
			return
		}
	} else {
		sql := "update tusers " +
			"set firstname = $1, lastname = $2, post = $3, phone = $4, " +
			"email = $5, roleid = $6 " +
			"where id = $7"
		_, err = postgresql.Db.Exec(sql, user.FirstName,
			user.LastName, user.Post, user.Phone, user.Email, user.Role.ID, user.ID)
		if err != nil {
			logs.Log.Error("postgresql.Db.Exec", err)
			return
		}
	}
	return
}

func (apiSql *SQL) delete(user User) (err error) {
	logs.Log.PushFuncName("users", "sql", "delete")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	logs.Log.Debug("user:", user)
	//sql := "delete from tusers where id = $1"
	query := "update tusers set isdeleted = true where id = $1"

	_, err = postgresql.Db.Exec(query, user.ID)
	if err != nil {
		logs.Log.Error("postgresql.Db.Exec", err)
		return
	}

	return
}

/* func userSelect(user *User) (err error) {
	logs.Log.PushFuncName("users", "usersSql", "userSelect")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("user:", user)
	sql := "select id, password, firstname, lastname, post, phone, email, registered, roleid " +
		"from tusers where (isdeleted = false) and (id = $1)"
	row := postgresql.Db.QueryRow(sql, &user.ID)
	/*if err != nil {
		logs.Log.Error("user.Select", user.ID, err)
		return
	}* /
	//defer rows.Close()

	//rows.Next()
	var roleid int
	err = row.Scan(&user.ID, &user.Password, &user.FirstName,
		&user.LastName, &user.Post, &user.Phone, &user.Email, &user.Registered, &roleid)
	if err != nil {
		logs.Log.Error("row.Scan", user.ID, err)
		return
	}
	user.RegisteredStr = core.TimeToA(user.Registered)
	user.Role = core.GetRole(core.UserRoleID(roleid))
	user.RoleInt = roleid
	return
} */

func (apiSql *SQL) exists(user *User) (err error) {
	logs.Log.PushFuncName("users", "sql", "exists")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN", "email:", user.Email, " pass: ", user.Password)
	//defer logs.Log.Debug("END")

	query := "select id from tusers where isdeleted = false and email = $1 and password = $2"
	row := postgresql.Db.QueryRow(query, &user.Email, &user.Password)
	/*if err != nil {
		logs.Log.Error("user.userExists", err)
		return
		//pa nic(err)
	}*/

	err = row.Scan(&user.ID)
	if err != nil {
		logs.Log.Warning("END urow.Scan", err)
		return
	}

	logs.Log.Debug("END user exists")
	return
}
