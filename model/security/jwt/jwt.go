package jwt

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// custom claims
type Claims struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

// jwt secret key
var token_pwd = []byte("secret")

// use it after database login check
func JwtGen(ac, pw string) (string, error) {
	now := time.Now()
	jwtId := ac + strconv.FormatInt(now.Unix(), 10)
	role := "Member"
	// set claims and sign
	claims := Claims{
		Account: ac,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			Audience:  ac,
			ExpiresAt: now.Add(10 * time.Minute).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(10 * time.Second).Unix(),
			Subject:   ac,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(token_pwd)
}

// validate JWT
func AuthRequired(token string) (*Claims, string, error) {

	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return token_pwd, nil
	})
	if err == nil && tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, "", nil
		}
		return nil, "token missing", nil
	}
	print(err.Error())
	return nil, jeterror(err), err
}

func jeterror(err error) (message string) {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			message = "token is malformed"
		} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
			message = "token could not be verified because of signing problems"
		} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
			message = "signature validation failed"
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			message = "token is expired"
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			message = "token is not yet valid before sometime"
		} else {
			message = "can not handle this token"
		}
	}
	return message
}
