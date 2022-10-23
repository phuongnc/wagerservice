package testgin

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"unsafe"

	"github.com/gin-gonic/gin"
)

type RequestFormOption func(writer *multipart.Writer) error

func init() {
	gin.SetMode(gin.TestMode)
}

func GetTestContext() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(recorder)
	return ctx, engine, recorder
}

func ExtractBody(body io.ReadCloser) string {
	bodyStr, _ := io.ReadAll(body)
	return *(*string)(unsafe.Pointer(&bodyStr))
}

func MustMakeRequest(method, path string, body map[string]interface{}) *http.Request {
	payload, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Content-Type", "application/json")
	return req
}

func WithFormFile(field, fileName string, content []byte) RequestFormOption {
	return func(writer *multipart.Writer) error {
		part, err := writer.CreateFormFile(field, fileName)
		if err != nil {
			return err // nolint: wrapcheck
		}
		if _, err := part.Write(content); err != nil {
			return err // nolint: wrapcheck
		}
		return nil
	}
}

func MustMakeRequestWithForm(method, path string, content map[string]string, opts ...RequestFormOption) *http.Request {
	buf := new(bytes.Buffer)
	mockForm := multipart.NewWriter(buf)
	for k, v := range content {
		if err := mockForm.WriteField(k, v); err != nil {
			panic(err)
		}
	}
	for _, opt := range opts {
		if err := opt(mockForm); err != nil {
			panic(err)
		}
	}
	if err := mockForm.Close(); err != nil {
		panic(err)
	}
	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", mockForm.FormDataContentType())
	return req
}

// JSONStr make json string from multiline string
func JSONStr(str string) string {
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")
	return str
}
