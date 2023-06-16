package config

import (
	"os"

	"github.com/aiteung/atdb"
)

var IteungIPAddress string = os.Getenv("ITEUNGBEV1")

var MongoString string = os.Getenv("MONGOSTRING")

var MariaStringAkademik string = os.Getenv("MARIASTRINGAKADEMIK")

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "db_dhs_tb",
}

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)
