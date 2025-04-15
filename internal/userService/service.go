package userService

type UserService struct {
	Repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return s.Repo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.Repo.DeleteUserByID(id)
}
