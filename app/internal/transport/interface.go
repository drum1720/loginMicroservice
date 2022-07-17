package transport

type Server interface {
	Run(cancel func())
	Shutdown()
}
