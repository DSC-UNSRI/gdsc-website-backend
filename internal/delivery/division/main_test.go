package delivery

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var rr *httptest.ResponseRecorder
var router *gin.Engine
var ginCtx *gin.Context

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	rr = httptest.NewRecorder()

	ginCtx, router = gin.CreateTestContext(rr)

	os.Exit(m.Run())
}