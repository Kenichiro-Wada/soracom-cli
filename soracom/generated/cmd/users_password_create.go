package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersPasswordCreateCmdOperatorId holds value of 'operator_id' option
var UsersPasswordCreateCmdOperatorId string

// UsersPasswordCreateCmdPassword holds value of 'password' option
var UsersPasswordCreateCmdPassword string

// UsersPasswordCreateCmdUserName holds value of 'user_name' option
var UsersPasswordCreateCmdUserName string

// UsersPasswordCreateCmdBody holds contents of request body to be sent
var UsersPasswordCreateCmdBody string

func init() {
	UsersPasswordCreateCmd.Flags().StringVar(&UsersPasswordCreateCmdOperatorId, "operator-id", "", TR("operator_id"))

	UsersPasswordCreateCmd.Flags().StringVar(&UsersPasswordCreateCmdPassword, "password", "", TR(""))

	UsersPasswordCreateCmd.Flags().StringVar(&UsersPasswordCreateCmdUserName, "user-name", "", TR("user_name"))

	UsersPasswordCreateCmd.Flags().StringVar(&UsersPasswordCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	UsersPasswordCmd.AddCommand(UsersPasswordCreateCmd)
}

// UsersPasswordCreateCmd defines 'create' subcommand
var UsersPasswordCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TR("users.create_user_password.post.summary"),
	Long:  TR(`users.create_user_password.post.description`),
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

		param, err := collectUsersPasswordCreateCmdParams()
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

func collectUsersPasswordCreateCmdParams() (*apiParams, error) {

	body, err := buildBodyForUsersPasswordCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForUsersPasswordCreateCmd("/operators/{operator_id}/users/{user_name}/password"),
		query:       buildQueryForUsersPasswordCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForUsersPasswordCreateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersPasswordCreateCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersPasswordCreateCmdUserName, -1)

	return path
}

func buildQueryForUsersPasswordCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForUsersPasswordCreateCmd() (string, error) {
	if UsersPasswordCreateCmdBody != "" {
		if strings.HasPrefix(UsersPasswordCreateCmdBody, "@") {
			fname := strings.TrimPrefix(UsersPasswordCreateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if UsersPasswordCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return UsersPasswordCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if UsersPasswordCreateCmdPassword != "" {
		result["password"] = UsersPasswordCreateCmdPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
