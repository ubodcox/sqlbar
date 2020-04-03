package core

import (
	"sqlbar/server/src/logs"
	"sqlbar/server/src/postgresql"
)

type (
	// UserRoleID type
	UserRoleID int

	// UserRoles struct
	UserRoles struct {
		ID      UserRoleID
		Caption string
		//Rights  map[UserRightID]bool
	}

	// Roles struct
	Roles struct {
		List map[UserRoleID]UserRoles
	}
)

const (
	// UroAdmin 1 "Администратор"
	UroAdmin UserRoleID = iota + 1
	// UroModerator 2 "Модератор"
	UroModerator
	// UroSysAdmin 3 "Системный администратор"
	UroSysAdmin
	// UroSuperAdmin 4 "Супер администратор"
	UroSuperAdmin
)

var roles Roles

// InitRoles func
func InitRoles() {
	logs.Log.PushFuncName("core", "roles", "init")
	defer logs.Log.PopFuncName()

	logs.Log.Info("IMPORTED")

	// make
	roles = Roles{
		List: make(map[UserRoleID]UserRoles),
	}
	// fill
	roles.refreshRoles()
}

// GetRole func
func GetRole(role UserRoleID) UserRoles {
	return roles.List[role]
}

func (r *Roles) refreshRoles() {
	logs.Log.PushFuncName("core", "roles", "refreshRoles")
	defer logs.Log.PopFuncName()

	//logs.Log.Debug("BEGIN")

	sql := "select id, caption from troles"
	rows, err := postgresql.Db.Query(sql)
	if err != nil {
		logs.Log.Error("postgresql.Db.Query", err)
		return
	}
	defer rows.Close()

	// clear
	for i := range r.List {
		delete(r.List, i)
	}

	// fill
	for rows.Next() {
		var roleid int
		var caption string
		err := rows.Scan(&roleid, &caption)
		if err != nil {
			logs.Log.Error("rows.Scan", err)
			continue
		}

		/*r.List[UserRoleID(roleid)] = UserRoles{
			ID:      UserRoleID(roleid),
			Caption: caption,
			Rights:  make(map[UserRightID]bool),
		}
		r.List[UserRoleID(roleid)].RefreshRights()*/
		//role.Rights[UserRightID(roleid)] = true
	}
	//logs.Log.Debug("END", roles)
	return
}
