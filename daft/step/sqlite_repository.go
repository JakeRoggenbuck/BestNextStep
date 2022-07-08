package step

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
    CREATE TABLE IF NOT EXISTS steps(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        desc TEXT NOT NULL,
		left INTEGER,
		right INTEGER,
		owner INTEGER NOT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(step Step) (*Step, error) {
	res, err := r.db.Exec("INSERT INTO steps(name, desc, left, right, owner) values(?,?,?,?,?)", step.Name, step.Desc, step.Left, step.Right, step.Owner)
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
	step.ID = id

	return &step, nil
}

func (r *SQLiteRepository) All() ([]Step, error) {
	rows, err := r.db.Query("SELECT * FROM steps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Step
	for rows.Next() {
		var step Step
		if err := rows.Scan(&step.ID, &step.Name, &step.Desc, &step.Left, &step.Right, &step.Owner); err != nil {
			return nil, err
		}
		all = append(all, step)
	}
	return all, nil
}

func (r *SQLiteRepository) GetByID(id int64) (*Step, error) {
	row := r.db.QueryRow("SELECT * FROM steps WHERE id = ?", id)

	var step Step
	if err := row.Scan(&step.ID, &step.Name, &step.Desc, &step.Left, &step.Right, &step.Owner); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &step, nil
}

func (r *SQLiteRepository) Update(id int64, updated Step) (*Step, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec("UPDATE steps SET name = ?, desc = ?, left = ?, right = ?, owner = ? WHERE id = ?", updated.Name, updated.Desc, updated.Left, updated.Right, updated.Owner, id)
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
	res, err := r.db.Exec("DELETE FROM steps WHERE id = ?", id)
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
