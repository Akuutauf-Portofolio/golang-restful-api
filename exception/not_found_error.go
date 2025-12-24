package exception

// membuat struct untuk menangani not found error
type NotFoundError struct {
	// karena error merupakan kontrak yang bersifat interface,
	// maka kita buat atribute error dengan tipe data apapun (string)
	Error string
}

// membuat function untuk membuat pesan not found error
func NewNotFoundError(error string) NotFoundError {
	// mengembalikan error yang diterima oleh server
	return NotFoundError{Error: error}
}