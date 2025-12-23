package web

// representasi request create, yang mana hanya mengirimkan atribut name saja
// juga untuk id akan otomatis di generate karena sifatnya adalah auto increment
type CategoryCreateRequest struct {
	Name string
}