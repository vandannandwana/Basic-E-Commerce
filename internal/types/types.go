package types

type Product struct {
	ProductId          int64  `json:"id"`
	ProductName        string `json:"name" validate:"required"`
	ProductPrice       int64  `json:"price" validate:"required"`
	ProductDescription string `json:"description" validate:"required"`
}
