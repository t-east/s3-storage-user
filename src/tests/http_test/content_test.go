package http

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"user/src/asserts"
	"user/src/core"
	"user/src/domains/entities"
	"user/src/interfaces/controllers"
	mock_port "user/src/mocks"

	"github.com/golang/mock/gomock"
)

func TestUserHandler_CreateUser(t *testing.T) { // nolint:gocognit
	t.Parallel()

	tests := []struct {
		name     string
		setup    func(contentCrypt mock_port.MockContentCrypt) *entities.Content
		body     func() (*entities.Param, *entities.Key, *bytes.Buffer, string,  error)
		wantCode int
		wantErr  bool
	}{
		{
			name: "Success",
			setup: func(contentCrypt mock_port.MockContentCrypt) *entities.Content {
				content := &entities.Content{
					Content:  entities.SampleData{},
					MetaData: [][]byte{},
				}
				contentCrypt.EXPECT().MakeMetaData(gomock.Any()).AnyTimes().Return(content, nil)
				return content
			},
			body: func() (*entities.Param, *entities.Key, *bytes.Buffer, string, error) {
				param, key, err := core.CreateParamMock()

				if err != nil {
					return nil, nil, nil,  "", err
				}
				f, err := core.UseFileRead("./linux_logo.jpg")
				if err != nil {
					t.Fatal(err)
				}
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)
				part, err := writer.CreateFormFile("content", "./linux_logo.jpg")
				if err != nil {
					t.Fatal(err)
				}
				if _, err := io.Copy(part, f); err != nil {
					t.Fatal(err)
				}
				err = writer.WriteField("privkey", string(key.PrivKey))
				if err != nil {
					t.Fatal(err)
				}
				err = writer.WriteField("address", "aaaa")
				if err != nil {
					t.Fatal(err)
				}
				err = writer.Close()
				if err != nil {
					t.Fatal(err)
				}
				header := writer.FormDataContentType()
				return param, key, body, header, nil
			},
			wantCode: 201,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			param, _, body, header, err := tt.body()
			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/content", body)
			req.Header.Set("Content-Type",header)

			rec := httptest.NewRecorder()
			cc := controllers.LoadContentController(param)
			cc.Post(rec, req)
			asserts.AssertEqual(t, http.StatusCreated, rec.Code, rec.Body.String())
		})
	}
}
