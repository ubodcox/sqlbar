package products

/* import (
	"core"
	"html/template"
	"images"
	"io/ioutil"
	"logs"
	"net/http"
	"strconv"
	"strings"
)

func productsListHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "productsHandler", "productsListHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("ReadCookie", err.Error())
		return
	}

	if !session.Role.HasProductRights() {
		logs.Log.Error("HasProductRights", "no rights")
		return
	}

	t, err := template.ParseFiles("templates/pages/products/data.html")
	if err != nil {
		logs.Log.Error("ParseFiles", err.Error())
		return
	}

	data, err := productsList(session.Role.HasProductRights())
	if err != nil {
		logs.Log.Error("productsList", err.Error())
		return
	}

	//logs.Log.Error(data)
	err = t.ExecuteTemplate(w, "products/data.html", data)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}
	//logs.Log.Error("productsListHandler.end")
}

func productsInsertHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "ordersSql", "init")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriProductAdd) {
		logs.Log.Error("HasRight UriProductAdd", err.Error())
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
	t, err := template.ParseFiles("templates/pages/products/insert.html",
		"templates/pages/products/insert_data.html",
		"templates/pages/products/api.html")
	if err != nil {
		logs.Log.Error("ParseFiles", err.Error())
		return
	}

	//data := productsAdd(sub)

	t.ExecuteTemplate(w, "products/insert.html", nil)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}
}

func productsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "productsHandlers", "productsUpdateHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriProductEdit) {
		logs.Log.Error("HasRight UriProductEdit", err.Error())
		return
	}

	prod := Product{}
	if _, ok := r.Form["id"]; ok {
		prod.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	err = productsSelect(&prod)
	if err != nil {
		logs.Log.Error("productsSelect", err.Error())
		return
	}

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
	}
	//----------------------------------------------------
	t, err := template.ParseFiles("templates/pages/products/update.html",
		"templates/pages/products/update_data.html",
		"templates/pages/products/api.html")
	if err != nil {
		logs.Log.Error("template.ParseFiles", err.Error())
		return
	}

	logs.Log.Error(prod)
	t.ExecuteTemplate(w, "products/update.html", prod)
	if err != nil {
		logs.Log.Error("t.ExecuteTemplate", err.Error())
		return
	}
}

func productsExecInsertHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "productsHandlers", "productsExecInsertHandler")
	defer logs.Log.PopFuncName()

	r.ParseMultipartForm(20 * 1024 * 1024 * 1024)

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriProductAdd) {
		logs.Log.Error("HasRight UriProductAdd", err.Error())
		return
	}

	prod := Product{}
	var action string
	if _, ok := r.PostForm["action"]; ok {
		action = r.PostForm["action"][0]
	}
	if _, ok := r.PostForm["title"]; ok {
		prod.Title = r.PostForm["title"][0]
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

	filename, err := images.ImportImage(r, "template", "products")
	if err == nil {
		prod.Template = filename
		images.SetTemplate(prod.Template)
	}

	if prod.Title != "" && prod.Price > 0 {
		err := productsAdd(&prod)
		if err != nil {
			logs.Log.Error("prod.Price error:", err)
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

	if action == "publish" {
		http.Redirect(w, r, "info.html?id="+strconv.FormatInt(prod.ID, 10), 303)
	} else if action == "save" {
		http.Redirect(w, r, "data.html", 303)
	}
}

func productsExecUpdateHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "productsHandlers", "productsExecUpdateHandler")
	defer logs.Log.PopFuncName()

	r.ParseMultipartForm(20 * 1024 * 1024 * 1024)

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriProductEdit) {
		logs.Log.Error("HasRight UriProductEdit", err.Error())
		return
	}

	prod := Product{}
	var action string
	if _, ok := r.PostForm["action"]; ok {
		action = r.PostForm["action"][0]
	}
	if _, ok := r.PostForm["id"]; ok {
		prod.ID, _ = strconv.ParseInt(r.PostForm["id"][0], 10, 64)
	}
	if _, ok := r.PostForm["title"]; ok {
		prod.Title = r.PostForm["title"][0]
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

	filename, err := images.ImportImage(r, "template", "templates")
	//logs.Log.Debug("ImportImage", filename, err)
	if (err == nil) || (err.Error() == "update") {
		prod.Template = filename
		images.SetTemplate(prod.Template)
	}
	//logs.Log.Error("TEMPLATE2 ", filename, err)

	if prod.Title != "" && prod.Price > 0 {
		err := productsUpdate(prod)
		if err != nil {
			logs.Log.Error("productsUpdate:", err)
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

	/*t, err := template.ParseFiles("templates/pages/products/data.html")
	if err != nil {
		logs.Log.Error("productsListHandler.ParseFiles", err.Error())
		return
	}

	data, err := productsList()
	if err != nil {
		logs.Log.Error("productsListHandler.productsList", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, "pages/products/data.html", data)
	if err != nil {
		logs.Log.Error("productsListHandler.ExecuteTemplate", err.Error())
		return
	}* /

	if action == "publish" {
		http.Redirect(w, r, "info.html?id="+strconv.FormatInt(prod.ID, 10), 303)
	} else if action == "save" {
		http.Redirect(w, r, "data.html", 303)
	}
}

func productsExecDeleteHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "productsHandlers", "productsExecDeleteHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	session, err := core.ReadCookie(w, r)
	if err != nil {
		logs.Log.Error("ReadCookie", err.Error())
		return
	}

	if !session.Role.HasRight(core.UriProductDel) {
		logs.Log.Error("HasRight UriProductDel", err.Error())
		return
	}

	prod := Product{}
	if _, ok := r.Form["id"]; ok {
		prod.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	err = productsDelete(prod)
	if err != nil {
		logs.Log.Error("productsDelete", err.Error())
	}

	t, err := template.ParseFiles("templates/pages/products/data.html")
	if err != nil {
		logs.Log.Error("ParseFiles", err.Error())
		return
	}

	data, err := productsList(session.Role.HasProductRights())
	if err != nil {
		logs.Log.Error("productsList", err.Error())
		return
	}

	err = t.ExecuteTemplate(w, "products/data.html", data)
	if err != nil {
		logs.Log.Error("ExecuteTemplate", err.Error())
		return
	}
}

func productsInfoHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "productsHandlers", "productsInfoHandler")
	defer logs.Log.PopFuncName()

	r.ParseForm()

	// no cookies

	prod := Product{}
	if _, ok := r.Form["id"]; ok {
		prod.ID, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}

	err := productsSelect(&prod)
	if err != nil {
		logs.Log.Error("productsSelect", err.Error())
		return
	}

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
	}

	if prod.Template == "" {
		// no template
		t, err := template.ParseFiles("templates/pages/products/info.html")
		if err != nil {
			logs.Log.Error("ParseFiles", err.Error())
			return
		}

		err = t.ExecuteTemplate(w, "products/info.html", prod)
		if err != nil {
			logs.Log.Error("ExecuteTemplate", err.Error())
			return
		}
	} else {
		// individual template
		t, err := template.ParseFiles("files/templates/" + prod.Template)
		if err != nil {
			logs.Log.Error("ParseFiles", err.Error())
			return
		}
		logs.Log.Error("test", prod)
		err = t.ExecuteTemplate(w, prod.Template, prod) //"pages/products/info.html", prod)
		if err != nil {
			logs.Log.Error("ExecuteTemplate", err.Error())
			return
		}
	}
}

func productsImagesHandler(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("products", "productsHandlers", "productsImagesHandler")
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
}
 */