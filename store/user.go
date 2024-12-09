package store

import (
	"api/entity"
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func (r *Repository) RegisterUser(ctx context.Context, db Execer, u *entity.User) error {
	u.Created = r.Clocker.Now()
	u.Modified = r.Clocker.Now()
	sql := `INSERT INTO users (name, password, role, created, modified) VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(ctx, sql, u.Name, u.Password, u.Role, u.Created, u.Modified)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == ErrCodeMySQLErrorDuplicateEntry {
			return fmt.Errorf("cannot create same name user: %w", ErrAlreadyEntry)
		}
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = entity.UserID(id)
	return nil
}
