package app

import (
	"github.com/kara9renai/yokattar-go/pkg/config"
	"github.com/kara9renai/yokattar-go/pkg/dao"
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
