package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersUnsetGroupCmdImsi string





func init() {
  SubscribersUnsetGroupCmd.Flags().StringVar(&SubscribersUnsetGroupCmdImsi, "imsi", "", "対象のSubscriberのIMSI")




  SubscribersCmd.AddCommand(SubscribersUnsetGroupCmd)
}

var SubscribersUnsetGroupCmd = &cobra.Command{
  Use: "unset-group",
  Short: TR("Unset Group to Subscriber"),
  Long: TR(`指定されたSubscriberのGroup指定を解除`),
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
    
    param, err := collectSubscribersUnsetGroupCmdParams()
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

func collectSubscribersUnsetGroupCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersUnsetGroupCmd("/subscribers/{imsi}/unset_group"),
    query: buildQueryForSubscribersUnsetGroupCmd(),
    
    
  }, nil
}

func buildPathForSubscribersUnsetGroupCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersUnsetGroupCmdImsi, -1)
  
  
  
  
  return path
}

func buildQueryForSubscribersUnsetGroupCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


