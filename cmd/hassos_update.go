package cmd

import (
	"fmt"

	helper "github.com/home-assistant/hassio-cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var hassosUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"up"},
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("args", args).Debug("hassos update")

		section := "hassos"
		command := "update"
		base := viper.GetString("endpoint")

		var options map[string]interface{}

		version, err := cmd.Flags().GetString("version")
		if version != "" {
			options = map[string]interface{}{"version": version}
		}

		resp, err := helper.GenericJSONPost(base, section, command, options)
		if err != nil {
			fmt.Println(err)
			ExitWithError = true
		} else {
			ExitWithError = !helper.ShowJSONResponse(resp)
		}

		return
	},
}

func init() {
	hassosUpdateCmd.Flags().StringP("version", "", "", "Version to update to")
	hassosCmd.AddCommand(hassosUpdateCmd)
}
