package service

import (
	"belajar-go-lang-restful-api/model/web"
	"context"
)

// service disebut juga sebagai bussiness logic
// di dalam service harus membuat kontrak
// membuat interface (kontrak)
type CategoryService interface {
	// biasanya kontrak ini berisi api sesuai dengan jumlah function
	// contoh: sebuah api hanya bisa memanggil 1 function di sebuah service

	// isi parameter berisi context, request (optional), dan response (optional)
	// untuk request bisa kita custom seperti category_create_request
	// dan untuk responsenya bisa kita custom seperti category_response

	// kalau domain itu ranahnya ke repository tidak ke service. maka dari itu nanti request akan-
	// representasikan ke dalam bentuk yang lain
	// juga untuk response nya direkomendasikan jangan menggunakan domain, karena biasanya kalau-
	// ditampilkan dalam bentuk table semua informasi tabel akan ditampilkan

	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, requst web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int) // kkarena delete cuman butuh id saja dan tidak ada return value
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context, ) []web.CategoryResponse
}