package interactor

import (
	entities "pairing_test/src/user/domains/entities"
	port "pairing_test/src/user/usecases/port"
)

type UserHandler struct {
	OutputPort port.UserOutputPort
	Repository port.UserRepository
	Crypt      port.UserCrypt
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, repository port.UserRepository, crypt port.UserCrypt) port.UserInputPort {
	return &UserHandler{
		OutputPort: outputPort,
		Repository: repository,
		Crypt:      crypt,
	}
}

func (uc *UserHandler) Create(user *entities.User) {
	// TODO: ユーザ情報を確認する -> 登録済みなのに新しく鍵を生成したら色々狂う
	key, err := uc.Crypt.KeyGen()
	if err != nil {
		uc.OutputPort.RenderError(err)
		return
	}
	// TODO: バリデーションを付ける
	user.PubKey = key.PubKey
	user.PrivKey = key.PrivKey
	user, err = uc.Repository.Create(user)
	if err != nil {
		uc.OutputPort.RenderError(err)
		return
	}
	uc.OutputPort.Render(user, 201)
}

func (uc *UserHandler) FindByID(id string) {
	user, err := uc.Repository.FindByID(id)
	if err != nil {
		uc.OutputPort.RenderError(err)
	}
	uc.OutputPort.Render(user, 200)
}

func (uc *UserHandler) KeyGen(id string) {
	user, err := uc.Repository.FindByID(id)
	if err != nil {
		uc.OutputPort.RenderError(err)
	}
	_, err = uc.Crypt.KeyGen()
	if err != nil {
		uc.OutputPort.RenderError(err)
	}
	// TODO; entitiesにkeyを加える処理を付ける
	updatedUser := user
	updatedUser, err = uc.Repository.Update(updatedUser)
	if err != nil {
		uc.OutputPort.RenderError(err)
	}
	uc.OutputPort.Render(updatedUser, 200)
}
