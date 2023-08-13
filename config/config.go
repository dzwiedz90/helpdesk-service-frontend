package config

import "log"

type Config struct {
	HTTPAddress      string
	HTTPPort         string
	Timeout          string
	UsersGRPCPort    string
	UsersGRPCAddress string
	InfoLog          *log.Logger
	ErrorLog         *log.Logger
}
