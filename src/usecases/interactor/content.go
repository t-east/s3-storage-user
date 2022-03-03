package interactor

import (
	entities "pairing_test/src/user/domains/entities"
	port "pairing_test/src/user/usecases/port"
)

type ContentHandler struct {
	OutputPort port.ContentOutputPort
	Repository port.ContentRepository
	Crypt      port.ContentCrypt
	ContentSP  port.ContentSP
}

func NewContentInputPort(outputPort port.ContentOutputPort, repository port.ContentRepository, cryptHandler port.ContentCrypt, spHandler port.ContentSP) port.ContentInputPort {
	return &ContentHandler{
		OutputPort: outputPort,
		Repository: repository,
		Crypt:      cryptHandler,
		ContentSP:  spHandler,
	}
}

func (c *ContentHandler) Upload(contentInput *entities.ContentInput) {
	// メタデータ作成
	content, err := c.Crypt.MakeMetaData(contentInput)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}
	// contentIDをデータベースに保存
	content, err = c.Repository.Create(content)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}


	// TODO　ブロックチェーンに登録する

	// SPにアップロードする
	receipt, err := c.OutputPort.UploadSP(content)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}
	c.OutputPort.Render(receipt, 201)
}

func (c *ContentHandler) FindByID(id string) {
	// content情報を取得
	content, err := c.Repository.Find(id)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}

	// SPからコンテント情報を取得
	receipt, err := c.ContentSP.GetContent(content.Id)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}

	c.OutputPort.Render(receipt, 201)
}
