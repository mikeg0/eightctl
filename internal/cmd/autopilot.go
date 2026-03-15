package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/steipete/eightctl/internal/client"
	"github.com/steipete/eightctl/internal/output"
)

var autopilotCmd = &cobra.Command{Use: "autopilot", Short: "Autopilot settings"}

var autopilotDetailsCmd = &cobra.Command{Use: "details", Short: "Show autopilot configuration", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	res, err := cl.Autopilot().Details(context.Background())
	if err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"details"}, []map[string]any{{"details": res}})
}}

func init() {
	autopilotCmd.AddCommand(autopilotDetailsCmd)
}
