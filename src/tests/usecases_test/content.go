package usecases_test

import (
	"testing"
	"user/src/domains/entities"
	"user/src/mocks"
	"user/src/usecases/interactor"
)

func TestContentCreate(t *testing.T) {
	outputPort := mocks.NewContentOutputPortMock()
	repository := mocks.NewContentRepositoryMock()
	crypt := mocks.NewContentCryptMock()
	sp := mocks.NewContentSPMock()
	contract := mocks.NewContentContractMock()
	inputPort := interactor.NewContentInputPort(outputPort, repository, crypt, sp, contract)

	contentInput := &entities.ContentInput{
		Content:     []byte{},
		ContentName: "コンテンツ1",
		Owner:       "オーナー1",
	}

	receipt, err := inputPort.Upload(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Id != "7" {
		t.Errorf("User.Create() should return entities.User.ID = 7, but got = %s", receipt.Id)
	}
	if receipt.Owner != contentInput.Owner {
		t.Errorf("User.Create() should return entities.User.Address = %s, but got = %s", contentInput.Owner, receipt.Owner)
	}
	if receipt.ContentName != contentInput.ContentName {
		t.Errorf("User.Create() should return entities.User.Address = %s, but got = %s", contentInput.ContentName, receipt.ContentName)
	}
}
