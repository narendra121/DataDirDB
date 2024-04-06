package filedata

import (
	"context"
	"os"

	"github.com/narendra121/data-dir-db/pkg/dao"
	"github.com/narendra121/data-dir-db/pkg/env"
	"github.com/narendra121/data-dir-db/pkg/model"
	logging "github.com/sirupsen/logrus"
)

var (
	FileNewDir string
	DataQueue  = make(chan string, env.EnvCfg.QueueSize)
)

func ProcessSavedData() {
	for {

		filename := <-DataQueue
		data, err := os.ReadFile(filename)
		if err != nil {
			logging.Errorln("Failed to read file ", filename, err)
			continue
		}
		if _, _, dbErr := dao.AddDatafiles(context.Background(), &model.Datafiles{Filename: filename, Data: string(data)}); dbErr != nil {
			logging.Errorln("Failed to save fil to db", filename, dbErr)
			continue
		}

		// if err := os.Remove(filename); err != nil {
		// 	logging.Errorln("Failed to remove file ", filename, err)
		// }
	}
}
