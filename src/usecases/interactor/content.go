package interactor

import (
	entities "user/src/domains/entities"
	"user/src/mocks"
	port "user/src/usecases/port"
)

type ContentHandler struct {
	OutputPort port.ContentOutputPort
	Crypt      port.ContentCrypt
}

func NewContentInputPort(outputPort port.ContentOutputPort, cryptHandler port.ContentCrypt) port.ContentInputPort {
	return &ContentHandler{
		OutputPort: outputPort,
		Crypt:      cryptHandler,
	}
}

func (c *ContentHandler) Upload(contentIn *entities.ContentIn) (*entities.Content, error) {
	//* メタデータ作成
	content, err := c.Crypt.MakeMetaData(contentIn)
	if err != nil {
		c.OutputPort.RenderError(err)
		return nil, err
	}
	c.OutputPort.Render(content, 200)
	return content, nil
}

func (c *ContentHandler) GetKey() (*entities.Key, error) {
	_, key, _ := mocks.CreateParamMock()
	c.OutputPort.RenderKey(key, 200)
	return key, nil
}
