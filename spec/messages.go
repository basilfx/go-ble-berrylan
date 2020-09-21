package spec

// Messages are based on the BerryLan specification from
// https://github.com/nymea/nymea-networkmanager

// WirelessCommandMessage message.
type WirelessCommandMessage struct {
	Command    WirelessCommand            `json:"c"`
	Parameters *WirelessCommandParameters `json:"p,omitempty"`
}

// WirelessCommandParameters struct.
type WirelessCommandParameters struct {
	Ssid       string `json:"e"`
	Passpharse string `json:"p"`
}

// WirelessResponseMessage message.
type WirelessResponseMessage struct {
	Command    WirelessCommand  `json:"c"`
	Response   WirelessResponse `json:"r"`
	Parameters interface{}      `json:"p,omitempty"`
}

// WirelessResponseParameters struct.
type WirelessResponseParameters struct {
	Ssid           string `json:"e"`
	Mac            string `json:"m"`
	SignalStrength int    `json:"s"`
	Protected      int    `json:"p"`
	IPAddress      string `json:"i"`
}
