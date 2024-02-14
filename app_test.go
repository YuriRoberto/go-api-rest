package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestApp(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "App Suite")
}

var _ = ginkgo.Describe("ShowBooks", func() {

	type BodyBook struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		MediumPrice float64 `json:"medium_price"`
		Author      string  `json:"author"`
		ImageURL    string  `json:"image_url"`
	}

	ginkgo.BeforeEach(func() {})

	ginkgo.It("Return status code 200 in get all", func() {
		urlData, _ := url.Parse("http://localhost:8000/api/v1/books/")
		request := &http.Request{
			Method: http.MethodGet,
			URL:    urlData,
		}
		response, err := http.DefaultClient.Do(request)
		fmt.Println(err)
		gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
		fmt.Println(response.StatusCode)
		gomega.Expect(response.StatusCode).To(gomega.Equal(http.StatusOK))

	})

	ginkgo.It("Return status code 200 and data in register of a book", func() {
		body := map[string]interface{}{
			"name":         "Livro sila",
			"description":  "livro n tao bom",
			"medium_price": 199.99,
			"author":       "Um ai",
			"image_url":    "https://algumblogdessesai.com",
		}
		jsonBody, err := json.Marshal(body)
		gomega.Expect(err).ShouldNot(gomega.HaveOccurred())

		// Criar um leitor de bytes para resolver o problema de antes
		reader := bytes.NewReader(jsonBody)

		resp, err := http.Post("http://localhost:8000/api/v1/books/", "application/json", reader)
		gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
		defer resp.Body.Close()

		gomega.Expect(resp.StatusCode).To(gomega.Equal(http.StatusOK))

		// Ler o corpo da resposta usando a lib ioutill, procurar mais sobre ela depois
		responseBody, err := ioutil.ReadAll(resp.Body)
		gomega.Expect(err).ShouldNot(gomega.HaveOccurred())

		// Remover o campo "id" do JSON retornado para a comparação
		var expectedJSON map[string]interface{}
		err = json.Unmarshal(jsonBody, &expectedJSON)
		gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
		delete(expectedJSON, "id")

		var actualJSON map[string]interface{}
		err = json.Unmarshal(responseBody, &actualJSON)
		gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
		delete(actualJSON, "id")

		gomega.Expect(actualJSON).To(gomega.Equal(expectedJSON))
	})

	ginkgo.It("Return error in get a book with id nonexistent", func() {
		urlData, _ := url.Parse("http://localhost:8000/api/v1/books/100")
		request := &http.Request{
			Method: http.MethodGet,
			URL:    urlData,
		}
		response, err := http.DefaultClient.Do(request)
		fmt.Println(err)
		gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
		fmt.Println(response.StatusCode)
		gomega.Expect(response.StatusCode).To(gomega.Equal(http.StatusBadRequest))
	})
})
