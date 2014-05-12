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

var response *httptest.ResponseRecorder

var _ = Describe("Testing Oauth App Create", func() {
  It("POST '/v1/apps/:app' will returns a 201 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Post("/:app", binding.Json(model.Application{}), handler.RegistryApp)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/apps/233", bytes.NewReader([]byte("{\"sign\":\"5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
    request.Header.Set("Accept", "application/json")
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

  It("POST '/v1/apps/:app' with a invalid app id,  will returns a 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Post("/:app", binding.Json(model.Application{}), handler.RegistryApp)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/apps/233oauth", bytes.NewReader([]byte("{\"sign\":\"5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    type Result struct {
      Error string
    }

    var result Result
    err := json.Unmarshal(response.Body.Bytes(), &result)

    Ω(err).Should(BeNil())
    Ω(len(result.Error)).Should(BeNumerically(">", 0))

    Expect(response.Code).To(Equal(http.StatusBadRequest))
  })

  It("POST '/v1/apps/:app' with a invalid json body,  will returns a 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Post("/:app", binding.Json(model.Application{}), handler.RegistryApp)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/apps/233", bytes.NewReader([]byte("{\"sign 5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    type Result struct {
      Error string
    }

    var result Result
    err := json.Unmarshal(response.Body.Bytes(), &result)

    Ω(err).Should(BeNil())
    Ω(len(result.Error)).Should(BeNumerically(">", 0))

    Expect(response.Code).To(Equal(http.StatusBadRequest))
  })

  It("POST '/v1/apps/:app' with a invalid json field,  will returns a 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Post("/:app", binding.Json(model.Application{}), handler.RegistryApp)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/apps/233oauth", bytes.NewReader([]byte("{\"grimm\":\"5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    type Result struct {
      Error string
    }

    var result Result
    err := json.Unmarshal(response.Body.Bytes(), &result)

    Ω(err).Should(BeNil())
    Ω(len(result.Error)).Should(BeNumerically(">", 0))

    Expect(response.Code).To(Equal(http.StatusBadRequest))
  })

})

var _ = Describe("Testing Oauth App PUT", func() {
  It("PUT '/v1/apps/:app' will returns a 200 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Put("/:app", binding.Json(model.Application{}), handler.UpdateApp)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/apps/233", bytes.NewReader([]byte("{\"sign\":\"5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
    request.Header.Set("Accept", "application/json")
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

    Expect(response.Code).To(Equal(http.StatusOK))
  })

  It("PUT '/v1/apps/:app' with non application id will returns a 404 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Put("/:app", binding.Json(model.Application{}), handler.UpdateApp)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/apps/999999", bytes.NewReader([]byte("{\"sign\":\"5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    type Result struct {
      Error string
    }

    var result Result
    err := json.Unmarshal(response.Body.Bytes(), &result)

    Ω(err).Should(BeNil())
    Ω(len(result.Error)).Should(BeNumerically(">", 0))

    Expect(response.Code).To(Equal(http.StatusNotFound))
  })

  It("PUT '/v1/apps/:app' with a invalid json body,  will returns a 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Put("/:app", binding.Json(model.Application{}), handler.UpdateApp)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/apps/233", bytes.NewReader([]byte("{\"sign 5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    type Result struct {
      Error string
    }

    var result Result
    err := json.Unmarshal(response.Body.Bytes(), &result)

    Ω(err).Should(BeNil())
    Ω(len(result.Error)).Should(BeNumerically(">", 0))

    Expect(response.Code).To(Equal(http.StatusBadRequest))
  })

  It("PUT '/v1/apps/:app' with a invalid json field,  will returns a 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Put("/:app", binding.Json(model.Application{}), handler.RegistryApp)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/apps/233oauth", bytes.NewReader([]byte("{\"grimm\":\"5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,TEST")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    type Result struct {
      Error string
    }

    var result Result
    err := json.Unmarshal(response.Body.Bytes(), &result)

    Ω(err).Should(BeNil())
    Ω(len(result.Error)).Should(BeNumerically(">", 0))

    Expect(response.Code).To(Equal(http.StatusBadRequest))
  })

})
