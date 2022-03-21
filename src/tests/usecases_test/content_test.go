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

	contentInput := &entities.ContentIn{
		Content:     []byte{},
		ContentName: "コンテンツ1",
		PrivKey:     []byte{},
	}

	receipt, err := inputPort.Upload(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.ContentName != contentInput.ContentName {
		t.Errorf("Content.Upload() should return entities.Content.ContentName = %s, but got = %s", contentInput.ContentName, receipt.ContentName)
	}
}
