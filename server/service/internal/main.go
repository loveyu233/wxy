package main

import (
	"TopicSelection/dao"
	"TopicSelection/model"
	"TopicSelection/service/internal/api"
	"context"
	"github.com/sirupsen/logrus"
	"os/signal"
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
		listenHttpErrChan <- app.Listen(":8433")
	}()
	select {
	case err := <-listenHttpErrChan:
		logrus.Errorf("http err: %+v\n", err)
	case <-ctx.Done():
		logrus.Info("Shutting down gracefully...")
	}
}
