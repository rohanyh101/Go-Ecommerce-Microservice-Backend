package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/roh4nyh/ecom/config"
	"github.com/roh4nyh/ecom/types"
	"github.com/roh4nyh/ecom/utils"
)

type ContextKey string

const UserKey ContextKey = "user_id"

func CreateJWT(secret []byte, user_id int) (string, error) {
	duration, _ := strconv.ParseInt(config.Env.JWTExpirationInSeconds, 10, 64)
	expiration := time.Second * time.Duration(duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    strconv.Itoa(user_id),
		"expired_at": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from user request...
		tokenString := getTokenFromRequest(r)

		// and validate the token...
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("falied to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		// fetch the user_id from the token...
		claims := token.Claims.(jwt.MapClaims)
		str := claims["user_id"].(string)

		userID, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert user_id to int: %v", err)
			return
		}

		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		// set the context with the user_id...
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)

		// assign modefied context to the request...
		r = r.WithContext(ctx)

		// and call the handlerFunc...
		handlerFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth != "" {
		return tokenAuth
	}

	return ""
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Env.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
