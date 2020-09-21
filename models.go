package berrylan

import "github.com/basilfx/go-ble-berrylan/spec"

// WirelessConnectionStatus enum type.
type WirelessConnectionStatus spec.WirelessConnectionStatus

// WirelessConnectionStatus enum values. This maps directly on the enum values
// of the specification.
const (
	WirelessConnectionStatusUnknown      = WirelessConnectionStatus(spec.WirelessConnectionStatusUnknown)
	WirelessConnectionStatusUnmanaged    = WirelessConnectionStatus(spec.WirelessConnectionStatusUnmanaged)
	WirelessConnectionStatusUnavailable  = WirelessConnectionStatus(spec.WirelessConnectionStatusUnavailable)
	WirelessConnectionStatusDisconnected = WirelessConnectionStatus(spec.WirelessConnectionStatusDisconnected)
	WirelessConnectionStatusPrepare      = WirelessConnectionStatus(spec.WirelessConnectionStatusPrepare)
	WirelessConnectionStatusConfig       = WirelessConnectionStatus(spec.WirelessConnectionStatusConfig)
	WirelessConnectionStatusNeedAuth     = WirelessConnectionStatus(spec.WirelessConnectionStatusNeedAuth)
	WirelessConnectionStatusIPConfig     = WirelessConnectionStatus(spec.WirelessConnectionStatusIPConfig)
	WirelessConnectionStatusIPCheck      = WirelessConnectionStatus(spec.WirelessConnectionStatusIPCheck)
	WirelessConnectionStatusSecondaries  = WirelessConnectionStatus(spec.WirelessConnectionStatusSecondaries)
	WirelessConnectionStatusActivated    = WirelessConnectionStatus(spec.WirelessConnectionStatusActivated)
	WirelessConnectionStatusDeactivating = WirelessConnectionStatus(spec.WirelessConnectionStatusDeactivating)
	WirelessConnectionStatusFailed       = WirelessConnectionStatus(spec.WirelessConnectionStatusFailed)
)

// NetworkStatus enum type.
type NetworkStatus spec.NetworkStatus

// NetworkStatus enum values. This maps directly on the enum values of the
// specification.
const (
	NetworkStatusUnknown         = NetworkStatus(spec.NetworkStatusUnknown)
	NetworkStatusAsleep          = NetworkStatus(spec.NetworkStatusAsleep)
	NetworkStatusDisconnected    = NetworkStatus(spec.NetworkStatusDisconnected)
	NetworkStatusDisconnecting   = NetworkStatus(spec.NetworkStatusDisconnecting)
	NetworkStatusConnecting      = NetworkStatus(spec.NetworkStatusConnecting)
	NetworkStatusLocal           = NetworkStatus(spec.NetworkStatusLocal)
	NetworkStatusConnectedSite   = NetworkStatus(spec.NetworkStatusConnectedSite)
	NetworkStatusConnectedGlobal = NetworkStatus(spec.NetworkStatusConnectedGlobal)
)

// ConnectionInfo struct.
type ConnectionInfo struct {
	Ssid           string `json:"ssid"`
	MACAddress     string `json:"macAddress"`
	SignalStrength int    `json:"signalStrength"`
	Protected      bool   `json:"protected"`
	IPAddress      string `json:"ipAddress"`
}

// NetworkInfo struct.
type NetworkInfo struct {
	Ssid           string `json:"ssid"`
	MACAddress     string `json:"macAddress"`
	SignalStrength int    `json:"signalStrength"`
	Protected      bool   `json:"protected"`
}
