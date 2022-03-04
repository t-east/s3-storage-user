package usecases_test

import (
	"testing"
	entities "user/src/domains/entities"
	"user/src/usecases/port"
)

type UserControllerMock struct {
	RepoFactory func() port.UserRepository
	CryptFactory func() port.UserCrypt
	OutputFactory func() port.UserOutputPort
	InputFactory func(
		o port.UserOutputPort,
		u port.UserRepository,
		cr port.UserCrypt,
	) port.UserInputPort
}

func NewUserRepositoryMock() port.UserRepository {
	return &userRepositoryMock{}
}

type userRepositoryMock struct {
}

type userCryptMock struct {
	FakeKeyGen func() (*entities.Key, error)
}

func (m *userRepositoryMock) Create(user *entities.User) (*entities.User, error) {
	created := &entities.User{ID: "7", Address: user.Address, PubKey: user.PubKey, PrivKey: user.PrivKey}
	return created, nil
}

func (m *userRepositoryMock) FindByID(id string) (*entities.User, error) {
	user := &entities.User{ID: id, Address: "user.Address", PubKey: "user.PubKey", PrivKey: "user.PrivKey"}
	return user, nil
}

func (m *userCryptMock) KeyGen() (*entities.Key, error) {
	return m.FakeKeyGen()
}

func TestUserCreate(t *testing.T) {
	userRepo := &userRepositoryMock{
		FakeCreate: func(user *entities.User) (*entities.User, error) {
			created := &entities.User{ID: "7", Address: user.Address, PubKey: user.PubKey, PrivKey: user.PrivKey}
			return created, nil
		},
	}
	userCrypt := &userCryptMock{
		FakeKeyGen: func() (*entities.Key, error) {
			created := &entities.Key{PubKey: "sdf", PrivKey: "sdf"}
			return created, nil
		},
	}

	key, err := userCrypt.KeyGen()
	if err != nil {
		t.Fatal(err)
	}
	userInput := entities.User{Address: "11", PubKey: key.PubKey, PrivKey: key.PrivKey}
	got, err := userRepo.Create(&userInput)
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != "7" {
		t.Errorf("User.Create() should return entities.User.ID = 7, but got = %s", got.ID)
	}
	if got.Address != userInput.Address {
		t.Errorf("User.Create() should return entities.User.Address = %s, but got = %s", userInput.Address, got.Address)
	}
	if got.PubKey != "sdf" {
		t.Errorf("User.Create() should return entities.User.PubKey = %s, but got = %s", userInput.PubKey, got.PubKey)
	}
	if got.PrivKey != "sdf" {
		t.Errorf("User.Create() should return entities.User.PrivKey = %s, but got = %s", userInput.PrivKey, got.PrivKey)
	}
}
