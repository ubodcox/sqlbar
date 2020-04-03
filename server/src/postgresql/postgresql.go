package postgresql

import (
	sql "database/sql"
	"io/ioutil"
	"os"
	"sqlbar/server/src/config"
	"sqlbar/server/src/logs"
)

type (
	// PostgreDb type
	PostgreDb struct {
		*sql.DB
	}
)

// Db var
var Db *PostgreDb

// DbInit func
func DbInit() {
	logs.Log.PushFuncName("postgresql", "postgresql", "DbInit")
	defer logs.Log.PopFuncName()

	connStr := "user=" + config.DbUser + " " +
		"password=" + config.DbPassword + " " +
		"dbname=" + config.DbName + " " +
		"sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		logs.Log.Error("sql.Open", err)
	}
	Db = &PostgreDb{conn}
}

// DbDeinit func
func DbDeinit() {
	Db.Close()
}

// DbUpdate func
func DbUpdate() {
	logs.Log.PushFuncName("postgresql", "postgresql", "DbUpdate")
	defer logs.Log.PopFuncName()

	//TODO: вынести похожий код в процедуру

	b, err := ioutil.ReadFile("./files/sql/procedures.sql")
	if err != nil {
		logs.Log.Error("ioutil.ReadFile 'procedures.sql'", err)
		os.Exit(1)
	}

	sql := string(b)
	_, err = Db.Exec(sql)
	if err != nil {
		logs.Log.Error("Db.Exec 'procedures.sql'", err)
	}
	//------------------------------------------------
	b, err = ioutil.ReadFile("./files/sql/create.sql")
	if err != nil {
		logs.Log.Error("ioutil.ReadFile 'create.sql'", err)
		os.Exit(1)
	}

	sql = string(b)
	_, err = Db.Exec(sql)
	if err != nil {
		logs.Log.Error("Db.Exec 'create.sql'", err)
		return
	}
	//------------------------------------------------
	b, err = ioutil.ReadFile("./files/sql/exec.sql")
	if err != nil {
		logs.Log.Error("ioutil.ReadFile 'exec.sql'", err)
		os.Exit(1)
	}

	sql = string(b)
	_, err = Db.Exec(sql)
	if err != nil {
		logs.Log.Error("Db.Exec 'exec.sql'", err)
		return
	}
	//- chat -----------------------------------------
	/* b, err = ioutil.ReadFile("./files/sql/chat.sql")
	if err != nil {
		logs.Log.Error("ioutil.ReadFile 'chat.sql'", err)
		os.Exit(1)
	}

	sql = string(b)
	_, err = Db.Exec(sql)
	if err != nil {
		logs.Log.Error("Db.Exec 'chat.sql'", err)
		return
	}	 */
}
