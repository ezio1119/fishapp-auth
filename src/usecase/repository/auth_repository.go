package repository

type AuthRepository interface {
	SAdd(t string) error
	SIsMember(t string) (bool, error)
}
