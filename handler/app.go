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

      result:=new(model.ReApplication)
      result.App=app.App
      result.Key=app.Key

      r.JSON(http.StatusCreated,result)
      return
    }
  }
}

func UpdateApp(orm *xorm.Engine, app model.Application, args martini.Params, r render.Render,redisClient redis.Client) {
  appId, err := strconv.ParseInt(args["app"], 10, 64)
  if err != nil {
    r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The application's id must be numerical."})
    return
  }

  if app.Sign == "" || len(app.Sign) != 32 {
    r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid json body "})
    return
  }
  //app存在，生成新的secretKey
  //加密的公式md5(sign(userkey)+app+时间戳)= Oauth 分配的 Secret Key 
  md5String := fmt.Sprintf("%v%v%v", app.Sign,args["app"], string(time.Now().Unix()))
  h := md5.New()
  h.Write([]byte(md5String))
  secretKey := hex.EncodeToString(h.Sum(nil))
  app.App = appId
  app.Key = secretKey
  //判断更新的Application是否存在，存在进行更新，不存在返回错误
  //orm.ShowSQL=true
  affect,updateErr:=orm.Cols("sign","key","updated","created").Where("app=?",appId).Update(app)
  if updateErr!=nil {
    r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Database update error"})
    return
  } else {
    if affect==0 {
      r.JSON(http.StatusNotFound, map[string]interface{}{"error": "Application is not exists"})
      return
    } else {
      applicationJson,jsonErr:=json.Marshal(app)
      if jsonErr!=nil {
         fmt.Println(jsonErr)
         r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Struct trans to json error"})
         return
      }
      // 更新redis库
      redisClient.Set(args["app"],applicationJson)

      result:=new(model.ReApplication)
      result.App=app.App
      result.Key=app.Key

      r.JSON(http.StatusOK,result)
      return
   }
 }
}

func GetAppKey(orm *xorm.Engine, args martini.Params, r render.Render,redisClient redis.Client) {
  appId, err := strconv.ParseInt(args["app"], 10, 64)
  if err != nil {
    r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "The application's id must be numerical."})
    return
  }
  //在redis查询记录是否存在,如果不存在，在mysql中继续检索
  if result,_:=redisClient.Get(args["app"]); result==nil {
    applicationModel:=new(model.Application)
    has,err:=orm.Where("app=?",appId).Get(applicationModel)
    //在redis未查询到记录，在mysql中查询有记录，将记录回插到redis中
    if has {
       jsonApplication,_:=json.Marshal(applicationModel)
       redisClient.Set(args["app"],jsonApplication)
       resultApp:=new(model.ReApplication)
       resultApp.App=applicationModel.App
       resultApp.Key=applicationModel.Key
       r.JSON(http.StatusOK,resultApp)
       return
    } else {
         if err !=nil {
            r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Mysql Database Get Data Error!"})
            return
         }else{
            r.JSON(http.StatusNotFound, map[string]interface{}{"error": "The Application is not exists!"})
            return
         }
    }
  }else{
      var appModel model.Application
      json.Unmarshal(result,&appModel)
      resultApp:=new(model.ReApplication)
      resultApp.App=appModel.App
      resultApp.Key=appModel.Key
      r.JSON(http.StatusOK,resultApp)
      return
  }

}

func ResetAppKey(orm *xorm.Engine, args martini.Params, r render.Render) {

}
