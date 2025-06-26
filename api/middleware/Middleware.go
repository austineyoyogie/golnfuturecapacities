package middleware

import "github.com/gin-gonic/gin"

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := VerifyToken(ctx)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

/*
import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"time"
)

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func RequestAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}

func generateCSRFToken() (string, error) {
	b := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func setCSRFCookieAndForm(w http.ResponseWriter, r *http.Request, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "csrfToken",
		Value:    token,
		HttpOnly: true,
		Secure:   true, // Set to true for HTTPS
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(time.Hour),
	})
	// Pass the token to the template
	// Example: <input type="hidden" name="csrfToken" value="{{.csrfToken}}">
}

func verifyCSRFToken(r *http.Request) bool {
	cookie, err := r.Cookie("csrfToken")
	if err != nil {
		return false
	}
	formToken := r.FormValue("csrfToken")
	return cookie.Value == formToken
}

func csrfMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if !verifyCSRFToken(r) {
				http.Error(w, "CSRF token mismatch", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
*/
