package skyline

import (
	"Skyline/internal/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "Skyline",
	Version: "0.0.0 unstable",
	Short:   "Skyline is a web app for air travels and developed by Mahdi Ghasemi",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		appSetting, err := utils.LoadAppConfig("./internal/configs")
		if err != nil {
			panic("AppSetting cannot be loaded !")
		}

		err = utils.InitDB(appSetting.DbConnection)
		if err != nil {
			panic("Cannot connected to the database!")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Skyline.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
