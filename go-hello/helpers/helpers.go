package helpers

type Protocol string

const (
	HTTP  Protocol = "http"
	HTTPS Protocol = "https"
)

type Port int

const (
	PORT_80 Port = 80
	PORT_90 Port = 90
)
