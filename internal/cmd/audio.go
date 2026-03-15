package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/steipete/eightctl/internal/client"
	"github.com/steipete/eightctl/internal/output"
)

var audioCmd = &cobra.Command{Use: "audio", Short: "Audio tracks and categories"}

var audioTracksCmd = &cobra.Command{Use: "tracks", Short: "List audio tracks", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	tracks, err := cl.Audio().Tracks(context.Background())
	if err != nil {
		return err
	}
	rows := make([]map[string]any, 0, len(tracks))
	for _, t := range tracks {
		rows = append(rows, map[string]any{"id": t.ID, "title": t.Title, "type": t.Type})
	}
	rows = output.FilterFields(rows, viper.GetStringSlice("fields"))
	headers := viper.GetStringSlice("fields")
	if len(headers) == 0 {
		headers = []string{"id", "title", "type"}
	}
	return output.Print(output.Format(viper.GetString("output")), headers, rows)
}}

var audioCategoriesCmd = &cobra.Command{Use: "categories", Short: "List audio categories", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	res, err := cl.Audio().Categories(context.Background())
	if err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"categories"}, []map[string]any{{"categories": res}})
}}

func init() {
	audioCmd.AddCommand(audioTracksCmd, audioCategoriesCmd)
}
