package icons

import (
	"github.com/mgorunuch/environment-launcher/bindata"
)

func GetShuttleIcon() ([]byte, error) {
	return bindata.Asset("static/icons/shuttle.png")
}

func GetExitIcon() ([]byte, error) {
	return bindata.Asset("static/icons/exit.png")
}
