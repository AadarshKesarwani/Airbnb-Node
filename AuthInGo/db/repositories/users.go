package repositories

type UserRepository interface {
    Create() error
}

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
    return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Create() error {
    return nil
}