package exception

import (
	"belajar-go-lang-restful-api/helper"
	"belajar-go-lang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// membuat function untuk penangan error handler pada endpoint
// function di samakan parameternya seperti router.PanicHandler

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	// ketika terjadi error maka akan menampilkan pesan error

	// menambahkan jenis erorr untuk not found error di error handler (404)
	if notFoundError(writer, request, err) {
	// kalau terjadi not found error, maka akan memberikan respon not found error ke user
		return
	} 

	// menambahkan error untuk bad request
	// mengecek kembali jika ada validation error / bad request dari user ke server
	if validationError(writer, request, err) {
		return
	}

	// menambahkan jenis error internal server error (500)
	internalServerError(writer, request, err)
}

// membuat function penanganan error untuk not found error
func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	// melakukan pengecekan error
	exception, ok := err.(NotFoundError)

	// jika dia ok/atau bisa dikonversi
	if ok {
		// mengubah data respon ke client
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound) // memberikan code error 404

		// membuat web response
		webResponse := web.WebReponse{
			Code: http.StatusNotFound,
			Status: "NOT FOUND ERROR",
			Data: exception.Error,
		}

		// menulis ke respon body
		helper.WriteToResponseBody(writer, webResponse)

		// jika bisa dikonversi 
		return true
	} else {
		// jika tidak bisa dikonversi
		return false
	}
}

// membuat function untuk penanganan error validation error / bad request dari user ke server
func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	// melakukan pengecekan ke validasi data request
	exception, ok := err.(validator.ValidationErrors)

	// jika dia ok/atau bisa dikonversi
	if ok {
		// mengubah data respon ke client
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound) // memberikan code error 404

		// membuat web response
		webResponse := web.WebReponse{
			Code: http.StatusNotFound,
			Status: "BAD REQUEST",
			Data: exception.Error(),
		}

		// menulis ke respon body
		helper.WriteToResponseBody(writer, webResponse)

		// jika bisa dikonversi 
		return true
	} else {
		// jika tidak bisa dikonversi
		return false
	}
}

// membuat function penanganan error untuk internal server error
func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	// mengubah data respon ke client
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError) // memberikan code error 500

	// membuat web response
	webResponse := web.WebReponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}

	// menulis ke respon body
	helper.WriteToResponseBody(writer, webResponse)
}
