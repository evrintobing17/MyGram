package authmiddleware

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	AuthorizeJWT(roles string) gin.HandlerFunc
	AuthorizeJWTWithUserContext() gin.HandlerFunc
	AuthorizeAccessToken(scope string) gin.HandlerFunc
}
