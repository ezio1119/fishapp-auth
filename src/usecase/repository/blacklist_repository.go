package repository

type BlackListRepository interface {
	SAdd(jti string) error
	SIsMember(jti string) (bool, error)
}
