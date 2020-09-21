# go-ble-berrylan
Golang implementation of the BerryLan BLE server for
[go-ble](https://github.com/go-ble/ble)

## Introduction
This library implements the [BerryLan](http://www.berrylan.org/) BLE server
[protocol](https://github.com/nymea/nymea-networkmanager). It can be used to
commission WiFi settings for a(n) (IoT) device.

Although BerryLan is tailored towards
[NetworkManager](https://wiki.archlinux.org/index.php/NetworkManager), it could
be used with other networking applications too.

## Dependencies
This library makes use of [go-ble](https://github.com/go-ble/ble). Currently,
it only runs on Linux.

## Examples
An example application is included that provides a 'dummy' implementation of
a backend. It can be used together with the BerryLan application (
[Apple App Store](https://apps.apple.com/us/app/berrylan/id1436156018),
[Google Play Store](https://play.google.com/store/apps/details?id=io.guh.berrylan)).

To run this example, navigate to the `examples/dummy` folder, and run
`go run main.go`. Root access may be necessary to access the Bluetooth
peripheral.

## License
See the [`LICENSE.md`](LICENSE.md) file (MIT license).