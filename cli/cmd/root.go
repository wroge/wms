package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var root = &cobra.Command{
	Use: "wms",
	Long: `This program helps you to generate images via web map services.

Configuration file: $HOME/wms-config/.wms.yaml
Further informations and examples: https://github.com/wroge/wms`,
	SilenceUsage: true,
}

var version = "No Version Provided"

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	root.PersistentFlags().Bool("help", false, "Help about any command")
	root.AddCommand(versionCommand)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.AddConfigPath(filepath.Join(home, "wms-config"))
	viper.SetConfigName(".wms")
	if err := viper.ReadInConfig(); err != nil {
		if _, err2 := os.Stat(filepath.Join(home, "wms-config", ".wms.yaml")); os.IsNotExist(err2) {
			err = os.Mkdir(filepath.Join(home, "wms-config"), 0777)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			_, err = os.Create(filepath.Join(home, "wms-config", ".wms.yaml"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}
	}
}

// Execute root command
func Execute(v string) {
	version = v
	err := root.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
