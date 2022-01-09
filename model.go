package main

import (
	"database/sql"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("select name, price from products where id = ?", p.ID).Scan(&p.Name, &p.Price)
	//return errors.New("Not implemented yet")
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err := db.Exec("Update products set name=?, price=? where id=?", p.Name, p.Price, p.ID)
	return err
	//return errors.New("Not implemented")
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("Delete from products where id=?", p.ID)
	return err
	//return errors.New("Not implemented")
}

func (p *product) createProduct(db *sql.DB) error {

	sqlStr := "INSERT INTO products(name, price) VALUES(?, ?);"
	stmt, err := db.Prepare(sqlStr)
	res, err := stmt.Exec(p.Name, p.Price)

	if err != nil {
		return err
	}

	prdID, err := res.LastInsertId()

	if err != nil {
		return err
	}

	p.ID = int(prdID)

	return nil
	//return errors.New("Not implemented")
}

func getProducts(db *sql.DB, start int, count int) ([]product, error) {

	rows, err := db.Query(
		"SELECT id, name,  price FROM products LIMIT ? OFFSET ?",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
