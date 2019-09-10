package jwt

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"

	"github.com/tingxin/bingo/service/auth/cmd"
	"github.com/tingxin/bingo/service/auth/dao"
)

var (
	key []byte
)

func init() {
	key := os.Getenv("Secret")
	if key == "" {
		panic("failed to get env Secret")
	}
}

// Entity used to login in system by jwt
type Entity struct {
	secret []byte
}

// NewEntity used to creat new jwt login process
func NewEntity() *Entity {

	j := Entity{}
	j.secret = []byte(key)
	return &j
}

// Sign by email and password
func (p *Entity) Sign(cmd cmd.SignCmd) (interface{}, int, error) {

	user, code, err := dao.GetUserInfoByMail(cmd.Eamil)
	if err != nil {
		return nil, code, err
	}

	sha := sha256.New()
	sha.Write([]byte(user.Password))

	passHash := fmt.Sprintf("%x", sha.Sum(nil))

	if passHash != cmd.Password {
		return nil, iris.StatusUnauthorized, fmt.Errorf("failed to sign by %s's password", cmd.Eamil)
	}

	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Eamil,
		"name":  user.Name,
		"role":  user.Role,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(p.secret)
	return tokenString, iris.StatusOK, nil
}

// SignUp by email and password
func (p *Entity) SignUp(cmd cmd.SignUpCmd) (interface{}, int, error) {

	code, err := dao.AddUser(cmd.Eamil, cmd.Name, cmd.Password)
	if err != nil {
		return nil, code, err
	}

	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": cmd.Eamil,
		"name":  cmd.Name,
		"role":  "normal",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(p.secret)
	return tokenString, iris.StatusOK, nil
}
