package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
)

type Configation interface {
	checkCorrectNess()
	appConfiguration()
	developmentConfig()
	prodcutionConfig()
}

type config struct {
	DBHost     string `arg:"env:DBHost, -D, --dbhost" help:"Host of the database, defaults to 127.0.0.1" placeholder:"DBHost" `
	DBPort     string `arg:"env:DBPort, -P, --dbport" help:"Port of the database, defaults to 5432 (PostgreSQL)" placeholder:"DBPort"`
	DBUser     string `arg:"env:DBUser, -U, --dbuser, required" help:"Databse User" placeholder:"DBUser" `
	DBName     string `arg:"env:DBName, -n, --dbname, required" help:"Database name" placeholder:"DBName"`
	DBPassword string `arg:"env:DBPassword, -W, --dbpass, required" help:"Database password" placeholder:"DBPass"`
	DBSslMode  string `arg:"env:DBSslMode, -S, --dbssl" help:"SSL mode for the database, defaults disabled" placeholder:"DBSSL"`
	DBTimeZone string `arg:"env:DBTimeZone, -t, --dbtime" help:"Timezone for your database" placeholder:"DBTimeZone" `

	DSN      string `arg:"-"`
	Settings string `arg:"env:Settings, -s, --settings, required" help:"Settings either dev, test or prod" placeholder:"Settings" `
	Debug    bool   `arg:"env:Debug, -d, --debug" help:"Debug level defaults to true, make sure you pass false in prod" placeholder:"Debug"`
	Port     string `arg:"env:Port, -p, --port" help:"Port of your application defaults to 8080" placeholder:"Port"`
	Host     string `arg:"env:Host, -h, --host" help:"Host of your application defaults to 127.0.0.1" placeholder:"Host"`
	AppName  string `arg:"env:AppName, -a, --appname" help:"Name of your microservice app, defaults gin-microservice" placeholder:"AppName"`

	WriteLog  bool   `arg:"env:WriteLog, -w, --writelog" help:"Log to be writed to file" placeholder:"WriteLog"`
	LogFormat string `arg:"env:LogFormat, -f, --logfotmat" help:"Log format either json or text defaults to json" placeholder:"LogFormat"`
	LogFile   string `arg:"env:LogFile, -l, --logfile" help:"Log file in order to keep logs defaults log.txt" placeholder:"LogFile"`
}

var (
	Config  config
	Testing bool
)

type Operation struct {
	app Configation
}

// appConfiguration configuration app itself
func (c *config) appConfiguration() {
	DatabaseSettings(c)
	ServiceSettings(c)
	LogFormatSettings(c)
}

// developmentConfig is used for development purpose
func (c *config) developmentConfig() {

	if c.DBSslMode == "" {
		c.DBSslMode = "disabled"
	}

	c.DSN = fmt.Sprintf("host=%s  dbname=%s  port=%s timezone=%s user=%s password=%s  sslmode=%s", c.DBHost, c.DBName, c.DBPort, c.DBTimeZone, c.DBUser, c.DBPassword, c.DBSslMode)
	c.Debug = true

}

// prodcutionConfig is used for production purpose
func (c *config) prodcutionConfig() {
	if c.DBSslMode == "" {
		c.DBSslMode = "enabled"
	}
	c.DSN = fmt.Sprintf("host=%s  dbname=%s  port=%s timezone=%s user=%s password=%s  sslmode=%s", c.DBHost, c.DBName, c.DBPort, c.DBTimeZone, c.DBUser, c.DBPassword, c.DBSslMode)
	c.Debug = false

}

// checkCorrectNess will check given variables over
func (c config) checkCorrectNess() {

	AppSettings(c.Settings)

}

func init() {

	Testing = strings.HasSuffix(os.Args[0], ".test")

	if Testing {
		p, _ := arg.NewParser(arg.Config{Program: "test"}, &Config)


		settings := []string{"--settings=test"}

		p.Parse(settings)


	} else {
		arg.MustParse(&Config)

	}

	service := Operation{
		app: &Config,
	}

	service.app.checkCorrectNess()
	service.app.appConfiguration()

	if Config.Settings == "dev" {
		service.app.developmentConfig()

	} else if Config.Settings == "prod" {
		service.app.prodcutionConfig()

	}

}

func AppSettings(settings string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occured", err)
			os.Exit(1)
		}
	}()

	if !(settings == "dev" || settings == "prod" || settings == "test") {
		panic("settings must be either dev or prod")
	}
}

func DatabaseSettings(c *config) {

	if c.DBHost == "" {
		c.DBHost = "127.0.0.1"

	}
	if c.DBPort == "" {
		c.DBPort = "5432"
	}
}

func ServiceSettings(c *config) {

	if c.Port == "" {
		c.Port = "8080"
	}
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.AppName == "" {
		c.AppName = "gin-microservice"
	}
}

func LogFormatSettings(c *config) {

	if c.LogFormat == "" {

		c.LogFormat = "json"
	}

}
