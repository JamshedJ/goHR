package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	App      app
	Database database
	Logger   logger
)

type app struct {
	URL     string `json:"url"`
	Version string `json:"version"`
}

type database struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type logger struct {
	FolderName  string `json:"folder_name"`
	FileDebug   string `json:"file_debug"`
	FileInfo    string `json:"file_info"`
	FileWarning string `json:"file_warning"`
	FileError   string `json:"file_error"`
	MaxSizeMB   int    `json:"max_size_mb"`
	MaxBackups  int    `json:"max_backups"`
	MaxAgeDays  int    `json:"max_age_days"`
	Compress    bool   `json:"compress"`
}

type file struct {
	App      app      `json:"app"`
	Database database `json:"database"`
	Logger   logger   `json:"logger"`
}

func Init() error {
	data, err := os.ReadFile("./configs.json")
	if err != nil {
		return fmt.Errorf("configs Init: error reading config file")
	}

	var f file
	if err = json.Unmarshal(data, &f); err != nil {
		return fmt.Errorf("configs Init: error unmarshalling settings")
	}

	App = f.App
	Database = f.Database
	Logger = f.Logger
	return nil
}
