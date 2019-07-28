package service

import "github.com/kataras/iris/core/router"

// Server interface
type Server interface {
	Run() error
	Stop() error
	register(r router.Party) error
}
