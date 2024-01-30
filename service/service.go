package service

import "database/sql"

type MyAppService struct {
	db *sql.DB
}

func NewMyApiService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
