package controllers

import "user/src/domains/entities"

type ContentReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type MetaDataReq struct {
	Content ContentReq `json:"content"`
	PrivKey string     `json:"priv_key"`
	Address string     `json:"address"`
}

func ContentAPISchemaToEntity(req *MetaDataReq) *entities.ContentCreateMetaData {
	return &entities.ContentCreateMetaData{
		Content: entities.Point{
			X: req.Content.X,
			Y: req.Content.Y,
		},
		PrivKey: []byte(req.PrivKey),
		Address: req.Address,
	}
}
