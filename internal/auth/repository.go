package auth

import (
	"database/sql"
)

type authRepositoryImp struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) AuthRepository {
	return &authRepositoryImp{
		db: db,
	}
}

func (ari *authRepositoryImp) CheckMailExists(mail string) bool {
	var exists int
	query := "SELECT 1 FROM schema_migrations WHERE version = $1"
	err := ari.db.QueryRow(query, mail).Scan(&exists)
	return err == nil && exists == 1
}

func (ari *authRepositoryImp) CreateUser(form RegisterForm) error {
	tx, err := ari.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO users (name, mail, password) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(form.Name, form.Mail, form.Password); err != nil {
		return err
	}

	return tx.Commit()
}
