package repo

import (
	"github.com/jinzhu/gorm"
	"iris_demo/models"
)

type UserRepository interface {
	GetUserList() *[]models.User
	SaveUser(book models.User) (err error)
	GetUserById(id uint) (book models.User, err error)
	DeleteUser(id uint) (err error)
	GetUserByName(name string) *[]models.User
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct{
	db *gorm.DB
}

func (repo userRepository) GetUserList() *[]models.User {
	user := new([]models.User)
	err := repo.db.Raw(`select * FROM user`).Scan(user).RowsAffected
	if err > 0 {
		return user
	} else {
		return nil
	}
}

func (repo userRepository) GetUserByName(name string) *[]models.User {
	user := new([]models.User)
	err := repo.db.Raw(`select * FROM user where user.name = ?`, name).Scan(user).RowsAffected
	if err > 0 {
		return user
	} else {
		return nil
	}
}

func (repo userRepository) SaveUser(user models.User) (err error) {
	if user.ID != 0 {
		err := repo.db.Save(&user).Error
		return err
	} else {
		err := repo.db.Create(&user).Error
		return err
	}
}

func (repo userRepository) GetUserById(id uint) (user models.User, err error) {
	err = repo.db.First(&user, id).Error
	return user, err
}

func (repo userRepository) DeleteUser(id uint) (err error) {
	user := new(models.User)
	user.ID = id
	err = repo.db.Unscoped().Delete(&user).Error
	return err
}
