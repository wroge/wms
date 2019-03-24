module github.com/wroge/wms

go 1.12

replace github.com/wroge/wms/getcap => ../getcap

replace github.com/wroge/wms/getmap => ../getmap

replace github.com/wroge/wms/content => ../content

replace github.com/wroge/wms/cli => ../cli

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/disintegration/imaging v1.6.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.3.0
	github.com/wroge/go-coo v0.0.4
)
