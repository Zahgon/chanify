//go:build !test
// +build !test

package cmd

import (
	"bytes"
	"mime/multipart"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	sendCmd := &cobra.Command{
		Use:   "send",
		Short: "Send message cli",
		Long:  "A send message command line tool.",
		RunE:  runSendCmd,
	}
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().String("endpoint", "https://api.chanify.net", "Node server endpoint.")
	sendCmd.Flags().String("token", "", "Send token.")
	sendCmd.Flags().String("sound", "", "Message sound code.")
	sendCmd.Flags().String("text", "", "Text message content.")
	sendCmd.Flags().String("link", "", "Link message content.")
	sendCmd.Flags().String("image", "", "Image file path.")
	sendCmd.Flags().String("audio", "", "Audio file path.")
	sendCmd.Flags().String("file", "", "File path.")
	sendCmd.Flags().String("title", "", "Message title.")
	sendCmd.Flags().String("copy", "", "Copy test for text message.")
	sendCmd.Flags().String("autocopy", "", "Auto copy text for text message.")
	sendCmd.Flags().StringArray("action", []string{}, "Action item for action message.")
	sendCmd.Flags().Int("priority", 0, "Message priority.")
	sendCmd.Flags().String("interruption-level", "", "Interruption level for message.")
	sendCmd.Flags().String("timeline.code", "", "Code for timeline message.")
	sendCmd.Flags().String("timeline.timestamp", "", "Timestamp for timeline message.")
	viper.BindPFlag("client.token", sendCmd.Flags().Lookup("token"))                           // nolint: errcheck
	viper.BindPFlag("client.sound", sendCmd.Flags().Lookup("sound"))                           // nolint: errcheck
	viper.BindPFlag("client.autocopy", sendCmd.Flags().Lookup("autocopy"))                     // nolint: errcheck
	viper.BindPFlag("client.priority", sendCmd.Flags().Lookup("priority"))                     // nolint: errcheck
	viper.BindPFlag("client.endpoint", sendCmd.Flags().Lookup("endpoint"))                     // nolint: errcheck
	viper.BindPFlag("client.interruption-level", sendCmd.Flags().Lookup("interruption-level")) // nolint: errcheck
}

func runSendCmd(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func parseTimelineMessage(cmd *cobra.Command, code string, w *multipart.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func parseOtherMessage(cmd *cobra.Command, w *multipart.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func sendMessage(in *bytes.Buffer, content string) error { _ = "STUB: not implemented"; return nil }

func readFile(filePath string) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func readInputFile(path string) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func setFieldValue(w *multipart.Writer, name string, value []byte) {
	_ = "STUB: not implemented"
	return
}

// nolint: errcheck

func setFieldValues(w *multipart.Writer, name string, value []string) {
	_ = "STUB: not implemented"
	return
}

// nolint: errcheck

func setFieldValueInt(w *multipart.Writer, name string, value int) {
	_ = "STUB: not implemented"
	return
}

// nolint: errcheck

func setFieldFile(w *multipart.Writer, name string, fname string, value []byte) {
	_ = "STUB: not implemented"
	return
}

// nolint: errcheck

func fixFilename(name string, defname string) string { _ = "STUB: not implemented"; return "" }
