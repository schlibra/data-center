package repository

import (
	"context"
	"data-center/internal/models"
	"database/sql"
	"time"
)

type MailDomain struct {
	DB *sql.DB
}

func NewMailDomain() (MailDomain, error) {
	db, err := LoadMySQL()
	return MailDomain{DB: db}, err
}
func (md MailDomain) Create() (sql.Result, error) {
	return md.DB.Exec("CREATE TABLE `mail-domain`  (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `domain` varchar(255) NOT NULL DEFAULT '',\n  `name` varchar(255) NOT NULL DEFAULT '',\n  `table_id` varchar(255) NOT NULL DEFAULT '',\n  PRIMARY KEY (`id`)\n);")
}
func (md MailDomain) Drop() (sql.Result, error) {
	return md.DB.Exec("DROP TABLE `mail-domain`;")
}
func (md MailDomain) Insert(domain string, name string, tableId string) (sql.Result, error) {
	return md.DB.Exec("INSERT INTO `mail-domain` (`domain`, `name`, `table_id`) VALUES (?, ?, ?);", domain, name, tableId)
}
func (md MailDomain) SelectById(id int) (row models.MailDomainTable, err error) {
	err = md.DB.QueryRow("SELECT * FROM `mail-domain` WHERE `id` = ?", id).Scan(&row.ID, &row.Domain, &row.Name, &row.TableId)
	return row, err
}
func (md MailDomain) SelectByDomain(domain string) (row models.MailDomainTable, err error) {
	err = md.DB.QueryRow("SELECT * FROM `mail-domain` WHERE `domain` = ?", domain).Scan(&row.ID, &row.Domain, &row.Name, &row.TableId)
	return row, err
}
func (md MailDomain) SelectAll() (results []models.MailDomainTable, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := md.DB.QueryContext(ctx, "SELECT * FROM `mail-domain`")
	if err != nil {
		return nil, err
	}
	defer Close(rows)
	for rows.Next() {
		var result models.MailDomainTable
		if err := rows.Scan(&result.ID, &result.Domain, &result.Name, &result.TableId); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
func (md MailDomain) UpdateById(id int, domain string, name string, tableId string) (sql.Result, error) {
	return md.DB.Exec("UPDATE `mail-domain` SET `domain` = ?, `name` = ?, `table_id` = ? WHERE `id` = ?", domain, name, tableId, id)
}
func (md MailDomain) DeleteById(id int) (sql.Result, error) {
	return md.DB.Exec("DELETE FROM `mail-domain` WHERE `id` = ?", id)
}
