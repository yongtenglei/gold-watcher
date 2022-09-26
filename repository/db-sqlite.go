package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func (r *SQLiteRepository) Migrate() error {
	query := `create table if not exists holdings(
		    id integer primary key autoincrement,
		    amount real not null,
		    purchase_date integer not null,
		    purchase_price integer not null);
		    `

	_, err := r.Conn.Exec(query)
	return err
}

func (r *SQLiteRepository) InsertHolding(h Holdings) (*Holdings, error) {
	stmt := "insert into holdings (amount, purchase_date, purchase_price) values (?, ?, ?)"
	res, err := r.Conn.Exec(stmt, h.Amount, h.PurchaseDate.Unix(), h.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	h.ID = id

	return &h, nil
}

func (r *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := "select id, amount, purchase_date, purchase_price from holdings order by purchase_date"
	rows, err := r.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Holdings

	for rows.Next() {
		var h Holdings
		var unixTime int64

		err := rows.Scan(
			&h.ID,
			&h.Amount,
			&unixTime,
			&h.PurchasePrice,
		)
		if err != nil {
			return nil, err
		}
		h.PurchaseDate = time.Unix(unixTime, 0)
		all = append(all, h)
	}

	return all, nil
}

func (r *SQLiteRepository) GetHoldingByID(id int) (*Holdings, error) {
	query := "select id, amount, purchase_date, purchase_price from holdings where id=?"
	row := r.Conn.QueryRow(query, id)

	var h Holdings
	var unixTime int64

	err := row.Scan(
		&h.ID,
		&h.Amount,
		&unixTime,
		&h.PurchasePrice,
	)
	if err != nil {
		return nil, err
	}

	h.PurchaseDate = time.Unix(unixTime, 0)

	return &h, nil
}

func (r *SQLiteRepository) UpdateHolding(id int64, updated Holdings) error {
	if id == 0 {
		return errors.New("invalid updated id")
	}

	stmt := "update Holdings set amount = ?, purchase_date = ?, purchase_price = ? where id = ?"
	res, err := r.Conn.Exec(stmt, updated.Amount, updated.PurchaseDate.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errUpdatedFailed
	}

	return nil
}

func (r *SQLiteRepository) DeleteHolding(id int64) error {
	stmt := "delete from holdings where id = ?"
	res, err := r.Conn.Exec(stmt, id)
	if err != nil {
		return err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errDeleteFailed
	}

	return nil
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{Conn: db}
}
