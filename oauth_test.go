package main

import (
  "bytes"
  "github.com/arkors/oauth/handler"
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
    m.Use(Db())
    m.Group("/v1/apps", func(r martini.Router) {
      r.Post("/:app", binding.Json(handler.RegistryApplication{}), handler.RegistryApp)
    })
    response := httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/apps/232", bytes.NewReader([]byte("{\"sign\":\"5024442115e7bd738354c1fac662aed5\"}")))
    request.Header.Set("Content-Type", "application/json")
    m.ServeHTTP(response, request)
    Expect(response.Code).To(Equal(http.StatusCreated))
  })

})
