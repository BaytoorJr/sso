package postgres

import (
	"context"
	"github.com/BaytoorJr/sso/src/config"
	"github.com/BaytoorJr/sso/src/repository"
	"github.com/go-kit/kit/log"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

type Store struct {
	db       *pgxpool.Pool
	logger   log.Logger
	UserRepo repository.UserRepository
}

func NewStore(db *pgxpool.Pool, logger log.Logger) (*Store, error) {
	repo := &Store{
		db:     db,
		logger: logger,
	}

	err := repo.migrate()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (s *Store) migrate() error {
	for i := 0; i < len(migrations); i++ {
		sql := strings.Replace(migrations[i], "$1", config.MainConfig.PostgresSchema, 1)
		_, err := s.db.Exec(context.Background(), sql)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Store) Users() repository.UserRepository {
	if s.UserRepo != nil {
		return s.UserRepo
	}

	s.UserRepo = &UserRepository{
		store: s,
	}

	return s.UserRepo
}
