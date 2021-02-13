package ib

const (
	host     = "127.0.0.1"
	port     = 4003
	clientID = 0
)

// similar function for Cloud / cluster config
func GetLocalIbConfig() IBConfig {
	return IBConfig{
		IbHost:   host,
		IbPort:   port,
		ClientID: clientID,
	}
}
