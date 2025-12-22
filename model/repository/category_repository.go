package repository

import (
	"belajar-go-lang-restful-api/model/domain"
	"context"
	"database/sql"
)

// membuat repository khusus untuk category

// membuat kontrak dalam bentuk interface (best practice)
type CategoryRepository interface {
	// membuat semua repository function untuk CRUD
	// functon custom yang membantu untuk melakukan operasi CRUD pada tabel category
	// parameter setiap function wajib berisi context, menggunakan transactional (database relation),-
	// lalu data yang asli (data category), terkecuali find by id dan find all
	// kemudian untuk return value function, semuanya adalah data category asli, terkecuali delete (tidak mereturn data)-
	// dan find all yang berisi kumpulan data category (slice category). Bila memerlukan tambahkan error (findById)
	// dan ini 5 kontrak function yang dibutuhkan nantinya
	
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}