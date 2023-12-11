package gookie

import "time"

func (c *Gookie) GetExpirationTime() time.Time {

	var duration = 2 * time.Minute

	c.age = -1 // La cookie expirará al cerrar el navegador

	switch {
	case c.OneHour:
		duration = 1 * time.Hour

	case c.OneDay:
		duration = 24 * time.Hour // Expira en un día

	case c.Forever:
		c.age = 0
		return time.Now().AddDate(9999, 0, 0)
	}

	c.age = int(duration.Seconds())

	return time.Now().Add(duration)
}

func (a Gookie) CookieAge() int {
	return a.age
}
