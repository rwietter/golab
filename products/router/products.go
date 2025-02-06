package router

import (
	"net/http"

	"api/controllers"
)

type ProductRouter struct {
	URL      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

var ProductRouters = []ProductRouter{
	{
		URL:      "/products",
		Method:   "GET",
		Function: controllers.GetProducts,
	},
}
