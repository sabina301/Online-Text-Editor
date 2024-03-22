package user

type Repository interface {
	Create(name string, passwordHash string) (int, error)
	Get()
}

type userService struct {
	userRepository Repository
}

func NewUserService(repo Repository) *userService {
	return &userService{
		userRepository: repo,
	}
}
