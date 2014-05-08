package model

import (
	_ "github.com/go-xorm/xorm"
	"time"
)

type Application struct {
	Id      int64
	App     int64     `xorm:"unique not null"`
	Sign    string    `json:"sign"xorm:"unique not null"`
	key     string    `xorm:"unique not null"`
	Updated time.Time `xorm:"updated"`
	Created time.Time `xorm:"created"`
	Version int       `xorm:"version"`
}
