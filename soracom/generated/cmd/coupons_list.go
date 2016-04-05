package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)










func init() {



  CouponsCmd.AddCommand(CouponsListCmd)
}

var CouponsListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List coupons"),
  Long: TR(`現在登録されているクーポン一覧を返します。`),
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
    
    param, err := collectCouponsListCmdParams()
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

func collectCouponsListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForCouponsListCmd("/coupons"),
    query: buildQueryForCouponsListCmd(),
    
    
  }, nil
}

func buildPathForCouponsListCmd(path string) string {
  
  
  
  return path
}

func buildQueryForCouponsListCmd() string {
  result := []string{}
  

  

  
  return strings.Join(result, "&")
}


