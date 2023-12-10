package gookie

import (
	"time"
)

type Gookie struct {
	Name       string
	Domain     string //localhost:8080 ,web.com
	Https      bool
	Expiration time.Duration //ej: 1*time.Minute
}
