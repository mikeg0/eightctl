package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/steipete/eightctl/internal/client"
	"github.com/steipete/eightctl/internal/output"
)

var travelCmd = &cobra.Command{Use: "travel", Short: "Travel / jetlag endpoints"}

var travelTripsCmd = &cobra.Command{Use: "trips", Short: "List travel trips", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	res, err := cl.Travel().Trips(context.Background())
	if err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"trips"}, []map[string]any{{"trips": res}})
}}

func init() {
	travelCmd.AddCommand(travelTripsCmd)
}
