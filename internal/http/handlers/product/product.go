package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/vandannandwana/Basic-E-Commerce/internal/storage"
	"github.com/vandannandwana/Basic-E-Commerce/internal/types"
	"github.com/vandannandwana/Basic-E-Commerce/internal/utils/response"
)



func DeleteProductById(storage storage.Storage) http.HandlerFunc{

	return func(writer http.ResponseWriter, request *http.Request){

		_id := request.PathValue("id")

		id, err := strconv.ParseInt(_id, 10, 64)

		if err != nil{
			response.WriteJson(writer, http.StatusBadRequest, response.GeneralError(err))
			return 
		}

		_, err = storage.DeleteProductById(id)

		if err != nil{
			response.WriteJson(writer, http.StatusInternalServerError, response.GeneralError(err))
			return 
		}

		result := map[string] string{"status":fmt.Sprintf("Product Deleted with id %s", _id)}

		response.WriteJson(writer, http.StatusOK, result)

	}

}



func GetProductById(storage storage.Storage) http.HandlerFunc{

	return func(writer http.ResponseWriter, request *http.Request){
		
		_id := request.PathValue("id")

		id, err := strconv.ParseInt(_id, 10, 64)

		if err != nil{
			response.WriteJson(writer, http.StatusBadRequest, response.GeneralError(err))
			return 
		}

		product, err := storage.GetProductById(id)

		if err != nil{
			slog.Error("error getting the product of ", slog.String("id: ", _id))
			response.WriteJson(writer, http.StatusInternalServerError, response.GeneralError(err))
			return 
		}

		slog.Info("Getting Product by ", slog.String("id: ", _id))
		response.WriteJson(writer, http.StatusOK, product)


	}

}

func GetProducts(storage storage.Storage) http.HandlerFunc{

	return func(writer http.ResponseWriter, request *http.Request){
		slog.Info("Getting All Products")

		products, err := storage.GetProducts()

		if err != nil{
			response.WriteJson(writer, http.StatusInternalServerError, response.GeneralError(err))
			return 
		}

		response.WriteJson(writer, http.StatusOK, products)

	}

}


func New(storage storage.Storage) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		slog.Info("Creating new product")

		var product types.Product

		err := json.NewDecoder(request.Body).Decode(&product)

		if errors.Is(err, io.EOF) {
			response.WriteJson(writer, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return 
		}

		if err != nil{
			response.WriteJson(writer, http.StatusBadRequest, response.GeneralError(err))
			return 
		}

		if err := validator.New().Struct(product); err != nil{
			validationErrs := err.(validator.ValidationErrors)
			response.WriteJson(writer, http.StatusBadRequest, response.ValidationError(validationErrs))
			return 
		}

		lastId, err := storage.CreateProduct(product.ProductName, product.ProductPrice, product.ProductDescription)

		if err != nil{
			response.WriteJson(writer, http.StatusInternalServerError, response.GeneralError(err))
			return 
		}

		slog.Info("Product Created Successfully", slog.String("Product ID: ", fmt.Sprintf("%d", lastId)))

		response.WriteJson(writer, http.StatusOK, lastId)

	}

}
