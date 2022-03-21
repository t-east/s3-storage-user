package usecases_test

import (
	"testing"
	"user/src/domains/entities"
	"user/src/mocks"
)

// メタデータ作成モックのテスト
func TestContentMakeMetaData(t *testing.T) {
	FakeCrypt := mocks.NewContentCryptMock()
	contentInput := &entities.ContentIn{
		Content:     []byte{},
		ContentName: "コンテンツ1",
		PrivKey:     []byte{},
	}
	content, err := FakeCrypt.MakeMetaData(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if content.ContentName != contentInput.ContentName {
		t.Errorf("Content.MakeMetaData() should return entities.Content.ContentName = %s, but got = %s", contentInput.ContentName, content.ContentName)
	}
}
