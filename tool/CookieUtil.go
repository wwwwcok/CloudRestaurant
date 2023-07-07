package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const CookieName = "cookie_user"
const CookieTimeLength = 10 * 60

func CookieAuth(context *gin.Context) (*http.Cookie, error) {
	cookie, err := context.Request.Cookie(CookieName)
	if err != nil {
		return nil, err
	} else {
		context.SetCookie(cookie.Name, cookie.Value, 10*60, "/", "localhost", true, true)
	}
	return cookie, nil
}
