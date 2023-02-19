package cmd

import (
	"probe/runner"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runProbe)
	runProbe.PersistentFlags().BoolP("test.v", "v", false, "Get verbose output of the tests.")
	runProbe.PersistentFlags().Bool("disableLogs", false, "Disable logs file entries.")
	runProbe.PersistentFlags().Bool("failfast", false, "Do not start new tests after the first test failure.")
	runProbe.PersistentFlags().Uint("parallel", 0, "Maximum number of tests to run simultaneously")
}

var runProbe = &cobra.Command{
	Use:   "run",
	Short: "Run probe tests",
	Long:  `Run probe test`,
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		verbose, err := cmd.Flags().GetBool("test.v")
		if err != nil {
			panic(err)
		}
		disableLogs, err := cmd.Flags().GetBool("disableLogs")
		if err != nil {
			panic(err)
		}
		failfast, err := cmd.Flags().GetBool("failfast")
		if err != nil {
			panic(err)
		}
		parallel, err := cmd.Flags().GetUint("parallel")
		if err != nil {
			panic(err)
		}
		flags := runner.Flags{
			Verbose:     verbose,
			DisableLogs: disableLogs,
			FailFast:    failfast,
			Parallel:    parallel,
		}

		probe := runner.NewProbe(path, flags, nil)
		probe.Execute()
	},
}
