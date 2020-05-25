package rating

import "sqlbar/server/src/logs"

type (
	// DB struct
	DB struct{}
)

var db DB

func (db *DB) ratingGet(mode string) (rating DBRating, err error) {
	logs.Log.PushFuncName("rating", "db", "ratingGet")
	defer logs.Log.PopFuncName()

	if mode == "all" {

	} else if mode == "month" {

	} else if mode == "week" {

	}

	//logs.Log.Error("user.List")
	/*sql := "select id, firstname, lastname, post, phone, email, registered, roleid, isdeleted from tusers"
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
	}*/

	return
}
