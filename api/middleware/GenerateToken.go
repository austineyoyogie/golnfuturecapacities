package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golnfuturecapacities/api/config"
	"golnfuturecapacities/api/models"
	"strconv"
	"strings"
	"time"
)

var cfs = config.LoadConfig()

type JWTClaim struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateAccessToken(user *models.User) (tokenStr string, err error) {
	claims := &JWTClaim{
		UserId: strconv.Itoa(int(user.ID)),
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "https://www.lts.co.uk",
			Subject:   user.Email,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(cfs.JWT.SecretTokenKey)
	return
}

func GenerateRefreshToken(user *models.User) (tokenStr string, err error) {
	claims := &JWTClaim{
		UserId: strconv.Itoa(int(user.ID)),
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "https://www.lts.co.uk",
			Subject:   user.Email,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(cfs.JWT.SecretTokenKey)
	return
}
func VerifyToken(c *gin.Context) error {
	tokenStr := ExtractToken(c)
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return cfs.JWT.SecretTokenKey, nil
		})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil
	}
	if claims.ExpiresAt > time.Now().Add(time.Minute*1).Unix() {
		err = errors.New("token has expired")
		return nil
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	useBearer := c.Request.Header.Get("Authorization")
	if useBearer == "" {
		c.JSON(401, gin.H{"error": "required an access token!!"})
		c.Abort()
		return ""
	}
	if len(strings.Split(useBearer, " ")) == 2 {
		return strings.Split(useBearer, " ")[1]
	}
	return ""
}

// https://reliasoftware.com/blog/golang-jwt-authentication
