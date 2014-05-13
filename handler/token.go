package handler

import (
	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func ExchangeAppToken(pool *redis.Pool, args martini.Params, r render.Render) {

}

func VerifyToken(pool *redis.Pool, args martini.Params, r render.Render) {

}
