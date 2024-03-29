package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jakeroggenbuck/BestNextStep/daft/col"
	"github.com/jakeroggenbuck/BestNextStep/daft/step"
	"github.com/jakeroggenbuck/BestNextStep/daft/user"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"time"
)

func getLogIn() gin.Accounts {
	return gin.Accounts{
		"Admin": GetEnvOrFatal("ADMIN_PASSWORD"),
	}
}

func getLocalIP() string {
	return GetEnvOrDefault("LOCAL_IP", "127.0.0.1")
}

func setupLogging() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

func dbExists() bool {
	if _, err := os.Stat("./sqlite.db"); err == nil {
		return true
	}
	return false
}

func corsMiddle() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		fmt.Println(c.Request.Method)
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, struct{}{})
		}
	})
}

func main() {
	setupLogging()

	dbExisted := dbExists()
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal("Database open failed")
	}

	stepRepository := step.NewSQLiteRepository(db)
	if err := stepRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	colRepository := col.NewSQLiteRepository(db)
	if err := colRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewSQLiteRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	// Create default items if db is new
	if !dbExisted {
		createDefaultElements(db)
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{getLocalIP()})
	router.LoadHTMLGlob("./web/templates/**/*")

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://127.0.0.1:8080", "http://localhost:8080"}
	// corsConfig.AddAllowMethods("OPTIONS", "POST", "DELETE")
	// corsConfig.AllowHeaders = []string{"Origin", "Authorization", "content-type", "body", "data", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"}
	// corsConfig.AllowCredentials = true

	cors := cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:8080", "http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "content-type", "body", "data", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	router.Use(cors)
	router.Use(corsMiddle())

	/*

		GET => 		/ 						homePage
		GET => 		/version 				version


		AUTHED USER		Group	/api/v1
		===========		=====	=======

		GET => 		/api/v1/				apiRootPage

		GET => 		/api/v1/step			allStep
		POST => 	/api/v1/step/			addStep
		PUT => 		/api/v1/step/			updateStep
		DELETE => 	/api/v1/step/			deleteStep

		GET => 		/api/v1/col				allCol
		POST => 	/api/v1/col/			addCol
		PUT => 		/api/v1/col/			updateCol
		DELETE => 	/api/v1/col/			deleteCol

		POST => 	/api/v1/user/			addUser
		PUT => 		/api/v1/user/			updateUser
		DELETE => 	/api/v1/user/			deleteUser

		ADMIN USER		Group	/api/v1/admin
		==========		=====	=============

		GET =>		/api/v1/admin/stats				allStats
		GET =>		/api/v1/admin/stats/user-count	userCount
		GET =>		/api/v1/admin/stats/step-count	stepCount

	*/

	router.GET("/", homePage)

	authAccount := getLogIn()
	authedSubRoute := router.Group("/api/v1/", gin.BasicAuth(authAccount))
	{

		stepSubRoute := authedSubRoute.Group("/step/")
		{
			stepSubRoute.GET("/aa", preflight)
			stepSubRoute.GET("/", func(c *gin.Context) { allStep(c, stepRepository) })
			stepSubRoute.POST("/", func(c *gin.Context) { addStep(c, stepRepository) })
			stepSubRoute.PUT("/:id", func(c *gin.Context) { updateStep(c, stepRepository) })
			stepSubRoute.DELETE("/:id", func(c *gin.Context) { deleteStep(c, stepRepository) })
		}

		colSubRoute := authedSubRoute.Group("/col/")
		{
			colSubRoute.GET("/", func(c *gin.Context) { allCol(c, colRepository) })
			colSubRoute.POST("/", func(c *gin.Context) { addCol(c, colRepository) })
			colSubRoute.PUT("/:id", func(c *gin.Context) { updateCol(c, colRepository) })
			colSubRoute.DELETE("/:id", func(c *gin.Context) { deleteCol(c, colRepository) })
		}

		userSubRoute := authedSubRoute.Group("/user/")
		{
			userSubRoute.POST("/", func(c *gin.Context) { addUser(c, userRepository) })
			userSubRoute.PUT("/:id", func(c *gin.Context) { updateUser(c, userRepository) })
			userSubRoute.DELETE("/:id", func(c *gin.Context) { deleteUser(c, userRepository) })
		}

		adminSubRoute := authedSubRoute.Group("/admin/")
		{

			statsSubRoute := adminSubRoute.Group("/stats/")
			{
				statsSubRoute.GET("/user-count", func(c *gin.Context) { userCountView(c, userRepository) })
				statsSubRoute.GET("/step-count", func(c *gin.Context) { stepCountView(c, stepRepository) })
			}
		}
	}

	listenPort := GetEnvOrDefault("PORT", "1357")
	fmt.Print(`
 ▄▄▄▄▄▄  ▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄
█      ██      █       █       █
█  ▄    █  ▄   █    ▄▄▄█▄     ▄█
█ █ █   █ █▄█  █   █▄▄▄  █   █  
█ █▄█   █      █    ▄▄▄█ █   █  
█       █  ▄   █   █     █   █  
█▄▄▄▄▄▄██▄█ █▄▄█▄▄▄█     █▄▄▄█`)

	fmt.Print("\n\nHosted at \"http://localhost:" + listenPort + "\"\n")
	router.Run(":" + listenPort)
}
