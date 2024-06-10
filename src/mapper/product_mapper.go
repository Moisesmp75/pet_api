package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func ProductRequestToModel(req request.ProductRequest) models.Product {
	return models.Product{
		Unit:     req.Unit,
		Quantity: req.Quantity,
		Name:     req.Name,
	}
}

func ProductModeltoResponse(product models.Product) response.ProductResponse {
	return response.ProductResponse{
		Unit:     product.Unit,
		Quantity: product.Quantity,
		Name:     product.Name,
	}
}

func ProductModelsToResponse(products []models.Product) []response.ProductResponse {
	resp := make([]response.ProductResponse, len(products))

	for i, v := range products {
		resp[i] = ProductModeltoResponse(v)
	}

	return resp
}

func ProductRequestsToModels(req []request.ProductRequest) []models.Product {
	models := make([]models.Product, len(req))

	for i, v := range req {
		models[i] = ProductRequestToModel(v)
	}

	return models
}
