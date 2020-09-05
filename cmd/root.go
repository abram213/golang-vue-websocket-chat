package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "Chat",
	Short: "Chat Web Application",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		/*logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		})*/
		/*if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}*/
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
		//logrus.Fatal(err)
	}
}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().BoolP("verbose", "v", false, "make output more verbose")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.example.yaml)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/chat")
		viper.AddConfigPath("$HOME/.chat")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Warnf("unable to read config from file")
	}
}
