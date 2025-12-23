package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// membuat interface untuk controller nya
type CategoryController interface {
	// karena kita sebelumnya punya 5 function, maka nanti akan di buat 5 API 
	// function sudah standar, jadi tinggal mengukuti handler milik HTTP
	// nama function bisa sama dengan CategoryService

	// karena kita saat ini menggunakan http router, maka butuh params
	// kalau http handler biasanya cuman membutuhkan writer dan request

	// ini adalah kontrak untuk category controller
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}