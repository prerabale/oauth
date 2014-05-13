package main

import (
  "bytes"
  "encoding/json"
  "net/http"
  "net/http/httptest"

  "github.com/arkors/oauth/handler"
  "github.com/arkors/oauth/model"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/render"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var response *httptest.ResponseRecorder

var _ = Describe("Testing Oauth App Create", func() {
  It("POST '/v1/apps/:app' will returns a 201 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
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

  It("POST '/v1/apps/:app' repetition application ID will returns a 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
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
      Error string
    }

    var result Result
    err := json.Unmarshal(response.Body.Bytes(), &result)

    Ω(err).Should(BeNil())
    Ω(len(result.Error)).Should(BeNumerically(">", 0))

    Expect(response.Code).To(Equal(http.StatusBadRequest))
  })

  It("POST '/v1/apps/:app' with a invalid app id,  will returns a 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
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
    m.Use(Pool())
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
    m.Use(Pool())
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
    m.Use(Pool())
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
    m.Use(Pool())
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
    m.Use(Pool())
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
    m.Use(Pool())
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

var _ = Describe("Testing Oauth App Key GET", func() {
  It("GET '/v1/apps/:app/key' will returns a 200 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Get("/:app/key", handler.GetAppKey)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("GET", "/v1/apps/233", nil)
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

  It("GET '/v1/apps/:app/key' with non application id will returns a 404 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Get("/:app/key", handler.GetAppKey)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("GET", "/v1/apps/999999", nil)
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

})

var _ = Describe("Testing Oauth App Key RESET", func() {
  It("PUT '/v1/apps/:app/key' will returns a 200 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Put("/:app/key", handler.ResetAppKey)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/apps/233", nil)
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

  It("PUT '/v1/apps/:app/key' with non application id will returns a 404 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Put("/:app/key", handler.ResetAppKey)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/apps/999999", nil)
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

})

var _ = Describe("Testing Oauth App Exchange For Token", func() {
  It("POST '/v1/apps/:app/sign' will returns a 201 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Post("/:app/sign", handler.ExchangeAppToken)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/apps/233/sign", nil)
    request.Header.Set("X-Arkors-Application-Id", "232")
    request.Header.Set("X-Arkors-Application-Sign", "cb21df532c6647383af7efa0fd8405f2,1389085779854")
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "3ad3ce877d6c42b131580748603f8d6a,ANDROID")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    type Result struct {
      Token string
    }

    var result Result
    err := json.Unmarshal(response.Body.Bytes(), &result)

    Ω(err).Should(BeNil())
    Ω(len(result.Token)).Should(BeNumerically("==", 32))

    Expect(response.Code).To(Equal(http.StatusCreated))
  })

  It("POST '/v1/apps/:app/sign' with non application id will returns a 404 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Post("/:app/sign", handler.ExchangeAppToken)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/apps/9999999/sign", nil)
    request.Header.Set("X-Arkors-Application-Id", "9999999")
    request.Header.Set("X-Arkors-Application-Sign", "cb21df532c6647383af7efa0fd8405f2,1389085779854")
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "3ad3ce877d6c42b131580748603f8d6a,ANDROID")
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

  It("POST '/v1/apps/:app/sign' with application id not equal header's will returns a 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Post("/:app/sign", handler.ExchangeAppToken)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/apps/233/sign", nil)
    request.Header.Set("X-Arkors-Application-Id", "232")
    request.Header.Set("X-Arkors-Application-Sign", "cb21df532c6647383af7efa0fd8405f2,1389085779854")
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "3ad3ce877d6c42b131580748603f8d6a,ANDROID")
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

var _ = Describe("Testing Oauth App Check Key", func() {
  It("GET '/v1/apps/token/verify/:token/:timestamp' will returns a 200 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Get("/token/verify/:token/:timestamp", handler.VerifyToken)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("GET", "/v1/apps/token/verify/cb21df532c6647383af7efa0fd8405f2/1389085779854", nil)
    request.Header.Set("X-Arkors-Application-Id", "233")
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "3ad3ce877d6c42b131580748603f8d6a,ANDROID")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    Expect(response.Code).To(Equal(http.StatusOK))
  })

  It("GET '/v1/apps/token/verify/:token/:timestamp' with non application id will returns a 404 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Get("/token/verify/:token/:timestamp", handler.VerifyToken)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("GET", "/v1/apps/token/verify/cb21df532c6647383af7efa0fd8405f2/1389085779854", nil)
    request.Header.Set("X-Arkors-Application-Id", "23323")
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "3ad3ce877d6c42b131580748603f8d6a,ANDROID")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    Expect(response.Code).To(Equal(http.StatusNotFound))
  })

  It("GET '/v1/apps/token/verify/:token/:timestamp' with wrong token value 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Get("/token/verify/:token/:timestamp", handler.VerifyToken)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("GET", "/v1/apps/token/verify/cb21df532c6647383af7efa0fd8405322f2/1389085779854", nil)
    request.Header.Set("X-Arkors-Application-Id", "233")
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "3ad3ce877d6c42b131580748603f8d6a,ANDROID")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    Expect(response.Code).To(Equal(http.StatusBadGateway))
  })

  It("GET '/v1/apps/token/verify/:token/:timestamp' with wrong timestamp value 400 status code", func() {

    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(Db())
    m.Use(Pool())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Get("/token/verify/:token/:timestamp", handler.VerifyToken)
    })

    response = httptest.NewRecorder()
    request, _ := http.NewRequest("GET", "/v1/apps/token/verify/cb21df532c6647383af7efa0fd8405f2/21389085779854", nil)
    request.Header.Set("X-Arkors-Application-Id", "233")
    request.Header.Set("X-Arkors-Application-Log", "5024442115e7bd738354c1fac662aed5")
    request.Header.Set("X-Arkors-Application-Client", "3ad3ce877d6c42b131580748603f8d6a,ANDROID")
    request.Header.Set("Accept", "application/json")
    m.ServeHTTP(response, request)

    Expect(response.Code).To(Equal(http.StatusBadGateway))
  })

})
