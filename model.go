package gookie

import "github.com/cdvelop/model"

type Gookie struct {
	Name   string
	Domain string //localhost:8080 ,web.com
	Https  bool

	model.CookieExpiration //ej: 1*time.Minute
	age                    int
}
