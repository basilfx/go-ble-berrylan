module github.com/basilfx/go-ble-berrylan/examples/dummy

go 1.13

replace github.com/basilfx/go-ble-berrylan => ../../

require (
	github.com/basilfx/go-ble-berrylan v0.0.0-00010101000000-000000000000
	github.com/basilfx/go-ble-device-information v0.0.0-20200921160719-46f83e527d78
	github.com/basilfx/go-ble-utilities v0.0.0-20200920114255-307344fe7cc5
	github.com/go-ble/ble v0.0.0-20200407180624-067514cd6e24
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
)
