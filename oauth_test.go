package main

import (
  "bytes"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "net/http"
  "net/http/httptest"
)

var _ = Describe("Todo", func() {
  It("POST '/v1/apps/:app' will returns a 201 status code", func() {
    m := martini.Classic()
    m.Use(render.Renderer())
    m.Use(DB())
    m.Post("/v1/apps/:app", binding.Json(RegistryApplication{}), RegistryApp)
    response := httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/apps/123", bytes.NewReader([]byte("{\"sign\":\"5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("Content-Type", "application/json")
    m.ServeHTTP(response, request)
    Expect(response.Code).To(Equal(http.StatusCreated))
  })

})
