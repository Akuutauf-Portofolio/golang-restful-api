package helper

import (
	"belajar-go-lang-restful-api/model/domain"
	"belajar-go-lang-restful-api/model/web"
)

// membuat function untuk mengkonversi data menjadi tipe Category Response
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

// membuat function untuk mengkonversi data slice menjadi slice Category Response
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	// membuat variabel baru (berjenis slice category response)
	var categoryResponses []web.CategoryResponse

	// melakukan konversi data categories (slice), ke category responses (slice)
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	
	// karena data category sudah dikonversi ke category responses (slice), maka langsung bisa di return
	return categoryResponses
}