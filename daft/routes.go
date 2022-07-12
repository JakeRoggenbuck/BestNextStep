package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jakeroggenbuck/BestNextStep/daft/col"
	"github.com/jakeroggenbuck/BestNextStep/daft/step"
	"net/http"
	"strconv"
)

func getUserId(c *gin.Context) int64 {
	return int64(1)
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "HomePage", nil)
}

func apiRootPage(c *gin.Context) {
	c.HTML(http.StatusOK, "ApiRootPage", nil)
}

func allStep(c *gin.Context, repo *step.SQLiteRepository) {
	owner := getUserId(c)

	all, err := repo.GetByOwner(owner)
	if err != nil {
		fmt.Print(err)
	}

	all_json, err := json.Marshal(all)
	if err != nil {
		fmt.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(all_json),
	})
}

func addStep(c *gin.Context, repo *step.SQLiteRepository) {
	owner := getUserId(c)

	collection := c.PostForm("collection")
	fmt.Println(collection)
	col, err := strconv.Atoi(collection)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Given Collection not found.",
		})
		return
	}

	stepToAdd := step.Step{
		Name:       c.PostForm("name"),
		Desc:       c.PostForm("desc"),
		Collection: int64(col),
		Owner:      owner,
	}

	_, err = repo.Create(stepToAdd)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not add step.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Added successfully.",
	})
}

func updateStep(c *gin.Context, db *sql.DB) {}

func deleteStep(c *gin.Context, repo *step.SQLiteRepository) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Given ID not found.",
		})
		return
	}

	err = repo.Delete(int64(i))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Given ID not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted " + id,
	})
}

func allCol(c *gin.Context, repo *col.SQLiteRepository) {
	owner := getUserId(c)

	all, err := repo.GetByOwner(owner)
	if err != nil {
		fmt.Print(err)
	}

	all_json, err := json.Marshal(all)
	if err != nil {
		fmt.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(all_json),
	})
}

func addCol(c *gin.Context, repo *col.SQLiteRepository) {
	owner := getUserId(c)

	colToAdd := col.Col{
		Name:  c.PostForm("name"),
		Desc:  c.PostForm("desc"),
		Owner: owner,
	}

	repo.Create(colToAdd)
}

func updateCol(c *gin.Context, db *sql.DB) {}

func deleteCol(c *gin.Context, repo *col.SQLiteRepository) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Given ID not found.",
		})
		return
	}

	err = repo.Delete(int64(i))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Given ID not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted " + id,
	})
}
