package repository

import (
	"fmt"
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	// tujuannya untuk membuka data dari db atau create filenya
	// TODO: replace this
	data, err := u.db.Load("users")
	if err != nil {
		data = [][]string{
			{"username", "password", "loggedin"},
		}

		if err := u.db.Save("users", data); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(data); i++ {
		status, err := strconv.ParseBool(data[i][2])
		if err != nil {
			return nil, err
		}

		user := User{
			Username: data[i][0],
			Password: data[i][1],
			Loggedin: status,
		}
		result = append(result, user)
	}

	return result, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	// TODO: replace this
	return u.LoadOrCreate()
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	// TODO: replace this
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			if err := u.changeStatus(username, true); err != nil {
				return nil, err
			}
			return &username, nil
		}

		if user.Username != username {
			u.Logout(user.Username)
		}
	}

	return nil, fmt.Errorf("Login Failed")
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	// TODO: replace this
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Loggedin {
			return &user.Username, nil
		}
	}
	return nil, fmt.Errorf("no user is logged in")
}

func (u *UserRepository) Logout(username string) error {
	userLogin, err := u.FindLoggedinUser()
	if err != nil {
		return err
	}

	return u.changeStatus(*userLogin, false)
}

func (u *UserRepository) Save(users []User) error {
	record := [][]string{{
		"username", "password", "loggedin",
	}}

	for _, user := range users {
		newRow := []string{
			user.Username, user.Password,
		}

		if user.Loggedin {
			newRow = append(newRow, "true")
		} else {
			newRow = append(newRow, "false")
		}

		record = append(record, newRow)
	}

	return u.db.Save("users", record)
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	_, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	data, err := u.db.Load("users")

	if err != nil {
		return err
	}

	for i := 1; i < len(data); i++ {
		if data[i][0] == username {
			if status {
				data[i][2] = "true"
			} else {
				data[i][2] = "false"
			}
		}
	}

	return nil
}

func (u *UserRepository) LogoutAll() error {
	users, err := u.SelectAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if err := u.Logout(user.Username); err != nil {
			return err
		}
	}
	return nil // TODO: replace this
}
