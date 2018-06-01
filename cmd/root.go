package cmd

import (
	"fmt"
	"os"

	"github.com/rajatjindal/wordpress-ios-debug/pkg/blog"
	"github.com/rajatjindal/wordpress-ios-debug/pkg/dnsserver"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wordpress-ios-debug",
	Short: "A brief description of your application",

	Run: func(cmd *cobra.Command, args []string) {
		go blog.ServeBlog()
		dnsserver.ServeDNS()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
