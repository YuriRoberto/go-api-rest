package tests

import (
	"github.com/YuriRoberto/go-api-rest/routes"
	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = ginkgo.Describe("ShowBooks", func() {
	var (
		router  *gin.Engine
		request *http.Request
		writer  *httptest.ResponseRecorder
		//context  *gin.Context
		response *http.Response
	)

	ginkgo.BeforeEach(func() {
		// Configuração inicial antes de cada teste
		router = gin.Default()
		router = routes.ConfigRoutes(router)

		// Criar uma solicitação HTTP de teste para a rota ShowBooks
		request, _ = http.NewRequest("GET", "/api/v1/books/", nil)
		writer = httptest.NewRecorder()
		//context, _ = gin.CreateTestContext(writer)

		// Executar o manipulador ShowBooks
		router.ServeHTTP(writer, request)
		response = writer.Result()
	})

	ginkgo.It("Return status code 200", func() {
		gomega.Expect(response.StatusCode).To(gomega.Equal(http.StatusOK))
	})
})
