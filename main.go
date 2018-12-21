package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	//helloHandler := func(w http.ResponseWriter, req *http.Request) {
	//	io.WriteString(w, "Hello, world!\n")
	//}
	//
	//http.HandleFunc("/", helloHandler)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Panic(err)
	//}
	//log.Println("http listen 8080 successful")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		osName, err := os.Hostname()
		if err != nil {
			log.Panic(err)
		}
		c.JSON(200, gin.H{"hostname": osName})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
