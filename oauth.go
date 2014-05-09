package main

import (
  "github.com/arkors/oauth/handler"
  "github.com/arkors/oauth/model"
  "github.com/go-martini/martini"
  _ "github.com/go-sql-driver/mysql"
  "github.com/go-xorm/xorm"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/render"
  "log"
  "net/http"
)

var orm *xorm.Engine

func init() {
  var err error

  orm, err = xorm.NewEngine("mysql", "arkors:arkors@/arkors_oauth?charset=utf8")

  if err != nil {
    log.Fatalf("Fail to create engine: %v\n", err)
  }

  if err = orm.Sync(new(model.Application)); err != nil {
    log.Fatalf("Fail to sync database: %v\n", err)
  }
}

func DB() martini.Handler {
  return func(c martini.Context) {
    c.Map(orm)
  }
}

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())
  m.Use(DB())

  m.Group("/v1/apps", func(r martini.Router) {
    r.Post("/:app", binding.Json(handler.RegistryApplication{}), handler.RegistryApp)
    r.Put("/:app", handler.UpdateApp)
    r.Get("/:app/key", handler.GetAppKey)
    r.Put("/:app/key", handler.ResetAppKey)
    r.Post("/:app/sign", handler.ExchangeAppToken)
    r.Get("/token/verify/:token/:timestamp", handler.VerifyToken)
  })

  http.ListenAndServe(":3000", m)
}
