package replicate

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/replicate/cli/internal/cmd"
	"github.com/replicate/cli/internal/cmd/model"
	"github.com/replicate/cli/internal/cmd/prediction"
	"github.com/replicate/cli/internal/cmd/training"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "replicate",
	Version: cmd.Version(),
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
	rootCmd.AddGroup(&cobra.Group{
		ID:    "core",
		Title: "Core commands:",
	})
	for _, cmd := range []*cobra.Command{
		model.RootCmd,
		prediction.RootCmd,
		training.RootCmd,
	} {
		rootCmd.AddCommand(cmd)
		cmd.GroupID = "core"
	}

	rootCmd.AddGroup(&cobra.Group{
		ID:    "alias",
		Title: "Alias commands:",
	})
	for _, cmd := range []*cobra.Command{
		cmd.RunCmd,
		cmd.TrainCmd,
	} {
		rootCmd.AddCommand(cmd)
		cmd.GroupID = "alias"
	}
}
