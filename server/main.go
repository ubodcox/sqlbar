//TODO: add user table to sql
//TODO: add role table to sql

package main

import (
	//_ "chat"
	"sqlbar/server/src/config"
	//_ "contacts"
	//_ "cutter"
	//_ "images"
	//_ "lists"
	"sqlbar/server/src/logs"
	"sqlbar/server/src/postgresqlx"

	"net/http"
	//_ "orders"
	//_ "products"
	"strconv"
	"time"

	//_ "users"
	//_ "other"
	//_ "tasksx"

	_ "github.com/lib/pq"
)

// color console
// go get github.com/fatih/color

// Server settings
const (
	isLogInFile = true
)

func addHandle(path string) {
	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir("."+path))))
}

func main() {
	if isLogInFile {
		if logs.Init() {
			defer logs.Deinit()
		}
	}
	logs.Log.PushFuncName("MAIN", "main", "main")
	defer logs.Log.PopFuncName()

	//addHandle("/assets/")
	//addHandle("/templates/")

	//addHandle("/files/templates/")

	/* addHandle("/templates/bower_components/bootstrap/dist/css/")
	addHandle("/templates/bower_components/font-awesome/css/")
	addHandle("/templates/bower_components/Ionicons/css/")
	addHandle("/templates/bower_components/datatables.net-bs/css/")
	addHandle("/templates/dist/css/")
	addHandle("/templates/dist/css/skins/")
	addHandle("/templates/dist/css/alt/")
	addHandle("/templates/mycss/")
	*/
	addHandle("/pages/")

	//addHandle("/templates/pages/toast/")

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(config.ServerPort),
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}
	logs.Log.Info("SERVER STARTED at port: " + strconv.Itoa(config.ServerPort))

	postgresqlx.DbInit()
	postgresqlx.DbUpdate()
	//core.InitRoles()
	defer postgresqlx.DbDeinit()

	logs.Log.Info("PostgreSQLx STARTED")

	err := s.ListenAndServe()
	if err != nil {
		logs.Log.Info("END")
	}
}
