package main

import (
  "github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()
  m.Group("/apps", func(r martini.Router) {
    r.Post("/:app", RegistryApp)
    r.Put("/:app", UpdateApp)
    r.Get("/:app/key", GetAppKey)
    r.Put("/:app/key", ResetAppKey)
    r.Post("/:app/sign", ExchangeAppToken)
    r.Get("token/verify/:token/:timestamp", VerifyToken)
  })
  m.Run()
}
