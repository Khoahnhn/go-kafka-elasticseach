package user

import (
	"errors"
	"github.com/Khoahnhn/go-kafka-elastichsearch/internal/user/request"
)

func CreateUserService(createUserRequest request.CreateUserRequest) (User, error) {
	user := User{
		Username: createUserRequest.Username,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	createUser, err := CreateUserRepository(user)
	if err != nil {
		return User{}, errors.New("error creating user")
	}

	return createUser, nil
}

func GetUsersService(page, pageSize int) ([]User, int64, error) {
	offset := (page - 1) * pageSize
	return GetUsersRepository(offset, pageSize)
}

func GetUserByIDService(id string) (User, error) {
	user, err := GetUserByIDRepository(id)
	if err != nil {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func UpdateUserService(id string, userData request.UpdateUserRequest) (User, error) {
	user, err := GetUserByIDRepository(id)
	if err != nil {
		return User{}, errors.New("user not found")
	}

	user.Username = userData.Username
	user.Email = userData.Email

	return UpdateUserRepository(user)
}

func DeleteUserService(id string) error {
	return DeleteUserRepository(id)
}

func SearchUserService(query string, filter map[string]string) ([]User, error) {
	return SearchUserRepository(query, filter)
}
