package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ShippingAddressesCreateCmdAddressLine1 holds value of 'addressLine1' option
var ShippingAddressesCreateCmdAddressLine1 string

// ShippingAddressesCreateCmdAddressLine2 holds value of 'addressLine2' option
var ShippingAddressesCreateCmdAddressLine2 string

// ShippingAddressesCreateCmdBuilding holds value of 'building' option
var ShippingAddressesCreateCmdBuilding string

// ShippingAddressesCreateCmdCity holds value of 'city' option
var ShippingAddressesCreateCmdCity string

// ShippingAddressesCreateCmdCompanyName holds value of 'companyName' option
var ShippingAddressesCreateCmdCompanyName string

// ShippingAddressesCreateCmdDepartment holds value of 'department' option
var ShippingAddressesCreateCmdDepartment string

// ShippingAddressesCreateCmdFullName holds value of 'fullName' option
var ShippingAddressesCreateCmdFullName string

// ShippingAddressesCreateCmdOperatorId holds value of 'operator_id' option
var ShippingAddressesCreateCmdOperatorId string

// ShippingAddressesCreateCmdPhoneNumber holds value of 'phoneNumber' option
var ShippingAddressesCreateCmdPhoneNumber string

// ShippingAddressesCreateCmdState holds value of 'state' option
var ShippingAddressesCreateCmdState string

// ShippingAddressesCreateCmdZipCode holds value of 'zipCode' option
var ShippingAddressesCreateCmdZipCode string

// ShippingAddressesCreateCmdBody holds contents of request body to be sent
var ShippingAddressesCreateCmdBody string

func init() {
	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdAddressLine1, "address-line1", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdAddressLine2, "address-line2", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdBuilding, "building", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdCity, "city", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdCompanyName, "company-name", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdDepartment, "department", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdFullName, "full-name", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdOperatorId, "operator-id", "", TR("operator_id"))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdPhoneNumber, "phone-number", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdState, "state", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdZipCode, "zip-code", "", TR(""))

	ShippingAddressesCreateCmd.Flags().StringVar(&ShippingAddressesCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	ShippingAddressesCmd.AddCommand(ShippingAddressesCreateCmd)
}

// ShippingAddressesCreateCmd defines 'create' subcommand
var ShippingAddressesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TR("shipping_addresses.create_shipping_address.post.summary"),
	Long:  TR(`shipping_addresses.create_shipping_address.post.description`),
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

		param, err := collectShippingAddressesCreateCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectShippingAddressesCreateCmdParams() (*apiParams, error) {

	body, err := buildBodyForShippingAddressesCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForShippingAddressesCreateCmd("/operators/{operator_id}/shipping_addresses"),
		query:       buildQueryForShippingAddressesCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForShippingAddressesCreateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", ShippingAddressesCreateCmdOperatorId, -1)

	return path
}

func buildQueryForShippingAddressesCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForShippingAddressesCreateCmd() (string, error) {
	if ShippingAddressesCreateCmdBody != "" {
		if strings.HasPrefix(ShippingAddressesCreateCmdBody, "@") {
			fname := strings.TrimPrefix(ShippingAddressesCreateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if ShippingAddressesCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return ShippingAddressesCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if ShippingAddressesCreateCmdAddressLine1 != "" {
		result["addressLine1"] = ShippingAddressesCreateCmdAddressLine1
	}

	if ShippingAddressesCreateCmdAddressLine2 != "" {
		result["addressLine2"] = ShippingAddressesCreateCmdAddressLine2
	}

	if ShippingAddressesCreateCmdBuilding != "" {
		result["building"] = ShippingAddressesCreateCmdBuilding
	}

	if ShippingAddressesCreateCmdCity != "" {
		result["city"] = ShippingAddressesCreateCmdCity
	}

	if ShippingAddressesCreateCmdCompanyName != "" {
		result["companyName"] = ShippingAddressesCreateCmdCompanyName
	}

	if ShippingAddressesCreateCmdDepartment != "" {
		result["department"] = ShippingAddressesCreateCmdDepartment
	}

	if ShippingAddressesCreateCmdFullName != "" {
		result["fullName"] = ShippingAddressesCreateCmdFullName
	}

	if ShippingAddressesCreateCmdPhoneNumber != "" {
		result["phoneNumber"] = ShippingAddressesCreateCmdPhoneNumber
	}

	if ShippingAddressesCreateCmdState != "" {
		result["state"] = ShippingAddressesCreateCmdState
	}

	if ShippingAddressesCreateCmdZipCode != "" {
		result["zipCode"] = ShippingAddressesCreateCmdZipCode
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
