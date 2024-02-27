package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowon-irc/go-gowon"
)

const (
	moduleName = "checkiday"
)

func main() {
	log.Printf("%s starting\n", moduleName)

	r := gin.Default()

	r.POST("/days/message", func(c *gin.Context) {
		var m gowon.Message

		if err := c.BindJSON(&m); err != nil {
			log.Println("Error: unable to bind message to json", err)
			return
		}

		out, err := checkiday()
		if err != nil {
			log.Println(err)
			m.Msg = "{red}Error when looking up days{clear}"
			c.IndentedJSON(http.StatusInternalServerError, &m)
		}

		m.Msg = out
		c.IndentedJSON(http.StatusOK, &m)
	})

	r.GET("/days/help", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, &gowon.Message{
			Module: moduleName,
			Msg:    "what days is it today?",
		})
	})

	r.POST("/mdays/message", func(c *gin.Context) {
		var m gowon.Message

		if err := c.BindJSON(&m); err != nil {
			log.Println("Error: unable to bind message to json", err)
			return
		}

		out, err := checkmday()
		if err != nil {
			log.Println(err)
			m.Msg = "{red}Error when looking up mdays{clear}"
			c.IndentedJSON(http.StatusInternalServerError, &m)
		}

		m.Msg = out
		c.IndentedJSON(http.StatusOK, &m)
	})

	r.GET("/mdays/help", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, &gowon.Message{
			Module: moduleName,
			Msg:    "what months is it today?",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
