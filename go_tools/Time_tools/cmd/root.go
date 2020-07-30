package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}



//编写完cmd/time.go文件里的time子命令后，再这里进行注册
func init(){
	rootCmd.AddCommand(timeCmd)
}
