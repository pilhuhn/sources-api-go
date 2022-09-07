package request

import (
	sourcesEcho "github.com/RedHatInsights/sources-api-go/util/echo"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"net/http/httptest"
)

var echoInstance = echo.New()

// CreateTestContext sets up a new echo context with the parameters given, and returns the context itself and the
// response recorder.
func CreateTestContext(method string, path string, body io.Reader, context map[string]interface{}) (echo.Context, *httptest.ResponseRecorder) {
	echoInstance.Binder = &sourcesEcho.NoUnknownFieldsBinder{}
	request := httptest.NewRequest(method, path, body)
	recorder := httptest.NewRecorder()
	echoContext := echoInstance.NewContext(request, recorder)

	for k, v := range context {
		echoContext.Set(k, v)
	}

	return echoContext, recorder
}

// EmptyTestContext returns an empty http context - for when we don't need much
// other than the recorder + context
func EmptyTestContext() (echo.Context, *httptest.ResponseRecorder) {
	return CreateTestContext(http.MethodGet, "/", nil, nil)
}
