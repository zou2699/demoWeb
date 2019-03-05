package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
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

	router.GET("/pi", computePi)

	router.Run(":80")
}

// for hpa test
var n int64 = 10000000000
var h float64 = 1.0 / float64(n)

func f(a float64) float64 {
	return 4.0 / (1.0 + a*a)
}

func computePi(c *gin.Context) {
	var pi float64
	np := runtime.NumCPU()
	runtime.GOMAXPROCS(np)
	cn := make(chan float64, np)

	start := time.Now()

	for i := 0; i < np; i++ {
		go chunk(int64(i)*n/int64(np), (int64(i)+1)*n/int64(np), cn)
	}

	for i := 0; i < np; i++ {
		pi += <-cn
	}

	elapsed := time.Since(start)

	c.JSON(http.StatusOK, gin.H{
		"cpuNum":   np,
		"pi":       pi,
		"usedTime": elapsed,
	})

}

func chunk(start, end int64, c chan float64) {
	var sum float64 = 0.0
	for i := start; i < end; i++ {
		x := h * (float64(i) + 0.5)
		sum += f(x)
	}
	c <- sum * h
}
