package test

import (
	"belajar-go-lang-restful-api/app"
	"belajar-go-lang-restful-api/controller"
	"belajar-go-lang-restful-api/helper"
	"belajar-go-lang-restful-api/middleware"
	"belajar-go-lang-restful-api/model/domain"
	"belajar-go-lang-restful-api/repository"
	"belajar-go-lang-restful-api/service"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// membuat function untuk setup database, direkomendasikan database nya berbeda dengan yang utama (sebagai test)
func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api_test")

	// mengecek error
	helper.PanicIfError(err)

	// kalau set connection pulling bisa dilakukan setelah mengecek error dibawah
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	
	return db
}

// membuat unit test dengan pendekatan integration test
func setupRouter(db *sql.DB) http.Handler {
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

	// mengembalikan router dalam bentuk handler yang sudah mengimplementasikan AuthMiddleware
	return middleware.NewAuthMiddleware(router)
}

// membuat function untuk menghapus seluruh data category ketika pengujian di running
func truncateCategory(db *sql.DB) {
	db.Exec("truncate category")
}

// membuat skenario pengujian
func TestCreateCategorySuccess(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request body dalam bentuk json
	requestBody := strings.NewReader(`{"name": "Hewan"}`)

	// membuat request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)

	// menambahkan header
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 200, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 200, int(responseBody["code"].(float64))) 
	assert.Equal(t, "OK", responseBody["status"]) 
	assert.Equal(t, "Hewan", responseBody["data"].(map[string]interface{})["name"]) 
}

func TestCreateCategoryFailed(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request body dalam bentuk json
	requestBody := strings.NewReader(`{"name": ""}`)

	// membuat request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)

	// menambahkan header
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 400, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 400, int(responseBody["code"].(float64))) 
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// mengaktifkan tx (transaction)
	tx, _ := db.Begin()

	// berhubung update, maka datanya harus ada terlebih dahulu
	// membuat repository
	categoryRepository := repository.NewCategoryRepository()

	// membuat data baru dengan repository dan transaction
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Hewan",
	})

	// melakukan commit untuk repository dengan transaction
	tx.Commit()

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request body dalam bentuk json
	requestBody := strings.NewReader(`{"name": "Hewan"}`)

	// membuat request
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)

	// menambahkan header
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 200, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 200, int(responseBody["code"].(float64))) 
	assert.Equal(t, "OK", responseBody["status"]) 
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64))) 
	assert.Equal(t, "Hewan", responseBody["data"].(map[string]interface{})["name"]) 
}

func TestUpdateCategoryFailed(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// mengaktifkan tx (transaction)
	tx, _ := db.Begin()

	// berhubung update, maka datanya harus ada terlebih dahulu
	// membuat repository
	categoryRepository := repository.NewCategoryRepository()

	// membuat data baru dengan repository dan transaction
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Hewan",
	})

	// melakukan commit untuk repository dengan transaction
	tx.Commit()

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request body dalam bentuk json
	requestBody := strings.NewReader(`{"name": ""}`)

	// membuat request
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)

	// menambahkan header
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 400, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 400, int(responseBody["code"].(float64))) 
	assert.Equal(t, "BAD REQUEST", responseBody["status"]) 
}

func TestGetCategorySuccess(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// mengaktifkan tx (transaction)
	tx, _ := db.Begin()

	// berhubung update, maka datanya harus ada terlebih dahulu
	// membuat repository
	categoryRepository := repository.NewCategoryRepository()

	// membuat data baru dengan repository dan transaction
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Hewan",
	})

	// melakukan commit untuk repository dengan transaction
	tx.Commit()

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)

	// menambahkan header
	request.Header.Add("X-API-Key", "SECRET")

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 200, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 200, int(responseBody["code"].(float64))) 
	assert.Equal(t, "OK", responseBody["status"]) 
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64))) 
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"]) 
}

func TestGetCategoryFailed(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// karena pengujian nya ditujukan untuk not found, maka data harus kosong atau benar benar tidak ada di database

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/100", nil)

	// menambahkan header
	request.Header.Add("X-API-Key", "SECRET")

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 404, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 404, int(responseBody["code"].(float64))) 
	assert.Equal(t, "NOT FOUND", responseBody["status"]) 
}

func TestDeleteCategorySuccess(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// mengaktifkan tx (transaction)
	tx, _ := db.Begin()

	// berhubung update, maka datanya harus ada terlebih dahulu
	// membuat repository
	categoryRepository := repository.NewCategoryRepository()

	// membuat data baru dengan repository dan transaction
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Hewan",
	})

	// melakukan commit untuk repository dengan transaction
	tx.Commit()

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)

	// menambahkan header
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 200, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 200, int(responseBody["code"].(float64))) 
	assert.Equal(t, "OK", responseBody["status"]) 
}

func TestDeleteCategoryFailed(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// karena pengujian nya ditujukan untuk delete (not found), maka data harus kosong atau benar benar tidak ada di database

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/100", nil)

	// menambahkan header
	request.Header.Add("X-API-Key", "SECRET")

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 404, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 404, int(responseBody["code"].(float64))) 
	assert.Equal(t, "NOT FOUND", responseBody["status"]) 
}

func TestListCategoriesSuccess(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// mengaktifkan tx (transaction)
	tx, _ := db.Begin()

	// berhubung mau get list, maka datanya harus ada terlebih dahulu
	// membuat repository
	categoryRepository := repository.NewCategoryRepository()

	// membuat data baru dengan repository dan transaction
	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Hewan",
	})
	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Tumbuhan",
	})

	// melakukan commit untuk repository dengan transaction
	tx.Commit()

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)

	// menambahkan header
	request.Header.Add("X-API-Key", "SECRET")

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 200, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 200, int(responseBody["code"].(float64))) 
	assert.Equal(t, "OK", responseBody["status"]) 

	// membuat variabel untuk slice, dan mengkonversi nya ke bentuk interface
	var categories = responseBody["data"].([]interface{})

	// kemudian konversi lagi untuk masing masing category, ke bentuk map
	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	// untuk membandingkan data pertama
	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64))) 
	assert.Equal(t, category1.Name, categoryResponse1["name"])
	
	// untuk membandingkan data berikutnya
	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64))) 
	assert.Equal(t, category2.Name, categoryResponse2["name"])
}

func TestUnauthorized(t *testing.T) {
	// membuat koneksi databse baru untuk pengujian
	db := setupTestDB()

	// menghapus seluruh data truncate
	truncateCategory(db)

	// setup router / membuat router
	router := setupRouter(db)

	// membuat request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)

	// tidak menggunakan header dibawah ini, karena yang diuji adalah unauthorized
	// request.Header.Add("X-API-Key", "SECRET")
	// request.Header.Add("X-API-Key", "FALSE") // atau bisa juga pakai value api key yang salah

	// membuat recorder untuk writer
	recorder := httptest.NewRecorder()

	// melakukan pengujian request
	router.ServeHTTP(recorder, request)

	// mendapatkan response (hasil)
	response := recorder.Result()

	// setelah mendapatkan response, bisa melakukan apapun di response nya, tergantung kebutuhan
	// contoh melakukan perbandingan dengan assert

	// membandingkan status kode dari response dengan ekpektasi hasil
	assert.Equal(t, 401, response.StatusCode) 

	// atau mau menampilkan seluruh response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} // karena json maka bisa menggunakan interface (tipe data bebas)

	// melakukan konversi data body agar menjadi bentuk map (sesuai tipe data responseBody)
	json.Unmarshal(body, &responseBody)

	// menampilkan hasil konversi response body
	fmt.Println(responseBody)

	// kemudian kita bisa cek ulang dengan response body yang sudah dikonversi sebelumnya
	assert.Equal(t, 401, int(responseBody["code"].(float64))) 
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"]) 
}
