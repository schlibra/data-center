package repository

import (
	"context"
	"data-center/internal/models"
	"database/sql"
	"time"
)

type FrpToken struct {
	DB *sql.DB
}

func NewFrpToken() (*FrpToken, error) {
	db, err := LoadMySQL()
	return &FrpToken{DB: db}, err
}
func (t *FrpToken) Create() (sql.Result, error) {
	return t.DB.Exec("CREATE TABLE `frp-token`  (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `name` varchar(255) NOT NULL,\n  `token` varchar(255) NOT NULL,\n  `user` int NOT NULL,\n  `enable` int NOT NULL,\n  PRIMARY KEY (`id`)\n);")
}
func (t *FrpToken) Insert(name string, token string, user int, enable int) (sql.Result, error) {
	return t.DB.Exec("INSERT INTO `frp-token` (`name`, `token`, `user`, `enable`) VALUES (?, ?, ?, ?);", name, token, user, enable)
}
func (t *FrpToken) UpdateById(id int, name string, token string, user int, enable int) (sql.Result, error) {
	return t.DB.Exec("UPDATE `frp-token` SET `name` = ?, `token` = ?, `user` = ?, `enable` = ? WHERE `id` = ?;", name, token, user, enable, id)
}
func (t *FrpToken) SelectById(id int) (token models.FrpTokenTable, err error) {
	user := User{DB: t.DB}
	err = t.DB.QueryRow("SELECT * FROM `frp-token` WHERE `id` = ?;", id).Scan(&token.ID, &token.Name, &token.Token, &token.User, &token.Enable)
	userInfo, err := user.SelectById(token.User)
	if err != nil {
		return token, err
	}
	userInfo.Password = "********"
	userInfo.TokenID = "********"
	token.UserInfo = userInfo
	return token, err
}
func (t *FrpToken) SelectByName(name string) (token models.FrpTokenTable, err error) {
	user := User{DB: t.DB}
	err = t.DB.QueryRow("SELECT * FROM `frp-token` WHERE `name` = ?;", name).Scan(&token.ID, &token.Name, &token.Token, &token.User, &token.Enable)
	userInfo, err := user.SelectById(token.User)
	if err != nil {
		return token, err
	}
	userInfo.Password = "********"
	userInfo.TokenID = "********"
	token.UserInfo = userInfo
	return token, err
}
func (t *FrpToken) SelectByUser(userId int) (tokens []models.FrpTokenTable, err error) {
	user := User{DB: t.DB}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := t.DB.QueryContext(ctx, "SELECT * FROM `frp-token` WHERE `user` = ?;", userId)
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var row models.FrpTokenTable
		if err := rows.Scan(&row.ID, &row.Name, &row.Token, &row.User, &row.Enable); err != nil {
			return nil, err
		}
		userInfo, err := user.SelectById(row.User)
		if err != nil {
			return nil, err
		}
		userInfo.Password = "********"
		userInfo.TokenID = "********"
		row.UserInfo = userInfo
		tokens = append(tokens, row)
	}
	if tokens == nil {
		tokens = []models.FrpTokenTable{}
	}
	return tokens, nil
}
func (t *FrpToken) SelectAll() (tokens []models.FrpTokenTable, err error) {
	user := User{DB: t.DB}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := t.DB.QueryContext(ctx, "SELECT * FROM `frp-token`;")
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var row models.FrpTokenTable
		if err := rows.Scan(&row.ID, &row.Name, &row.Token, &row.User, &row.Enable); err != nil {
			return nil, err
		}
		userInfo, err := user.SelectById(row.User)
		if err != nil {
			return nil, err
		}
		userInfo.Password = "********"
		userInfo.TokenID = "********"
		row.UserInfo = userInfo
		tokens = append(tokens, row)
	}
	if tokens == nil {
		tokens = []models.FrpTokenTable{}
	}
	return tokens, nil
}
func (t *FrpToken) DeleteById(id int) (sql.Result, error) {
	return t.DB.Exec("DELETE FROM `frp-token` WHERE `id` = ?;", id)
}
func (t *FrpToken) Drop() (sql.Result, error) {
	return t.DB.Exec("DROP TABLE `frp-token`;")
}
