package core

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port        int
	URL         string
	Compress    bool
	Domain      string
	Dev         bool
	Ssl         bool
	MailDomain  string
	MailPKey    string
	SQLDBName   string
	SQLUser     string
	SQLPassword string
	SQLHost     string
	SQLPort     string
	StripeKey   string
	AppEmail    string
	AppName     string
	LocalHost   string
}

var Secure bool

var Config ServerConfig

func SetupConfig() {
	if flag.Lookup("test.v") == nil {
		viper.AddConfigPath(BinPath)
		viper.SetConfigName("config")
	} else {
		log.Println("Testing.......")
		viper.AddConfigPath("./")
		viper.SetConfigName("test-config")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&Config)

	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	//Set env
	if Config.Dev {
		Secure = false

	} else {
		Secure = true
	}

}

var SKey = "JPcT1k6SwyqA2JX-oyGjGfOfHzKsN2BQdI4Cr56KG9M="
var StaticUrl = "/static"

type Paths struct {
	DirPath   string
	StaticUrl string
}

func CleanPath(path string) string {
	path = strings.Replace(path, "//", "/", -1)
	return path
}

func AbsolutePath(path string) string {
	if flag.Lookup("test.v") == nil {
		return CleanPath(filepath.FromSlash(BinPath + "/" + path))
	} else {
		return CleanPath(filepath.FromSlash(path))
	}
}

func BinaryPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("Error getting file path")
	}
	return dir
}
func FixPathSlash(path string) string {
	return filepath.FromSlash(path)
}
func PagePath(dirPath string, spage string) string {
	return FixPathSlash(dirPath + "/html/" + spage)
}

var BinPath = BinaryPath()
var StaticDir = AbsolutePath("static")

func Start() {
	SetupConfig()
}
