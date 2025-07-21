package sqlite

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vandannandwana/Basic-E-Commerce/internal/config"
	"github.com/vandannandwana/Basic-E-Commerce/internal/types"
)

type Sqlite struct {
	Db *sql.DB
}

func (s Sqlite) DeleteProductById(productId int64) (bool, error) {

	stmt, err := s.Db.Prepare("DELETE FROM products WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(productId)

	if err != nil {
		return false, err
	}

	slog.Info("Product deleted with ", slog.String("id: ", fmt.Sprint(productId)))

	return true, nil

}

func (s Sqlite) GetProductById(productId int64) (types.Product, error) {

	stmt, err := s.Db.Prepare("SELECT id, name, price, description FROM products WHERE id = ? LIMIT 1")

	if err != nil {
		return types.Product{}, err
	}

	defer stmt.Close()

	var product types.Product

	err = stmt.QueryRow(productId).Scan(&product.ProductId, product.ProductName, product.ProductPrice, product.ProductDescription)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Product{}, fmt.Errorf("no student found with the id %s", fmt.Sprint(productId))
		}
		return types.Product{}, fmt.Errorf("query Error: %w", err)
	}

	return product, nil

}

func (s Sqlite) GetProducts() ([]types.Product, error) {

	stmt, err := s.Db.Prepare("SELECT id, name, price, description FROM products")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []types.Product

	for rows.Next() {
		var product types.Product

		err := rows.Scan(&product.ProductId, &product.ProductName, &product.ProductPrice, &product.ProductDescription)

		if err != nil {
			return nil, err
		}

		products = append(products, product)

	}

	return products, nil

}

func (s Sqlite) CreateProduct(name string, price int64, description string) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO products (name, price, description) VALUES(?,?,?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, price, description)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil

}

func New(cfg *config.Config) (*Sqlite, error) {

	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	price INTEGER,
	description TEXT)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil

}
