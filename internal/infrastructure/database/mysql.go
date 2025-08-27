package database

import (
	"fmt"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/infrastructure/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.MySQL.User,
		cfg.Database.MySQL.Password,
		cfg.Database.MySQL.Host,
		cfg.Database.MySQL.Port,
		cfg.Database.MySQL.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// Jangan panggil log.Fatalf di sini, kembalikan error-nya
		return nil, err
	}

	// Kembalikan db dan nil untuk error jika sukses
	return db, nil
}
