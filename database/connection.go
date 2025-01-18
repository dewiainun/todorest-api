package database
import (
	"fmt"
	"github.com/dewi/todo-app/config"
	"github.com/dewi/todo-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connect() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB.AutoMigrate(&models.ToDo{})
}