package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/9506hqwy/redmine-mcp-server/pkg/redmine"
)

var version = "<version>"
var commit = "<commit>"

func fromArgument(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, redmine.UrlKey{}, viper.GetString("url"))
	ctx = context.WithValue(ctx, redmine.UserKey{}, viper.GetString("user"))
	ctx = context.WithValue(ctx, redmine.PasswordKey{}, viper.GetString("password"))
	ctx = context.WithValue(ctx, redmine.ApiKeyKey{}, viper.GetString("apikey"))
	return ctx
}

var rootCmd = &cobra.Command{
	Use:     "redmine-mcp-server",
	Short:   "Redmine MCP Server",
	Long:    "Redmine MCP Server",
	Version: fmt.Sprintf("%s\nCommit: %s", version, commit),
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewMCPServer(
			"Redmine MCP Server",
			"0.1.0",
			server.WithToolCapabilities(false),
		)

		redmine.RegisterTools(s, viper.GetBool("readonly"))

		if err := server.ServeStdio(s, server.WithStdioContextFunc(fromArgument)); err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Fatalf("Server error: %v", err)
			}
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().String("url", "http://127.0.0.1:3000", "Redmine server URL.")
	rootCmd.PersistentFlags().String("user", "", "Redmine server username.")
	rootCmd.PersistentFlags().String("password", "", "Redmine server password.")
	rootCmd.PersistentFlags().String("apikey", "", "Redmine server API key.")
	rootCmd.PersistentFlags().Bool("readonly", true, "HTTP GET method only.")

	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))
	viper.BindPFlag("readonly", rootCmd.PersistentFlags().Lookup("readonly"))
}

func initConfig() {
	viper.SetEnvPrefix("redmine")
	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
