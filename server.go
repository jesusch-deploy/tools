package tools

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jesusch-deploy/tools/config"

	"github.com/jesusch-deploy/db"
)

type AppConfig struct {
	Port string
	DB   *sql.DB
}

func StartService(configPath string, name string, routes func(AppConfig) http.Handler) {
	cfg := config.LoadConfig(func() string {
		if configPath == "" {
			log.Fatal("Error: .env Ruta es requerida")
		}
		return configPath
	}(), name)

	dsn := cfg.GetURL(name)
	conn := db.GetInstance(dsn)
	defer conn.Close()

	if conn == nil {
		log.Panic("No es posible conectar a la base de datos")
	}
	app := AppConfig{
		Port: cfg.App.Port,
		DB:   conn.GetConnection(),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Port),
		Handler: routes(app),
	}
	fmt.Println("Iniciando servicio en ->", app.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Panic("error: ", err)
	}
}
