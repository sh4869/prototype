package main

import (
	"context"
	"log"
	"os"
	"server/evaluation"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func getClient() *github.Client {
	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	client := getClient()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3030"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/score/:name/:repo/ci", func(c *gin.Context) {
		name := c.Param("name")
		repo := c.Param("repo")
		sha := c.Query("sha")
		result, err := evaluation.GetCIScore(client, name, repo, sha)
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, result)
	})

	r.GET("/score/:name/:repo/test", func(c *gin.Context) {
		name := c.Param("name")
		repo := c.Param("repo")
		sha := c.Query("sha")
		result, err := evaluation.GetTestScore(client, name, repo, sha)
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, result)
	})

	r.GET("/score/:name/:repo/code", func(c *gin.Context) {
		name := c.Param("name")
		repo := c.Param("repo")
		sha := c.Query("sha")
		result, err := evaluation.GetCodeScore(client, name, repo, sha)
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, result)
	})

	r.GET("/commit/:name/:repo/commit_point", func(c *gin.Context) {
		name := c.Param("name")
		repo := c.Param("repo")
		result, err := evaluation.GetCommitPoint(client, name, repo)
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, result)
	})
	r.GET("/commit/:name/:repo/commit_status/:count", func(c *gin.Context) {
		name := c.Param("name")
		repo := c.Param("repo")
		count, _ := strconv.Atoi(c.Param("count"))
		result, err := evaluation.GetCommitStatus(client, name, repo, count)
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, result)
	})

	r.Run()
}
