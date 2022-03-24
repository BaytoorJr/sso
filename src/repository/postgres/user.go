package postgres

import (
	"context"
	"time"

	"github.com/BaytoorJr/sso/src/domain"
	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	store Store
}

const (
	userTable   = "users_test"
	fieldsTable = "users_fields_test"
)

func (u *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	conn, err := u.store.db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "insert into "+userTable+" ("+
		"id, "+
		"login, "+
		"password, "+
		"created_at, "+
		"updated_at) values ("+
		"$1, $2, $3, $4, $5)",
		user.ID,
		user.Login,
		user.Password,
		user.Created_at,
		user.Updated_at,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUser(ctx context.Context, login string) (*domain.User, error) {
	var user domain.User
	conn, err := u.store.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, "select "+
		"id, "+
		"login, "+
		"password, "+
		"created_at, "+
		"updated_at "+
		"from "+userTable+
		" where login = $1;", login).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.Created_at,
		&user.Updated_at)
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, "select field_name, field_value from "+fieldsTable+
		" where id = $1", user.ID)
	if err != nil {
		if err.Error() != "no rows in result set" {
			return nil, err
		}
	}
	defer rows.Close()

	data := make(map[string]string)

	for rows.Next() {
		var name, value string
		err = rows.Scan(&name, &value)
		if err != nil {
			if err.Error() != "no rows in result set" {
				return nil, err
			}
		}

		data[name] = value
	}

	user.Data = data

	return &user, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	conn, err := u.store.db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	batch := &pgx.Batch{}
	batch.Queue("update "+userTable+" set "+
		"login = $1, "+
		"password = $2, "+
		"updated_at = $3 "+
		"where id = $4;",
		user.Login,
		user.Password,
		time.Now(),
		user.ID)

	_, err = conn.SendBatch(ctx, batch).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) AddProfileFields(ctx context.Context, user *domain.User) error {
	conn, err := u.store.db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	for name, value := range user.Data {
		batch.Queue("insert into "+fieldsTable+" ( "+
			"id, "+
			"field_name, "+
			"field_value) values ("+
			"$1, $2, $3 )",
			user.ID,
			name,
			value)
	}

	_, err = conn.SendBatch(ctx, batch).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, ID string) error {
	conn, err := u.store.db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	batch.Queue("delete from "+userTable+" where id = $1", ID)
	batch.Queue("delete from "+fieldsTable+" where id = $1", ID)

	_, err = conn.SendBatch(ctx, batch).Exec()
	if err != nil {
		return err
	}

	return nil
}