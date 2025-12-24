package helper

import (
	"encoding/json"
	"net/http"
)

// membuat function untuk mengkonversi ke json
func ReadFromRequestBody(request *http.Request, result interface{}) {
	// membuat decoder dari request body
	// tenpa perlu membuat string, karena request body sudah mengandung json-
	// sehingga tinggal melakukan konversi saya ke bentuk byte
	decoder := json.NewDecoder(request.Body)

	// melakukan decode
	err := decoder.Decode(result)

	// mengecek error
	PanicIfError(err)
}

// membuat function untuk menulis ke Response body
func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	// karena data yang dikirim nanti dalam bentuk json, maka perlu tambahakan header berikut pada writer
	writer.Header().Add("Content-Type", "application/json") // untuk memberi tahu bahwa data yang dikirim berbentuk json

	// mengirim web  dengan mengkonversi ke json kemudian dikirim ke writer
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)

	// mengecek error
	PanicIfError(err)
}