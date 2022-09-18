package user_test

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/jpbmdev/payment-api/controllers"
	mocks_test "github.com/jpbmdev/payment-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetTargetSchemas(t *testing.T) {
	//Cannot load .env file in tests
	t.Setenv("DB_CONNECTION_STRING", mocks_test.DB_CONNECTION_STRING)
	t.Setenv("DB_NAME", mocks_test.DB_NAME)

	//Create mock router
	r := mocks_test.GetRouter()

	//Create controller
	//This shold be a mock controller o the original controller with stubs
	//But for the lack of time I will use the original controller
	//So this test will do a db call
	testController := controllers.NewTargetSchemaController()

	//Create the request in the router
	r.GET("/target-schema", testController.GetTargetSchemas)

	//Generate the http request to the server
	req, _ := http.NewRequest("GET", "/target-schema", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//Test validation
	assert.Equal(t, http.StatusOK, w.Code)
}
