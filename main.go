package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// docs "github.com/evrintobing17/docs"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/evrintobing17/MyGram/app/middlewares/authmiddleware"

	usersauthdelivery "github.com/evrintobing17/MyGram/app/modules/users/delivery"
	usersrepository "github.com/evrintobing17/MyGram/app/modules/users/repository"
	usersusecase "github.com/evrintobing17/MyGram/app/modules/users/usecase"

	photodelivery "github.com/evrintobing17/MyGram/app/modules/photo/delivery"
	photorepository "github.com/evrintobing17/MyGram/app/modules/photo/repository"
	photousecase "github.com/evrintobing17/MyGram/app/modules/photo/usecase"

	commentdelivery "github.com/evrintobing17/MyGram/app/modules/comment/delivery"
	commentrepository "github.com/evrintobing17/MyGram/app/modules/comment/repository"
	commentusecase "github.com/evrintobing17/MyGram/app/modules/comment/usecase"

	socialmediadelivery "github.com/evrintobing17/MyGram/app/modules/socialmedia/delivery"
	socialmediarepo "github.com/evrintobing17/MyGram/app/modules/socialmedia/repository"
	socialmediausecase "github.com/evrintobing17/MyGram/app/modules/socialmedia/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func Initialize(Dbdriver, DbUser, DbPassword, DbHost, DbName string, DbPort int) (DB *gorm.DB) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbName, DbPassword)
		fmt.Println(DBURL)
		DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Println("Cannot connect to database")
			log.Fatal("Error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
	}

	return DB

}

func run() {

	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error opening env, %v", err)
	} else {
		fmt.Println(".env file loaded")
	}
	db_driver := os.Getenv("DB_DRIVER")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	db_port, _ := strconv.Atoi(port)
	DB := Initialize(db_driver, db_user, db_password, db_host, db_name, db_port)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaulTableName string) string {
		return "" + defaulTableName
	}

	// orderRepo := orderRepository.NewOrderRepository(DB)

	// orderUC := orderUC.NewOrderUsecase(orderRepo)
	// docs.SwaggerInfo.BasePath = ""

	r := gin.New()
	// orderDelivery.NewHttpDelivery(r, orderUC)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Repository
	usersRepository := usersrepository.NewUserRepository(DB)
	photoRepository := photorepository.NewPhotosRepository(DB)
	commentRepo := commentrepository.NewCommentsRepository(DB)
	socialMediaRepo := socialmediarepo.NewsocialMediaRepository(DB)

	//Usecase
	usersUsecase := usersusecase.NewUserUsecase(usersRepository)
	photoUsecase := photousecase.NewPhotoUsecase(photoRepository)
	commentUC := commentusecase.NewCommentUsecase(commentRepo, photoRepository)
	socialMediaUseCase := socialmediausecase.NewsocialMediaUsecase(socialMediaRepo, photoRepository)

	//Middleware
	authMiddleware := authmiddleware.NewAuthMiddleware(usersRepository)

	//Presenter
	usersauthdelivery.NewAuthHTTPHandler(r, usersUsecase, authMiddleware)
	photodelivery.NewPhotoHTTPHandler(r, photoUsecase, authMiddleware)
	commentdelivery.NewcommentHTTPHandler(r, commentUC, authMiddleware)
	socialmediadelivery.NewSocialMediaHandler(r, socialMediaUseCase, authMiddleware)

	fmt.Println("Listening to port 8081")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}

func main() {

	run()
}
