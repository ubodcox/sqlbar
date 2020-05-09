package tasksx

import (
	sqlr "database/sql"
	"errors"
	"sqlbar/server/src/logs"
	"sqlbar/server/src/postgresqlx"
	"strings"
)

type (
	// Product struct
	/* Product struct {
		ID                 int64
		Title              string
		ShortDescription   string
		FullDescription    string
		Price              float64
		Registered         time.Time
		Template           string
		Image1             string
		Image2             string
		Image3             string
		Image4             string
		Image5             string
		IsDeleted          bool
		RegisteredStr      string
		ShortDescrStr      string
		FullDescriptionStr string
		LastInsertId       int64
		ListID             int64
		ListName           string
	}
	// Products struct
	Products struct {
		List []Product
	} */

	/*//Answer struct
	Answer struct {
		Col1 string `json:"col1"`
		Col2 string `json:"col2"`
	}

	//AnswerList struct
	AnswerList struct {
		List []Answer `json:"list"`
	}*/

	// Task struct
	Task struct {
		ID    int64  `json:"id"`
		Text  string `json:"text"`
		Cols  string `json:"cols"`
		Stars int    `json:"stars"`
	}

	//ResultList struct
	ResultList struct {
		IsCorrect bool                     `json:"iscorrect"`
		List      []map[string]interface{} `json:"list"`
	}
)

type (
	// SQL struct
	SQL struct{}
)

var sql SQL

/* func (product *Product) toData() (data Data) {
	data.ID = product.ID
	data.Title = product.Title
	data.ShortDescription = product.ShortDescrStr
	data.FullDescription = product.FullDescriptionStr
	data.Price = product.Price
	data.Registered = product.RegisteredStr
	data.Template = product.Template
	data.Image1 = product.Image1
	data.Image2 = product.Image2
	data.Image3 = product.Image3
	data.Image4 = product.Image4
	data.Image5 = product.Image5
	data.IsDeleted = product.IsDeleted
	data.ListID = product.ListID
	data.ListName = product.ListName
	return
} */

/* func (products *Products) toData() (datas []Data) {
	datas = make([]Data, 0)
	for _, user := range products.List {
		data := user.toData()
		datas = append(datas, data)
	}
	return
} */

// data func
/* func (apiSql *SQL) lists() (lists []DataList, err error) {
	logs.Log.PushFuncName("products", "sql", "lists")
	defer logs.Log.PopFuncName()

	query := "select id, name from tlists where isdeleted = false"
	query += " order by id"
	rows, err := postgresqlx.Db.Query(query)
	if err != nil {
		logs.Log.Error("postgresqlx.Db.Query", err)
		return
	}
	defer rows.Close()

	lists = make([]DataList, 0)
	for rows.Next() {
		list := DataList{}
		err := rows.Scan(&list.ID, &list.Name)
		if err != nil {
			logs.Log.Error("rows.Scan", err)
			continue
		}

		lists = append(lists, list)
	}

	return
} */

// data func
func (apiSql *SQL) getTask(id int64) (tasks []Task, err error) {
	logs.Log.PushFuncName("tasks", "sql", "getTask")
	defer logs.Log.PopFuncName()

	query := "select id, text, col1, col2, stars from ttasks"
	var rows *sqlr.Rows
	if id > 0 {
		query += " where id = $1"
		rows, err = postgresqlx.Db.Query(query, id)
	} else {
		rows, err = postgresqlx.Db.Query(query)
	}
	if err != nil {
		logs.Log.Error("postgresqlx.Db.Query", err)
		return
	}
	defer rows.Close()

	tasks = make([]Task, 0)
	var col1, col2 string
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.ID, &task.Text, &col1, &col2, &task.Stars)
		if err != nil {
			logs.Log.Error("rows.Scan", err)
			continue
		}

		task.Cols = col1
		if col2 != "" {
			task.Cols += " " + col2
		}

		tasks = append(tasks, task)
	}

	return
}

func (apiSql *SQL) cols(id int64) (count int, col1, col2 string) {
	logs.Log.PushFuncName("tasks", "sql", "cols")
	defer logs.Log.PopFuncName()

	query := "select col1, col2 from ttasks where id = $1"
	row := postgresqlx.Db.QueryRow(query, id)
	if row == nil {
		logs.Log.Error("postgresqlx.Db.QueryRow", "empty row")
		return
	}

	err := row.Scan(&col1, &col2)
	if err != nil {
		logs.Log.Error("row.Scan", err)
		return
	}

	count = 1
	if col2 != "" {
		count++
	}

	return
}

func (apiSql *SQL) checkSQL(sql string) bool {
	if strings.Contains(sql, "insert") {
		return false
	}

	if strings.Contains(sql, "delete") {
		return false
	}

	if strings.Contains(sql, "update") {
		return false
	}

	if strings.Contains(sql, "drop") {
		return false
	}

	if strings.Contains(sql, "create") {
		return false
	}

	return true
}

func (apiSql *SQL) compareSQL(baseSQL, userSQL string) bool {
	/*"SELECT d.* FROM "+baseSql+" d
	    LEFT OUTER JOIN "+userSql+" r ON r.id = d.id
		WHERE r.id IS NULL"


		"SELECT d.* FROM dump_data d
	    LEFT OUTER JOIN real_data r ON r.id = d.id
		WHERE r.id IS NULL"*/
	return false
}

// CheckTask checks tasks sql and returning result
func (apiSql *SQL) CheckTask(id int64, sql string) (results ResultList, err error) {
	logs.Log.PushFuncName("tasks", "sql", "checkTask")
	defer logs.Log.PopFuncName()

	if !apiSql.checkSQL(sql) {
		err = errors.New("bad sql")
		return
	}

	//colCount, _, _ := apiSql.cols(id) //colCount, col1, col2 := apiSql.cols(id)

	rows, err := postgresqlx.Dbx.Queryx(sql)
	for rows.Next() {
		result := make(map[string]interface{})
		err = rows.MapScan(result)
		if err != nil {
			logs.Log.Error("postgresql.Db.Query", err)
			return
		}

		logs.Log.Warning(result)

		results.List = append(results.List, result)
		//jsonString, _ := json.Marshal(datas)
	}

	/* 	query := sql
	   	rows, err := postgresql.Db.Query(query)
	   	if err != nil {
	   		logs.Log.Error("postgresql.Db.Query", err)
	   		return
	   	}
	   	defer rows.Close()

	   	answers = AnswerList{false, make([]Answer, 0)}
	   	for rows.Next() {
	   		answer := Answer{}
	   		if colCount == 1 {
	   			err = rows.Scan(&answer.Col1)
	   		} else {
	   			err = rows.Scan(&answer.Col1, &answer.Col2)
	   		}
	   		if err != nil {
	   			logs.Log.Error("rows.Scan", err)
	   			continue
	   		}

	   		answers.List = append(answers.List, answer)
	   	} */

	logs.Log.Debug("5")

	return
}
