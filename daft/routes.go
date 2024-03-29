package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jakeroggenbuck/BestNextStep/daft/col"
	"github.com/jakeroggenbuck/BestNextStep/daft/step"
	"github.com/jakeroggenbuck/BestNextStep/daft/user"
	"net/http"
	"strconv"
)

func getUserId(c *gin.Context) int64 {
	return int64(1)
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "HomePage", nil)
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

func preflight(c *gin.Context) {
	fmt.Println("preflight hit!!!!")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, method")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.JSON(http.StatusOK, struct{}{})
}

func addStep(c *gin.Context, repo *step.SQLiteRepository) {
	owner := getUserId(c)

	collection := c.PostForm("collection")

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

func updateStep(c *gin.Context, repo *step.SQLiteRepository) {
	owner := getUserId(c)

	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Given ID not found.",
		})
		return
	}

	collection := c.PostForm("collection")
	fmt.Println(collection)
	col, err := strconv.Atoi(collection)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Given Collection not found.",
		})
		return
	}

	stepToUpdate := step.Step{
		Name:       c.PostForm("name"),
		Desc:       c.PostForm("desc"),
		Collection: int64(col),
		Owner:      owner,
	}

	_, err = repo.Update(int64(i), stepToUpdate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not update.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully.",
	})
}

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
		"message": "Deleted successfully.",
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

	_, err := repo.Create(colToAdd)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not create.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Created successfully.",
	})
}

func updateCol(c *gin.Context, repo *col.SQLiteRepository) {
	owner := getUserId(c)

	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Given ID not found.",
		})
		return
	}

	colToUpdate := col.Col{
		Name:  c.PostForm("name"),
		Desc:  c.PostForm("desc"),
		Owner: owner,
	}

	_, err = repo.Update(int64(i), colToUpdate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not update.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully.",
	})
}

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
		"message": "Deleted successfully.",
	})
}

func addUser(c *gin.Context, repo *user.SQLiteRepository) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	if name == "" || password == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Name or Password empty.",
		})
	}

	hash, _ := HashPassword(password)

	userToAdd := user.User{
		Name:         name,
		PasswordHash: hash,
	}

	_, err := repo.Create(userToAdd)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not create.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Created successfully.",
	})
}

func updateUser(c *gin.Context, repo *user.SQLiteRepository) {
	owner := getUserId(c)

	name := c.PostForm("name")
	password := c.PostForm("password")

	if name == "" || password == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Name or Password empty.",
		})
	}

	hash, _ := HashPassword(password)

	userToUpdate := user.User{
		Name:         name,
		PasswordHash: hash,
	}

	_, err := repo.Update(int64(owner), userToUpdate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could not update.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully.",
	})
}

func deleteUser(c *gin.Context, repo *user.SQLiteRepository) {
	owner := getUserId(c)
	err := repo.Delete(int64(owner))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Given ID not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully.",
	})
}

func stepCountView(c *gin.Context, repo *step.SQLiteRepository) {
	c.JSON(http.StatusOK, gin.H{
		"message": stepCount(repo),
	})
}

func userCountView(c *gin.Context, repo *user.SQLiteRepository) {
	c.JSON(http.StatusOK, gin.H{
		"message": userCount(repo),
	})
}
