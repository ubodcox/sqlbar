package postgresql

/*type
cnfitylpge (
	params struct {
		count int
	}
)*/
/*
// Subscribe func
func Subscribe(firstname, lastname, phone, email string) (err error) {
	logs.Log.PushFuncName("postgresql", "sql", "Subscribe")
	defer logs.Log.PopFuncName()

	registered := time.Now()
	sql := "insert into torders (firstname, lastname, phone, email, registered) " +
		"values ($1, $2, $3, $4, $5) returning id"
	//"values ('" + firstname + "', '" + lastname + "', '" + phone + "' , '" + email + "', '" + registered + "') returning id"
	//db.Exec("INSERT INTO 'UserAccount' ('email", "login_time") VALUES ($1, $2)',"human@example.com",time.Now())
	_, err = Db.Exec(sql, firstname, lastname, phone, email, registered)
	//log.Println("insert")
	if err != nil {
		logs.Log.Error("Db.Exec", err)
		return
	}
	logs.Log.Debug("Subscribed:", firstname, lastname, phone, email, registered)
	return
}

func subsCount() int {
	logs.Log.PushFuncName("postgresql", "sql", "subsCount")
	defer logs.Log.PopFuncName()

	sql := "select count(*) from torders"
	rows := Db.QueryRow(sql)
	//if err != nil {
	//	pa nic(err)
	//}
	//defer rows.Close()
	//products := []product{}
	param := params{}
	err := rows.Scan(&param.count)
	if err != nil {
		logs.Log.Error("rows.Scan", err)
		return -1
	}
	return param.count
}*/
