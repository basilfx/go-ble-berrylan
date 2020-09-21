package main

import (
	"errors"

	berrylan "github.com/basilfx/go-ble-berrylan"
	log "github.com/sirupsen/logrus"
)

// DummyNetworkInterface represents a dummy network interface.
type DummyNetworkInterface struct {
	networkStatusUpdateHandler   berrylan.NetworkStatusUpdateHandler
	networkingStateUpdateHandler berrylan.NetworkingStateUpdateHandler
	wirelessStateUpdateHandler   berrylan.WirelessStateUpdateHandler
}

// NewDummyNetworkInterface initializes a new instance of DummyNetworkInterface
// that does nothing.
func NewDummyNetworkInterface() *DummyNetworkInterface {
	return &DummyNetworkInterface{}
}

func (n *DummyNetworkInterface) onStatusUpdate(s berrylan.NetworkStatus) {
	log.Infof("Changing network status to '%s'.", s.String())

	if n.networkStatusUpdateHandler != nil {
		n.networkStatusUpdateHandler(s)
	}
}

func (n *DummyNetworkInterface) onNetworkingEnabledUpdate(e bool) {
	if n.networkingStateUpdateHandler != nil {
		n.networkingStateUpdateHandler(e)
	}
}

func (n *DummyNetworkInterface) onWirelessEnabledUpdate(e bool) {
	if n.wirelessStateUpdateHandler != nil {
		n.wirelessStateUpdateHandler(e)
	}
}

// EnableNetworking implements berrylan.NetworkInterface.EnableNetworking by
// returning an error because it is not supported.
func (n *DummyNetworkInterface) EnableNetworking(enabled bool) error {
	return errors.New("not supported")
}

// EnableWireless implements berrylan.NetworkInterface.EnableWireless by
// returning an error because it is not supported.
func (n *DummyNetworkInterface) EnableWireless(enabled bool) error {
	return errors.New("not supported")
}

// HandleNetworkStatusUpdate implements
// berrylan.NetworkInterface.HandleNetworkStatusUpdate by storing a reference
// to an event handler function.
func (n *DummyNetworkInterface) HandleNetworkStatusUpdate(f berrylan.NetworkStatusUpdateHandler) {
	n.networkStatusUpdateHandler = f
}

// HandleNetworkingStateUpdate implements
// berrylan.NetworkInterface.HandleNetworkingStateUpdate by storing a
// reference to an event handler function.
func (n *DummyNetworkInterface) HandleNetworkingStateUpdate(f berrylan.NetworkingStateUpdateHandler) {
	n.networkingStateUpdateHandler = f
}

// HandleWirelessStateUpdate implements
// berrylan.NetworkInterface.HandleWirelessStateUpdate by storing a reference
// to an event handler function.
func (n *DummyNetworkInterface) HandleWirelessStateUpdate(f berrylan.WirelessStateUpdateHandler) {
	n.wirelessStateUpdateHandler = f
}
