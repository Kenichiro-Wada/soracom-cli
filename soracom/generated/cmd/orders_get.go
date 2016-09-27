package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OrdersGetCmdOrderId holds value of 'order_id' option
var OrdersGetCmdOrderId string

func init() {
	OrdersGetCmd.Flags().StringVar(&OrdersGetCmdOrderId, "order-id", "", TR("order_id"))

	OrdersCmd.AddCommand(OrdersGetCmd)
}

// OrdersGetCmd defines 'get' subcommand
var OrdersGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("orders.get_order.get.summary"),
	Long:  TR(`orders.get_order.get.description`),
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

		param, err := collectOrdersGetCmdParams()
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

func collectOrdersGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForOrdersGetCmd("/orders/{order_id}"),
		query:  buildQueryForOrdersGetCmd(),
	}, nil
}

func buildPathForOrdersGetCmd(path string) string {

	path = strings.Replace(path, "{"+"order_id"+"}", OrdersGetCmdOrderId, -1)

	return path
}

func buildQueryForOrdersGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
