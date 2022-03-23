package repository

type MainRepo interface {
	Users() UserRepository
}
