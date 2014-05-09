package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "io"
  "net/http"
  "net/http/httptest"
  "testing"
)

var (
  response *httptest.ResponseRecorder
)

func TestOauth(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Oauth Suite")
}

func Request(method string, route string, handler martini.Handler) {
  m := martini.Classic()
  m.Get(route, handler)
  m.Use(render.Renderer())
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}

func POSTRequest(method string, route string, handler martini.Handler, body io.Reader) {
  m := martini.Classic()
  m.Post(route, binding.Json(Todo{}), handler)
  m.Use(render.Renderer())
  request, _ := http.NewRequest(method, route, body)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}
