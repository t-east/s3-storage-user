package mocks

import (
	entities "user/src/domains/entities"
	"user/src/usecases/port"
)

type ContentControllerMock struct {
	CryptFactory  func() port.ContentCrypt
	OutputFactory func() port.ContentOutputPort
	InputFactory  func(
		o port.ContentOutputPort,
		cr port.ContentCrypt,
	) port.ContentInputPort
}

func NewContentCryptMock() port.ContentCrypt {
	return &ContentCryptMock{}
}

func NewContentOutputPortMock() port.ContentOutputPort {
	return &ContentOutputPortMock{}
}

type ContentRepositoryMock struct {
}

type ContentCryptMock struct {
}

type ContentOutputPortMock struct {
}

type ContentSPMock struct {
}

type ContentContractMock struct {
}

func (m *ContentCryptMock) MakeMetaData(input *entities.ContentInput) (*entities.Content, error) {
	content := &entities.Content{
		Content:     input.Content,
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: input.ContentName,
		SplitCount:  0,
	}
	return content, nil
}

func (m *ContentOutputPortMock) Render(c *entities.Content, n int) {
}

func (m *ContentOutputPortMock) RenderError(error) {
}
