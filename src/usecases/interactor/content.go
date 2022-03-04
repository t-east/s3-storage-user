package interactor

import (
	entities "user/src/domains/entities"
	port "user/src/usecases/port"
)

type ContentHandler struct {
	OutputPort      port.ContentOutputPort
	Repository      port.ContentRepository
	Crypt           port.ContentCrypt
	ContentSP       port.ContentSP
	ContentContract port.ContentContract
}

func NewContentInputPort(outputPort port.ContentOutputPort, repository port.ContentRepository, cryptHandler port.ContentCrypt, spHandler port.ContentSP, contract port.ContentContract) port.ContentInputPort {
	return &ContentHandler{
		OutputPort:      outputPort,
		Repository:      repository,
		Crypt:           cryptHandler,
		ContentSP:       spHandler,
		ContentContract: contract,
	}
}

func (c *ContentHandler) Upload(contentInput *entities.ContentInput) (*entities.ReceiptFromSP, error) {
	//* メタデータ作成
	content, err := c.Crypt.MakeMetaData(contentInput)
	if err != nil {
		c.OutputPort.RenderError(err)
		return nil, err
	}
	//* contentIDをデータベースに保存
	contentInDB, err := c.Repository.Create(content)
	if err != nil {
		c.OutputPort.RenderError(err)
		return nil, err
	}
	content.Id = contentInDB.Id

	//* ブロックチェーンに登録
	err = c.ContentContract.Register(content.ContentName, content.Id)
	if err != nil {
		c.OutputPort.RenderError(err)
		return nil, err
	}

	//* SPにアップロードする
	receipt, err := c.ContentSP.UploadSP(content)
	if err != nil {
		c.OutputPort.RenderError(err)
		return nil, err
	}
	c.OutputPort.Render(receipt, 201)
	return receipt, nil
}

func (c *ContentHandler) FindByID(id string) {
	//* content情報を取得
	content, err := c.Repository.Find(id)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}

	//* SPからコンテント情報を取得
	receipt, err := c.ContentSP.GetContent(content.Id)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}

	c.OutputPort.Render(receipt, 201)
}
