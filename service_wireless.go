package berrylan

import (
	"encoding/json"

	"github.com/basilfx/go-ble-berrylan/spec"
	"github.com/basilfx/go-ble-utilities/characteristics"
	"github.com/basilfx/go-utilities/observable"

	"github.com/go-ble/ble"

	log "github.com/sirupsen/logrus"
)

// WirelessService represents the BLE wireless service.
type WirelessService struct {
	wirelessInterface WirelessInterface

	commanderResponse *observable.Observable
	connectionStatus  *observable.Observable
	mode              *observable.Observable
}

// NewWirelessService initializes a new WirelessService. It wraps an interface
// to the actual wireless service.
func NewWirelessService(wirelessInterface WirelessInterface) *WirelessService {
	service := WirelessService{
		wirelessInterface: wirelessInterface,

		connectionStatus: observable.NewWithValue([]byte{
			byte(spec.WirelessConnectionStatusUnknown),
		}),
		mode: observable.NewWithValue([]byte{
			byte(spec.WirelessModeUnknown),
		}),
	}

	// Register handlers on the wireless interface.
	wirelessInterface.HandleConnectionStatusUpdate(func(s WirelessConnectionStatus) {
		service.connectionStatus.SetValue([]byte{byte(s)})
	})

	return &service
}

// Create returns a ble.Service with all the required characteristics for a
// wireless service.
func (w *WirelessService) Create() *ble.Service {
	s := ble.NewService(spec.ServiceWireless)

	writer, notifier :=
		characteristics.CommandResponseHandlerFunc(w.handleCommand)

	// Commander.
	char1 := s.NewCharacteristic(spec.CharacteristicWirelessCommander)

	char1.HandleWrite(writer)

	// Commander Response.
	char2 := s.NewCharacteristic(spec.CharacteristicWirelessCommanderResponse)

	char2.HandleNotify(notifier)

	// Connection status.
	char3 := s.NewCharacteristic(spec.CharacteristicWirelessConnectionStatus)

	char3.HandleRead(
		characteristics.ObservableReadHandlerFunc(w.connectionStatus))
	char3.HandleNotify(
		characteristics.ObservableNotifyHandlerFunc(w.connectionStatus))

	// Mode.
	char4 := s.NewCharacteristic(spec.CharacteristicWirelessMode)

	char4.HandleRead(characteristics.ObservableReadHandlerFunc(w.mode))
	char4.HandleNotify(characteristics.ObservableNotifyHandlerFunc(w.mode))

	return s
}

func writeResponse(data interface{}) *[]byte {
	response, err := json.Marshal(data)

	if err != nil {
		log.Errorf("Unable to marshall response: %v", err)
		return nil
	}

	log.Debugf("Outgoing response: %s", string(response))

	return &response
}

func (w *WirelessService) handleCommand(command []byte) *[]byte {
	// Unmarshall command.
	var request spec.WirelessCommandMessage

	log.Debugf("Incoming request: %s", string(command))

	err := json.Unmarshal([]byte(command), &request)

	if err != nil {
		log.Warnf("Unable to unmarshall request: %v", err)
		return nil
	}

	// Handle command.
	log.Debugf("Wireless service command 0x%02x.", request.Command)

	switch request.Command {
	case spec.WirelessCommandGetNetworks:
		networks := w.wirelessInterface.GetNetworks()

		parameters := make([]spec.WirelessResponseParameters, len(networks))

		for k, v := range networks {
			parameters[k] = spec.WirelessResponseParameters{
				Ssid:           v.Ssid,
				Mac:            v.MACAddress,
				SignalStrength: v.SignalStrength,
				Protected:      boolToInt(v.Protected),
			}
		}

		return writeResponse(spec.WirelessResponseMessage{
			Command:    spec.WirelessCommandGetNetworks,
			Response:   spec.WirelessResponseSuccess,
			Parameters: parameters,
		})
	case spec.WirelessCommandConnect:
		if request.Parameters == nil {
			log.Errorf("Missing parameters.")

			return writeResponse(spec.WirelessResponseMessage{
				Command:  spec.WirelessCommandConnect,
				Response: spec.WirelessResponseInvalidParameter,
			})
		}

		err := w.wirelessInterface.Connect(
			request.Parameters.Ssid,
			request.Parameters.Passpharse,
			false)

		if err != nil {
			log.Errorf("Error while connecting: %v.", err)

			return writeResponse(spec.WirelessResponseMessage{
				Command:  spec.WirelessCommandConnect,
				Response: spec.WirelessResponseUnknown,
			})
		}

		return writeResponse(spec.WirelessResponseMessage{
			Command:  spec.WirelessCommandConnect,
			Response: spec.WirelessResponseSuccess,
		})
	case spec.WirelessCommandConnectHidden:
		if request.Parameters == nil {
			log.Errorf("Missing parameters.")

			return writeResponse(spec.WirelessResponseMessage{
				Command:  spec.WirelessCommandConnectHidden,
				Response: spec.WirelessResponseInvalidParameter,
			})
		}

		err := w.wirelessInterface.Connect(
			request.Parameters.Ssid,
			request.Parameters.Passpharse,
			true)

		if err != nil {
			log.Errorf("Error while connecting hidden: %v.", err)

			return writeResponse(spec.WirelessResponseMessage{
				Command:  spec.WirelessCommandConnectHidden,
				Response: spec.WirelessResponseUnknown,
			})
		}

		return writeResponse(spec.WirelessResponseMessage{
			Command:  spec.WirelessCommandConnectHidden,
			Response: spec.WirelessResponseSuccess,
		})
	case spec.WirelessCommandDisconnect:
		err := w.wirelessInterface.Disconnect()

		if err != nil {
			log.Errorf("Unable to disconnect: %v", err)

			return writeResponse(spec.WirelessResponseMessage{
				Command:  spec.WirelessCommandDisconnect,
				Response: spec.WirelessResponseUnknown,
			})
		}

		return writeResponse(spec.WirelessResponseMessage{
			Command:  spec.WirelessCommandDisconnect,
			Response: spec.WirelessResponseSuccess,
		})
	case spec.WirelessCommandScan:
		w.wirelessInterface.ScanNetwork()

		return writeResponse(spec.WirelessResponseMessage{
			Command:  spec.WirelessCommandScan,
			Response: spec.WirelessResponseSuccess,
		})
	case spec.WirelessCommandGetConnection:
		var response *spec.WirelessResponseParameters

		connection := w.wirelessInterface.GetConnection()

		if connection != nil {
			response = &spec.WirelessResponseParameters{
				Ssid:           connection.Ssid,
				Mac:            connection.MACAddress,
				SignalStrength: connection.SignalStrength,
				Protected:      boolToInt(connection.Protected),
				IPAddress:      connection.IPAddress,
			}
		}

		return writeResponse(spec.WirelessResponseMessage{
			Command:    spec.WirelessCommandGetConnection,
			Response:   spec.WirelessResponseSuccess,
			Parameters: response,
		})
	case spec.WirelessCommandStartAccessPoint:
		if request.Parameters == nil {
			log.Errorf("Missing parameters.")

			return writeResponse(spec.WirelessResponseMessage{
				Command:  spec.WirelessCommandStartAccessPoint,
				Response: spec.WirelessResponseInvalidParameter,
			})
		}

		err := w.wirelessInterface.StartAccessPoint(
			request.Parameters.Ssid,
			request.Parameters.Ssid)

		if err != nil {
			log.Errorf("Unable to start access point: %v", err)

			return writeResponse(spec.WirelessResponseMessage{
				Command:  spec.WirelessCommandStartAccessPoint,
				Response: spec.WirelessResponseUnknown,
			})
		}

		return writeResponse(spec.WirelessResponseMessage{
			Command:  spec.WirelessCommandStartAccessPoint,
			Response: spec.WirelessResponseSuccess,
		})
	default:
		log.Warnf(
			"Wireless service command 0x%02x not supported.",
			request.Command)

		return writeResponse(spec.WirelessResponseMessage{
			Command:  request.Command,
			Response: spec.WirelessResponseInvalidCommand,
		})
	}
}
