## Oauth API Specification

### API Specification

#### Oauth API Specification

##### POST /v1/apps/:app

由 Board 调用向 Oauth 模块注册 Application，获得 Oauth 分配的 Secret Key 。

###### Example Request

###### Example Response

###### Status Codes

##### GET /v1/apps/:app/keys

由 Board 调用获取 Application 的 Secret Key 。

###### Example Request

###### Example Response

###### Status Codes

##### PUT /v1/apps/:app/keys

由 Board 调用重置 Application 的 Secret Key 。

###### Example Request

###### Example Response

###### Status Codes

##### POST /v1/apps/token

由 Application 客户端调用，使用 User Key 和 Application Secret Key 根据算法生成的 Token 换取 Arkors 平台运行的 Token 。

###### Example Request

###### Example Response

###### Status Codes

##### POST /v1/apps/token/verify/:token/:time

其它模块在获得 Application 客户端发出的访问请求中的 Token 信息向 Oauth 模块进行验证。

###### Example Request

###### Example Response

###### Status Codes
