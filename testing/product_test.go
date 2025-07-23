package product_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/vandannandwana/Basic-E-Commerce/internal/config"
	"github.com/vandannandwana/Basic-E-Commerce/internal/http/handlers/product"
	"github.com/vandannandwana/Basic-E-Commerce/internal/storage/sqlite"
)

func setupRouter() http.Handler{
	cfg := &config.Config{
		Env:         "test",
		StoragePath: "file::memory:?cache=shared",
		HttpServer: config.HttpServer{
			Address: ":8082",
		},
	}

	storage, _ := sqlite.New(cfg)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/products", product.New(storage))
	mux.HandleFunc("GET /api/products/{id}", product.GetProductById(storage))
	mux.HandleFunc("GET /api/products", product.GetProducts(storage))
	mux.HandleFunc("DELETE /api/products/{id}", product.DeleteProductById(storage))

	return mux

}

// func TestNew (t *testing.T){

// 	router := setupRouter()

// 	productPayLoad := map[string] interface{}{
// 		"name":"Product 1",
// 		"price":2000,
// 		"description":"Description of product 1",
// 	}

// 	body, _ := json.Marshal(productPayLoad)

// 	req := httptest.NewRequest("POST", "/api/products", bytes.NewReader(body))

// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	if w.Code != http.StatusCreated{
// 		t.Errorf("Expected status 201 Created, got %d", w.Code)
// 	}

// 	var respBody map[string] interface{}

// 	if err := json.Unmarshal(w.Body.Bytes(), &respBody); err != nil{
// 		t.Errorf("failed to parese response %v", err)
// 	}

// }

func createTestProduct(router http.Handler, t *testing.T) int {
	body := map[string]any{
		"name":        "Test Product",
		"price":       9000,
		"description": "For GET and DELETE test",
	}

	data, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/products", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201 Created, got %d", w.Code)
	}

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Failed to parse product creation response: %v", err)
	}

	idFloat, ok := resp["id"].(float64)
	if !ok {
		t.Fatalf("Invalid ID format in response")
	}

	return int(idFloat)
}


func TestGetProductById(t *testing.T){

	router := setupRouter()

	createTestProduct(router, t)

	req := httptest.NewRequest("GET", "/api/products/1", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404 Not Found, got %d", w.Code)
	}

}

func TestGetProducts(t * testing.T){

	router := setupRouter()

	createTestProduct(router, t)

	req := httptest.NewRequest("GET", "/api/products", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", w.Code)
	}

	var products []map[string] any

	err := json.Unmarshal(w.Body.Bytes(), &products)
	if err != nil {
		t.Fatalf("Failed to parse products list: %v", err)
	}

	if len(products) == 0 {
		t.Error("Expected at least one product in the list")
	}


}

func TestDeleteProductById(t *testing.T) {
	router := setupRouter()

	//Creating testProduct to perform delete operation
	productID := createTestProduct(router, t)

	req := httptest.NewRequest("DELETE", "/api/products/"+ fmt.Sprintf("%d", productID), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code == http.StatusNoContent {
		t.Errorf("Expected status 204 No Content, got %d", w.Code)
	}

	verifyReq := httptest.NewRequest("GET", "/api/products/"+strconv.Itoa(productID), nil)
	verifyRes := httptest.NewRecorder()
	router.ServeHTTP(verifyRes, verifyReq)

	if verifyRes.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found after delete, got %d", verifyRes.Code)
	}
}

