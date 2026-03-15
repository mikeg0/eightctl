package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/steipete/eightctl/internal/client"
	"github.com/steipete/eightctl/internal/output"
)

var alarmCmd = &cobra.Command{
	Use:   "alarm",
	Short: "Manage alarms",
}

var alarmListCmd = &cobra.Command{
	Use:   "list",
	Short: "List alarms/routines",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireAuthFields(); err != nil {
			return err
		}
		cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
		res, err := cl.ListAlarms(context.Background())
		if err != nil {
			return err
		}
		return output.Print(output.Format(viper.GetString("output")), []string{"routines"}, []map[string]any{{"routines": res}})
	},
}

func init() {
	alarmCmd.AddCommand(alarmListCmd, alarmSnoozeCmd, alarmDismissCmd, alarmDismissAllCmd, alarmVibeCmd)
}

// snooze
var alarmSnoozeCmd = &cobra.Command{Use: "snooze <id>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	return cl.Alarms().Snooze(context.Background(), args[0])
}}

var alarmDismissCmd = &cobra.Command{Use: "dismiss <id>", Args: cobra.ExactArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	if err := cl.Alarms().Dismiss(context.Background(), args[0]); err != nil {
		return err
	}
	fmt.Printf("Alarm %s dismissed.\n", args[0])
	return nil
}}

var alarmDismissAllCmd = &cobra.Command{Use: "dismiss-all", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	if err := cl.Alarms().DismissAll(context.Background()); err != nil {
		return err
	}
	fmt.Println("All alarms dismissed.")
	return nil
}}

var alarmVibeCmd = &cobra.Command{Use: "vibration-test", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	return cl.Alarms().VibrationTest(context.Background())
}}

