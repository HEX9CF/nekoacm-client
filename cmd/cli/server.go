package cli

import (
	"github.com/spf13/cobra"
	"nekoacm-client/cmd/bootstrap"
)

// 服务器命令
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the server.",
	Long:  "Run the server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return bootstrap.InitServer()
	},
}
