package repository

import (
	"context"
	"data-center/internal/models"
	"database/sql"
	"time"
)

type AutoLogin struct {
	Repository
	DB *sql.DB
}

func (al *AutoLogin) New() error {
	db, err := LoadMySQL()
	al.DB = db
	return err
}

func (al *AutoLogin) Close() {
	_ = al.DB.Close()
}
func (al *AutoLogin) Create() (sql.Result, error) {
	return al.DB.Exec("CREATE TABLE `autologin`  (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `url` varchar(255) NOT NULL DEFAULT '',\n  `site` varchar(255) NOT NULL DEFAULT '',\n  `icon` varchar(255) NOT NULL DEFAULT '',\n  `name` varchar(255) NOT NULL DEFAULT '',\n  `username` varchar(255) NOT NULL DEFAULT '',\n  `password` varchar(255) NOT NULL DEFAULT '',\n  PRIMARY KEY (`id`)\n);")
}
func (al *AutoLogin) Drop() (sql.Result, error) {
	return al.DB.Exec("DROP TABLE `autologin`;")
}
func (al *AutoLogin) Insert(url string, site string, icon string, name string, username string, password string) (sql.Result, error) {
	return al.DB.Exec("INSERT INTO `autologin` (`url`, `site`, `icon`, `name`, `username`, `password`) VALUES (?, ?, ?, ?, ?, ?);", url, site, icon, name, username, password)
}
func (al *AutoLogin) SelectById(id int) (autologin models.AutoLoginTable, err error) {
	err = al.DB.QueryRow("SELECT * FROM `autologin` WHERE `id` = ?", id).Scan(&autologin.ID, &autologin.URL, &autologin.Site, &autologin.Icon, &autologin.Name, &autologin.Username, &autologin.Password)
	return autologin, err
}
func (al *AutoLogin) SelectAll() (autoLogins []models.AutoLoginTable, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := al.DB.QueryContext(ctx, "SELECT * FROM `autologin`")
	if err != nil {
		return autoLogins, err
	}
	defer Close(rows)
	for rows.Next() {
		var autologin models.AutoLoginTable
		if err := rows.Scan(&autologin.ID, &autologin.URL, &autologin.Site, &autologin.Icon, &autologin.Name, &autologin.Username, &autologin.Password); err != nil {
			return autoLogins, err
		}
		autoLogins = append(autoLogins, autologin)
	}
	return autoLogins, nil
}
func (al *AutoLogin) UpdateById(id int, url string, site string, icon string, name string, username string, password string) (sql.Result, error) {
	return al.DB.Exec("UPDATE `autologin` SET `url` = ?, `site` = ?, `icon` = ?, `name` = ?, `username` = ?, `password` = ? WHERE `id` = ?", url, site, icon, name, username, password, id)
}
func (al *AutoLogin) DeleteById(id int) (sql.Result, error) {
	return al.DB.Exec("DELETE FROM `autologin` WHERE `id` = ?", id)
}
