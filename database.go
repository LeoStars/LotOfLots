package main

import (
	"database/sql"
	"fmt"
)

type database struct {
	db *sql.DB
}

func (d *database) newLot (name string, price string) error {
	_, err := d.db.Exec("INSERT INTO allLots (name, status, price) VALUES($1, $2, $3)", name, true, price)
	return err
}

func (d *database) getAllLots() ([]*lot, error) {
	rows, err := d.db.Query("SELECT * FROM allLots WHERE allLots.Status = true")
	if err != nil {
		return nil, err
	}
	lots := []*lot{}
	for rows.Next(){
		p := &lot{}
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Status)
		if err != nil{
			fmt.Println(err)
			continue
		}
		lots = append(lots, p)
	}
	return lots, err
}

func (d *database) closeLot (id int) error {
	_, err := d.db.Exec("UPDATE allLots SET allLots.status = $1 WHERE allLots.id = $2", false, id)
	return err
}

func (d *database) updatePrice (id int, price string) error {
	_, err := d.db.Exec("UPDATE allLots SET allLots.price = $1 WHERE allLots.id = $2", price, id)
	return err
}