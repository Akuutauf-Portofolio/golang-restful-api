package web

// membuat standar web response
type WebReponse struct {
	Code int
	Status string
	Data interface{} // agar datanya bisa di isi tipe data apapun
}