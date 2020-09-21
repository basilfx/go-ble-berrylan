package berrylan

import (
	"github.com/basilfx/go-ble-berrylan/spec"
	"github.com/basilfx/go-ble-utilities/characteristics"
	"github.com/basilfx/go-utilities/observable"

	"github.com/go-ble/ble"

	log "github.com/sirupsen/logrus"
)

// NetworkService represents the BLE wireless service.
type NetworkService struct {
	networkInterface NetworkInterface

	status            *observable.Observable
	commanderResponse *observable.Observable
	networkingEnabled *observable.Observable
	wirelessEnabled   *observable.Observable
}

// NewNetworkService initializes a new NetworkService. It wraps an interface
// to the actual network service.
func NewNetworkService(networkInterface NetworkInterface) *NetworkService {
	service := NetworkService{
		networkInterface: networkInterface,

		status:            observable.NewWithValue([]byte{byte(spec.NetworkStatusUnknown)}),
		commanderResponse: observable.New(),
		networkingEnabled: observable.NewWithValue([]byte{byte(spec.NetworkingStateDisabled)}),
		wirelessEnabled:   observable.NewWithValue([]byte{byte(spec.WirelessStateDisabled)}),
	}

	// Register handlers on the network interface.
	networkInterface.HandleNetworkStatusUpdate(func(s NetworkStatus) {
		service.status.SetValue([]byte{byte(s)})
	})
	networkInterface.HandleNetworkingStateUpdate(func(e bool) {
		if e {
			service.networkingEnabled.SetValue([]byte{byte(spec.NetworkingStateEnabled)})
		} else {
			service.networkingEnabled.SetValue([]byte{byte(spec.NetworkingStateDisabled)})
		}
	})
	networkInterface.HandleWirelessStateUpdate(func(e bool) {
		if e {
			service.wirelessEnabled.SetValue([]byte{byte(spec.WirelessStateEnabled)})
		} else {
			service.wirelessEnabled.SetValue([]byte{byte(spec.WirelessStateDisabled)})
		}
	})

	return &service
}

// Create returns a ble.Service with all the required characteristics for a
// network service.
func (n *NetworkService) Create() *ble.Service {
	s := ble.NewService(spec.ServiceNetwork)

	// Network Status.
	char1 := s.NewCharacteristic(spec.CharacteristicNetworkStatus)

	char1.HandleRead(characteristics.ObservableReadHandlerFunc(n.status))
	char1.HandleNotify(characteristics.ObservableNotifyHandlerFunc(n.status))

	// Network Commander.
	char2 := s.NewCharacteristic(spec.CharacteristicNetworkCommander)

	char2.HandleWrite(
		ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
			n.handleCommand(req.Data())
		}))

	// Network Commander Response.
	char3 := s.NewCharacteristic(spec.CharacteristicNetworkCommanderResponse)

	char3.HandleNotify(characteristics.ObservableNotifyHandlerFunc(n.commanderResponse))

	// Networking Enabled.
	char4 := s.NewCharacteristic(spec.CharacteristicNetworkEnabled)

	char4.HandleRead(characteristics.ObservableReadHandlerFunc(n.networkingEnabled))
	char4.HandleNotify(characteristics.ObservableNotifyHandlerFunc(n.networkingEnabled))

	// Wireless Enabled.
	char5 := s.NewCharacteristic(spec.CharacteristicNetworkWirelessEnabled)

	char5.HandleRead(characteristics.ObservableReadHandlerFunc(n.wirelessEnabled))
	char5.HandleNotify(characteristics.ObservableNotifyHandlerFunc(n.wirelessEnabled))

	return s
}

func (n *NetworkService) writeResponse(r spec.NetworkResponse) {
	n.commanderResponse.SetValue([]byte{byte(spec.NetworkResponseInvalidValue)})
}

func (n *NetworkService) handleCommand(req []byte) {
	// Unmarshall command.
	if len(req) != 1 {
		log.Warnf("Expected command to be one byte, got %d bytes.", len(req))
		return
	}

	command := spec.NetworkCommand(req[0])

	// Handle command.
	log.Debugf("Network service command 0x%02x.", command)

	switch command {
	case spec.NetworkCommandEnableNetworking:
		err := n.networkInterface.EnableNetworking(true)

		if err != nil {
			log.Errorf("Unable to enable networking: %v", err)

			n.writeResponse(spec.NetworkResponseUnknown)
			return
		}

		n.writeResponse(spec.NetworkResponseSuccess)
	case spec.NetworkCommandDisableNetworking:
		err := n.networkInterface.EnableNetworking(false)

		if err != nil {
			log.Errorf("Unable to disable networking: %v", err)

			n.writeResponse(spec.NetworkResponseUnknown)
			return
		}

		n.writeResponse(spec.NetworkResponseSuccess)
	case spec.NetworkCommandEnableWireless:
		err := n.networkInterface.EnableWireless(true)

		if err != nil {
			log.Errorf("Unable to enable wireless: %v", err)

			n.writeResponse(spec.NetworkResponseUnknown)
			return
		}

		n.writeResponse(spec.NetworkResponseSuccess)
	case spec.NetworkCommandDisableWireless:
		err := n.networkInterface.EnableWireless(false)

		if err != nil {
			log.Errorf("Unable to disable wireless: %v", err)

			n.writeResponse(spec.NetworkResponseUnknown)
			return
		}

		n.writeResponse(spec.NetworkResponseSuccess)
	default:
		log.Warnf("Network service command 0x%02x not supported.", command)

		n.writeResponse(spec.NetworkResponseUnknown)
	}
}
