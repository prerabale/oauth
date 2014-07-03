package handler

import (
  "github.com/arkors/oauth/model"
  "github.com/go-martini/martini"
  "github.com/go-xorm/xorm"
  "github.com/martini-contrib/render"
  "net/http"
  "strconv"
  "crypto/md5"
  "encoding/hex"
  "encoding/json"
  "fmt"
  "time"
  "github.com/hoisie/redis"
)

func RegistryApp(orm *xorm.Engine, app model.Application, args martini.Params, r render.Render,redisClient redis.Client) {
  appId, err := strconv.ParseInt(args["app"], 10, 64)
  if err != nil {
    r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The application's id must be numerical."})
    return
  }

  if app.Sign == "" || len(app.Sign) != 32 {
    r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid json body "})
    return
  }

  applicationModel:=new(model.Application)
    //判断app已经被注册
  has,errDb:=orm.Where("app=?",appId).Get(applicationModel)
  if errDb!=nil {
    r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Database connect error"})
    return
  } else {
    if has {
      r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The App already registry!"})
      return
    } else {
      //加密的公式md5(sign(userkey)+app+时间戳)= Oauth 分配的 Secret Key 
      md5String := fmt.Sprintf("%v%v%v", app.Sign,args["app"], string(time.Now().Unix()))
      h := md5.New()
      h.Write([]byte(md5String))
      secretKey := hex.EncodeToString(h.Sum(nil))
      app.App = appId
      app.Key = secretKey
      _,errInsert:=orm.Insert(app)

      if errInsert !=nil {
        r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Database Insert Error"})
        return
      }
      //将对象转换成JSON格式，存入redis redis中的格式   key=appId  value=对象JSON
      applicationJson,jsonErr:=json.Marshal(app)
      if jsonErr!= nil {
        r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Struct trans to json error"})
        return
      }
      redisClient.Set(args["app"],applicationJson)
      r.JSON(http.StatusCreated,app)
    }
  }
}

func UpdateApp(orm *xorm.Engine, app model.Application, args martini.Params, r render.Render) {

}

func GetAppKey(orm *xorm.Engine, args martini.Params, r render.Render) {

}

func ResetAppKey(orm *xorm.Engine, args martini.Params, r render.Render) {

}
