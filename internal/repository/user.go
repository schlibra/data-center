package repository

import (
	"context"
	"data-center/internal/models"
	"database/sql"
	"time"
)

type User struct {
	DB *sql.DB
}

func (u *User) New() error {
	db, err := LoadMySQL()
	u.DB = db
	return err
}
func (u *User) Close() {
	_ = u.DB.Close()
}
func (u *User) Create() (result sql.Result, err error) {
	return u.DB.Exec("CREATE TABLE `user`  (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `username` varchar(255) NOT NULL,\n  `password` varchar(255) NOT NULL,\n  `nickname` varchar(255) NOT NULL,\n  `group` int NOT NULL,\n  `enable` int NOT NULL,\n  `token_id` varchar(255) NULL DEFAULT '',\n `api_id` varchar(255) NULL DEFAULT '',\n  PRIMARY KEY (`id`)\n);")
}
func (u *User) Insert(username string, password string, nickname string, group int, enable int) (result sql.Result, err error) {
	return u.DB.Exec("INSERT INTO `user` (`username`, `password`, `nickname`, `group`, `enable`) VALUES (?, ?, ?, ?, ?)", username, password, nickname, group, enable)
}
func (u *User) SelectById(id int) (result models.UserTable, err error) {
	group := Group{DB: u.DB}
	if err := u.DB.QueryRow("SELECT * FROM `user` WHERE `id` = ?", id).Scan(&result.ID, &result.Username, &result.Password, &result.Nickname, &result.Group, &result.Enable, &result.TokenID, &result.ApiId); err != nil {
		return result, err
	}
	groupInfo, err := group.SelectById(result.Group)
	if err != nil {
		return result, err
	}
	result.GroupInfo = groupInfo
	return result, err
}
func (u *User) SelectByUsername(username string) (result models.UserTable, err error) {
	group := Group{DB: u.DB}
	if err := u.DB.QueryRow("SELECT * FROM `user` WHERE `username` = ?", username).Scan(&result.ID, &result.Username, &result.Password, &result.Nickname, &result.Group, &result.Enable, &result.TokenID, &result.ApiId); err != nil {
		return result, err
	}
	groupInfo, err := group.SelectById(result.Group)
	if err != nil {
		return result, err
	}
	result.GroupInfo = groupInfo
	return result, err
}
func (u *User) SelectByGroup(groupId int) (users []models.UserTable, err error) {
	group := Group{DB: u.DB}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := u.DB.QueryContext(ctx, "SELECT * FROM `user` WHERE `group` = ?", groupId)
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var row models.UserTable
		if err := rows.Scan(&row.ID, &row.Username, &row.Password, &row.Nickname, &row.Group, &row.Enable, &row.TokenID, &row.ApiId); err != nil {
			return nil, err
		}
		groupInfo, err := group.SelectById(row.Group)
		if err != nil {
			return users, err
		}
		row.GroupInfo = groupInfo
		users = append(users, row)
	}
	if users == nil {
		users = []models.UserTable{}
	}
	return users, nil
}
func (u *User) SelectAll() (users []models.UserTable, err error) {
	group := Group{DB: u.DB}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := u.DB.QueryContext(ctx, "SELECT * FROM `user`")
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var row models.UserTable
		if err := rows.Scan(&row.ID, &row.Username, &row.Password, &row.Nickname, &row.Group, &row.Enable, &row.TokenID, &row.ApiId); err != nil {
			return nil, err
		}
		groupInfo, err := group.SelectById(row.Group)
		if err != nil {
			return users, err
		}
		row.GroupInfo = groupInfo
		users = append(users, row)
	}
	if users == nil {
		users = []models.UserTable{}
	}
	return users, nil
}
func (u *User) DeleteById(id int) (result sql.Result, err error) {
	return u.DB.Exec("DELETE FROM `user` WHERE `id` = ?", id)
}
func (u *User) UpdateById(id int, password string, nickname string, group int, enable int, tokenId string, apiId string) (result sql.Result, err error) {
	return u.DB.Exec("UPDATE `user` SET `password` = ?, `nickname` = ?, `group` = ?, `enable` = ?, `token_id` = ?, `api_id` = ? WHERE `id` = ?", password, nickname, group, enable, tokenId, apiId, id)
}
func (u *User) Drop() (result sql.Result, err error) {
	return u.DB.Exec("DROP TABLE IF EXISTS `user`")
}
