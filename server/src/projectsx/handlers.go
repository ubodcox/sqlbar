package products

import (
	"sqlbar/server/src/logs"
)

func init() {
	logs.Log.PushFuncName("products", "handlers", "init")
	defer logs.Log.PopFuncName()

	logs.Log.Info("IMPORTED")

	/* http.HandleFunc("/products/data.html", productsListHandler)
	http.HandleFunc("/products/insert.html", productsInsertHandler)
	http.HandleFunc("/products/update.html", productsUpdateHandler)

	http.HandleFunc("/products/exec_insert.html", productsExecInsertHandler)
	http.HandleFunc("/products/exec_update.html", productsExecUpdateHandler)
	http.HandleFunc("/products/exec_delete.html", productsExecDeleteHandler) */

	/* http.HandleFunc("/products/info.html", api.info)

	http.HandleFunc("/products/images", api.images)

	http.HandleFunc("/api/v1/product", api.product)

	chainMiddleWare := core.ChainMiddleWare(core.WithLogging, core.WithAuth)
	// html
	http.HandleFunc("/products/data.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "products", "data")
	}))
	http.HandleFunc("/products/insert.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "products", "insert")
	}))
	/* http.HandleFunc("/products/update.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "products", "update")
	})) */
	/*http.HandleFunc("/products/update.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.IDHandler(w, r, "products", "update")
	}))
	//http.HandleFunc("/products/api.html", chainMiddleWare(api.api))
	http.HandleFunc("/products/api.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.IDHandler(w, r, "products", "api")
	})) */
	/* http.HandleFunc("/products/api.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "products", "api")
	})) */
	// api
	/* http.HandleFunc("/api/v1/products/data", chainMiddleWare(api.data))
	http.HandleFunc("/api/v1/products/insert", chainMiddleWare(api.insert))
	http.HandleFunc("/api/v1/products/update", chainMiddleWare(api.update))
	http.HandleFunc("/api/v1/products/delete", chainMiddleWare(api.delete)) */
}
