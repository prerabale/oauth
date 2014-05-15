package main

import (
  "encoding/json"
  "github.com/arkors/oauth/handler"
  "github.com/arkors/oauth/model"
  "github.com/garyburd/redigo/redis"
  "github.com/go-martini/martini"
  _ "github.com/go-sql-driver/mysql"
  "github.com/go-xorm/xorm"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/render"
  "io/ioutil"
  "log"
  "net/http"
)

var orm *xorm.Engine
var pool *redis.Pool

func init() {
  var err error

  orm, err = xorm.NewEngine("mysql", "arkors:arkors@/arkors_oauth?charset=utf8")

  if err != nil {
    log.Fatalf("Fail to create engine: %v\n", err)
  }

  if err = orm.Sync(new(model.Application)); err != nil {
    log.Fatalf("Fail to sync database: %v\n", err)
  }

  pool = redis.NewPool(func() (redis.Conn, error) {
    c, err := redis.Dial("tcp", ":6379")

    if err != nil {
      return nil, err
    }

    return c, err
  }, 10)

}

func Db() martini.Handler {
  return func(c martini.Context) {
    c.Map(orm)
  }
}

func Pool() martini.Handler {
  return func(c martini.Context) {
    c.Map(pool)
  }
}

func VerifyJSONBody() martini.Handler {
  return func(c martini.Context, w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadAll(r.Body)

    if len(data) == 0 {
      return
    }

    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Header().Set("Content-Type", "application/json")
      w.Write([]byte("{\"error\":\"Invalid request body.\"}"))
    }

    var app model.Application
    err = json.Unmarshal(data, &app)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Header().Set("Content-Type", "application/json")
      w.Write([]byte("{\"error\":\"Invalid request body, it should be JSON format.\"}"))
    }

  }
}

func VerifyHTTPHeader() martini.Handler {
  return func(c martini.Context, w http.ResponseWriter, r *http.Request) {
    _, log := r.Header["X-Arkors-Application-Log"]
    _, client := r.Header["X-Arkors-Application-Client"]
    if log != true || client != true {
      w.WriteHeader(http.StatusBadRequest)
      w.Header().Set("Content-Type", "application/json")
      w.Write([]byte("{\"error\":\"Invalid request header, it should be include 'X-Arkors-Application-Log'  and 'X-Arkors-Application-Client'.\"}"))
    }
  }
}

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())
  m.Use(Db())
  m.Use(Pool())
  m.Use(VerifyJSONBody())
  m.Use(VerifyHTTPHeader())

  m.Group("/v1/apps", func(r martini.Router) {
    r.Post("/:app", binding.Json(model.Application{}), handler.RegistryApp)
    r.Put("/:app", binding.Json(model.Application{}), handler.UpdateApp)
    r.Get("/:app/key", handler.GetAppKey)
    r.Put("/:app/key", handler.ResetAppKey)
    r.Post("/:app/sign", handler.ExchangeAppToken)
    r.Get("/token/verify/:token/:timestamp", handler.VerifyToken)
  })

  http.ListenAndServe(":3000", m)
}
