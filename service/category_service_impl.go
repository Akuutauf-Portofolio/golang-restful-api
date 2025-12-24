package service

import (
	"belajar-go-lang-restful-api/exception"
	"belajar-go-lang-restful-api/helper"
	"belajar-go-lang-restful-api/model/domain"
	"belajar-go-lang-restful-api/model/web"
	"belajar-go-lang-restful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	// di category service implementation membutuhkan repository, maka perlu di tambahkan di atribut
	CategoryRepository repository.CategoryRepository

	// butuh koneksi database juga, maka tambahkan attribute sql
	// db bentuknya adalah struct bukan interface, maka di set sebagai pointer
	DB *sql.DB

	// menambahkan attribute validate
	// validate ditambahkan di function create dan update
	// di function yang lain tidak dibutuhkan validasi, karena tidak ada payload (data asli) di parameter nya
	Validate *validator.Validate 
}

// mendefinisikan construktor dari controller

// membuat untuk router dengan return CategoryService (interface)
func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: DB,
		Validate: validate,
	}
}

// mengimplementasi category service (membuat method yang ada di category  agar dimiliki oleh CategoryServiceImpl)
// karena database yang kita gunakan adalah transactional (mysql), maka requestnya nanti dalam bentuk transactonal

func(service CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// mengimplementasi validation sebelum melakukan transaction
	errorValidation := service.Validate.Struct(request)

	// mengecek error
	helper.PanicIfError(errorValidation)

	// memulai koneksi database transactional
	tx, err := service.DB.Begin()

	// mengecek error
	helper.PanicIfError(err)

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer helper.CommitOrRollback(tx)

	// membuat data category 
	category := domain.Category{
		Name: request.Name,
	}

	// kemudian lakukan proses untuk create/save data
	// yang diambil dari implementation CategoryRepository melalaui parameter service, dan attribute CategoryRepository
	// karena return dari function save adalah domain category, maka di set ulang untuk menyimpan hasilnya
	category = service.CategoryRepository.Save(ctx, tx, category)

	// mengkonversi data category dari object menjadi response
	return helper.ToCategoryResponse(category)
}

func(service CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// mengimplementasi validation sebelum melakukan transaction
	errorValidation := service.Validate.Struct(request)

	// mengecek error
	helper.PanicIfError(errorValidation)
	
	// memulai koneksi database transactional
	tx, err := service.DB.Begin()

	// mengecek error
	helper.PanicIfError(err)

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer helper.CommitOrRollback(tx)

	// melakukan pencarian data category terlebih dahulu sebelum dilakukan  dengan function FindById
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)

	// melakukan pengecekan ketika data category tidak ditemukan
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// mengecek error
	helper.PanicIfError(err)

	// kalau data category ada, maka ubah datanya
	category.Name = request.Name

	// kemudian lakukan proses untuk update data
	// yang diambil dari implementation CategoryRepository melalaui parameter service, dan attribute CategoryRepository
	// karena return dari function update adalah domain category, maka di set ulang untuk menyimpan hasilnya
	category = service.CategoryRepository.Update(ctx, tx, category)

	// mengkonversi data category dari object menjadi response
	return helper.ToCategoryResponse(category)
}

func(service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// memulai koneksi database transactional
	tx, err := service.DB.Begin()

	// mengecek error
	helper.PanicIfError(err)

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer helper.CommitOrRollback(tx)

	// melakukan pencarian data category terlebih dahulu sebelum dilakukan  dengan function FindById
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	// melakukan pengecekan ketika data category tidak ditemukan
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// mengecek error
	helper.PanicIfError(err)

	// kemudian lakukan proses untuk delete data
	// yang diambil dari implementation CategoryRepository melalaui parameter service, dan attribute CategoryRepository
	// karena return dari function delete adalah domain category, maka di set ulang untuk menyimpan hasilnya
	service.CategoryRepository.Delete(ctx, tx, category)
}

func(service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	// memulai koneksi database transactional
	tx, err := service.DB.Begin()

	// mengecek error
	helper.PanicIfError(err)

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer helper.CommitOrRollback(tx)

	// melakukan pencarian data category terlebih dahulu sebelum dilakukan  dengan function FindById
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	// melakukan pengecekan ketika data category tidak ditemukan
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// mengecek error
	helper.PanicIfError(err)
	
	// mengkonversi data category dari object menjadi response
	return helper.ToCategoryResponse(category)
}

func(service CategoryServiceImpl) FindAll(ctx context.Context, ) []web.CategoryResponse {
	// memulai koneksi database transactional
	tx, err := service.DB.Begin()

	// mengecek error
	helper.PanicIfError(err)

	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
	defer helper.CommitOrRollback(tx)

	// melakukan pencarian data category terlebih dahulu sebelum dilakukan  dengan function FindById
	categories := service.CategoryRepository.FindAll(ctx, tx)

	// melakkan return dan konversi ke bentuk slice category responses
	return helper.ToCategoryResponses(categories)
}
