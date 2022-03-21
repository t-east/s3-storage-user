package port

import (
	entities "user/src/domains/entities"
)

type ContentInputPort interface {
	Upload(content *entities.ContentIn) (*entities.Content, error)
}

type ContentOutputPort interface {
	Render(*entities.Content, int)
	RenderError(error)
}

type ContentCrypt interface {
	MakeMetaData(contentInput *entities.ContentIn) (*entities.Content, error)
}
