package database

import (
	"github.com/EdsonJunio/api_rest_with_go_gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaCoimBancoDeDados() {
	stringDeConexao := "host=127.0.0.1 user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
