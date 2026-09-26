package main

import (
	"data-center/cmd"
	"embed"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed frontend/dist
var embedFS embed.FS

//go:embed i18n/*
var i18nFS embed.FS

// @title data-center API
// @version 1.0
// @description data-center 服务接口文档
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description 格式：Bearer {token}
func main() {
	cmd.Execute(embedFS, i18nFS)
}
