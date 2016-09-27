package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	BillsCmd.AddCommand(BillsListCmd)
}

// BillsListCmd defines 'list' subcommand
var BillsListCmd = &cobra.Command{
	Use:   "list",
	Short: TR("bills.get_billing_history.get.summary"),
	Long:  TR(`bills.get_billing_history.get.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			Endpoint: getSpecifiedEndpoint(),
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectBillsListCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result != "" {
			return prettyPrintStringAsJSON(result)
		} else {
			return nil
		}
	},
}

func collectBillsListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForBillsListCmd("/bills"),
		query:  buildQueryForBillsListCmd(),
	}, nil
}

func buildPathForBillsListCmd(path string) string {

	return path
}

func buildQueryForBillsListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
