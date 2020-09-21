package main

import (
	"errors"
	"math/rand"
	"time"

	berrylan "github.com/basilfx/go-ble-berrylan"

	log "github.com/sirupsen/logrus"
)

// DummyWirelessInterface represents a wireless interface.
type DummyWirelessInterface struct {
	connected bool
	ssid      string

	connectionStatusUpdateHandler berrylan.ConnectionStatusUpdateHandler
}

// NewDummyWirelessInterface initializes a new instance of
// DummyWirelessInterface that simulates a wireless interface.
func NewDummyWirelessInterface() *DummyWirelessInterface {
	return &DummyWirelessInterface{}
}

func (d *DummyWirelessInterface) onConnectionUpdate(s berrylan.WirelessConnectionStatus) {
	log.Infof("Changing network status to '%s'.", s.String())

	if d.connectionStatusUpdateHandler != nil {
		d.connectionStatusUpdateHandler(s)
	}
}

// StartAccessPoint implements berrylan.WirelessInterface.StartAccessPoint by
// returning an error because it is not supported.
func (d *DummyWirelessInterface) StartAccessPoint(ssid string, passphrase string) error {
	return errors.New("not supported")
}

// GetConnection implements berrylan.WirelessInterface.GetConnection by
// returning dummy connection information (if connected).
func (d *DummyWirelessInterface) GetConnection() *berrylan.ConnectionInfo {
	if d.connected {
		return &berrylan.ConnectionInfo{
			Ssid:           d.ssid,
			MACAddress:     "00:00:00:00:00:00",
			IPAddress:      "127.0.0.1",
			Protected:      true,
			SignalStrength: 50,
		}
	}

	return nil
}

// GetNetworks implements berrylan.WirelessInterface.GetNetworks by returning
// two networks.
func (d *DummyWirelessInterface) GetNetworks() []berrylan.NetworkInfo {
	return []berrylan.NetworkInfo{
		{
			Ssid:           "Test Network 1",
			MACAddress:     "00:00:00:00:00:00",
			Protected:      true,
			SignalStrength: rand.Intn(100),
		},
		{
			Ssid:           "Test Network 2",
			MACAddress:     "00:00:00:00:00:00",
			Protected:      false,
			SignalStrength: rand.Intn(100),
		},
	}
}

// ScanNetwork implements berrylan.WirelessInterface.ScanNetwork by retuning
// nothing because it is not implemented.
func (d *DummyWirelessInterface) ScanNetwork() {
	return
}

// Connect implements berrylan.WirelessInterface.Connect by simulating a
// connection request to a given network.
func (d *DummyWirelessInterface) Connect(ssid string, passphrase string, hidden bool) error {
	log.Infof(
		"Connecting to '%s' with passpharse of length %d.",
		ssid,
		len(passphrase))

	d.ssid = ssid

	go func() {
		d.onConnectionUpdate(berrylan.WirelessConnectionStatusPrepare)
		time.Sleep(2 * time.Second)
		d.onConnectionUpdate(berrylan.WirelessConnectionStatusSecondaries)
		time.Sleep(2 * time.Second)

		d.connected = true

		d.onConnectionUpdate(berrylan.WirelessConnectionStatusActivated)
	}()

	return nil
}

// Disconnect implements berrylan.WirelessInterface.Disconnect by simulating a
// disconnection request.
func (d *DummyWirelessInterface) Disconnect() error {
	log.Infof("Disconnecting from network.")

	d.ssid = ""

	go func() {
		d.connectionStatusUpdateHandler(
			berrylan.WirelessConnectionStatusDeactivating)
		time.Sleep(1 * time.Second)

		d.connected = false

		d.connectionStatusUpdateHandler(
			berrylan.WirelessConnectionStatusDisconnected)
	}()

	return nil
}

// HandleConnectionStatusUpdate implements
// berrylan.WirelessInterface.HandleConnectionStatusUpdate by storing a
// reference to an event handler function.
func (d *DummyWirelessInterface) HandleConnectionStatusUpdate(f berrylan.ConnectionStatusUpdateHandler) {
	d.connectionStatusUpdateHandler = f
}
