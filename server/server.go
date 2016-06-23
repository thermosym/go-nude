// main.go
package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thermosym/go-nude"
)

type CheckResponse struct {
	IsNude  bool   `json:"is_nude"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func OkResponse(is_nude bool) *CheckResponse {
	f := new(CheckResponse)
	f.IsNude = is_nude
	f.Success = true
	f.Message = "ok"
	return f
}

func ErrResponse(reason string) *CheckResponse {
	f := new(CheckResponse)
	f.IsNude = false
	f.Success = false
	f.Message = reason
	return f
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	// parse the program input args
	default_usage_info := "Usage: $./server [-p port_number]"
	var port = 8080     // default port number
	args := os.Args[1:] // args without program
	if len(args) <= 0 {
		fmt.Printf("Default binds to http port to %d!\n", port)
		fmt.Printf("If you want to manually setup the port, please refer %s \n\n", default_usage_info)
	} else if len(args) == 2 && args[0] == "-p" {
		config_port, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("port number is illegal")
			fmt.Println(default_usage_info)
			log.Fatal(err)
		}
		port = config_port
		fmt.Printf("Configure binds to http port to %d!\n", port)

	} else {
		fmt.Println(default_usage_info)
		log.Fatal("Cannot parse the input args")
	}

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.POST("/check", func(c *gin.Context) {

		file, header, err := c.Request.FormFile("file")
		filename := header.Filename
		fmt.Println(header.Filename)
		log.Println("received file:" + filename)

		img, _, err := image.Decode(file)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, ErrResponse("fails to decode image"))
			return
		}

		flag, err := nude.IsImageNude(img)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, ErrResponse("fails to detect the image"))
		} else {
			c.JSON(http.StatusOK, OkResponse(flag))
		}

	})

	router.Run(":" + strconv.Itoa(port))
}
