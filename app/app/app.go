package app

import (
	"github.com/kara9renai/yokattar-go/app/config"
	"github.com/kara9renai/yokattar-go/app/dao"
)

type App struct {
	Dao dao.Dao
}

func NewApp() (*App, error) {
	daoCfg := config.MySQLConfig()

	dao, err := dao.New(daoCfg)
	if err != nil {
		return nil, err
	}

	return &App{Dao: dao}, nil
}
