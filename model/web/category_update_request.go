package web

// representasi request update, yang mana mengirimkan atribut name dan id
// meskipun id akan otomatis di generate karena sifatnya adalah auto increment-
// namun update membutuhkan id sebagai parameter
// proses validasi dibutuhkan pada saat update

type CategoryUpdateRequest struct {
	// meskipun secara tidak langsung data yang diubah adalah name saja, namun id tetap diperlukan
	Id int `validate:"required"`
	Name string `validate:"required,max:200,min:1"`
}