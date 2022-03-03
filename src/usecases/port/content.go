package port

import (
	entities "pairing_test/src/user/domains/entities"
)

type ContentInputPort interface {
	Upload(content *entities.ContentInput)
	FindByID(id string)
}

type ContentOutputPort interface {
	Render(*entities.ReceiptFromSP, int)
	RenderError(error)
	UploadSP(*entities.Content) (*entities.ReceiptFromSP, error)
}

type ContentRepository interface {
	Create(user *entities.Content) (*entities.Content, error)
	Find(id string) (*entities.Content, error)
}

type ContentCrypt interface {
	MakeMetaData(contentInput *entities.ContentInput) (*entities.Content, error)
}

type ContentContract interface {
	Register(string, string) error
}

type ContentSP interface {
	UploadSP(*entities.Content) (*entities.ReceiptFromSP, error)
	GetContent(id string) (*entities.ReceiptFromSP, error)
}
