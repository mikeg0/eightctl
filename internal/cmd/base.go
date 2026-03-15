package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/steipete/eightctl/internal/client"
	"github.com/steipete/eightctl/internal/output"
)

var baseCmd = &cobra.Command{Use: "base", Short: "Adjustable base controls"}

var baseInfoCmd = &cobra.Command{Use: "info", Short: "Get base state", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	res, err := cl.Base().Info(context.Background())
	if err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"info"}, []map[string]any{{"info": res}})
}}

var baseAngleCmd = &cobra.Command{Use: "angle", Short: "Set base angles", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	head, _ := cmd.Flags().GetInt("head")
	foot, _ := cmd.Flags().GetInt("foot")
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	return cl.Base().SetAngle(context.Background(), head, foot)
}}

func init() {
	baseAngleCmd.Flags().Int("head", 0, "head angle")
	baseAngleCmd.Flags().Int("foot", 0, "foot angle")
	baseCmd.AddCommand(baseInfoCmd, baseAngleCmd)
}
