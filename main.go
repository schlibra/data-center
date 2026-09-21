package main

import (
	"data-center/cmd"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cmd.Execute()
}
