package interactor

import (
	entities "user/src/domains/entities"
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

func (c *ContentHandler) Upload(contentInput *entities.ContentInput) (*entities.Content, error) {
	//* メタデータ作成
	content, err := c.Crypt.MakeMetaData(contentInput)
	if err != nil {
		c.OutputPort.RenderError(err)
		return nil, err
	}
	return content, nil
}
