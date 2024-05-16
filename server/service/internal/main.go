package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os/signal"
	"server/dao"
	"server/model"
	"server/service/internal/api"
	"syscall"
)

var listenHttpErrChan = make(chan error)

func AutoMigrate() {
	err := dao.DB.AutoMigrate(&model.User{})
	if err != nil {
		logrus.Error("11111")
	}
	err = dao.DB.AutoMigrate(&model.Topic{})
	if err != nil {
		logrus.Error("11111")
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	dao.PgsqlInit()
	AutoMigrate()
	go func() {
		app := api.RouterInit()
		listenHttpErrChan <- app.Listen("127.0.0.1:8433")
	}()
	select {
	case err := <-listenHttpErrChan:
		logrus.Errorf("http err: %+v\n", err)
	case <-ctx.Done():
		logrus.Info("Shutting down gracefully...")
	}
}
