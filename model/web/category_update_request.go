package web

type CategoryUpdateRequest struct {
	// meskipun secara tidak langsung data yang diubah adalah name saja, namun id tetap diperlukan
	Id int
	Name string
}