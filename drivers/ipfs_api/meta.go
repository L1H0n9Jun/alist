package ipfs

import (
	"alist/internal/driver"
	"alist/internal/op"
)

type Addition struct {
	// Usually one of two
	driver.RootPath
	Endpoint string `json:"endpoint" default:"http://127.0.0.1:5001"`
	Gateway  string `json:"gateway" default:"https://ipfs.io"`
}

var config = driver.Config{
	Name:        "IPFS API",
	DefaultRoot: "/",
	LocalSort:   true,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &IPFS{}
	})
}
