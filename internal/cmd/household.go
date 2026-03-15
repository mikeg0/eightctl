package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/steipete/eightctl/internal/client"
	"github.com/steipete/eightctl/internal/output"
)

var householdCmd = &cobra.Command{Use: "household", Short: "Household info"}

var householdSummaryCmd = &cobra.Command{Use: "summary", Short: "Show household summary", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	res, err := cl.Household().Summary(context.Background())
	if err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"summary"}, []map[string]any{{"summary": res}})
}}

func init() {
	householdCmd.AddCommand(householdSummaryCmd)
}
