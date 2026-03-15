package cmd

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/steipete/eightctl/internal/client"
	"github.com/steipete/eightctl/internal/output"
)

var metricsCmd = &cobra.Command{Use: "metrics", Short: "Sleep metrics and insights"}

var metricsTrendsCmd = &cobra.Command{Use: "trends", Short: "Fetch daily sleep trends", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")
	tz := viper.GetString("timezone")
	if tz == "local" {
		tz = time.Local.String()
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	var out any
	if err := cl.Metrics().Trends(context.Background(), from, to, tz, &out); err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"trends"}, []map[string]any{{"trends": out}})
}}

var metricsIntervalsCmd = &cobra.Command{Use: "intervals", Short: "Fetch recent sleep intervals with timeseries", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cursor, _ := cmd.Flags().GetString("cursor")
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	var out any
	if err := cl.Metrics().Intervals(context.Background(), cursor, &out); err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"intervals"}, []map[string]any{{"intervals": out}})
}}

var metricsSummaryCmd = &cobra.Command{Use: "summary", Short: "Fetch metrics summary", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	var out any
	if err := cl.Metrics().Summary(context.Background(), &out); err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"summary"}, []map[string]any{{"summary": out}})
}}

var metricsAggregateCmd = &cobra.Command{Use: "aggregate", Short: "Fetch aggregated metrics", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	var out any
	if err := cl.Metrics().Aggregate(context.Background(), &out); err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"aggregate"}, []map[string]any{{"aggregate": out}})
}}

var metricsInsightsCmd = &cobra.Command{Use: "insights", Short: "Fetch daily sleep insights", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	date, _ := cmd.Flags().GetString("date")
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	var out any
	if err := cl.Metrics().Insights(context.Background(), date, &out); err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"insights"}, []map[string]any{{"insights": out}})
}}

var metricsLLMInsightsCmd = &cobra.Command{Use: "llm-insights", Short: "Fetch AI-generated sleep insights", RunE: func(cmd *cobra.Command, args []string) error {
	if err := requireAuthFields(); err != nil {
		return err
	}
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")
	cl := client.New(viper.GetString("email"), viper.GetString("password"), viper.GetString("user_id"), viper.GetString("client_id"), viper.GetString("client_secret"))
	var out any
	if err := cl.Metrics().LLMInsights(context.Background(), from, to, &out); err != nil {
		return err
	}
	return output.Print(output.Format(viper.GetString("output")), []string{"insights"}, []map[string]any{{"insights": out}})
}}

func init() {
	metricsTrendsCmd.Flags().String("from", "", "from date YYYY-MM-DD")
	metricsTrendsCmd.Flags().String("to", "", "to date YYYY-MM-DD")

	metricsIntervalsCmd.Flags().String("cursor", "", "pagination cursor from previous response")

	metricsInsightsCmd.Flags().String("date", "", "date YYYY-MM-DD (optional)")

	metricsLLMInsightsCmd.Flags().String("from", "", "from date YYYY-MM-DD")
	metricsLLMInsightsCmd.Flags().String("to", "", "to date YYYY-MM-DD")

	metricsCmd.AddCommand(metricsTrendsCmd, metricsIntervalsCmd, metricsSummaryCmd, metricsAggregateCmd, metricsInsightsCmd, metricsLLMInsightsCmd)
}
