//go:build !test
// +build !test

package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chanify/chanify/core"
	"github.com/chanify/chanify/logic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Launch chanify api server",
		Long:  `Launch service for chanify api server.`,
		Run: func(cmd *cobra.Command, args []string) {
			srv := &http.Server{
				Addr:           fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")),
				ReadTimeout:    parseDuration(viper.GetString("server.http.readtimeout")),
				WriteTimeout:   parseDuration(viper.GetString("server.http.writetimeout")),
				MaxHeaderBytes: 1 << 20,
			}
			log.Println("Launching service...")
			if c := viper.ConfigFileUsed(); len(c) > 0 {
				log.Println("Using config file:", c)
			}
			go func() {
				c := core.New()
				if c == nil {
					log.Fatalln("Create service failed!")
					return
				}
				defer c.Close()
				endpoint := getEndpoint()
				opts := &logic.Options{
					Name:       getName(),
					Version:    Version,
					Endpoint:   endpoint,
					DataPath:   getExpandPath("server.datapath"),
					FilePath:   getExpandPath("server.filepath"),
					PluginPath: getExpandPath("server.pluginpath"),
					DBUrl:      viper.GetString("server.dburl"),
					Secret:     viper.GetString("server.secret"),
					WebHooks:   getWebhooks(),
				}
				opts.Registerable, opts.RegUsers = getUserWhitlist(cmd)
				if err := c.Init(opts); err != nil {
					log.Fatalln("Init service failed:", err)
					return
				}
				srv.Handler = c.APIHandler()
				log.Println("Launch service", srv.Addr)
				log.Println("Node server endpoint:", endpoint)
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Println("Launch service failed:", err)
				}
			}()
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutting down server...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Println("Shutdown service failed:", err)
			}
			log.Println("Shutdown service success.")
		},
	}
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().String("host", "0.0.0.0", "Http restful service hostname")
	serveCmd.Flags().Int("port", 80, "Http restful service port")
	serveCmd.Flags().String("endpoint", "", "Http restful service endpoint")
	serveCmd.Flags().String("name", "", "Http service name")
	serveCmd.Flags().String("datapath", "~/.chanify", "Data file path")
	serveCmd.Flags().String("filepath", "", "Store file path")
	serveCmd.Flags().String("pluginpath", "~/.chanify/plugin", "Plugin file path")
	serveCmd.Flags().String("dburl", "", "Databse dsn uri")
	serveCmd.Flags().String("secret", "", "Secret key for serverless mode")
	serveCmd.Flags().String("readtimeout", "10s", "Http restful service read timeout.")
	serveCmd.Flags().String("writetimeout", "10s", "Http restful service read timeout.")
	serveCmd.Flags().Bool("registerable", true, "Enable register user")
	serveCmd.Flags().String("whitelist", "", "Whitelist for register users")
	viper.BindPFlag("server.host", serveCmd.Flags().Lookup("host"))                      // nolint: errcheck
	viper.BindPFlag("server.port", serveCmd.Flags().Lookup("port"))                      // nolint: errcheck
	viper.BindPFlag("server.endpoint", serveCmd.Flags().Lookup("endpoint"))              // nolint: errcheck
	viper.BindPFlag("server.name", serveCmd.Flags().Lookup("name"))                      // nolint: errcheck
	viper.BindPFlag("server.datapath", serveCmd.Flags().Lookup("datapath"))              // nolint: errcheck
	viper.BindPFlag("server.filepath", serveCmd.Flags().Lookup("filepath"))              // nolint: errcheck
	viper.BindPFlag("server.pluginpath", serveCmd.Flags().Lookup("pluginpath"))          // nolint: errcheck
	viper.BindPFlag("server.dburl", serveCmd.Flags().Lookup("dburl"))                    // nolint: errcheck
	viper.BindPFlag("server.secret", serveCmd.Flags().Lookup("secret"))                  // nolint: errcheck
	viper.BindPFlag("server.http.readtimeout", serveCmd.Flags().Lookup("readtimeout"))   // nolint: errcheck
	viper.BindPFlag("server.http.writetimeout", serveCmd.Flags().Lookup("writetimeout")) // nolint: errcheck
	viper.BindPFlag("server.register.enable", serveCmd.Flags().Lookup("registerable"))   // nolint: errcheck
}

func getName() string { _ = "STUB: not implemented"; return "" }

func getExpandPath(key string) string { _ = "STUB: not implemented"; return "" }

func getEndpoint() string { _ = "STUB: not implemented"; return "" }

func getUserWhitlist(cmd *cobra.Command) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func getWebhooks() []map[string]interface{} { _ = "STUB: not implemented"; return nil }

func parseDuration(dur string) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }
