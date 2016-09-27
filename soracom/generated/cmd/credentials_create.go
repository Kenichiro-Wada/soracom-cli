package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// CredentialsCreateCmdCredentialsId holds value of 'credentials_id' option
var CredentialsCreateCmdCredentialsId string

// CredentialsCreateCmdDescription holds value of 'description' option
var CredentialsCreateCmdDescription string

// CredentialsCreateCmdType holds value of 'type' option
var CredentialsCreateCmdType string

// CredentialsCreateCmdBody holds contents of request body to be sent
var CredentialsCreateCmdBody string

func init() {
	CredentialsCreateCmd.Flags().StringVar(&CredentialsCreateCmdCredentialsId, "credentials-id", "", TR("credentials_id"))

	CredentialsCreateCmd.Flags().StringVar(&CredentialsCreateCmdDescription, "description", "", TR(""))

	CredentialsCreateCmd.Flags().StringVar(&CredentialsCreateCmdType, "type", "", TR(""))

	CredentialsCreateCmd.Flags().StringVar(&CredentialsCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	CredentialsCmd.AddCommand(CredentialsCreateCmd)
}

// CredentialsCreateCmd defines 'create' subcommand
var CredentialsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TR("credentials.create_credential.post.summary"),
	Long:  TR(`credentials.create_credential.post.description`),
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

		param, err := collectCredentialsCreateCmdParams()
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

func collectCredentialsCreateCmdParams() (*apiParams, error) {

	body, err := buildBodyForCredentialsCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForCredentialsCreateCmd("/credentials/{credentials_id}"),
		query:       buildQueryForCredentialsCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForCredentialsCreateCmd(path string) string {

	path = strings.Replace(path, "{"+"credentials_id"+"}", CredentialsCreateCmdCredentialsId, -1)

	return path
}

func buildQueryForCredentialsCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForCredentialsCreateCmd() (string, error) {
	if CredentialsCreateCmdBody != "" {
		if strings.HasPrefix(CredentialsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(CredentialsCreateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if CredentialsCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return CredentialsCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if CredentialsCreateCmdDescription != "" {
		result["description"] = CredentialsCreateCmdDescription
	}

	if CredentialsCreateCmdType != "" {
		result["type"] = CredentialsCreateCmdType
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
