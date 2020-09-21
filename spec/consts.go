package spec

// Constant values are based on the BerryLan specification from
// https://github.com/nymea/nymea-networkmanager

// WirelessCommand enum type.
type WirelessCommand byte

// WirelessCommand enum values.
const (
	WirelessCommandGetNetworks      = WirelessCommand(0x00)
	WirelessCommandConnect          = WirelessCommand(0x01)
	WirelessCommandConnectHidden    = WirelessCommand(0x02)
	WirelessCommandDisconnect       = WirelessCommand(0x03)
	WirelessCommandScan             = WirelessCommand(0x04)
	WirelessCommandGetConnection    = WirelessCommand(0x05)
	WirelessCommandStartAccessPoint = WirelessCommand(0x06)
)

// WirelessResponse enum type.
type WirelessResponse uint

// WirelessResponse enum values.
const (
	WirelessResponseSuccess                    = WirelessResponse(0x00)
	WirelessResponseInvalidCommand             = WirelessResponse(0x01)
	WirelessResponseInvalidParameter           = WirelessResponse(0x02)
	WirelessResponseNetworkManagerNotAvailable = WirelessResponse(0x03)
	WirelessResponseWirelessNotAvailable       = WirelessResponse(0x04)
	WirelessResponseNetworkingDisabled         = WirelessResponse(0x05)
	WirelessResponseWirelessDisabled           = WirelessResponse(0x06)
	WirelessResponseUnknown                    = WirelessResponse(0x07)
)

// WirelessConnectionStatus enum type.
type WirelessConnectionStatus byte

// WirelessConnectionStatus enum values.
const (
	WirelessConnectionStatusUnknown      = WirelessConnectionStatus(0x00)
	WirelessConnectionStatusUnmanaged    = WirelessConnectionStatus(0x01)
	WirelessConnectionStatusUnavailable  = WirelessConnectionStatus(0x02)
	WirelessConnectionStatusDisconnected = WirelessConnectionStatus(0x03)
	WirelessConnectionStatusPrepare      = WirelessConnectionStatus(0x04)
	WirelessConnectionStatusConfig       = WirelessConnectionStatus(0x05)
	WirelessConnectionStatusNeedAuth     = WirelessConnectionStatus(0x06)
	WirelessConnectionStatusIPConfig     = WirelessConnectionStatus(0x07)
	WirelessConnectionStatusIPCheck      = WirelessConnectionStatus(0x08)
	WirelessConnectionStatusSecondaries  = WirelessConnectionStatus(0x09)
	WirelessConnectionStatusActivated    = WirelessConnectionStatus(0x0a)
	WirelessConnectionStatusDeactivating = WirelessConnectionStatus(0x0b)
	WirelessConnectionStatusFailed       = WirelessConnectionStatus(0x0c)
)

// WirelessMode enum type.
type WirelessMode byte

// WirelessMode enum values.
const (
	WirelessModeUnknown        = WirelessMode(0x00)
	WirelessModeAdhoc          = WirelessMode(0x01)
	WirelessModeInfrastructure = WirelessMode(0x02)
	WirelessModeAccessPoint    = WirelessMode(0x03)
)

// NetworkStatus enum type.
type NetworkStatus byte

// NetworkStatus enum values.
const (
	NetworkStatusUnknown         = NetworkStatus(0x00)
	NetworkStatusAsleep          = NetworkStatus(0x01)
	NetworkStatusDisconnected    = NetworkStatus(0x02)
	NetworkStatusDisconnecting   = NetworkStatus(0x03)
	NetworkStatusConnecting      = NetworkStatus(0x04)
	NetworkStatusLocal           = NetworkStatus(0x05)
	NetworkStatusConnectedSite   = NetworkStatus(0x06)
	NetworkStatusConnectedGlobal = NetworkStatus(0x07)
)

// NetworkCommand enum type.
type NetworkCommand byte

// NetworkCommand enum values.
const (
	NetworkCommandEnableNetworking  = NetworkCommand(0x00)
	NetworkCommandDisableNetworking = NetworkCommand(0x01)
	NetworkCommandEnableWireless    = NetworkCommand(0x02)
	NetworkCommandDisableWireless   = NetworkCommand(0x03)
)

// NetworkResponse enum type.
type NetworkResponse byte

// NetworkResponse enum values.
const (
	NetworkResponseSuccess                    = NetworkResponse(0x00)
	NetworkResponseInvalidValue               = NetworkResponse(0x01)
	NetworkResponseNetworkManagerNotAvailable = NetworkResponse(0x02)
	NetworkResponseWirelessNotAvailable       = NetworkResponse(0x03)
	NetworkResponseUnknown                    = NetworkResponse(0x04)
)

// NetworkingState enum type.
type NetworkingState byte

// NetworkingState enum values.
const (
	NetworkingStateDisabled = NetworkingState(0x00)
	NetworkingStateEnabled  = NetworkingState(0x01)
)

// WirelessState enum type.
type WirelessState byte

// WirelessState enum values.
const (
	WirelessStateDisabled = WirelessState(0x00)
	WirelessStateEnabled  = WirelessState(0x01)
)
