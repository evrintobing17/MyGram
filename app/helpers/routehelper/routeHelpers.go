package routehelper

import (
	"errors"
	"strconv"

	"github.com/evrintobing17/MyGram/app/middlewares/authmiddleware"
	"github.com/evrintobing17/MyGram/app/models"

	"github.com/gin-gonic/gin"
)

var (
	ErrorContextNotExist  = errors.New("user context not exist")
	ErrorParsingUserModel = errors.New("error parsing user model")
)

func GetUserFromJWTContext(c *gin.Context) (*models.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, ErrorContextNotExist
	}
	if user == authmiddleware.ErrUserContextNotSet.Error() {
		return nil, ErrorContextNotExist
	}

	userModel, ok := user.(*models.User)
	if !ok {
		return nil, ErrorParsingUserModel
	}

	return userModel, nil
}

func GetAccessTokenClaims(c *gin.Context) (interface{}, error) {
	claims, exists := c.Get("access_token_claims")
	if !exists {
		return nil, ErrorContextNotExist
	}
	if claims == authmiddleware.ErrUserContextNotSet.Error() {
		return nil, ErrorContextNotExist
	}

	return claims, nil
}

func GetPageRequestFromHeader(c *gin.Context) (int, error) {
	var err error
	page := 1
	pageHeader := c.GetHeader("x-page")
	if pageHeader != "" {
		page, err = strconv.Atoi(pageHeader)
		if err != nil {
			return 0, err
		}
	}
	return page, nil
}

func GetPageLimitRequestFromHeader(c *gin.Context) (int, error) {
	var err error
	pageLimit := 250
	pageLimitHeader := c.GetHeader("x-per-page-limit")
	if pageLimitHeader != "" {
		pageLimit, err = strconv.Atoi(pageLimitHeader)
		if err != nil {
			return 0, err
		}
	}
	return pageLimit, nil
}

func GetOrderFieldFromHeader(c *gin.Context) string {
	return c.GetHeader("x-order-field")
}
