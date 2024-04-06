package main

import (
	"flag"
	"os"
	"path/filepath"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4"
	"github.com/narendra121/data-dir-db/internal/controllers"
	"github.com/narendra121/data-dir-db/internal/svc/filedata"
	"github.com/narendra121/data-dir-db/pkg/dao"
	"github.com/narendra121/data-dir-db/pkg/env"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func main() {
	appEnv := flag.String("env", "", "")
	flag.Parse()
	env.InitEnv(appEnv)
	tmpDir := os.TempDir()

	// Create a directory named "example" in the temporary directory
	fileNewDir := filepath.Join(tmpDir, env.EnvCfg.DataDir)

	if err := os.MkdirAll(fileNewDir, 0755); err != nil {
		logging.Fatalf("Failed to create data directory: %v", err)
	}
	filedata.FileNewDir = fileNewDir

	for i := 0; i < env.EnvCfg.WorkerCount; i++ {
		go filedata.ProcessSavedData()
	}

	d := new(dao.DBS)
	d.InitDBEnv()

	var err error

	dao.DB, err = d.Connect()
	if err != nil {
		logging.Errorln("error in d.Connect()", err)
		return
	}
	runDBMigration()
	router := gin.Default()

	router.GET("/status", controllers.GetServerStatus)
	router.POST("savedata", controllers.Savedata)
	logging.Info("Hey Guys Server Is Up On Running At Port: ", env.EnvCfg.AppPort)
	router.Run(":" + env.EnvCfg.AppPort)
}

func runDBMigration() {
	migration, err := migrate.New(env.EnvCfg.MigrationUrl, env.EnvCfg.DatabaseSource)
	if err != nil {
		logging.Errorln("cannot create new migrate instance:", err)
		return
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		logging.Errorln("failed to run migrate up:", err)
		return
	}
	logging.Infoln("db migrated successfully")
}
