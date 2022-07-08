package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jakeroggenbuck/BestNextStep/daft/step"
	"github.com/jakeroggenbuck/BestNextStep/daft/user"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
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

func getLocalIP() string {
	ip := os.Getenv("LOCAL_IP")
	if ip == "" {
		fmt.Printf("LOCAL_IP not set")
		log.Fatal("LOCAL_IP not set")
	}

	return ip
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
		Desc:  "The first step.",
		Left:  -1,
		Right: 2,
		Owner: 1,
	}
	stepTwo := step.Step{
		Name:  "Step Two",
		Desc:  "The second step.",
		Left:  1,
		Right: -1,
		Owner: 1,
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

func dbExists() bool {
	if _, err := os.Stat("./sqlite.db"); err == nil {
		return true
	}
	return false
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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

	router.GET("/", homePage)

	authAccount := getLogIn()
	authedSubRoute := router.Group("/api/v1/", gin.BasicAuth(authAccount))

	authedSubRoute.GET("/", apiRootPage)

	authedSubRoute.GET("/all", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprint(stepRepository.All()))
	})

	authedSubRoute.POST("/new-user", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")

		if name != "" && password != "" {
			hash, _ := HashPassword(password)
			userRepository.Create(user.User{Name: name, PasswordHash: hash})

			c.String(http.StatusOK, fmt.Sprint(userRepository.All()))
		} else {
			c.String(http.StatusNotAcceptable, "name or password empty"+name+password)
		}

	})

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "1357"
	}

	fmt.Print("\nHosted at http://localhost:" + listenPort + "\n")
	router.Run(":" + listenPort)
}
