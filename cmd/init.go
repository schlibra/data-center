package cmd

import (
	"data-center/internal/repository"
	"data-center/pkg/utils"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize data-center database",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Open("install.lock")
		if err == nil {
			log.Fatal("Already initialized. Remove `install.lock` to reinitialize.")
		}
		log.Println("Loading database...")
		user := new(repository.User)
		if err := user.New(); err != nil {
			log.Fatalf("Failed to initialize user repository: %v", err)
		}
		defer user.Close()
		settings := new(repository.Settings)
		if err := settings.New(); err != nil {
			log.Fatalf("Failed to initialize settings repository: %v", err)
		}
		defer settings.Close()
		group := new(repository.Group)
		if err := group.New(); err != nil {
			log.Fatalf("Failed to initialize group repository: %v", err)
		}
		defer group.Close()
		permission := new(repository.Permission)
		if err := permission.New(); err != nil {
			log.Fatalf("Failed to initialize permission repository: %v", err)
		}
		defer permission.Close()
		frpRule := new(repository.FrpRule)
		if err := frpRule.New(); err != nil {
			log.Fatalf("Failed to initialize frp rule repository: %v", err)
		}
		defer frpRule.Close()
		frpToken := new(repository.FrpToken)
		if err := frpToken.New(); err != nil {
			log.Fatalf("Failed to initialize frp token repository: %v", err)
		}
		defer frpToken.Close()
		mailDomain := new(repository.MailDomain)
		if err := mailDomain.New(); err != nil {
			log.Fatalf("Failed to initialize mail domain repository: %v", err)
		}
		defer mailDomain.Close()
		log.Println("Creating tables...")
		if _, err := user.Create(); err != nil {
			log.Fatalf("Failed to create user table: %v", err)
		}
		if _, err := settings.Create(); err != nil {
			log.Fatalf("Failed to create settings table: %v", err)
		}
		if _, err := group.Create(); err != nil {
			log.Fatalf("Failed to create group table: %v", err)
		}
		if _, err := permission.Create(); err != nil {
			log.Fatalf("Failed to create permission table: %v", err)
		}
		if _, err := frpRule.Create(); err != nil {
			log.Fatalf("Failed to create frp rule table: %v", err)
		}
		if _, err := frpToken.Create(); err != nil {
			log.Fatalf("Failed to create frp token table: %v", err)
		}
		if _, err := mailDomain.Create(); err != nil {
			log.Fatalf("Failed to create mail domain table: %v", err)
		}
		log.Println("Inserting default data...")
		if _, err := group.Insert("普通用户", 0, "[]"); err != nil {
			log.Fatalf("Failed to insert default group: %v", err)
		}
		adminGroupRowResult, err := group.Insert("管理员", 1, "[]")
		if err != nil {
			log.Fatalf("Failed to insert default group: %v", err)
		}
		adminGroupId, err := adminGroupRowResult.LastInsertId()
		if err != nil {
			log.Fatalf("Failed to get last insert id: %v", err)
		}
		adminPwd, err := utils.GenerateRandomString(4)
		if err != nil {
			log.Fatalf("Failed to generate random password: %v", err)
		}
		pwdHash, err := utils.PasswordHash(adminPwd)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}
		_, err = user.Insert("admin", pwdHash, "管理员", int(adminGroupId), 1)
		if err != nil {
			log.Fatalf("Failed to insert default user: %v", err)
		}
		log.Printf("Created admin user with username 'admin' and password '%s'\n", adminPwd)
		log.Println("Initialization complete!")
		_, _ = os.Create("install.lock")
	},
}
