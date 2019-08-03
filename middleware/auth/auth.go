package auth

import (
	"os"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/tingxin/bingo/setting"
)

var j *jwt.Middleware

func init() {
	key := os.Getenv("Secret")
	if key == "" {
		panic("failed to get env Secret")
	}
	secret := []byte(key)

	j = jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},

		// Extract by the "token" url.
		// There are plenty of options.
		// The default jwt's behavior to extract a token value is by
		// the `Authentication: Bearer $TOKEN` header.
		Extractor: jwt.FromParameter(setting.AuthKey),
		// When set, the middleware verifies that tokens are
		// signed with the specific signing algorithm
		// If the signing method is not constant the `jwt.Config.ValidationKeyGetter` callback
		// can be used to implement additional checks
		// Important to avoid security issues described here:
		// https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})
}

// Run will check the auth status in one request
func Run(ctx iris.Context) {
	j.Serve(ctx)
}
