package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "jj",
	Short: "Jujutsu (An experimental VCS)",
	Long:  "https://github.com/martinvonz/jj",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("at-operation", "", "Operation to load the repo at")
	rootCmd.PersistentFlags().String("color", "", "When to colorize output (always, never, auto)")
	rootCmd.PersistentFlags().StringSlice("config-toml", []string{}, "Additional configuration options (can be repeated)")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().Bool("ignore-working-copy", false, "Don't snapshot the working copy, and don't update it")
	rootCmd.PersistentFlags().Bool("no-pager", false, "Disable the pager")
	rootCmd.PersistentFlags().StringP("repository", "R", "", "Path to repository to operate on")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"at-operation": carapace.ActionValues(), // TODO
		"color":        carapace.ActionValues("always", "never", "auto").StyleF(style.ForKeyword),
		"config-toml":  carapace.ActionFiles(),
		"repository":   carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(rootCmd.Flag("repository")))
	})
}
