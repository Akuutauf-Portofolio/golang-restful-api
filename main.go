package main

import (
	"belajar-go-lang-restful-api/app"
	"belajar-go-lang-restful-api/controller"
	"belajar-go-lang-restful-api/exception"
	"belajar-go-lang-restful-api/helper"
	"belajar-go-lang-restful-api/middleware"
	"belajar-go-lang-restful-api/repository"
	"belajar-go-lang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

// mendefinisikan main program
func main() {
	// membuat koneksi databse bar
	db := app.NewDB()

	// membuat validator
	validate := validator.New()

	// membuat repository
	categoryRepository := repository.NewCategoryRepository()

	// membuat service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	// membuat category controller
	categoryController := controller.NewCategoryController(categoryService)
	
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

	// membuat server
	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),

		// ketika auth middleware sudah ditambahkan, maka handler yang digunakan adalah-
		// router yang dibungkus dengan auth middleware
	}

	// menjalankan server
	err := server.ListenAndServe()

	// mengecek error
	helper.PanicIfError(err)
}