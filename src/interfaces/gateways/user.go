package gateways

import (
	"pairing_test/src/user/domains/entities"
	"pairing_test/src/user/usecases/port"

	"gorm.io/gorm"
)

type SQLHandler interface {
	Find(interface{}, ...interface{}) (*entities.User, error)
	First(interface{}, ...interface{}) (*entities.User, error)
	Create(interface{}) error
	Save(interface{}) error
	Delete(interface{}) *entities.User
	Where(interface{}, ...interface{}) *entities.User
}

type userRepository struct {
	Conn *gorm.DB
	SQLHandler
}

func NewUserRepository(conn *gorm.DB) port.UserRepository {
	return &userRepository{
		Conn: conn,
	}
}

func (ur *userRepository) FindByID(id string) (user *entities.User, err error) {
	userInDB, err := ur.SQLHandler.Find(&user, id)
	if err != nil {
		return nil, err
	}
	return userInDB, nil
}

func (ur *userRepository) Create(u *entities.User) (user *entities.User, err error) {
	err = ur.SQLHandler.Create(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Update(u *entities.User) (user *entities.User, err error) {
	err = ur.SQLHandler.Save(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}
