package handler

import (
  "github.com/arkors/oauth/model"
  "github.com/garyburd/redigo/redis"
  "github.com/go-martini/martini"
  "github.com/go-xorm/xorm"
  "github.com/martini-contrib/render"
  "net/http"
  "strconv"
)

func RegistryApp(orm *xorm.Engine, app model.Application, args martini.Params, r render.Render) {
  _, err := strconv.ParseInt(args["app"], 10, 64)
  if err != nil {
    r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The application's id must be numerical."})
  }
}

func UpdateApp(pool *redis.Pool, orm *xorm.Engine, app model.Application, args martini.Params, r render.Render) {

}

func GetAppKey(orm *xorm.Engine, args martini.Params, r render.Render) {

}

func ResetAppKey(pool *redis.Pool, orm *xorm.Engine, args martini.Params, r render.Render) {

}
