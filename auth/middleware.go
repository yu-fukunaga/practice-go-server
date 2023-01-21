package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserAuth struct {
	UserID uuid.UUID
	Token  string
}

type contextKey struct {
	name string
}

type UserClaims struct {
	UserID uuid.UUID `Json:"user_id"`
	jwt.StandardClaims
}

var issuer = []byte(os.Getenv("SIGNINGKEY"))
var userCtxKey = &contextKey{name: "user"}

// JwtMiddleware is middleware for grpc server
func JwtMiddleware() (func(http.Handler) http.Handler){
  return func(next http.Handler) http.Handler {
		return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      token := ExtractTokenFromHTTPRequest(r)
			userID := ExtractUserIDFromToken(token)
			userAuth := UserAuth{UserID: userID}
			ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
			r = r.WithContext(ctx)
			next.ServeHTTP(w,r)
		})
	}
}


func ExtractTokenFromHTTPRequest(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	var token string
	splits := strings.Split(reqToken, "Barer ")
	if len(splits) > 1 {
		token = splits[1]
	}
	return token
}

func ExtractUserIDFromToken(tokeNString string) uuid.UUID{
	UserClaims := UserClaims{}
	token, err := jwt.ParseWithClaims(tokeNString, &UserClaims, func (token *jwt.Token) (interface{}, error) {
		return issuer, nil
	})
	if err != nil {
		return uuid.Nil
	}
	if !(token.Valid){
		return uuid.Nil
	}
	return UserClaims.UserID
}

func GetAuthFromContext(ctx context.Context) *UserAuth {
	raw := ctx.Value(userCtxKey)
	if raw == nil {
		return nil
	}
	return raw.(*UserAuth)
}

func JWTGenerate(userID uuid.UUID, expiredAt int64) string {
	claims := UserClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer: string(issuer),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	signedToken, _ := token.SignedString(issuer)
	return signedToken
}