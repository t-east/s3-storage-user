package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user/src/asserts"
	"user/src/core"
	"user/src/domains/entities"
	"user/src/interfaces/controllers"

	"github.com/Nik-U/pbc"
)

//* 登録済みユーザのuserIdでコンテンツを作成
func TestCreateContent(t *testing.T) {

	p := pbc.GenerateA(uint32(160), uint32(512))
	pairing := p.NewPairing()
	g := pairing.NewG1().Rand()
	u := pairing.NewG1().Rand()
	param := &entities.Param{
		Pairing: p.String(),
		G:       g.Bytes(),
		U:       u.Bytes(),
	}
	//* メタデータ作成
	f, err := core.UseFileRead("./linux_logo.jpg")
	if err != nil {
		t.Fatal(err)
	}
	c := &entities.ContentInput{
		Content:     f,
		ContentName: "たろう",
		Owner:       "aaa",
		Param: param,
		Key: &entities.Key{
			PubKey:  "a",
			PrivKey: "U9CymreQT4/4ah1iYYBbCc30odE",
		},
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(c); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/content", &buf)
	rec := httptest.NewRecorder()
	cc := controllers.LoadContentController()
	cc.Post(rec, req)
	asserts.AssertEqual(t, http.StatusOK, rec.Code, rec.Result().Status)
	res := &entities.Content{}
	err = json.NewDecoder(rec.Body).Decode(&res)
	if err != nil {
		t.Errorf("Failed to Get Content: %v", err)
	}
	t.Fatal(res)
}
