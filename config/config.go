package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfighcl"
	"log"
	"sync"
)

type Config struct {
	DBHost     string `hcl:"database_host" env:"DBHOST" default:"postgres"`
	DBPort     string `hcl:"database_port" env:"DBPORT" default:"5434"`
	DBUser     string `hcl:"database_user" env:"DBUSER" default:"postgres"`
	DBPassword string `hcl:"database_password" env:"DBPASSWD" default:"postgres"`
	DBName     string `hcl:"database_name" env:"DBNAME" default:"transaction_service"`
	SSLMode    string `hcl:"database_sslmode" env:"SSLMODE" default:"disable"`
	APPPort    string `hcl:"app_port" env:"PORT" default:"8080"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		loader := aconfig.LoaderFor(&cfg, aconfig.Config{
			EnvPrefix: "NFB",
			Files:     []string{"./config.hcl", "./config.local.hcl"},
			FileDecoders: map[string]aconfig.FileDecoder{
				".hcl": aconfighcl.New(),
			},
		})

		if err := loader.Load(); err != nil {
			log.Printf("[ERROR] failed to load config: %v", err)
		}

	})

	return cfg
}
