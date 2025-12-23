package controller

import (
	"belajar-go-lang-restful-api/helper"
	"belajar-go-lang-restful-api/model/web"
	"belajar-go-lang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// membuat struct
type CategoryControllerImpl struct {
	// menambahkan category service
	// CategoryService adalah sebuah interface, maka tidak perlu menggunakan pointer
	CategoryService service.CategoryService
}

// ini adalah kontrak untuk category controller
func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// membuat request create baru
	// dengan melampirkan attribute dan validasinya
	categoryCreateRequest := web.CategoryCreateRequest{}

	// memanggil function untuk menagkap data dari request body
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	// memanggil function create pada kontrak CategoryService
	// di dalam request sudah ada context, sehingga kita bisa memanfaatkannya
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	// membuat web response
	webResponse := web.WebReponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	// memanggil function write ke json
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// membuat request update baru
	// dengan melampirkan attribute dan validasinya
	categoryUpdateRequest := web.CategoryUpdateRequest{}

	// memanggil function untuk menagkap data dari request body
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	// mengambil parameter category id melalui parameter
	categoryId := params.ByName("categoryId")

	// melakukan konversi categoryId menjadi string
	atoi, err := strconv.Atoi(categoryId)

	// mengecek error
	helper.PanicIfError(err)

	// mengambil category id yang sudah dikonversi ke string
	categoryUpdateRequest.Id = atoi

	// memanggil function update pada kontrak CategoryService
	// di dalam request sudah ada context, sehingga kita bisa memanfaatkannya
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	// membuat web response
	webResponse := web.WebReponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	// memanggil function write ke json
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// delete tidak membutuhkan request body, maka bisa langsung ke tahapan pengambilan parameter category id
	// mengambil parameter category id melalui parameter
	categoryId := params.ByName("categoryId")

	// melakukan konversi categoryId menjadi string
	id, err := strconv.Atoi(categoryId)

	// mengecek error
	helper.PanicIfError(err)

	// memanggil function delete pada kontrak CategoryService
	// di dalam request sudah ada context, sehingga kita bisa memanfaatkannya
	// delete tidak membutuhkan respon, sehingga bisa langsung mengakses function delete nya
	controller.CategoryService.Delete(request.Context(), id)

	// membuat web response
	webResponse := web.WebReponse{
		Code: 200,
		Status: "OK",
	}

	// memanggil function write ke json
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// findById tidak membutuhkan request body, maka bisa langsung ke tahapan pengambilan parameter category id
	// mengambil parameter category id melalui parameter
	categoryId := params.ByName("categoryId")

	// melakukan konversi categoryId menjadi string
	id, err := strconv.Atoi(categoryId)

	// mengecek error
	helper.PanicIfError(err)

	// memanggil function findById pada kontrak CategoryService
	// di dalam request sudah ada context, sehingga kita bisa memanfaatkannya
	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	// membuat web response
	webResponse := web.WebReponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	// memanggil function write ke json
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// memanggil function findAll pada kontrak CategoryService
	// di dalam request sudah ada context, sehingga kita bisa memanfaatkannya
	categoryResponses := controller.CategoryService.FindAll(request.Context())

	// membuat web response
	webResponse := web.WebReponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponses,
	}

	// memanggil function write ke json
	helper.WriteToResponseBody(writer, webResponse)
}


