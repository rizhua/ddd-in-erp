package etc

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/koding/multiconfig"
)

var (
	C    = new(Config)
	once sync.Once
)

func MustLoad(fpaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
		}
		for _, fpath := range fpaths {
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}
		// Can be overwritten config with environment variable
		loaders = append(loaders, &multiconfig.EnvironmentLoader{})
		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(C)
	})
}

func PrintWithJSON() {
	if C.PrintConfig {
		b, err := json.MarshalIndent(C, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

type Config struct {
	Casbin      Casbin
	Email       Email
	HTTP        HTTP
	Log         Log
	LogXormHook LogXormHook
	Monitor     Monitor
	Postgres    Postgres
	PrintConfig bool
	Redis       Redis
	Root        Root
	RPC         RPC
	RunMode     string
	Swagger     bool
	Xorm        Xorm
	File        File
}

func (c *Config) DebugMode() bool {
	return c.RunMode == "debug"
}

type HTTP struct {
	Host             string
	Port             int
	CertFile         string
	KeyFile          string
	ShutdownTimeout  int
	MaxContentLength int64
	MaxLoggerLength  int `default:"4096"`
}

func (h HTTP) Addr() string {
	return h.Host + ":" + fmt.Sprintf("%d", h.Port)
}

type RPC struct {
	Host string
	Port int
}

func (r RPC) Addr() string {
	return r.Host + ":" + fmt.Sprintf("%d", r.Port)
}

type Monitor struct {
	Enable    bool
	Addr      string
	ConfigDir string
}

type Xorm struct {
	Debug             bool
	DBType            string
	MaxLifetime       int
	MaxOpenConns      int
	MaxIdleConns      int
	TablePrefix       string
	EnableAutoMigrate bool
}

type Casbin struct {
	Enable           bool
	Debug            bool
	Model            string
	AutoLoad         bool
	AutoLoadInternal int
}

type Log struct {
	Level         int
	Format        string
	Output        string
	OutputFile    string
	EnableHook    bool
	HookLevels    []string
	Hook          LogHook
	HookMaxThread int
	HookMaxBuffer int
}

type LogXormHook struct {
	DBType       string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	Table        string
}

type LogHook string

func (h LogHook) IsXorm() bool {
	return h == "xorm"
}

type Redis struct {
	Addr     string
	Password string
}

type Root struct {
	UserName string
	Password string
	Nickname string
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (d Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		d.Host, d.Port, d.User, d.DBName, d.Password, d.SSLMode)
}

type Email struct {
	User     string
	Password string
	Host     string
}

type File struct {
	Path   string
	Prefix string
	Size   int64
}
