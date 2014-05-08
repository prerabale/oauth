package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  m.Group("/v1/apps", func(r martini.Router) {
    r.Post("/:app", binding.Json(Application{}), RegistryApp)
    r.Put("/:app", UpdateApp)
    r.Get("/:app/key", GetAppKey)
    r.Put("/:app/key", ResetAppKey)
    r.Post("/:app/sign", ExchangeAppToken)
    r.Get("/token/verify/:token/:timestamp", VerifyToken)
  })
  m.Run()
}
