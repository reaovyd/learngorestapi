package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/reaovyd/learngorestapi/api/controllers"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestRootController(t *testing.T) {
	r := setupRouter()
	t.Run("testing GET endpoint", func(t *testing.T) {
		r.GET("/", controllers.DisplayRoot)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		r.ServeHTTP(w, req)

		wantCode := 200
		wantString := w.Body.String()
		assert.Equal(t, wantCode, w.Code)
		assert.Equal(t, wantString, "{\"message\":\"hello to my page\"}")
	})

	t.Run("testing POST endpoint", func(t *testing.T) {
		r := setupRouter()
		r.POST("/", controllers.ProcessRoot)

		var mockJsonData = MockLinkData{
			Link: "jonuts",
		}

		jsonVal, _ := json.Marshal(mockJsonData)

		postData := bytes.NewBuffer(jsonVal)
		req, _ := http.NewRequest("POST", "/", postData)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		got := MockLinkData{}
		gotCode := w.Code
		json.Unmarshal(w.Body.Bytes(), &got)

		assert.Equal(t, http.StatusCreated, gotCode)
		assert.Equal(t, mockJsonData, got)
	})
	//t.Run("testing POST fail", func(t *testing.T) {
	//	r := setupRouter()
	//	r.POST("/", controllers.ProcessRoot)
	//	mockJsonData := `{
	//		Lin : "awd",
	//	}`
	//})
}

type MockLinkData struct {
	Link string `json:"link"`
}
