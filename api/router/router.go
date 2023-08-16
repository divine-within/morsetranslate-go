package router

import (
	"fmt"
	"net/http"

	"github.com/divine-within/morsetranslate-go/api/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func GenerateJWT(s string) (*jwtauth.JWTAuth, string) {
	tokenAuth := jwtauth.New("HS256", []byte(s), nil)
	// TODO: Add DB with users and their respective JWT tokens
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	return tokenAuth, tokenString
}

func init() {
	// TODO: Change the secret
	var tokenString string
	tokenAuth, tokenString = GenerateJWT("secret")
	fmt.Printf("DEBUG: JWT sample token is %s\n\n", tokenString)
}

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Accept", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(tokenAuth))
		router.Use(jwtauth.Authenticator)

		router.Get("/api/v1/dictionary", controllers.GetDictionary)
		router.Get("/api/v1/translate/text", controllers.GetMorseCode)
		router.Get("/api/v1/translate/morsecode", controllers.GetText)
	})
	return router
}
