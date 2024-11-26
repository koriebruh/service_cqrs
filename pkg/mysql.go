package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"koriebruh/cqrs/config"
	"log/slog"
)

func MysqlClient(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cnf.DataBase.User,
		cnf.DataBase.Pass,
		cnf.DataBase.Host,
		cnf.DataBase.Port,
		cnf.DataBase.Name, // <-- THIS NAME DB
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		slog.Error("failed make connection to database", err)
	}

	// AUTO MIGRATE
	if err = db.AutoMigrate(
	// INSERT HERE IF U WANT AUTO MIGRATE
	//&domain.User{},
	); err != nil {
		slog.Error("failed auto migrate ", err)
	}

	slog.Info("success migrate")
	return db

}
