package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

var VpgRegisterGatePeerCmdInnerIpAddress string

var VpgRegisterGatePeerCmdOuterIpAddress string

var VpgRegisterGatePeerCmdVpgId string

var VpgRegisterGatePeerCmdBody string

func init() {
	VpgRegisterGatePeerCmd.Flags().StringVar(&VpgRegisterGatePeerCmdInnerIpAddress, "inner-ip-address", "", TR(""))

	VpgRegisterGatePeerCmd.Flags().StringVar(&VpgRegisterGatePeerCmdOuterIpAddress, "outer-ip-address", "", TR(""))

	VpgRegisterGatePeerCmd.Flags().StringVar(&VpgRegisterGatePeerCmdVpgId, "vpg-id", "", TR("virtual_private_gateway.register_virtual_private_gateway_peer.post.parameters.vpg_id.description"))

	VpgRegisterGatePeerCmd.Flags().StringVar(&VpgRegisterGatePeerCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgRegisterGatePeerCmd)
}

var VpgRegisterGatePeerCmd = &cobra.Command{
	Use:   "register-gate-peer",
	Short: TR("virtual_private_gateway.register_virtual_private_gateway_peer.post.summary"),
	Long:  TR(`virtual_private_gateway.register_virtual_private_gateway_peer.post.description`),
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

		param, err := collectVpgRegisterGatePeerCmdParams()
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

func collectVpgRegisterGatePeerCmdParams() (*apiParams, error) {

	body, err := buildBodyForVpgRegisterGatePeerCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgRegisterGatePeerCmd("/virtual_private_gateways/{vpg_id}/gate/peers"),
		query:       buildQueryForVpgRegisterGatePeerCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgRegisterGatePeerCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgRegisterGatePeerCmdVpgId, -1)

	return path
}

func buildQueryForVpgRegisterGatePeerCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgRegisterGatePeerCmd() (string, error) {
	if VpgRegisterGatePeerCmdBody != "" {
		if strings.HasPrefix(VpgRegisterGatePeerCmdBody, "@") {
			fname := strings.TrimPrefix(VpgRegisterGatePeerCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if VpgRegisterGatePeerCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return VpgRegisterGatePeerCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if VpgRegisterGatePeerCmdInnerIpAddress != "" {
		result["innerIpAddress"] = VpgRegisterGatePeerCmdInnerIpAddress
	}

	if VpgRegisterGatePeerCmdOuterIpAddress != "" {
		result["outerIpAddress"] = VpgRegisterGatePeerCmdOuterIpAddress
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}