package repository

import (
	"context"
	"data-center/internal/models"
	"database/sql"
	"time"
)

type Settings struct {
	DB *sql.DB
}

func NewSettings() (*Settings, error) {
	db, err := LoadMySQL()
	return &Settings{DB: db}, err
}
func (s Settings) Create() (result sql.Result, err error) {
	return s.DB.Exec("CREATE TABLE `settings`  (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `key` varchar(255) NOT NULL DEFAULT '',\n  `name` varchar(255) NOT NULL DEFAULT '',\n  `value` longtext NULL,\n  PRIMARY KEY (`id`)\n);")
}
func (s Settings) Insert(key string, name string, value string) (result sql.Result, err error) {
	return s.DB.Exec("INSERT INTO `settings` (`key`, `name`, `value`) VALUES (?, ?, ?);", key, name, value)
}
func (s Settings) SelectById(id int) (row models.SettingsTable, err error) {
	err = s.DB.QueryRow("SELECT * FROM `settings` WHERE `id` = ?", id).Scan(&row.ID, &row.Key, &row.Name, &row.Value)
	return row, err
}
func (s Settings) SelectByKey(key string) (row models.SettingsTable, err error) {
	err = s.DB.QueryRow("SELECT * FROM `settings` WHERE `key` = ?", key).Scan(&row.ID, &row.Key, &row.Name, &row.Value)
	return row, err
}
func (s Settings) SelectAll() (rows []models.SettingsTable, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row, err := s.DB.QueryContext(ctx, "SELECT * FROM `settings`")
	defer Close(row)
	if err != nil {
		return rows, err
	}
	for row.Next() {
		var item models.SettingsTable
		if err := row.Scan(&item.ID, &item.Key, &item.Name, &item.Value); err != nil {
			return rows, err
		}
		rows = append(rows, item)
	}
	if rows == nil {
		rows = []models.SettingsTable{}
	}
	return rows, nil
}
func (s Settings) UpdateById(id int, key string, name string, value string) (result sql.Result, err error) {
	return s.DB.Exec("UPDATE `settings` SET `key` = ?, `name` = ?, `value` = ? WHERE `id` = ?", key, name, value, id)
}
func (s Settings) DeleteById(id int) (result sql.Result, err error) {
	return s.DB.Exec("DELETE FROM `settings` WHERE `id` = ?", id)
}
func (s Settings) Drop() (result sql.Result, err error) {
	return s.DB.Exec("DROP TABLE `settings`")
}
