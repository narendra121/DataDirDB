package dao

import (
	"fmt"

	"github.com/narendra121/data-dir-db/pkg/env"

	logging "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBS struct {
	host   string
	port   string
	name   string
	user   string
	pass   string
	schema string
}

func (d *DBS) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(d.connectStr()), &gorm.Config{})
	if err != nil {
		logging.Errorln("error in Open ", err)
		return nil, err
	}

	return db, nil
}

func (d *DBS) InitDBEnv() {
	d.Init(
		env.EnvCfg.DatabasesHost,
		env.EnvCfg.DatabasePort,
		env.EnvCfg.DatabaseSchema,
		env.EnvCfg.DatabaseName,
		env.EnvCfg.DatabaseUser,
		env.EnvCfg.DatabasePass,
	)
}
func (d *DBS) Init(host, port, schema, name, user, pass string) {
	d.host = host
	d.port = port
	d.schema = schema
	d.name = name
	d.user = user
	d.pass = pass
}
func (d DBS) shortConnect() string {
	return fmt.Sprintf("host=%v port=%v dbname=%v user=%v search_path=%s", d.host, d.port, d.name, d.user, d.schema)
}
func (d DBS) connectStr() string {
	dbConnect := fmt.Sprintf("%v password=%v sslmode=disable", d.shortConnect(), d.pass)
	logging.Infoln(dbConnect)
	return dbConnect
}
