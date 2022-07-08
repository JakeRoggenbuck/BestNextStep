package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jakeroggenbuck/BestNextStep/daft/step"
	_ "github.com/mattn/go-sqlite3"
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

func createDefaultElements(db *sql.DB) {
	stepRepository := step.NewSQLiteRepository(db)

	if err := stepRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	stepOne := step.Step{
		Name:  "Step One",
		Left:  -1,
		Right: 2,
	}
	stepTwo := step.Step{
		Name:  "Step Two",
		Left:  1,
		Right: -1,
	}

	createdStepOne, err := stepRepository.Create(stepOne)
	if err != nil {
		fmt.Println(err)
	}

	createdStepTwo, err := stepRepository.Create(stepTwo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(createdStepOne)
	fmt.Println(createdStepTwo)
}

func main() {
	setupLogging()

	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal("Database open failed")
	}

	createDefaultElements(db)

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
