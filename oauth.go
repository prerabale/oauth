package main

import (
  "encoding/json"
  "github.com/arkors/oauth/handler"
  "github.com/arkors/oauth/model"
  "github.com/go-martini/martini"
  "github.com/go-xorm/xorm"
  "github.com/martini-contrib/render"
  "io/ioutil"
  "log"
  "net/http"
  "github.com/hoisie/redis"
  _ "github.com/go-sql-driver/mysql"
)

var orm *xorm.Engine
var redisClient redis.Client

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

func Db() martini.Handler {
  return func(c martini.Context) {
    c.Map(orm)
  }
}

func RedisDb() martini.Handler {
  return func(c martini.Context) {
    redisClient.Addr="127.0.0.1:6379"
    c.Map(redisClient)
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
      return
    }

    var app model.Application
    err = json.Unmarshal(data, &app)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Header().Set("Content-Type", "application/json")
      w.Write([]byte("{\"error\":\"Invalid request body, it should be JSON format.\"}"))
      return
    }else{
      c.Map(app)
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
      return
    }
  }
}

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())
  m.Use(Db())
  m.Use(RedisDb())
  m.Use(VerifyJSONBody())
  m.Use(VerifyHTTPHeader())

  m.Group("/v1/apps", func(r martini.Router) {
    r.Post("/:app",handler.RegistryApp)
    r.Put("/:app",handler.UpdateApp)
    r.Get("/:app/key", handler.GetAppKey)
    r.Put("/:app/key", handler.ResetAppKey)
    r.Post("/:app/sign", handler.ExchangeAppToken)
    r.Get("/token/verify/:token/:timestamp", handler.VerifyToken)
  })

  http.ListenAndServe(":3000", m)
}
