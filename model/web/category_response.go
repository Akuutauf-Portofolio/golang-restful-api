package web

// membuat representasi response, agar data yang ditampilkan untuk respon (hasil)-
// tidak terekspose semua
// juga mengikuti apispec yang sudah dibuat sebelumnya
type CategoryResponse struct {
	Id int
	Name string
}