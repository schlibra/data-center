package repository

import (
	"context"
	"data-center/internal/models"
	"database/sql"
	"time"
)

type Permission struct {
	DB *sql.DB
}

func NewPermission() (*Permission, error) {
	db, err := LoadMySQL()
	return &Permission{DB: db}, err
}
func (p *Permission) Create() (sql.Result, error) {
	return p.DB.Exec("CREATE TABLE IF NOT EXISTS `permission`  (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `key` varchar(255) NOT NULL,\n  `name` varchar(255) NOT NULL,\n  `parent` int NOT NULL,\n  PRIMARY KEY (`id`)\n);")
}
func (p *Permission) Drop() (sql.Result, error) {
	return p.DB.Exec("DROP TABLE IF EXISTS `permission`")
}
func (p *Permission) Insert(key, name string, parent int) (sql.Result, error) {
	return p.DB.Exec("INSERT INTO `permission` (`key`, `name`, `parent`) VALUES (?, ?, ?)", key, name, parent)
}
func (p *Permission) UpdateByID(id int, key string, name string, parent int) (sql.Result, error) {
	return p.DB.Exec("UPDATE `permission` SET `key` = ?, `name` = ?, `parent` = ? WHERE `id` = ?", key, name, parent, id)
}
func (p *Permission) SelectById(id int) (permission models.PermissionTable, err error) {
	err = p.DB.QueryRow("SELECT * FROM `permission` WHERE `id` = ?", id).Scan(&permission.ID, &permission.Key, &permission.Name, &permission.Parent)
	return permission, err
}
func (p *Permission) SelectByKey(key string) (permission models.PermissionTable, err error) {
	err = p.DB.QueryRow("SELECT * FROM `permission` WHERE `key` = ?", key).Scan(&permission.ID, &permission.Key, &permission.Name, &permission.Parent)
	return permission, err
}
func (p *Permission) SelectAll() (permissions []models.PermissionTable, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := p.DB.QueryContext(ctx, "SELECT * FROM `permission`")
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var row models.PermissionTable
		if err := rows.Scan(&row.ID, &row.Key, &row.Name, &row.Parent); err != nil {
			return nil, err
		}
		permissions = append(permissions, row)
	}
	if permissions == nil {
		permissions = []models.PermissionTable{}
	}
	return permissions, nil
}
func (p *Permission) DeleteById(id int) (sql.Result, error) {
	return p.DB.Exec("DELETE FROM `permission` WHERE `id` = ?", id)
}
