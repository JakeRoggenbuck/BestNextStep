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
	"os"
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

	router.Use(cors.Default())

	/*

		GET => 		/ 						homePage
		GET => 		/version 				version


		AUTHED USER		Group	/api/v1
		===========		=====	=======

		GET => 		/api/v1/				apiRootPage

		GET => 		/api/v1/step			allStep
		POST => 	/api/v1/step/add		addStep
		PUT => 		/api/v1/step/update		updateStep
		DELETE => 	/api/v1/step/delete		deleteStep

		GET => 		/api/v1/col				allCol
		POST => 	/api/v1/col/add			addCol
		PUT => 		/api/v1/col/update		updateCol
		DELETE => 	/api/v1/col/delete		deleteCol

		GET => 		/api/v1/user			allUser
		POST => 	/api/v1/user/add		addUser
		PUT => 		/api/v1/user/update		updateUser
		DELETE => 	/api/v1/user/delete		deleteUser

		ADMIN USER		Group	/api/v1/admin
		==========		=====	=============

		GET =>		/api/v1/admin/stats				allStats
		GET =>		/api/v1/admin/stats/user-count	userCount
		GET =>		/api/v1/admin/stats/step-count	stepCount

	*/

	router.GET("/", homePage)

	// New auth for normal users in userRepository
	// https://github.com/yasaricli/gah
	// https://chenyitian.gitbooks.io/gin-tutorials/content/tdd/8.html
	authAccount := getLogIn()
	authedSubRoute := router.Group("/api/v1/", gin.BasicAuth(authAccount))

	authedSubRoute.GET("/", apiRootPage)

	stepSubRoute := authedSubRoute.Group("/step/")
	stepSubRoute.GET("/", func(c *gin.Context) { allStep(c, stepRepository) })

	// stepSubRoute.POST("/add", func(c *gin.Context) { addStep(c, stepRepository) })
	// stepSubRoute.PUT("/update", func(c *gin.Context) { updateStep(c, stepRepository) })
	// stepSubRoute.DELETE("/delete", func(c *gin.Context) { deleteStep(c, stepRepository) })

	colSubRoute := authedSubRoute.Group("/col/")
	colSubRoute.GET("/", func(c *gin.Context) { allCol(c, colRepository) })
	// colSubRoute.POST("/add", func(c *gin.Context) { addCol(c, colRepository) })
	// colSubRoute.PUT("/update", func(c *gin.Context) { updateCol(c, colRepository) })
	// colSubRoute.DELETE("/delete", func(c *gin.Context) { deleteCol(c, colRepository) })

	// authedSubRoute.POST("/new-user", func(c *gin.Context) {
	// 	name := c.PostForm("name")
	// 	password := c.PostForm("password")

	// 	if name != "" && password != "" {
	// 		hash, _ := HashPassword(password)
	// 		userRepository.Create(user.User{Name: name, PasswordHash: hash})

	// 		c.String(http.StatusOK, fmt.Sprint(userRepository.All()))
	// 	} else {
	// 		c.String(http.StatusNotAcceptable, "name or password empty")
	// 	}

	// })

	listenPort := GetEnvOrDefault("PORT", "1357")
	fmt.Print("\nHosted at http://localhost:" + listenPort + "\n")
	router.Run(":" + listenPort)
}
