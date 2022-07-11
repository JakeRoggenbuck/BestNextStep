package col

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS cols(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        desc TEXT NOT NULL,
		owner INTEGER NOT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(col Col) (*Col, error) {
	insert := "INSERT INTO cols(name, desc, owner) values(?,?,?)"
	res, err := r.db.Exec(insert, col.Name, col.Desc, col.Owner)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	col.ID = id

	return &col, nil
}

func (r *SQLiteRepository) All() ([]Col, error) {
	rows, err := r.db.Query("SELECT * FROM cols")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Col
	for rows.Next() {
		var col Col
		if err := rows.Scan(&col.ID, &col.Name, &col.Desc, &col.Owner); err != nil {
			return nil, err
		}
		all = append(all, col)
	}
	return all, nil
}

func (r *SQLiteRepository) GetByID(id int64) (*Col, error) {
	row := r.db.QueryRow("SELECT * FROM cols WHERE id = ?", id)

	var col Col
	if err := row.Scan(&col.ID, &col.Name, &col.Desc, &col.Owner); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &col, nil
}

func (r *SQLiteRepository) Update(id int64, updated Col) (*Col, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	update := "UPDATE cols SET name = ?, desc = ?, owner = ? WHERE id = ?"
	res, err := r.db.Exec(update, updated.Name, updated.Desc, updated.Owner, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM cols WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
