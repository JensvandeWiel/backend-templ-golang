package main

import (
	"backend-templ-golang/app"
	"backend-templ-golang/services"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viperInit()
	Entrypoint()

}

func viperInit() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func Entrypoint() {
	app := app.NewApp()
	setupRoutes(app)
	services.NewMigrationService(app).Migrate()
	app.Run()

}
