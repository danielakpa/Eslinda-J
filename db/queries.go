package db

import "database/sql"

// Product represents one cake, matching the "products" table.
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	HasVariants bool    `json:"has_variants"`
}

// GetAllProducts fetches every cake from the database.
func GetAllProducts(conn *sql.DB) ([]Product, error) {
	rows, err := conn.Query("SELECT id, name, price, image, category, has_variants FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	// Go through each row returned and build a Product from it
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Image, &p.Category, &p.HasVariants)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}