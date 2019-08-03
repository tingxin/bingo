package service

// Server interface
type Server interface {
	Run() error
	Stop() error
}
