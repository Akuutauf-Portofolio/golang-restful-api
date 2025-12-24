package middleware

import (
	"belajar-go-lang-restful-api/helper"
	"belajar-go-lang-restful-api/model/web"
	"net/http"
)

// membuat middleware untuk authentication
// membuat struct sesuai dengan kontrak handler, jadi middleware harus dalam bentuk handler

type AuthMiddleware struct {
	// auth middleware nanti akan ditempatkan diatas middleware,
	// sehingga otomatis dia perlu meneruskan request ke handler berikutnya
	// maka perlu menambahkan attribute handler
	Handler http.Handler
}

// membuat function constructor
func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

// membuat method milik auth middleware agar sesuai dengan kontrak handler
func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, requst *http.Request) {
	// karena kita sebelumnya sudah membuat apispec, dengan menentukan key API nya adalah "X-API-Key"
	// melakukan pengecekan api key, dengan "SECRET" sebagai contoh
	if "SECRET" == requst.Header.Get("X-API-Key") {
		// ok
		// kalau misalnya ada x api key di header, maka,
		// tinggal meneruskan ke handler selanjutnya
		
		middleware.Handler.ServeHTTP(writer, requst)
	} else {
		// error: kalau misalnya dia tidak punya x api key di header maka,
		// akan mengirimkan erorr unauthorized / tidak dizinkan 

		// mengubah data respon ke client
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized) // memberikan code error 401

		// membuat web response
		webResponse := web.WebReponse{
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		// menulis ke respon body
		helper.WriteToResponseBody(writer, webResponse)
	}
}