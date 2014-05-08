package main

import (
  "fmt"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "strconv"
)

func RegistryApp(app Application, args martini.Params, r render.Render) {
  fmt.Println(app.Sign)
  appId, err := strconv.ParseInt(args["app"], 10, 64)
  if err == nil {
    r.JSON(201, map[string]interface{}{"app": appId, "key": "cb21df532c6647383af7efa0fd8405f2"})
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
