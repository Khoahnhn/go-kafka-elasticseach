package user

import (
	"github.com/Khoahnhn/go-kafka-elastichsearch/pkg/database"
)

func CreateUserRepository(user User) (User, error) {
	if err := database.DB.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUsersRepository(offset, pageSize int) ([]User, int64, error) {
	var users []User
	var total int64

	if err := database.DB.Model(&User{}).
		Where("deleted_at IS NULL").
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func GetUserByIDRepository(id string) (User, error) {
	var user User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateUserRepository(user User) (User, error) {
	if err := database.DB.Save(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func DeleteUserRepository(id string) error {
	if err := database.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
