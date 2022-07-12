package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jakeroggenbuck/BestNextStep/daft/col"
	"github.com/jakeroggenbuck/BestNextStep/daft/step"
	"net/http"
)

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "HomePage", nil)
}

func apiRootPage(c *gin.Context) {
	c.HTML(http.StatusOK, "ApiRootPage", nil)
}

func allStep(c *gin.Context, repo *step.SQLiteRepository) {
	owner := int64(1)

	all, err := repo.GetByOwner(owner)
	if err != nil {
		fmt.Print(err)
	}

	all_json, err := json.Marshal(all)
	if err != nil {
		fmt.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": string(all_json),
	})
}

func addStep(c *gin.Context, db *sql.DB)    {}
func updateStep(c *gin.Context, db *sql.DB) {}
func deleteStep(c *gin.Context, db *sql.DB) {}

func allCol(c *gin.Context, repo *col.SQLiteRepository) {
	owner := int64(1)

	all, err := repo.GetByOwner(owner)
	if err != nil {
		fmt.Print(err)
	}

	all_json, err := json.Marshal(all)
	if err != nil {
		fmt.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": string(all_json),
	})
}

func addCol(c *gin.Context, db *sql.DB)    {}
func updateCol(c *gin.Context, db *sql.DB) {}
func deleteCol(c *gin.Context, db *sql.DB) {}
