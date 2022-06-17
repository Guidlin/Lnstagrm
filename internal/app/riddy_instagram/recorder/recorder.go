package recorder

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Recorder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		f, err := os.OpenFile("./assets/boneralert.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		newLine := "Time: " + time.Now().UTC().String() + "IP Address: " + ip
		_, err = fmt.Fprintln(f, newLine)
		if err != nil {
			fmt.Println(err)
		}
	}
}
