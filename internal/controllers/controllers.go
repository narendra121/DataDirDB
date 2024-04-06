package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	logging "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/narendra121/data-dir-db/internal/svc/filedata"
	"github.com/narendra121/data-dir-db/pkg/env"
)

func GetServerStatus(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, "server is Up :)")
}

func Savedata(c *gin.Context) {
	if len(filedata.DataQueue) >= env.EnvCfg.QueueSize {
		c.AbortWithStatusJSON(http.StatusBadRequest, "No resources to process further files")
		return
	}
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logging.Errorln("error in reading request body ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	filename := fmt.Sprintf("%s/%v.txt", filedata.FileNewDir, uuid.New().String())

	err = os.WriteFile(filename, requestBody, 0644)
	if err != nil {
		logging.Errorln("error in writing to file ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data to file"})
		return
	}
	filedata.DataQueue <- filename
	c.AbortWithStatusJSON(http.StatusCreated, filename+" created")
}
