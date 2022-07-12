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
		collection INTEGER NOT NULL,
		owner INTEGER NOT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(step Step) (*Step, error) {
	insert := "INSERT INTO steps(name, desc, collection, owner) values(?,?,?,?)"
	res, err := r.db.Exec(insert, step.Name, step.Desc, step.Collection, step.Owner)
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
		if err := rows.Scan(&step.ID, &step.Name, &step.Desc, &step.Collection, &step.Owner); err != nil {
			return nil, err
		}
		all = append(all, step)
	}
	return all, nil
}

func (r *SQLiteRepository) GetByID(id int64) (*Step, error) {
	row := r.db.QueryRow("SELECT * FROM steps WHERE id = ?", id)

	var step Step
	if err := row.Scan(&step.ID, &step.Name, &step.Desc, &step.Collection, &step.Owner); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &step, nil
}

func (r *SQLiteRepository) GetByOwner(id int64) ([]Step, error) {
	rows, err := r.db.Query("SELECT * FROM steps WHERE owner = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Step
	for rows.Next() {
		var step Step
		if err := rows.Scan(&step.ID, &step.Name, &step.Desc, &step.Collection, &step.Owner); err != nil {
			return nil, err
		}
		all = append(all, step)
	}
	return all, nil
}

func (r *SQLiteRepository) GetByCollection(col int64) (*Step, error) {
	row := r.db.QueryRow("SELECT * FROM steps WHERE collection = ?", col)

	var step Step
	if err := row.Scan(&step.ID, &step.Name, &step.Desc, &step.Collection, &step.Owner); err != nil {
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
	update := "UPDATE steps SET name = ?, desc = ?, collection = ?, owner = ? WHERE id = ?"
	res, err := r.db.Exec(update, updated.Name, updated.Desc, updated.Collection, updated.Owner, id)
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
