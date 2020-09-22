package main

import (
	"context"
	"database/sql"
	"fmt"
)

type database struct {
	db *sql.DB
}

func (d *database) newLot (name string, price string) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO allLots (name, status, price) VALUES($1, $2, $3)", name, true, price)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func (d *database) getAllLots() ([]*lot, error) {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("SELECT * FROM allLots WHERE allLots.Status = true")
	if err != nil {
		tx.Rollback()
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
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return lots, err
}

func (d *database) closeLot (id int) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE allLots SET status = $1 WHERE id = $2", false, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func (d *database) updatePrice (id int, price string, name string) error {
	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx,"UPDATE allLots SET price = $1, name = $3 WHERE id = $2", price, id, name)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO history (customer_name, lot_id, new_price, time_now) VALUES ($1, $2, $3, clock_timestamp())", name, id, price)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func (d *database) getHistory(id int) ([]*history_id, error) {

	ctx := context.Background()
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("SELECT * FROM history WHERE lot_id = $1", id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	lots := []*history_id{}
	for rows.Next(){
		p := &history_id{}
		err := rows.Scan(&p.ID, &p.CustomerName, &p.LotID, &p.NewPrice, &p.TimeNow)
		if err != nil{
			fmt.Println(err)
			continue
		}
		lots = append(lots, p)
	}
	fmt.Println(lots)
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return lots, err
}