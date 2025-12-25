package main

import (
	"belajar-go-lang-restful-api/app"
	"belajar-go-lang-restful-api/controller"
	"belajar-go-lang-restful-api/helper"
	"belajar-go-lang-restful-api/middleware"
	"belajar-go-lang-restful-api/repository"
	"belajar-go-lang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// mendefinisikan main program
func main() {
	// membuat koneksi databse baru
	db := app.NewDB()

	// membuat validator
	validate := validator.New()

	// membuat repository
	categoryRepository := repository.NewCategoryRepository()

	// membuat service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	// membuat category controller
	categoryController := controller.NewCategoryController(categoryService)
	
	// implementasi router
	router := app.NewRouter(categoryController)

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