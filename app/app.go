package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"os"
)

type App struct {
	Server *echo.Echo
	Db     *gorm.DB
	Logger *Logger
}

func NewApp() *App {
	a := &App{}
	a.setupLogger()
	a.setupServer()
	a.setupDatabase()
	return a
}

func (a *App) Run() {
	err := a.Server.Start(viper.GetString("Server.host") + ":" + viper.GetString("Server.port"))
	if err != nil {
		a.Logger.Fatal(err)
	}
}

func (a *App) RunNonBlocking() {
	go func() {
		err := a.Server.Start(viper.GetString("Server.host") + ":" + viper.GetString("Server.port"))
		if err != nil {
			a.Logger.Fatal(err)
		}
	}()
}

func (a *App) setupLogger() {
	a.Logger = NewLogger(Debug, log.New(os.Stdout, "[APP] ", log.LstdFlags))
}

func (a *App) setupDatabase() {
	conf := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		IgnoreRelationshipsWhenMigrating:         true,
		Logger:                                   gormlogger.New(log.New(a.Logger.Writer(), "[DATABASE] ", log.LstdFlags), gormlogger.Config{LogLevel: 4, Colorful: false}),
	}

	// Make connection
	db, err := gorm.Open(*createConnection(false), &conf)
	if err != nil {
		a.Server.Logger.Fatal(err)
	}

	// Create database if not exists
	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + viper.GetString("database.dbname") + ";")

	a.Db, err = gorm.Open(*createConnection(true), &conf)
	if err != nil {
		a.Logger.Fatal(err)
	}
}

func (a *App) setupServer() {
	a.Server = echo.New()
	a.Server.Use(middleware.Logger())
	a.Server.Use(middleware.Recover())
	a.Server.Logger.SetPrefix("[SERVER] ")
	a.Server.Logger.SetOutput(a.Logger.Writer())
	a.Server.Logger.SetLevel(1)
	a.Server.Logger.SetHeader("${prefix} ${time_rfc3339_nano} ${short_file}:${line} ${message}")
	a.setupViewRenderer()
}

func createDSN(withDb bool) string {
	if withDb {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.dbname"))
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local", viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.host"), viper.GetString("database.port"))
}

func createConnection(withDb bool) *gorm.Dialector {
	var conn gorm.Dialector
	switch viper.GetString("database.driver") {
	case "mysql":
		conn = mysql.Open(createDSN(withDb))
	default:
		panic("Database driver not supported: " + viper.GetString("database.driver"))
	}
	return &conn
}
