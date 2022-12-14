/*
Copyright © 2022 Atom Pi <coder.atompi@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"gitee.com/autom-studio/alert-feishu/api/webhook/router"
	"gitee.com/autom-studio/alert-feishu/internal/config"
	logkit "gitee.com/autom-studio/go-kits/log"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "alert-feishu",
	Short: "Sent alert message to Feishu",
	Long:  `Integration of Feishu-robot for Prometheus Alertmanager via webhook.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var rootConfig config.RootConfigStruct
		err := viper.Unmarshal(&rootConfig)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unmarshal config failed: ", err)
			os.Exit(1)
		}

		logPath := rootConfig.Main.Log.Path
		logLevel := rootConfig.Main.Log.Level
		logger := logkit.InitLogger(logPath, logLevel)
		defer logger.Sync()
		undo := zap.ReplaceGlobals(logger)
		defer undo()

		listenAddr := viper.GetString("main.listenAddress")
		gin.SetMode(viper.GetString("main.ginMode"))
		r := gin.New()
		r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
		r.Use(ginzap.RecoveryWithZap(logger, true))
		router.Register(r, rootConfig)
		r.Run(listenAddr)
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./alert_feishu.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".alert-feishu" (without extension).
		viper.AddConfigPath("./")
		viper.SetConfigType("yaml")
		viper.SetConfigName("alert_feishu")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Fprintln(os.Stderr, "Init config file failed:", err)
	}
}
