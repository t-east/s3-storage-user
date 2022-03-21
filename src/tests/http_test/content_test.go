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
	"user/src/mocks"
)

//* 登録済みユーザのuserIdでコンテンツを作成
func TestCreateContent(t *testing.T) {

	param, key, _ := mocks.CreateParamSample()
	//* メタデータ作成
	f, err := core.UseFileRead("./linux_logo.jpg")
	if err != nil {
		t.Fatal(err)
	}
	c := &entities.ContentIn{
		Content:     f,
		ContentName: "aaaaaa",
		PrivKey:     key.PrivKey,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(c); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/content", &buf)
	rec := httptest.NewRecorder()
	cc := controllers.LoadContentController(param)
	cc.Post(rec, req)
	asserts.AssertEqual(t, http.StatusOK, rec.Code, rec.Result().Status)
	res := &entities.Content{}
	err = json.NewDecoder(rec.Body).Decode(&res)
	if err != nil {
		t.Errorf("Failed to Get Content: %v", err)
	}
}

func TestGetKey(t *testing.T) {
	param, _ , _ := mocks.CreateParamSample()
	req := httptest.NewRequest(http.MethodGet, "/api/content", nil)
	rec := httptest.NewRecorder()
	cc := controllers.LoadContentController(param)
	cc.Get(rec, req)
	asserts.AssertEqual(t, http.StatusOK, rec.Code, rec.Result().Status)
	res := &entities.Key{}
	err := json.NewDecoder(rec.Body).Decode(&res)
	if err != nil {
		t.Errorf("Failed to Get Content: %v", err)
	}
	t.Fatal(res)
}
