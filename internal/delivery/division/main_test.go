package delivery

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


func getRouter() (*httptest.ResponseRecorder, *gin.Engine) {
	rr := httptest.NewRecorder()
	_, router := gin.CreateTestContext(rr)
	return rr, router
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		os.Exit(1)
	}

	validations.InitValidations(validator)

	os.Exit(m.Run())
}
