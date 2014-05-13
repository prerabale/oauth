## Oauth API Specification

### API Specification

#### Oauth API Specification

##### POST /v1/apps/:app

由 Board 调用向 Oauth 注册 Application，获得 Oauth 分配的 Secret Key 。

> `sign` 参数是 Arkors 用户 User Key 的 MD5 值。

###### Example Request
```
POST /v1/apps/232 HTTP/1.1
Host: oauth.arkors.com
X-Arkors-Application-Log: 5024442115e7bd738354c1fac662aed5
X-Arkors-Application-Client: xxx.xx.xx.xxx,BOARD
Accept: application/json
{
  "sign": "5024442115e7bd738354c1fac662aed5"
}
```

###### Example Response
```
HTTP/1.1 201 OK
X-Arkors-Application-Log: cb21df532c6647383af7efa0fd8405f2
Content-Type: application/json
{
  "app": 232,
  "key": "cb21df532c6647383af7efa0fd8405f2"
}
```

###### Status Codes
* 201 - 创建 Application 记录成功
* 400 - Errors (invalid json, missing, duplication or invalid fields, etc)

##### PUT /v1/apps/:app

由 Board 调用更新 Application 的 User Key 。 User Key 更新后，Application 的 Secret Key 也会重新生成， 该 Application 的 Token 全部失效。

###### Example Request
```
PUT /v1/apps/232 HTTP/1.1
Host: oauth.arkors.com
X-Arkors-Application-Log: 5024442115e7bd738354c1fac662aed5
X-Arkors-Application-Client: xxx.xx.xx.xxx,BOARD
Accept: application/json
{
  "sign": "5024442115e7bd738354c1fac662aed5"
}
```

###### Example Response
```
HTTP/1.1 200 OK
X-Arkors-Application-Log: cb21df532c6647383af7efa0fd8405f2
Content-Type: application/json
{
  "app": 232,
  "key": "cb21df532c6647383af7efa0fd8405f2"
}
```

###### Status Codes
* 200 - 更新 Application 的 User Key 成功
* 400 - Errors (invalid json, missing or invalid fields, etc)
* 404 - 没有找到 Application 的记录

##### GET /v1/apps/:app/key

由 Board 调用获取 Application 的 Secret Key 。

###### Example Request
```
GET /v1/apps/232/key HTTP/1.1
Host: oauth.arkors.com
X-Arkors-Application-Log: 5024442115e7bd738354c1fac662aed5
X-Arkors-Application-Client: xxx.xx.xx.xxx,BOARD
Accept: application/json
```

###### Example Response
```
HTTP/1.1 200 OK
X-Arkors-Application-Log: cb21df532c6647383af7efa0fd8405f2
Content-Type: application/json
{
  "app": 232,
  "key": "cb21df532c6647383af7efa0fd8405f2"
}
```

###### Status Codes
* 200 - 返回 Application 的 User Key
* 400 - Errors (invalid json, missing or invalid fields, etc)
* 404 - 没有找到 Application 的记录

##### PUT /v1/apps/:app/key

由 Board 调用重置 Application 的 Secret Key 。Application 的 Secret Key 重新生成后该 Application 的 Token 全部失效。

###### Example Request
```
PUT /v1/apps/232/key HTTP/1.1
Host: oauth.arkors.com
X-Arkors-Application-Log: 5024442115e7bd738354c1fac662aed5
X-Arkors-Application-Client: xxx.xx.xx.xxx,BOARD
Accept: application/json
```

###### Example Response
```
HTTP/1.1 200 OK
X-Arkors-Application-Log: cb21df532c6647383af7efa0fd8405f2
Content-Type: application/json
{
  "app": 232,
  "key": "cb21df532c6647383af7efa0fd8405f2"
}
```

###### Status Codes
* 200 - 更新 Application 的 User Key 成功
* 400 - Errors (invalid json, missing or invalid fields, etc)
* 404 - 没有找到 Application 的记录

##### POST /v1/apps/:app/sign

由 Application 客户端调用，使用 User Key 和 Application Secret Key 根据算法生成的签名换取 Arkors 平台运行的 Token 。

> `timestamp` - 客户端产生本次请求的 Unix 时间戳，精确到毫秒。
> `sign` - User Key 的 MD5 值字符串加上 Application Secret Key 和 timestamp 再计算 MD5 值。 
> 将 `sign` 和 `timestamp` 两个字符串用 `,` 连接作为 sign 的值。

###### Example Request
```
POST /v1/apps/233/sign HTTP/1.1
Host: oauth.arkors.com
X-Arkors-Application-Id: 232
X-Arkors-Application-Sign: cb21df532c6647383af7efa0fd8405f2,1389085779854
X-Arkors-Application-Log: 5024442115e7bd738354c1fac662aed5
X-Arkors-Application-Client: 3ad3ce877d6c42b131580748603f8d6a,ANDROID
Accept: application/json
```

###### Example Response
```
HTTP/1.1 201 OK
X-Arkors-Application-Log: cb21df532c6647383af7efa0fd8405f2
Content-Type: application/json
{
  "token": "cb21df532c6647383af7efa0fd8405f2"
}
```

###### Status Codes
* 201 - 生成 Token 记录成功
* 400 - Errors (invalid json, missing, duplication or invalid fields, etc)
* 404 - 没有找到 Application 的记录

##### GET /v1/apps/token/verify/:token/:timestamp

其它模块在获得 Application 客户端访问请求中的 Token 信息向 Oauth 模块进行验证。

> `timestamp` - 客户端产生本次请求的 Unix 时间戳，精确到毫秒。
> `sign` - 客户端通过 /v1/apps/sign 获得的 Token 加上 Application Secret Key 和 timestamp 计算 MD5 值。 
> 将 `sign` 和 `timestamp` 两个字符串用 `,` 连接作为 Token 的值。


###### Example Request
```
GET /v1/apps/token/verify/cb21df532c6647383af7efa0fd8405f2/1389085779854 HTTP/1.1
Host: oauth.arkors.com
X-Arkors-Application-Id: 232
X-Arkors-Application-Log: 5024442115e7bd738354c1fac662aed5
X-Arkors-Application-Client: xxx.xx.xx.xxx,AUTH
Accept: application/json
```

###### Example Response
```
HTTP/1.1 200 OK
X-Arkors-Application-Log: cb21df532c6647383af7efa0fd8405f2
Content-Type: application/json
```

###### Status Codes
* 200 - Token 验证成功
* 400 - Errors (invalid json, missing, duplication or invalid fields, etc)
* 404 - 没有找到 Application 的记录
