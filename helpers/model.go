package helpers

import (
	"github.com/adityaqudaedah/go_restful_api/model/domain"
	"github.com/adityaqudaedah/go_restful_api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse{
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
		CreatedAt: category.CreatedAt,
	}
} 

func ToCategoryResponses(categories ...domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse

	for _,domainCategory := range categories {
		categoryResponses = append(categoryResponses,ToCategoryResponse(domainCategory))
	}

	return categoryResponses
}