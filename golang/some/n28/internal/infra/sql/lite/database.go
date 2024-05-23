package lite

import (
	"context"
	"database/sql"
	"errors"
	"n28/internal/app/dto"
	"n28/internal/domain"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type SqlDatabase struct {
	db *sql.DB
	mx sync.Mutex
}

func New() (*SqlDatabase, error) {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		return nil, errors.New("failed to open database")
	}

	if err = initTables(db); err != nil {
		return nil, err
	}

	return &SqlDatabase{
		db: db,
		mx: sync.Mutex{},
	}, nil
}

// public
// -----------------------------------------------------------------------

func (s *SqlDatabase) InsertProduct(ctx context.Context, data dto.InsertProduct) (uint64, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	result, err := s.db.Exec("INSERT INTO products(name, description, price) VALUES(?, ?, ?);",
		data.Name, data.Desc, data.Price)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func (s *SqlDatabase) GetProductById(ctx context.Context, id uint64) (*domain.Product, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	row := s.db.QueryRow("SELECT id, name, description, price FROM products WHERE id = ?;", id)

	p := domain.Product{}
	err := row.Scan(&p.Id, &p.Name, &p.Desc, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *SqlDatabase) GetProductCount(ctx context.Context) (uint64, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	var count uint64
	row := s.db.QueryRow("SELECT COUNT(*) FROM products;")
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// private
// -----------------------------------------------------------------------

func initTables(db *sql.DB) error {
	queryText := `
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT,
			price REAL
		);`

	_, err := db.Exec(queryText)
	if err != nil {
		return err
	}

	return nil
}
