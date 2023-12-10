package gookie

import (
	"net/http"
)

func (g Gookie) Get(r *http.Request) (c *http.Cookie, err string) {

	const this = "gookie Get error "

	cookie, er := r.Cookie(g.Name)
	if er != nil {
		return nil, this + er.Error() + " " + g.Name
	}

	er = cookie.Valid()
	if er != nil {
		return nil, this + er.Error()
	}

	return cookie, ""
}
