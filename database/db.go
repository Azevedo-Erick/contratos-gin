package database

import (
	"log"
	"os"

	"test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDB() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=root password=123 dbname=test port=5432",
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Conexão inválida\n", err.Error())
		os.Exit(2)
	}

	log.Println("Conectado com sucesso")
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.Usuario{})

	Database = DbInstance{Db: db}

}
