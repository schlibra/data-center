package repository

import (
	"context"
	"data-center/internal/models"
	"database/sql"
	"time"
)

type FrpRule struct {
	DB *sql.DB
}

func NewFrpRule() (*FrpRule, error) {
	db, err := LoadMySQL()
	return &FrpRule{DB: db}, err
}

func (r *FrpRule) Create() (sql.Result, error) {
	return r.DB.Exec("CREATE TABLE `frp-rule`  (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `min` int NOT NULL,\n  `max` int NOT NULL,\n  `token` int NOT NULL,\n  `user` int NOT NULL,\n  PRIMARY KEY (`id`)\n);")
}
func (r *FrpRule) Insert(min int, max int, token int, user int) (sql.Result, error) {
	return r.DB.Exec("INSERT INTO `frp-rule` (`min`, `max`, `token`, `user`) VALUES (?, ?, ?, ?);", min, max, token, user)
}
func (r *FrpRule) UpdateById(id int, min int, max int, token int, user int) (sql.Result, error) {
	return r.DB.Exec("UPDATE `frp-rule` SET `min`=?, `max`=?, `token`=?, `user`=? WHERE `id`=?;", min, max, token, user, id)
}
func (r *FrpRule) SelectById(id int) (rule models.FrpRuleTable, err error) {
	frpToken := FrpToken{DB: r.DB}
	err = r.DB.QueryRow("SELECT * FROM `frp-rule` WHERE `id`=?;", id).Scan(&rule.ID, &rule.Min, &rule.Max, &rule.Token, &rule.User)
	tokenRow, err := frpToken.SelectById(rule.Token)
	if err != nil {
		return rule, err
	}
	rule.TokenInfo = tokenRow
	rule.UserInfo = tokenRow.UserInfo
	return rule, err
}
func (r *FrpRule) SelectByToken(token int) (rules []models.FrpRuleTable, err error) {
	frpToken := FrpToken{DB: r.DB}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := r.DB.QueryContext(ctx, "SELECT * FROM `frp-rule` WHERE `token` = ?;", token)
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var row models.FrpRuleTable
		if err := rows.Scan(&row.ID, &row.Min, &row.Max, &row.Token, &row.User); err != nil {
			return nil, err
		}
		tokenRow, err := frpToken.SelectById(row.Token)
		if err != nil {
			return nil, err
		}
		row.TokenInfo = tokenRow
		row.UserInfo = tokenRow.UserInfo
		rules = append(rules, row)
	}
	return rules, nil
}
func (r *FrpRule) SelectByUser(user int) (rules []models.FrpRuleTable, err error) {
	frpToken := FrpToken{DB: r.DB}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := r.DB.QueryContext(ctx, "SELECT * FROM `frp-rule` WHERE `user` = ?;", user)
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var row models.FrpRuleTable
		if err := rows.Scan(&row.ID, &row.Min, &row.Max, &row.Token, &row.User); err != nil {
			return nil, err
		}
		tokenRow, err := frpToken.SelectById(row.Token)
		if err != nil {
			return nil, err
		}
		row.TokenInfo = tokenRow
		row.UserInfo = tokenRow.UserInfo
		rules = append(rules, row)
	}
	return rules, nil
}
func (r *FrpRule) SelectAll() (rules []models.FrpRuleTable, err error) {
	frpToken := FrpToken{DB: r.DB}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := r.DB.QueryContext(ctx, "SELECT * FROM `frp-rule`;")
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var row models.FrpRuleTable
		if err := rows.Scan(&row.ID, &row.Min, &row.Max, &row.Token, &row.User); err != nil {
			return nil, err
		}
		tokenRow, err := frpToken.SelectById(row.Token)
		if err != nil {
			return nil, err
		}
		row.TokenInfo = tokenRow
		row.UserInfo = tokenRow.UserInfo
		rules = append(rules, row)
	}
	return rules, nil
}
func (r *FrpRule) DeleteById(id int) (sql.Result, error) {
	return r.DB.Exec("DELETE FROM `frp-rule` WHERE `id`=?;", id)
}
func (r *FrpRule) Drop() (sql.Result, error) {
	return r.DB.Exec("DROP TABLE `frp-rule`;")
}
