package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserRepo interface {
	Get(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}

type userRepo struct {
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Get(id int) (*User, error) {
	return nil, nil
}

func (r *userRepo) GetByUsername(username string) (*User, error) {
	return nil, nil
}

func (r *userRepo) Create(user *User) error {
	return nil
}

func (r *userRepo) Update(user *User) error {
	return nil
}

func (r *userRepo) Delete(id int) error {
	return nil
}
