gopackage database

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"

	"github.com/kartmatias/cdwk-pos-agent/dao/model"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	password string
	port     int
	server   string
	user     string
	dbname   string
)

var Database *gorm.DB

var connString string

func Open(logger *zap.Logger) {
	var err error

	if Database != nil {
		sqlDB, err := Database.DB()
		// Ping
		err = sqlDB.Ping()
		if err != nil {
			logger.Info("Database connection is not active.", zap.Error(err))
		} else {
			logger.Info("Database connection is already active.")
			return
		}
	}

	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	server = os.Getenv("DB_HOST")
	dbname = os.Getenv("DB_NAME")

	connString = fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, server, port, dbname)
	Database, err = gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		logger.Error("Error opening database", zap.Error(err))
	}

	sqlDB, _ := Database.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(20)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = Database.AutoMigrate(
		&model.IntegracaoProduto{},
		&model.IntegracaoGrupo{},
		&model.IntegracaoVariacao{},
		&model.IntegracaoAtributo{},
		&model.IntegracaoPedido{},
	)

	if err != nil {
		logger.Error("Error when running migrations", zap.Error(err))
	}

}
