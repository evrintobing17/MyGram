package registry

import (
	"github.com/evrintobing17/MyGram/app/helpers/dsnbuilder"
	"github.com/evrintobing17/MyGram/app/middlewares/authmiddleware"
	usersauthdelivery "github.com/evrintobing17/MyGram/app/modules/users/delivery"
	usersrepository "github.com/evrintobing17/MyGram/app/modules/users/repository"
	usersusecase "github.com/evrintobing17/MyGram/app/modules/users/usecase"

	"fmt"
	"os"
	"strconv"

	"github.com/evrintobing17/MyGram/app/registry/serverhandler"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	// Application configuration variable
	appVersion string
	appPort    string
	wsPort     string

	// DB configuration variable
	dbHost     string
	dbPort     string
	dbDatabase string
	dbUsername string
	dbPassword string
	dbDialect  string
	dbDsn      string
)

type AppRegistry struct {
	dbConn *gorm.DB

	httpHandler      *serverhandler.HttpHandler
	websocketHandler *serverhandler.WebsocketHandler
}

func initializeEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error().Msg("Failed read configuration database")
		return
	}

	if os.Getenv("project_env") == "DEV" || os.Getenv("project_env") == "" { //project_env is Development

		appVersion = os.Getenv("application.version")
		appPort = os.Getenv("application.port")
		wsPort = os.Getenv("application.websocket.port")

		//specify db dialect
		dbDialect = os.Getenv("db.dialect")
		dbDialect = "postgres"

		dbHost = os.Getenv("db." + dbDialect + ".host")
		dbPort = os.Getenv("db." + dbDialect + ".port")
		dbUsername = os.Getenv("db." + dbDialect + ".username")
		dbPassword = os.Getenv("db." + dbDialect + ".password")
		dbDatabase = os.Getenv("db." + dbDialect + ".dbname")

	} else if os.Getenv("project_env") == "docker" { //project_env is STAGING
		appVersion = os.Getenv("application.version")
		appPort = os.Getenv("application.port")
		wsPort = os.Getenv("application.websocket.port")

		//specify db dialect
		dbDialect = os.Getenv("db.dialect")

		dbHost = os.Getenv("db." + dbDialect + ".docker.host")
		dbPort = os.Getenv("db." + dbDialect + ".docker.port")
		dbUsername = os.Getenv("db." + dbDialect + ".docker.username")
		dbPassword = os.Getenv("db." + dbDialect + ".docker.password")
		dbDatabase = os.Getenv("db." + dbDialect + ".docker.dbname")
	}

	//create DSN Builder
	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		panic(err)
	}

	dbDsn, err = dsnbuilder.New(dbHost, dbPortInt, dbUsername, dbPassword, dbDatabase).Build(dbDialect)
	if err != nil {
		panic(err)
	}
}

//NewAppRegistry will return new object for App Registry
func NewAppRegistry() *AppRegistry {
	return &AppRegistry{}
}

//init will initialize all the needed component for Go
func (registry *AppRegistry) initializeAppRegistry() {

	appASCIIArt := `
Bogadelivery Backend Service
2020 by Chrombit Digital
`

	fmt.Println(appASCIIArt)

	//Initialize Logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stdin,
			NoColor: false,
		},
	)

	initializeEnv()
	registry.initializeHandler()
	err := registry.initializeDependency()
	if err != nil {
		panic("app exit on error")
	}

	registry.initializeDomainFeature()
}

func (reg *AppRegistry) StartServer() {
	reg.initializeAppRegistry()

	//Run Websocket Server
	log.Info().Msg("Websocket API Service Running version " + appVersion + " at port : " + wsPort)
	go func() {
		if errWS := reg.websocketHandler.RunWebsocketServer(); errWS != nil {
			log.Error().Msg(errWS.Error())

		}
	}()

	//Run HTTP Server
	log.Info().Msg("REST API Service Running version " + appVersion + " at port : " + appPort)
	if errHTTP := reg.httpHandler.RunHttpServer(); errHTTP != nil {
		log.Error().Msg(errHTTP.Error())
	}

	//Close connection
	defer func() {
		err := reg.dbConn.Close()
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}()

}

func (reg *AppRegistry) initializeHandler() {

	/*
		Register HTTP Server Handler
	*/
	reg.httpHandler = serverhandler.NewHTTPHandler(":" + appPort)
	reg.httpHandler.RunSwaggerMiddleware() //Run Swagger UI Docs

	/*
		Register Websocket Server Handler
	*/
	reg.websocketHandler = serverhandler.NewWebsocketHandler(":" + wsPort)

}

func (reg *AppRegistry) initializeDependency() error {
	err := reg.initializeMySQLDatabase()
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

func (reg *AppRegistry) initializeMySQLDatabase() error {
	// DB Connection Configuration
	// Handles GORM
	db, err := gorm.Open(dbDialect, dbDsn)
	if err != nil {
		log.Error().Msg("Failed Connecting to database " + err.Error())
		return err
	}

	log.Info().Msg("Connected to Database with following configuration:" +
		"\n Database Dialect \t: " + dbDialect +
		"\n Database Host \t\t: " + dbHost + ":" + dbPort +
		"\n Database Name \t\t: " + dbDatabase)

	//Table Migration

	//case if using postgres
	if dbDialect == "postgres" {
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return "bogadelivery_core." + defaultTableName
		}
	}
	//db.AutoMigrate(&models.DriverLicense{})
	reg.dbConn = db
	return nil
}

func (reg *AppRegistry) initializeDomainFeature() {

	//Repository
	usersRepository := usersrepository.NewUserRepository(reg.dbConn)

	//Usecase
	usersUsecase := usersusecase.NewUserUsecase(usersRepository)

	//Middleware
	authMiddleware := authmiddleware.NewAuthMiddleware(usersRepository)
	// reg.httpHandler.GetRouteEngine().Use(audittrail.PreRequest())
	// reg.httpHandler.GetRouteEngine().Use(audittrail.PostRequest(userService, auditTrailService))

	//Presenter
	usersauthdelivery.NewAuthHTTPHandler(reg.httpHandler.GetRouteEngine(), usersUsecase, authMiddleware)
}
