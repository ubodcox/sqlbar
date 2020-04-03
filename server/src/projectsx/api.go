package products

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sqlbar/server/src/config"
	"sqlbar/server/src/core"

	//"sqlbar/server/src/images"
	"sqlbar/server/src/logs"
	"strconv"
	//"strings"
)

type (
	// API struct
	API struct {
	}

	// DataList struct
	DataList struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	// Data struct
	Data struct {
		ID               int64   `json:"id"`
		Title            string  `json:"title"`
		ShortDescription string  `json:"short"`
		FullDescription  string  `json:"full"`
		Price            float64 `json:"price"`
		Registered       string  `json:"registered"`
		Template         string  `json:"template"`
		Image1           string  `json:"Image1"`
		Image2           string  `json:"Image2"`
		Image3           string  `json:"Image3"`
		Image4           string  `json:"Image4"`
		Image5           string  `json:"Image5"`
		IsDeleted        bool    `json:"isdeleted"`
		ListID           int64   `json:"listid"`
		ListName         string  `json:"listname"`
	}

	// SettingsData struct
	SettingsData struct {
		Host      string `json:"host"`
		Port      int    `json:"port"`
		ProductID int64  `json:"productid"`
	}
	//TemplateData struct
	TemplateData struct {
		Product  Product      `json:"product"`
		Settings SettingsData `json:"host"`
	}
	//IDData struct
	IDData struct {
		ID   int64 `json:"id"`
		Port int   `json:"port"`
	}

	// response struct
	response struct {
		core.DefaultResponse
		Datas []Data     `json:"datas"`
		Lists []DataList `json:"lists"`
	}
)

//DefaultTemplate var
var DefaultTemplate = "mojo"

var api API

// makeResponse func
func makeResponse(code int, msg string, s *core.Session, d []Data, l []DataList) *response {
	return &response{
		*core.MakeDefaultResponse(code, msg, s),
		d, l,
	}
}

func (api *API) product(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "api", "product")
	defer logs.Log.PopFuncName()

	//log.Println("api.product BEGIN")
	r.ParseForm()

	// no cookies

	core.EnableCors(&w)

	prod := Product{}
	if _, ok := r.Form["id"]; ok {
		prod.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	_, err := sql.data(prod.ID, false)
	if err != nil {
		logs.Log.Error("productsSelect", err.Error())
		return
	}
	/*
		imgs, err := images.List(prod.ID)
		if err != nil {
			logs.Log.Error("images.List", err.Error())
			return
		}
		//----------------------------------------------------
		if len(imgs.List) > 0 {
			img := imgs.List[0]
			prod.Image1 = "./images?id=" + img.FileName
		}
		if len(imgs.List) > 1 {
			img := imgs.List[1]
			prod.Image2 = "./images?id=" + img.FileName
		}
		if len(imgs.List) > 2 {
			img := imgs.List[2]
			prod.Image3 = "./images?id=" + img.FileName
		}
		if len(imgs.List) > 3 {
			img := imgs.List[3]
			prod.Image4 = "./images?id=" + img.FileName
		}
		if len(imgs.List) > 4 {
			img := imgs.List[4]
			prod.Image5 = "./images?id=" + img.FileName
		}*/

	/* if prod.Template == "" {
		t, err := template.ParseFiles("templates/pages/products/info.html")
		if err != nil {
			log.Println("productsModalsHandler.ParseFiles", err.Error())
			return
		}

		err = t.ExecuteTemplate(w, "products/info.html", prod)
		if err != nil {
			log.Println("productsModalsHandler.ExecuteTemplate", err.Error())
			return
		}
	} else {
		t, err := template.ParseFiles("files/templates/" + prod.Template)
		if err != nil {
			log.Println("productsModalsHandler.ParseFiles", err.Error())
			return
		}
		log.Println("test", prod)
		err = t.ExecuteTemplate(w, prod.Template, prod) //"pages/products/info.html", prod)
		if err != nil {
			log.Println("productsModalsHandler.ExecuteTemplate", err.Error())
			return
		}
	} */

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(prod)
}

func (api *API) api(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "api", "DefaultHandler")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeDefaultResponse(-1, "error parse form", nil))
		return
	}

	var id int64 = -1
	if _, ok := r.Form["id"]; ok {
		id, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	core.EnableCors(&w)

	t, err := template.ParseFiles("templates/pages/products/api.html")
	if err != nil {
		logs.Log.Error("ParseFiles", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, "products/api.html", IDData{ID: id, Port: config.ServerPort})
	if err != nil {
		logs.Log.Error("ExecuteTemplate", err.Error())
		return
	}
}

/*
func (api *API) data(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "api", "data")
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

	var id int64 = -1 //list := List{}
	if _, ok := r.Form["id"]; ok {
		id, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	core.EnableCors(&w)
	products, err := sql.data(id, session.Role.HasProductRights())
	if err != nil {
		logs.Log.Error("sql.Data", err.Error())
		json.NewEncoder(w).Encode(core.MakeDefaultResponse(-4, "error read data from db", session))
		return
	}

	lists, err := sql.lists()
	if err != nil {
		logs.Log.Error("sql.Data", err.Error())
		json.NewEncoder(w).Encode(core.MakeDefaultResponse(-4, "error read data from db", session))
		return
	}

	json.NewEncoder(w).Encode(
		makeResponse(1, "", session, products.toData(), lists),
	)

	//core.MakeResponse(w, r, 1, "")
	logs.Log.Debug("code: 1") /* ,
	ResponseList{
		Code: 1,
		Msg:  "",
		Lists: lists,
	}) */ /*
}*/
/*
func (api *API) insert(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "api", "insert")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("BEGIN")
	defer logs.Log.Debug("END")

	//session := r.Context().Value("Session").(*core.Session)

	//err := r.ParseForm()
	err := r.ParseMultipartForm(20 * 1024 * 1024 * 1024)
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "error parse form"))
		//json.NewEncoder(w).Encode(core.MakeResponse(-1, "error parse form"))
		return
	}

	core.EnableCors(&w)
	prod := Product{}
	/* var action string
	if _, ok := r.PostForm["action"]; ok {
		action = r.PostForm["action"][0]
	} */ /*
	if _, ok := r.PostForm["title"]; ok {
		prod.Title = r.PostForm["title"][0]
	}
	if _, ok := r.PostForm["shortdescription"]; ok {
		prod.ShortDescription = r.PostForm["shortdescription"][0]
	}
	if _, ok := r.PostForm["fulldescription"]; ok {
		prod.FullDescription = r.PostForm["fulldescription"][0]
	}
	if _, ok := r.Form["listid"]; ok {
		prod.ListID, _ = strconv.ParseInt(r.Form["listid"][0], 10, 64)
	}

	if _, ok := r.PostForm["price"]; ok {
		var err error
		price := r.PostForm["price"][0]
		logs.Log.Warning("price", price)
		prod.Price, err = strconv.ParseFloat(strings.Replace(price, ",", ".", -1), 64)
		if err != nil {
			logs.Log.Error("strconv.ParseFloat", err)
			return
		}
	}

	filename, err := images.ImportImage(r, "template", "products")
	if err == nil {
		prod.Template = filename
		images.SetTemplate(prod.Template)
	}

	if prod.Title != "" && prod.Price > 0 {
		err := sql.insert(&prod)
		if err != nil {
			logs.Log.Error("sql.insert", err.Error())
			json.NewEncoder(w).Encode(core.MakeEmptyResponse(-4, "error write to db"))
			return
		}

		// images --------------------------------------------
		filename, err := images.ImportImage(r, "image1", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		filename, err = images.ImportImage(r, "image2", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		filename, err = images.ImportImage(r, "image3", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		filename, err = images.ImportImage(r, "image4", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		filename, err = images.ImportImage(r, "image5", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		//--------------------------------------------
	}

	url := "../../../products/data.html"
	json.NewEncoder(w).Encode(core.MakeEmptyResponse(1, url))
	logs.Log.Debug("code: 1", *core.MakeEmptyResponse(1, url))
}*/
/*
func (api *API) update(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "api", "update")
	defer logs.Log.PopFuncName()

	//logs.Log.Debug("BEGIN")
	//defer logs.Log.Debug("END")

	//session := r.Context().Value("Session").(*core.Session)

	err := r.ParseMultipartForm(20 * 1024 * 1024 * 1024)
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "error parse form"))
		//json.NewEncoder(w).Encode(core.MakeResponse(-1, "error parse form"))
		return
	}

	/* err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "error parse form"))
		//json.NewEncoder(w).Encode(core.MakeResponse(-1, "error parse form"))
		return
	} */
/*
	core.EnableCors(&w)
	prod := Product{}
	var action string
	if _, ok := r.PostForm["action"]; ok {
		action = r.PostForm["action"][0]
	}

	logs.Log.Warning("id", r.PostForm["id"])
	logs.Log.Warning("title", r.PostForm["title"])

	if _, ok := r.PostForm["id"]; ok {
		prod.ID, _ = strconv.ParseInt(r.PostForm["id"][0], 10, 64)
	}
	if _, ok := r.PostForm["title"]; ok {
		prod.Title = r.PostForm["title"][0]
		logs.Log.Warning("title1", prod.Title)
		logs.Log.Warning("title2", r.PostForm["title"][0])
	}
	if _, ok := r.PostForm["shortdescription"]; ok {
		prod.ShortDescription = r.PostForm["shortdescription"][0]
	}
	if _, ok := r.PostForm["fulldescription"]; ok {
		prod.FullDescription = r.PostForm["fulldescription"][0]
	}
	if _, ok := r.PostForm["price"]; ok {
		var err error
		prod.Price, err = strconv.ParseFloat(strings.Replace(r.PostForm["price"][0], ",", ".", -1), 64)
		if err != nil {
			logs.Log.Error("strconv.ParseFloat", err)
			return
		}
	}
	if _, ok := r.PostForm["listid"]; ok {
		prod.ListID, _ = strconv.ParseInt(r.PostForm["listid"][0], 10, 64)
	}

	if prod.ID < 0 {
		logs.Log.Error("check id", "no id")
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "no id"))
		return
	}

	filename, err := images.ImportImage(r, "template", "templates")
	//logs.Log.Debug("ImportImage", filename, err)
	if (err == nil) || (err.Error() == "update") {
		prod.Template = filename
		images.SetTemplate(prod.Template)
	}
	//logs.Log.Error("TEMPLATE2 ", filename, err)

	if prod.Title != "" && prod.Price > 0 {
		err := sql.update(prod)
		if err != nil {
			logs.Log.Error("sql.update:", err)
			return
		}

		// images --------------------------------------------
		filename, err := images.ImportImage(r, "image1", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		filename, err = images.ImportImage(r, "image2", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		filename, err = images.ImportImage(r, "image3", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		filename, err = images.ImportImage(r, "image4", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		filename, err = images.ImportImage(r, "image5", "products")
		if err == nil {
			image := images.Image{}
			image.FileName = filename
			image.ProductID = prod.ID
			err = images.Add(&image)
		}
		//--------------------------------------------
	}

	if prod.Title == "" {
		logs.Log.Error("check name", "no name")
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "no name"))
		return
	}

	/* err = sql.update(id, prod.Title)
	if err != nil {
		logs.Log.Error("sql.update", err.Error())
		core.MakeResponse(w, r, session, -4, "error write to db")
		return
	} */

/*url := "../../../products/data.html"
json.NewEncoder(w).Encode(core.MakeEmptyResponse(1, url))
logs.Log.Debug("code: 1")*/

//logs.Log.Warning("action: "+action)/*
/*	if action == "publish" {
		//logs.Log.Warning("publish", "info.html?id="+strconv.FormatInt(prod.ID, 10))
		url := "../../../products/info.html?id=" + strconv.FormatInt(prod.ID, 10)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(1, url))
		//http.Redirect(w, r, "info.html?id="+strconv.FormatInt(prod.ID, 10), 303)
	} else if action == "save" {
		//logs.Log.Warning("save", "data.html")
		url := "../../../products/data.html"
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(1, url))
		//http.Redirect(w, r, "data.html", 303)
	}
}*/

func (api *API) delete(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "api", "delete")
	defer logs.Log.PopFuncName()

	//logs.Log.Debug("BEGIN")
	//defer logs.Log.Debug("END")

	//session := r.Context().Value("Session").(*core.Session)

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "error parse form"))
		//json.NewEncoder(w).Encode(core.MakeResponse(-1, "error parse form"))
		return
	}

	core.EnableCors(&w)
	prod := Product{}
	if _, ok := r.Form["id"]; ok {
		prod.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}
	if prod.ID < 0 {
		logs.Log.Error("check id", "no id")
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-2, "no id"))
		return
	}

	/* order := Order{}
	if _, ok := r.Form["id"]; ok {
		order.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	} */

	err = sql.delete(prod)
	if err != nil {
		logs.Log.Error("sql.delete", err.Error())
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-4, "error write to db"))
		return
	}

	url := "../../../products/data.html"
	json.NewEncoder(w).Encode(core.MakeEmptyResponse(1, url))
	logs.Log.Debug("code: 1")
}

func (api *API) info(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "api", "info")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	// no cookies

	logs.Log.Warning("info")

	prod := Product{}
	if _, ok := r.Form["id"]; ok {
		prod.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	prods, err := sql.data(prod.ID, false)
	if err != nil {
		logs.Log.Error("productsSelect", err.Error())
		return
	}

	if len(prods.List) == 0 {
		logs.Log.Error("Length prods == 0", err.Error())
		return
	}

	prod = prods.List[0]
	/*
		imgs, err := images.List(prod.ID)
		if err != nil {
			logs.Log.Error("images.List", err.Error())
			return
		}
		//----------------------------------------------------
		if len(imgs.List) > 0 {
			img := imgs.List[0]
			prod.Image1 = "./images?id=" + img.FileName
		}
		if len(imgs.List) > 1 {
			img := imgs.List[1]
			prod.Image2 = "./images?id=" + img.FileName
		}
		if len(imgs.List) > 2 {
			img := imgs.List[2]
			prod.Image3 = "./images?id=" + img.FileName
		}
		if len(imgs.List) > 3 {
			img := imgs.List[3]
			prod.Image4 = "./images?id=" + img.FileName
		}
		if len(imgs.List) > 4 {
			img := imgs.List[4]
			prod.Image5 = "./images?id=" + img.FileName
		}*/

	//TODO: !!!
	if prod.Template == "" {
		prod.Template = DefaultTemplate

	}
	// no template
	//logs.Log.Error("ParseFiles", "no template")
	//return
	/*t, err := template.ParseFiles("templates/pages/products/info.html")
	if err != nil {
		logs.Log.Error("ParseFiles", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, "products/info.html", prod)
	if err != nil {
		logs.Log.Error("ExecuteTemplate", err.Error())
		return
	}*/
	//} else {
	prod.ShortDescription = prod.ShortDescrStr
	prod.FullDescription = prod.FullDescriptionStr

	settingsData := SettingsData{Host: config.ServerHost, Port: config.ServerPort, ProductID: prod.ID}
	templateData := TemplateData{prod, settingsData}

	// individual template
	logs.Log.Warning("info2")
	t, err := template.ParseFiles("files/templates/" + prod.Template + "/index.html")
	if err != nil {
		logs.Log.Error("ParseFiles", err.Error())
		return
	}
	logs.Log.Warning("info3", prod)
	err = t.ExecuteTemplate(w, prod.Template, templateData) //"pages/products/info.html", prod)
	if err != nil {
		logs.Log.Error("ExecuteTemplate", err.Error())
		return
	}
}

/* func (api *API) images(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "api", "images")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	//core.ReadCookie(w, r)

	var filename string
	if _, ok := r.Form["id"]; ok {
		filename = r.Form["id"][0]
	}

	logs.Log.Info("Read request: " + filename)
	file, err := ioutil.ReadFile("./files/products/" + filename)
	if err != nil {
		logs.Log.Error("Cann't open file: " + filename)
		return
	}

	w.Write(file)
} */
