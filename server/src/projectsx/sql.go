package products

import (
	"errors"
	"sqlbar/server/src/core"
	"sqlbar/server/src/logs"
	"sqlbar/server/src/postgresql"
	"strconv"
	"time"
)

type (
	// Product struct
	Product struct {
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
	}
)

type (
	// SQL struct
	SQL struct{}
)

var sql SQL

func (product *Product) toData() (data Data) {
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
}

func (products *Products) toData() (datas []Data) {
	datas = make([]Data, 0)
	for _, user := range products.List {
		data := user.toData()
		datas = append(datas, data)
	}
	return
}

// data func
func (apiSql *SQL) lists() (lists []DataList, err error) {
	logs.Log.PushFuncName("products", "sql", "lists")
	defer logs.Log.PopFuncName()

	query := "select id, name from tlists where isdeleted = false"
	query += " order by id"
	rows, err := postgresql.Db.Query(query)
	if err != nil {
		logs.Log.Error("postgresql.Db.Query", err)
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
}

// data func
func (apiSql *SQL) data(id int64, isadmin bool) (prods Products, err error) {
	logs.Log.PushFuncName("products", "sql", "data")
	defer logs.Log.PopFuncName()

	//logs.Log.Debug("prods", prods)
	query := "select p.id, p.title, p.shortdescription, p.fulldescription, p.price, p.template, p.registered, " +
		"coalesce(p.listid, -1), coalesce(l.name, ''), p.isdeleted from tproducts p left join tlists l on l.id = p.listid"
	var where = ""
	if !isadmin {
		where += " (p.isdeleted = false)"
	}
	if id >= 0 {
		if where != "" {
			where += " and "
		}
		where += " (p.id = " + strconv.FormatInt(id, 10) + ")"
	}

	if where != "" {
		query += " where " + where
	}
	query += " order by p.id"
	logs.Log.Debug("query", query)
	rows, err := postgresql.Db.Query(query)
	if err != nil {
		logs.Log.Error("postgresql.Db.Query", err)
		return
	}
	defer rows.Close()

	prods = Products{List: []Product{}}
	for rows.Next() {
		prod := Product{}
		err := rows.Scan(&prod.ID, &prod.Title, &prod.ShortDescription, &prod.FullDescription,
			&prod.Price, &prod.Template, &prod.Registered, &prod.ListID, &prod.ListName, &prod.IsDeleted)
		if err != nil {
			logs.Log.Error("rows.Scan", err)
			continue
		}

		//logs.Log.Error("DDD", prod.FullDescription)
		prod.ShortDescrStr = core.TrimToLen(prod.ShortDescription, core.MaxShortDescrLen)
		prod.FullDescriptionStr = core.TrimToLen(prod.FullDescription, core.MaxShortDescrLen)
		prod.RegisteredStr = core.TimeToA(prod.Registered)
		prods.List = append(prods.List, prod)
	}

	return
}

// insert func
func (apiSql *SQL) insert(prod *Product) (err error) {
	logs.Log.PushFuncName("products", "sql", "insert")
	defer logs.Log.PopFuncName()

	//logs.Log.Debug("prod", prod)
	prod.Registered = time.Now()
	query := "insert into tproducts (title, shortdescription, fulldescription, price, template, listid, registered) " +
		"values ($1, $2, $3, $4, $5, $6, $7) returning id"
	row := postgresql.Db.QueryRow(query, prod.Title, prod.ShortDescription,
		prod.FullDescription, prod.Price, prod.Template, prod.ListID, prod.Registered)
	if err != nil {
		logs.Log.Error("postgresql.Db.QueryRow", err)
		return
	}

	err = row.Scan(&prod.ID)
	if err != nil {
		logs.Log.Error("row.Scan", err)
		return
	}

	if prod.ID <= 0 {
		err = errors.New("no id")
	}

	return
}

// update func
func (apiSql *SQL) update(prod Product) (err error) {
	logs.Log.PushFuncName("products", "sql", "update")
	defer logs.Log.PopFuncName()

	//logs.Log.Error("productsUpdate", prod)
	query := "update tproducts " +
		"set title = $1, shortdescription = $2, fulldescription = $3, price = $4, template = $5, listid = $6" +
		"where id = $7"

	_, err = postgresql.Db.Exec(query, prod.Title, prod.ShortDescription, prod.FullDescription,
		prod.Price, prod.Template, prod.ListID, prod.ID)
	if err != nil {
		logs.Log.Error("postgresql.Db.Exec", err)
		return
	}

	/*sub.ID, err = res.LastInsertId()
	if err != nil {
		logs.Log.Error("productsAdd", err)
		return
	}

	//logs.Log.Error("Subscribed:", firstname, lastname, phone, email, registered)
	if sub.ID <= 0 {
		err = errors.New("no id")
	}*/

	return
}

// delete func
func (apiSql *SQL) delete(prod Product) (err error) {
	logs.Log.PushFuncName("products", "sql", "delete")
	defer logs.Log.PopFuncName()

	logs.Log.Error("productsDelete", prod)
	//sql := "delete from tproducts where id = $1"
	sql := "update tproducts set isdeleted = true where id = $1"

	_, err = postgresql.Db.Exec(sql, prod.ID)
	if err != nil {
		logs.Log.Error("postgresql.Db.Exec", err)
		return
	}

	/*sub.ID, err = res.LastInsertId()
	if err != nil {
		logs.Log.Error("productsUpdate", err)
		return
	}

	if sub.ID <= 0 {
		logs.Log.Error("productsUpdate", err)
		return
	}*/

	return
}

/* func productsSelect(prod *Product) (err error) {
	logs.Log.PushFuncName("products", "productsSql", "productsSelect")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("prod:", prod)
	sql := "select id, title, shortdescription, fulldescription, price, template, registered from tproducts where isdeleted = false and id = $1"
	row := postgresql.Db.QueryRow(sql, &prod.ID)
	if (err != nil) && (err.Error() != "sql: no rows in result set") {
		logs.Log.Error("postgresql.Db.QueryRow", err)
		return
	}

	err = row.Scan(&prod.ID, &prod.Title, &prod.ShortDescription, &prod.FullDescription,
		&prod.Price, &prod.Template, &prod.Registered)
	if (err != nil) && (err.Error() != "sql: no rows in result set") {
		logs.Log.Error("row.Scan", err)
		return
	}
	logs.Log.Debug("select prod:", prod)
	return
} */

// updateTemplate func
func updateTemplate(prod *Product) (err error) {
	logs.Log.PushFuncName("products", "sql", "updateTemplate")
	defer logs.Log.PopFuncName()

	//logs.Log.Debug("prod:", prod)
	sql := "update tproducts " +
		"set title = $1, shortdescription = $2, fulldescription = $3, price = $4, template = $5" +
		"where id = $6"

	_, err = postgresql.Db.Exec(sql, prod.Title, prod.ShortDescription, prod.FullDescription,
		prod.Price, prod.Template, prod.ID)
	if err != nil {
		logs.Log.Error("postgresql.Db.Exec", err)
		return
	}

	return
}
