package berrylan

// WirelessConnectionStatusToString returns a string representation of a given
// WirelessConnectionStatus.
func (s *WirelessConnectionStatus) String() string {
	if *s > WirelessConnectionStatusFailed {
		panic("argument out of range.")
	}

	return [...]string{
		"Unknown",
		"Unmanaged",
		"Unavailable",
		"Disconnected",
		"Prepare",
		"Config",
		"NeedAuth",
		"IpConfig",
		"IpCheck",
		"Secondaries",
		"Activated",
		"Deactivating",
		"Failed",
	}[*s]
}

// NetworkStatusToString returns a string representation of a given
// NetworkStatus.
func (s *NetworkStatus) String() string {
	if *s > NetworkStatusConnectedGlobal {
		panic("argument out of range.")
	}

	return [...]string{
		"Unknown",
		"Asleep",
		"Disconnected",
		"Disconnecting",
		"Connecting",
		"Local",
		"ConnectedSite",
		"ConnectedGlobal",
	}[*s]
}
