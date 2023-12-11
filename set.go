package gookie

import (
	"net/http"
)

// domain ej: "example.com"
func (c Gookie) Set(value string, w http.ResponseWriter) {

	cookie := http.Cookie{
		Name:       c.Name,
		Value:      value,
		Path:       "/",      // / La cookie se aplica a todo el sitio
		Domain:     c.Domain, //"example.com"  La cookie se aplica solo a ese dominio
		Expires:    c.GetExpirationTime(),
		RawExpires: "",                      // No se utiliza un valor personalizado en bruto para la fecha de vencimiento
		MaxAge:     c.CookieAge(),           // 0 La cookie se eliminará al cerrar el navegador
		Secure:     c.Https,                 //true La cookie solo se enviará a través de una conexión segura (HTTPS)
		HttpOnly:   true,                    //true La cookie no está disponible para scripts del lado del cliente (XSS)
		SameSite:   http.SameSiteStrictMode, // 3 Restricción estricta en el envío de la cookie en solicitudes cruzadas (CSRF)
		Raw:        "",                      // ej: rawCookie := name + "=" + value + "; Path=/; Expires=Wed, 15 Jun 2023 12:00:00 GMT; Secure; HttpOnly"
		Unparsed:   []string{},              //solo recurrir a Unparsed cuando sea absolutamente necesario ej: "SameSite=None", "Priority=High"
	}

	http.SetCookie(w, &cookie)

}
