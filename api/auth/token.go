package auth

import (
	"github.com/obasajujoshua31/blogos/api/models"
	"github.com/obasajujoshua31/blogos/config"
	"github.com/robbert229/jwt"
)

type UserId struct {
	userId int
}

func Algorithm() (algorithm jwt.Algorithm) {
	algorithm = jwt.HmacSha256(config.SECRETKEY)
	return
}

func GenerateToken(user models.User) (token string, err error) {
	claims := jwt.NewClaim()
	claims.Set("userId", user.ID)

	algorithm := Algorithm()

	token, err = algorithm.Encode(claims)
	return
}

// func GetUserId(token string) (userId interface) {
// 	algorithm := Algorithm()

// 	// claims, err := algorithm.Decode(token)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// userId, _ = claims.Get("userId")
// }
