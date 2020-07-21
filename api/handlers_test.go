package api

import (
	"github.com/devdimidved/int-bst-http/bst"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_searchHandler(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		param        string
		expectedCode int
	}{
		{
			name:         "method POST is not allowed",
			method:       http.MethodPost,
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "method DELETE is not allowed",
			method:       http.MethodDelete,
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "invalid GET",
			method:       http.MethodGet,
			param:        "foo=bar",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "valid GET for existing val",
			method:       http.MethodGet,
			param:        "val=8",
			expectedCode: http.StatusOK,
		},
		{
			name:         "valid GET for non-existing val",
			method:       http.MethodGet,
			param:        "val=100",
			expectedCode: http.StatusNotFound,
		},
	}

	input := []int{8, 5, 6, 10, 12, -5, 0, 20}
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	bstSrv := bst.NewService(input)
	app := NewApplication(logger, bstSrv)
	url := "/search?"

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, url+tc.param, nil)
			app.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func Test_application_insertHandler(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		param        string
		expectedCode int
	}{
		{
			name:         "method GET is not allowed",
			method:       http.MethodGet,
			param:        `{"val":42}`,
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "method DELETE is not allowed",
			method:       http.MethodDelete,
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "valid POST",
			method:       http.MethodPost,
			param:        `{"val":42}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid POST valid json",
			method:       http.MethodPost,
			param:        `{"foo":42}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid POST invalid json",
			method:       http.MethodPost,
			param:        `{"val"=42}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	input := []int{8, 5, 6, 10, 12, -5, 0, 20}
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	bstSrv := bst.NewService(input)
	app := NewApplication(logger, bstSrv)
	url := "/insert"

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, url, strings.NewReader(tc.param))
			app.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func Test_application_deleteHandler(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		param        string
		expectedCode int
	}{
		{
			name:         "method GET is not allowed",
			method:       http.MethodGet,
			param:        "val=8",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "method POST is not allowed",
			method:       http.MethodPost,
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "invalid DELETE",
			method:       http.MethodDelete,
			param:        "foo=bar",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "valid DELETE for existing val",
			method:       http.MethodDelete,
			param:        "val=8",
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "valid DELETE for non-existing val",
			method:       http.MethodDelete,
			param:        "val=100",
			expectedCode: http.StatusNoContent,
		},
	}

	input := []int{8, 5, 6, 10, 12, -5, 0, 20}
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	bstSrv := bst.NewService(input)
	app := NewApplication(logger, bstSrv)
	url := "/delete?"

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, url+tc.param, nil)
			app.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
