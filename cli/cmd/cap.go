package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wroge/wms/getmap"
)

var getcapCommand = &cobra.Command{
	Use:     "cap",
	Aliases: []string{"getcap"},
	Args:    cobra.RangeArgs(0, 1),
	Short:   "Get the capabilities of a WMS",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		w := &getmap.Service{}
		var service = "default"
		if len(args) == 1 {
			service = args[0]
		}
		if cmd.Flag("user").Changed {
			user, err := cmd.Flags().GetString("user")
			if err != nil {
				return err
			}
			w.User = user
		}
		if cmd.Flag("password").Changed {
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				return err
			}
			w.Password = password
		}
		url := viper.GetString(service + ".url")
		if cmd.Flag("url").Changed {
			url, err = cmd.Flags().GetString("url")
			if err != nil {
				return
			}
		}
		err = w.AddURL(url)
		if err != nil {
			return
		}
		version := viper.GetString(service + ".version")
		if cmd.Flag("version").Changed {
			version, err = cmd.Flags().GetString("version")
			if err != nil {
				return
			}
		}
		err = w.AddVersion(version)
		if err != nil {
			return
		}
		f, err := cmd.Flags().GetBool("formats")
		if err != nil {
			return
		}
		l, err := cmd.Flags().GetBool("layers")
		if err != nil {
			return
		}
		e, err := cmd.Flags().GetBool("epsg")
		if err != nil {
			return
		}
		c, err := w.GetCapabilities(w.User, w.Password)
		if err != nil {
			return
		}
		if !f && !l && !e {
			fmt.Println(c)
		}
		if f {
			fmt.Println(c.Formats)
		}
		if l {
			fmt.Println(c.GetLayerNames())
		}
		if e {
			fmt.Println(c.GetBBoxes().GetEPSG())
		}
		return
	},
}

func init() {
	getcapCommand.Flags().StringP("url", "u", "", "Set url")
	getcapCommand.Flags().StringP("version", "v", "", "Set version")

	getcapCommand.Flags().BoolP("formats", "f", false, "Get available formats")
	getcapCommand.Flags().BoolP("layers", "l", false, "Get available layers")
	getcapCommand.Flags().BoolP("epsg", "e", false, "Get available epsg-codes")

	getcapCommand.Flags().StringP("user", "", "", "Set user")
	getcapCommand.Flags().StringP("password", "", "", "Set password")

	root.AddCommand(getcapCommand)
}
