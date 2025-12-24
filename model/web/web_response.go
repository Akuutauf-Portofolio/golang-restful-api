package web

// membuat standar web response
type WebReponse struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Data interface{} `json:"data"` // agar datanya bisa di isi tipe data apapun
}