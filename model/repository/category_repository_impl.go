package repository

import (
	"belajar-go-lang-restful-api/helper"
	"belajar-go-lang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

// membuat implementasi dari repository (kontrak sebelumnya) untuk data category

// membuat struct category implementation
type CategoryRepositoryImpl struct {

}

// membuat method milik struct CategoryRepositoryImpl yang mana menerapkan kontrak sebelumnya
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// membuat  insert
	sql := "insert into category(name) values (?)"

	// mengeksekusi query
	// menggunakan ExecContext, karena data category di manipulasi
	result, err := tx.ExecContext(ctx, sql, category.Name)

	// mengecek error
	helper.PanicIfError(err)

	// mendapatkan id data yang terakhir (baru saja di tambahkan)
	id, err := result.LastInsertId()

	// mengecek error
	helper.PanicIfError(err)

	// melakukan set category.id sesuai dengan id yang sudah kita tambahkan (auto increment)
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// membuat query update
	sql := "update category set name = ? where id = ?"

	// mengeksekusi query
	// menggunakan ExecContext, karena data category di manipulasi
	// result di skip, karena tidak perlu menampilkan datanya
	_, err := tx.ExecContext(ctx, sql, category.Id)

	// mengecek error
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	// membuat query delete
	sql := "delete from category where id = ?"

	// mengeksekusi query
	// menggunakan ExecContext, karena data category di manipulasi
	_, err := tx.ExecContext(ctx, sql, category.Id)

	// mengecek error
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	// membuat query find by id
	sql := "select id, name from category where id = ?"

	// mengeksekusi query 
	// menggunakan QueryContext, karena data category hanya diambil, tanpa di manipulasi
	rows, err := tx.QueryContext(ctx, sql, categoryId)

	// mengecek error
	helper.PanicIfError(err)

	// membuat data category kosong
	category := domain.Category{}

	// melakukan pengecekan sebuah data
	if rows.Next() {
		// mengambil data hasil qeury sebelumnya
		err := rows.Scan(&category.Id, &category.Name)

		// mengecek error
		helper.PanicIfError(err)

		return category, nil
	} else {
		// kalau misalnya tidak ada data yang ditemukan
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	// membuat query find all
	sql := "select id, name from category"

	// mengeksekusi query 
	// menggunakan QueryContext, karena data category hanya diambil, tanpa di manipulasi
	rows, err := tx.QueryContext(ctx, sql)

	// mengecek error
	helper.PanicIfError(err)

	// membuat data category kosong
	var categories []domain.Category

	// melakukan perulangan
	for rows.Next() {
		// membuat data category (satu object) untuk setiap perulangan
		category := domain.Category{}

		// mengambil data hasil qeury sebelumnya
		err := rows.Scan(&category.Id, &category.Name)

		// mengecek error
		helper.PanicIfError(err)

		// kemudian menggabungkan data categori di setiap iterasi ke kumpulan categories ([]categories)
		categories = append(categories, category)
	}

	return categories
}