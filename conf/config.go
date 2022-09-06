package conf

import "gorm.io/gorm"
import "gorm.io/driver/mysql"

type Config struct {
	DbDsn string
}

type ServiceContext struct {
	Db *gorm.DB
}

var SvcCtx *ServiceContext

func InitServiceContext(config Config) {
	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	// Find the last '/' (since the password or the net addr might contain a '/')
	db, err := gorm.Open(mysql.Open(config.DbDsn))
	if err != nil {
		panic(err)
	}
	SvcCtx = &ServiceContext{
		Db: db,
	}
}
