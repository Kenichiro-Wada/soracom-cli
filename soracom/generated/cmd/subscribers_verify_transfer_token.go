package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersVerifyTransferTokenCmdToken string




var SubscribersVerifyTransferTokenCmdBody string


func init() {
  SubscribersVerifyTransferTokenCmd.Flags().StringVar(&SubscribersVerifyTransferTokenCmdToken, "token", "", "")



  SubscribersVerifyTransferTokenCmd.Flags().StringVar(&SubscribersVerifyTransferTokenCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  SubscribersCmd.AddCommand(SubscribersVerifyTransferTokenCmd)
}

var SubscribersVerifyTransferTokenCmd = &cobra.Command{
  Use: "verify-transfer-token",
  Short: TR("Verify Subscriber Transfer Token"),
  Long: TR(`Subscriber移管用のトークンを確認して、移管を実施する。このAPIは移管先のOperatorで呼び出しを行う。`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "/v1",
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
    
    param, err := collectSubscribersVerifyTransferTokenCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    return prettyPrintJSON(result)
  },
}

func collectSubscribersVerifyTransferTokenCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForSubscribersVerifyTransferTokenCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersVerifyTransferTokenCmd("/subscribers/transfer_token/verify"),
    query: buildQueryForSubscribersVerifyTransferTokenCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForSubscribersVerifyTransferTokenCmd(path string) string {
  
  
  
  
  
  return path
}

func buildQueryForSubscribersVerifyTransferTokenCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForSubscribersVerifyTransferTokenCmd() (string, error) {
  if SubscribersVerifyTransferTokenCmdBody != "" {
    if strings.HasPrefix(SubscribersVerifyTransferTokenCmdBody, "@") {
      fname := strings.TrimPrefix(SubscribersVerifyTransferTokenCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if SubscribersVerifyTransferTokenCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return SubscribersVerifyTransferTokenCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if SubscribersVerifyTransferTokenCmdToken != "" {
    result["token"] = SubscribersVerifyTransferTokenCmdToken
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

