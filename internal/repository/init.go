package repository

import (
	"data-center/pkg/utils"
	"database/sql"
	"fmt"
)

func LoadMySQL() (*sql.DB, error) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		return nil, err
	}
	username := cfg.MySQL.Username
	password := cfg.MySQL.Password
	host := cfg.MySQL.Host
	port := cfg.MySQL.Port
	database := cfg.MySQL.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	mysql, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := mysql.Ping(); err != nil {
		return nil, err
	}
	return mysql, nil
}
func Close(row *sql.Rows) {
	_ = row.Close()
}
