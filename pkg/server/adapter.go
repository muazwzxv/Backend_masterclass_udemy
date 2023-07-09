package server

type IServer interface {
  Start(address string) error
  Stop() error
}
