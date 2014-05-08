package main

import (
  "github.com/go-martini/martini"
  "github.com/go-xorm/xorm"
  "github.com/martini-contrib/render"
  "net/http"
  "strconv"
)

func RegistryApp(orm *xorm.Engine, app Application, args martini.Params, r render.Render) {
  appId, err := strconv.ParseInt(args["app"], 10, 64)
  if err == nil {
    r.JSON(http.StatusCreated, map[string]interface{}{"app": appId, "key": "cb21df532c6647383af7efa0fd8405f2"})
  } else {
    r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The application's id must be numerical."})
  }
}

func UpdateApp(args martini.Params, r render.Render) {

}

func GetAppKey(args martini.Params, r render.Render) {

}

func ResetAppKey(args martini.Params, r render.Render) {

}

func ExchangeAppToken(args martini.Params, r render.Render) {

}

func VerifyToken(args martini.Params, r render.Render) {
}
