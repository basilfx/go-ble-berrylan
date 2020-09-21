package berrylan

// ConnectionStatusUpdateHandler event handler function.
type ConnectionStatusUpdateHandler func(status WirelessConnectionStatus)

// WirelessInterface is the interface for managing the wireless connection.
type WirelessInterface interface {
	StartAccessPoint(ssid string, passphrase string) error
	GetConnection() *ConnectionInfo
	GetNetworks() []NetworkInfo
	ScanNetwork()
	Connect(ssid string, passphrase string, hidden bool) error
	Disconnect() error

	HandleConnectionStatusUpdate(f ConnectionStatusUpdateHandler)
}

// NetworkStatusUpdateHandler event handler function.
type NetworkStatusUpdateHandler func(status NetworkStatus)

// WirelessStateUpdateHandler event handler function.
type WirelessStateUpdateHandler func(enabled bool)

// NetworkingStateUpdateHandler event handler function
type NetworkingStateUpdateHandler func(enabled bool)

// NetworkInterface is the interface for communicating with the (system)
// networking service. It is responsible for enabling the networkint service
// and/or the wireless service.
type NetworkInterface interface {
	EnableNetworking(enabled bool) error
	EnableWireless(enabled bool) error

	HandleNetworkStatusUpdate(f NetworkStatusUpdateHandler)
	HandleWirelessStateUpdate(f WirelessStateUpdateHandler)
	HandleNetworkingStateUpdate(f NetworkingStateUpdateHandler)
}

// Interface is a combined interface of WirelessInterface and NetworkInterface,
// for convenience reasons.
type Interface interface {
	WirelessInterface
	NetworkInterface
}
