package dao

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kara9renai/yokattar-go/app/domain/repository"
)

type (
	Dao interface {
		// Get Account repository
		Account() repository.Account

		// Get Status repository
		Status() repository.Status

		// Get Timeline repository
		Timeline() repository.Timeline

		// Clear ALl date in DB
		InitAll() error
	}

	dao struct {
		db *sqlx.DB
	}
)

func New(config DBConfig) (Dao, error) {
	db, err := initDb(config)
	if err != nil {
		return nil, err
	}

	return &dao{db: db}, nil
}

func (d *dao) Account() repository.Account {
	return NewAccount(d.db)
}

func (d *dao) Status() repository.Status {
	return NewStatus(d.db)
}

func (d *dao) Timeline() repository.Timeline {
	return NewTimeline(d.db)
}

func (d *dao) InitAll() error {
	if err := d.exec("SET FOREIGN_KEY_CHECKS=0"); err != nil {
		return fmt.Errorf("can't disable FOREIGN_KEY_CHECKS: %w", err)
	}

	defer func() {
		err := d.exec("SET FOREIGN_KEY_CHECKS=0")
		if err != nil {
			log.Printf("can't resotre FOREIGN_KEY_CHECKS: %+v", err)
		}
	}()

	for _, table := range []string{"account", "status"} {
		if err := d.exec("TRUNCATE TABLE" + table); err != nil {
			return fmt.Errorf("can't truncate table "+table+": %w", err)
		}
	}

	return nil
}

func (d *dao) exec(query string, args ...interface{}) error {
	_, err := d.db.Exec(query, args...)
	return err
}
