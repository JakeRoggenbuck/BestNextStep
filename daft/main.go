package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func getLogIn() gin.Accounts {
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		fmt.Printf("ADMIN_PASSWORD not set")
		log.Fatal("ADMIN_PASSWORD not set")
	}

	return gin.Accounts{
		"Admin": password,
	}
}

func setupLogging() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

func main() {
	setupLogging()
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.86.245"})
	router.LoadHTMLGlob("./web/templates/**/*")

	router.GET("/", homePage)

	authAccount := getLogIn()
	authedSubRoute := router.Group("/api/v1/", gin.BasicAuth(authAccount))

	authedSubRoute.GET("/", apiRootPage)

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "1357"
	}

	fmt.Print("\nHosted at http://localhost:" + listenPort + "\n")
	router.Run(":" + listenPort)
}
