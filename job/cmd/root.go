package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rtb-job",
	Short: "rtb job",
	Long:  `rtb job`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.example.yaml)")
}

func initConfig() {
	//err := config.Load(
	//	file.NewSource(
	//		file.WithPath("./config/"+env.Mode()+"/mysql.yaml"),
	//		source.WithEncoder(yaml.NewEncoder()),
	//	),
	//)
	//if err != nil {
	//	log.Fatalf("load config file fail: %s", err.Error())
	//}
}
