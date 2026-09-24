package repository

import (
	"context"
	"data-center/internal/models"
	"database/sql"
	"time"
)

type Group struct {
	DB *sql.DB
}

func NewGroup() (*Group, error) {
	db, err := LoadMySQL()
	return &Group{DB: db}, err
}
func (g *Group) Create() (sql.Result, error) {
	return g.DB.Exec("CREATE TABLE IF NOT EXISTS `group`  (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `name` varchar(255) NOT NULL,\n  `admin` int NOT NULL,\n  `permission` longtext NOT NULL,\n  PRIMARY KEY (`id`)\n);")
}
func (g *Group) Insert(name string, admin int, permission string) (sql.Result, error) {
	return g.DB.Exec("INSERT INTO `group` (`name`, `admin`, `permission`) VALUES (?, ?, ?)", name, admin, permission)
}
func (g *Group) SelectById(id int) (group models.GroupTable, err error) {
	err = g.DB.QueryRow("SELECT * FROM `group` WHERE `id` = ?", id).Scan(&group.ID, &group.Name, &group.Admin, &group.Permission)
	return group, err
}
func (g *Group) SelectByName(name string) (group models.GroupTable, err error) {
	err = g.DB.QueryRow("SELECT * FROM `group` WHERE `name` = ?", name).Scan(&group.ID, &group.Name, &group.Admin, &group.Permission)
	return group, err
}
func (g *Group) SelectAll() (groups []models.GroupTable, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := g.DB.QueryContext(ctx, "SELECT * FROM `group`")
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var row models.GroupTable
		if err := rows.Scan(&row.ID, &row.Name, &row.Admin, &row.Permission); err != nil {
			return nil, err
		}
		groups = append(groups, row)
	}
	if groups == nil {
		groups = []models.GroupTable{}
	}
	return groups, nil
}
func (g *Group) UpdateById(id int, name string, admin int, permission string) (sql.Result, error) {
	return g.DB.Exec("UPDATE `group` SET `name` = ?, `admin` = ?, `permission` = ? WHERE `id` = ?", name, admin, permission, id)
}
func (g *Group) DeleteById(id int) (sql.Result, error) {
	return g.DB.Exec("DELETE FROM `group` WHERE `id` = ?", id)
}
func (g *Group) Drop() (sql.Result, error) {
	return g.DB.Exec("DROP TABLE IF EXISTS `group`")
}
