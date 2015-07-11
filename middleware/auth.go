package middleware

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/techjanitor/pram-post/config"
	e "github.com/techjanitor/pram-post/errors"
	u "github.com/techjanitor/pram-post/utils"
)

// checks for session cookie and handles permissions
func Auth(perms Permissions) gin.HandlerFunc {
	return func(c *gin.Context) {

		// set default anonymous user
		user := u.User{
			Id:    1,
			Group: 0,
		}

		// parse jwt token if its there
		token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {

			// check alg
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.Settings.Session.Secret), nil
		})
		// if the error is anything but no token
		if err != jwt.ErrNoTokenInRequest {
			c.JSON(e.ErrorMessage(e.ErrUnauthorized))
			c.Error(e.ErrUnauthorized)
			c.Abort()
			return
		}

		// if the token is valid set the data
		if err == nil && token.Valid {
			user.Id = token.Claims["user_id"]
			user.Group = token.Claims["user_group"]

		}

		fmt.Println(user.Id, user.Group)

		// check if user meets set permissions
		if user.Group < perms.Minimum {
			c.JSON(e.ErrorMessage(e.ErrUnauthorized))
			c.Error(e.ErrUnauthorized)
			c.Abort()
			return
		}

		// set user data
		c.Set("userdata", user)

		c.Next()

	}

}

// permissions data
type Permissions struct {
	Minimum uint
}

func SetAuthLevel() Permissions {
	return Permissions{}
}

// All users
func (p Permissions) All() Permissions {
	p.Minimum = 0
	return p
}

// registered users
func (p Permissions) Registered() Permissions {
	p.Minimum = 1
	return p
}

// moderators
func (p Permissions) Moderators() Permissions {
	p.Minimum = 2
	return p
}

// admins
func (p Permissions) Admins() Permissions {
	p.Minimum = 3
	return p
}
