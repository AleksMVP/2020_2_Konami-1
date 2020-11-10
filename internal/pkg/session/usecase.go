package session

type UseCase interface {
	GetUserId(token string) (userId int, err error)
	CreateSession(userId int) (token string, err error)
	RemoveSession(token string) error
}