package serverhandler

import (
	"os"
	"regexp"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type WebsocketHandler struct {
	r       *gin.Engine
	address string
}

//NewWebsocketHandler will return new handler for websocket
func NewWebsocketHandler(address string) *WebsocketHandler {

	//initialize new Gin Router
	router := gin.Default()
	var rxURL = regexp.MustCompile(`^/regexp\d*`)
	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.=
	subLog := zerolog.New(os.Stdout).With().Logger()

	router.Use(logger.SetLogger(logger.Config{
		Logger:         &subLog,
		UTC:            true,
		SkipPath:       []string{"/skip"},
		SkipPathRegexp: rxURL,
	}))

	router.Use(gin.Recovery())

	return &WebsocketHandler{r: router, address: address}
}

//GetWSRouter will return pointer of WebsocketHandler struct gin router
func (websocket *WebsocketHandler) GetWSRouter() *gin.Engine {
	return websocket.r
}

//RunWebsocketServer will run gin router as websocket server handler
func (websocket *WebsocketHandler) RunWebsocketServer() error {
	err := websocket.r.Run(websocket.address)
	if err != nil {
		return err
	}
	return nil
}
