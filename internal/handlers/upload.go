package handlers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const bufferSize = 64 * 1024 // 64 KB

func UploadCheck(c *gin.Context) {

	defer c.Request.Body.Close()

	buffer := make([]byte, bufferSize)

	var totalBytes int64

	for {

		n, err := c.Request.Body.Read(buffer)

		if n > 0 {
			totalBytes += int64(n)
		}

		if err == io.EOF {
			break
		}

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"bytes": totalBytes,
	})
}
