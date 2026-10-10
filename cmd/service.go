package cmd

import (
	"data-center/internal/server"
	"log"
	"os"
	"path/filepath"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

const (
	serviceName        = "data-center"
	serviceDisplayName = "Data-Center Service"
	serviceDescription = "Data-Center Service"
)

type SystemService struct{}

func (ss *SystemService) Start(_ service.Service) error {
	log.Println("Starting Data-Center Service...")
	go ss.run()
	return nil
}

func (ss *SystemService) Stop(_ service.Service) error {
	log.Println("Stopping Data-Center Service...")
	server.Shutdown()
	log.Println("Data-Center Service stopped")
	return nil
}

func (ss *SystemService) run() {
	log.Println("Data-Center Service run...")
	server.Run(EmbedFS, I18nFS, Version, Commit, BuildTime)
}

func newService() (service.Service, error) {
	ss := &SystemService{}
	return service.New(ss, &service.Config{
		Name:        serviceName,
		DisplayName: serviceDisplayName,
		Description: serviceDescription,
	})
}

func RunService() {
	if err := chdirToExeDir(); err != nil {
		log.Println("change working directory failed:", err)
	}
	if err := redirectLogToFile(); err != nil {
		log.Println("open log file failed:", err)
	}

	s, err := newService()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

func chdirToExeDir() error {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return os.Chdir(filepath.Dir(exe))
}

func redirectLogToFile() error {
	logPath := filepath.Join(exeDir(), "data-center.log")
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logPath = filepath.Join(os.TempDir(), "data-center.log")
		if f, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err != nil {
			log.Fatal(err)
		}
	}
	log.SetOutput(f)
	return nil
}

func exeDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(exe)
}

var serviceCmd = &cobra.Command{
	Use:       "service [install|uninstall|start|stop|restart]",
	Short:     "Manage services",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"install", "uninstall", "start", "stop", "restart"},
	RunE: func(cmd *cobra.Command, args []string) error {
		action := args[0]
		s, err := newService()
		if err != nil {
			log.Fatal(err)
		}
		return service.Control(s, action)
	},
}
