package spec

// UUIDs are based on the BerryLan specification from
// https://github.com/nymea/nymea-networkmanager

import "github.com/go-ble/ble"

// UUIDs of the BerryLan services.
var (
	ServiceWireless = ble.MustParse("e081fec0-f757-4449-b9c9-bfa83133f7fc")
	ServiceNetwork  = ble.MustParse("ef6d6610-b8af-49e0-9eca-ab343513641c")
)

// UUIDs of the characteristics.
var (
	CharacteristicWirelessCommander         = ble.MustParse("e081fec1-f757-4449-b9c9-bfa83133f7fc")
	CharacteristicWirelessCommanderResponse = ble.MustParse("e081fec2-f757-4449-b9c9-bfa83133f7fc")
	CharacteristicWirelessConnectionStatus  = ble.MustParse("e081fec3-f757-4449-b9c9-bfa83133f7fc")
	CharacteristicWirelessMode              = ble.MustParse("e081fec4-f757-4449-b9c9-bfa83133f7fc")
	CharacteristicNetworkStatus             = ble.MustParse("ef6d6611-b8af-49e0-9eca-ab343513641c")
	CharacteristicNetworkCommander          = ble.MustParse("ef6d6612-b8af-49e0-9eca-ab343513641c")
	CharacteristicNetworkCommanderResponse  = ble.MustParse("ef6d6613-b8af-49e0-9eca-ab343513641c")
	CharacteristicNetworkEnabled            = ble.MustParse("ef6d6614-b8af-49e0-9eca-ab343513641c")
	CharacteristicNetworkWirelessEnabled    = ble.MustParse("ef6d6615-b8af-49e0-9eca-ab343513641c")
)
