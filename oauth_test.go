package main

import (
  "bytes"
  "encoding/json"
  "github.com/arkors/oauth/handler"
  "github.com/arkors/oauth/model"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/render"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "net/http"
  "net/http/httptest"
)

var _ = Describe("Testing For Oauth App And Token", func() {
  It("POST '/v1/apps/:app' will returns a 201 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Post("/:app", binding.Json(model.Application{}), handler.RegistryApp)
    })
    response := httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/apps/233", bytes.NewReader([]byte("{\"sign\":\"5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("Content-Type", "application/json")
    m.ServeHTTP(response, request)

    type Result struct {
      App int
      Key string
    }

    var result Result
    err := json.Unmarshal(response.Body.Bytes(), &result)

    Ω(err).Should(BeNil())
    Ω(result.App).Should(BeEquivalentTo(233))
    Ω(len(result.Key)).Should(BeNumerically("==", 32))

    Expect(response.Code).To(Equal(http.StatusCreated))
  })

})
