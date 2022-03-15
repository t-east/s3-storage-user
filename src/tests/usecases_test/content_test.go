package usecases_test

import (
	"testing"
	"user/src/domains/entities"
	"user/src/mocks"
	"user/src/usecases/interactor"
)

func TestContentUpload(t *testing.T) {
	outputPort := mocks.NewContentOutputPortMock()
	crypt := mocks.NewContentCryptMock()
	inputPort := interactor.NewContentInputPort(outputPort, crypt)

	contentInput := &entities.ContentInput{
		Content:     []byte{},
		ContentName: "コンテンツ1",
		Owner:       "オーナー1",
		Param: &entities.Param{
			Pairing: "a",
			G:       "d",
			U:       "f",
		},
	}

	receipt, err := inputPort.Upload(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Id != "7" {
		t.Errorf("Content.Upload() should return entities.Content.Id = 7, but got = %s", receipt.Id)
	}
	if receipt.Owner != contentInput.Owner {
		t.Errorf("Content.Upload() should return entities.Content.Owner = %s, but got = %s", contentInput.Owner, receipt.Owner)
	}
	if receipt.ContentName != contentInput.ContentName {
		t.Errorf("Content.Upload() should return entities.Content.ContentName = %s, but got = %s", contentInput.ContentName, receipt.ContentName)
	}
}
