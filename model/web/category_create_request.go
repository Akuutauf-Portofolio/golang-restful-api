package web

// representasi request create, yang mana hanya mengirimkan atribut name saja
// juga untuk id akan otomatis di generate karena sifatnya adalah auto increment
// proses validasi dibutuhkan pada saat create

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}