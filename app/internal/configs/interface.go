package configs

//Configure ...
type Configure interface {
	GetDsnPG() string
	GetUrl() string
	GetKeyJWT() string
}
