package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type RESTServer struct {
	game *Game
}

func NewRESTServer(game *Game) *RESTServer {
	return &RESTServer{game: game}
}

func (s *RESTServer) Start() error {
	gin.SetMode(gin.ReleaseMode)

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = f

	router := gin.Default()
	defer f.Close()

	// very stupid move handler
	router.GET("/player1/move/up", s._move)
	router.GET("/player1/move/down", s._move)
	router.GET("/player2/move/up", s._move)
	router.GET("/player2/move/down", s._move)

	go router.Run(":8887")
	return nil
}

func (s *RESTServer) _move(c *gin.Context) {
	req := c.Request
	url := req.URL.String()
	racket := s.game.lracket

	if strings.Contains(url, "player2") {
		racket = s.game.rracket
	}

	if strings.Contains(url, "down") {
		racket.Move(s.game.ticker, 1)
	} else {
		racket.Move(s.game.ticker, -1)
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
