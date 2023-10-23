package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/HavvokLab/true-solar-monitoring/api"
	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gofiber/fiber/v2"
)

func init() {
	config.InitConfig()
	util.SetTimezone()
}

func init() {
	logger.InitLogger(constant.API_LOG_NAME)
}

func init() {
	infra.InitGormDB()
}

func main() {
	app := fiber.New()
	app.Get("", func(c *fiber.Ctx) error { return c.SendStatus(200) })
	api.InitAPI(app)

	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
		syscall.SIGTERM, // kill -SIGTERM XXXX
	)

	serverShutdown := make(chan struct{})
	go func() {
		<-c
		logger.GetLogger().Info("Gracefully shutting down...")
		app.Shutdown()
		serverShutdown <- struct{}{}
	}()

	addr := getAddress()
	if err := app.Listen(addr); err != nil {
		logger.GetLogger().Fatal(err)
	}

	<-serverShutdown
	logger.GetLogger().Info("Running cleanup tasks...")
}

func getAddress() string {
	conf := config.GetConfig().API
	addr := ":8000"
	if !util.EmptyString(conf.Port) {
		addr = fmt.Sprintf(":%v", conf.Port)
	}

	return addr
}
