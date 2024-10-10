package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App  App
		Db   Db
		Grpc Grpc
	}

	App struct {
		Name string
		Url  string
		Port string
	}

	Db struct {
		UrlCartera  string
		UrlSecurity string
		UrlConta    string
	}

	Grpc struct {
		AuthUrl        string
		AuthPort       string
		PromotionUrl   string
		PromotionPort  string
		AccountUrl     string
		AccountPort    string
		PaymentUrl     string
		PaymentPort    string
		BackofficeUrl  string
		BackofficePort string
		AccountingUrl  string
		AccountingPort string
	}
)

func LoadConfig(path string, name string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	return Config{
		App: App{
			Name: os.Getenv("APP_" + name),
			Url:  os.Getenv("APP_URL_" + name),
			Port: os.Getenv("APP_PORT_" + name),
		},
		Db: Db{
			UrlCartera:  os.Getenv("DB_URL_CARTERA"),
			UrlSecurity: os.Getenv("DB_URL_SECURITY"),
			UrlConta:    os.Getenv("DB_URL_CONTABILIDAD"),
		},
		Grpc: Grpc{
			AuthUrl:        os.Getenv("GRPC_AUTH_URL"),
			AuthPort:       os.Getenv("GRPC_AUTH_PORT"),
			PromotionUrl:   os.Getenv("GRPC_PROMOTION_URL"),
			PromotionPort:  os.Getenv("GRPC_PROMOTION_PORT"),
			AccountUrl:     os.Getenv("GRPC_ACCOUNT_URL"),
			AccountPort:    os.Getenv("GRPC_ACCOUNT_PORT"),
			PaymentUrl:     os.Getenv("GRPC_PAYMENT_URL"),
			PaymentPort:    os.Getenv("GRPC_PAYMENT_PORT"),
			BackofficeUrl:  os.Getenv("GRPC_BACKOFFICE_URL"),
			BackofficePort: os.Getenv("GRPC_BACKOFFICE_PORT"),
			AccountingUrl:  os.Getenv("GRPC_CONTABILIDAD_URL"),
			AccountingPort: os.Getenv("GRPC_CONTABILIDAD_PORT"),
		},
	}
}

func (d *Config) GetURL(name string) string {
	switch name {
	case "security":
		return d.Db.UrlSecurity
	case "cartera":
		return d.Db.UrlCartera
	case "contabilidad":
		return d.Db.UrlConta
	default:
		return ""
	}

}
