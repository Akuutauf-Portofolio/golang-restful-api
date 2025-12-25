package app

import (
	"belajar-go-lang-restful-api/controller"
	"belajar-go-lang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

// membuat function untuk membuat route baru
func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	// mengimplementasikan router
	router := httprouter.New()

	// membuat endpoint
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// membuat panic handler, agar ketika error si pengguna juga mendapatkan respon error,-
	// contoh error seperti error validasi, error not found dan lain lain
	router.PanicHandler = exception.ErrorHandler

	// mengembalikan router
	return router
}