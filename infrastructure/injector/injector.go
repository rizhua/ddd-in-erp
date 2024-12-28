package injector

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/gops/agent"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"

	"rizhua.com/infrastructure/adapter"
	"rizhua.com/infrastructure/etc"
	"rizhua.com/interface/http"
)

func NewInjector(engine *gin.Engine) *Injector {
	return &Injector{
		engine: engine,
	}
}

type Injector struct {
	engine *gin.Engine
}

func InitGinEngine(r *http.Router) *gin.Engine {
	cfg := etc.C

	gin.SetMode(cfg.RunMode)
	engine := gin.Default()

	r.Register(engine)

	return engine
}

type Options struct {
	ConfigFile string
	ModelFile  string
	Version    string
}

type Option func(*Options)

func SetConfigFile(s string) Option {
	return func(o *Options) {
		o.ConfigFile = s
	}
}

func SetModelFile(s string) Option {
	return func(o *Options) {
		o.ModelFile = s
	}
}

func SetVersion(s string) Option {
	return func(o *Options) {
		o.Version = s
	}
}

func InitLogger() (func(), error) {
	c := etc.C.Log
	logrus.SetLevel(logrus.Level(c.Level))

	switch c.Format {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	default:
		logrus.SetFormatter(new(logrus.TextFormatter))
	}

	// log output
	var file *os.File
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			logrus.SetOutput(os.Stdout)
		case "stderr":
			logrus.SetOutput(os.Stderr)
		case "file":
			if name := c.OutputFile; name != "" {
				_ = os.MkdirAll(filepath.Dir(name), 0777)

				f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					return nil, err
				}
				logrus.SetOutput(f)
				file = f
			}
		}
	}

	return func() {
		if file != nil {
			file.Close()
		}
	}, nil
}

func InitCasbin(db *xorm.Engine) (e *casbin.Enforcer, err error) {
	cfg := etc.C.Casbin
	if cfg.Model == "" {
		return
	}
	a, err := xormadapter.NewAdapterByEngine(db)
	if err != nil {
		return
	}
	e, err = casbin.NewEnforcer(cfg.Model, a)

	return
}

func initMonitor() func() {
	if c := etc.C.Monitor; c.Enable {
		// ShutdownCleanup set false to prevent automatically closes on os.Interrupt
		// and close agent manually before service shutting down
		err := agent.Listen(agent.Options{Addr: c.Addr, ConfigDir: c.ConfigDir, ShutdownCleanup: false})
		if err != nil {
			logrus.Errorf("Agent monitor error: %s", err.Error())
		}
		return func() {
			agent.Close()
		}
	}
	return func() {}
}

func initServer(opts ...Option) (func(), error) {
	cfg := etc.C
	var o Options

	for _, opt := range opts {
		opt(&o)
	}
	etc.MustLoad(o.ConfigFile)
	if v := o.ModelFile; v != "" {
		etc.C.Casbin.Model = v
	}
	etc.PrintWithJSON()

	loggerCleanFunc, err := InitLogger()
	if err != nil {
		return nil, err
	}

	monitorCleanFunc := initMonitor()

	db, _ := adapter.NewDb("postgres")
	defer db.Close()

	injector, cleanup, err := BuildInjector(db)
	if err != nil {
		return nil, err
	}
	injector.engine.Run(cfg.HTTP.Addr())

	return func() {
		loggerCleanFunc()
		monitorCleanFunc()
		cleanup()
	}, nil
}

func RunServer(ctx context.Context, opts ...Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := initServer(opts...)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logrus.WithContext(ctx).Infof("catched signal[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logrus.WithContext(ctx).Infof("stopping server")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
