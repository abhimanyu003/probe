package cmd

import (
	"fmt"
	"github.com/abhimanyu003/probe/runner"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(runProbe)
	runProbe.PersistentFlags().BoolP("test.v", "v", false, "Get verbose output of the tests.")
	runProbe.PersistentFlags().Bool("disableLogs", false, "Disable logs file entries.")
	runProbe.PersistentFlags().Bool("failfast", false, "Do not start new tests after the first test failure.")
	runProbe.PersistentFlags().Uint("parallel", 0, "Maximum number of tests to run simultaneously")
	runProbe.PersistentFlags().String("env-file", "", "environment file to read and use as env in the containers (default .env)")
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
		envPath, err := cmd.Flags().GetString("env-file")
		if err != nil {
			cliutil.Redln(err)
			os.Exit(1)
		}
		loadEnvFromGivenPath(envPath)

		verbose, err := cmd.Flags().GetBool("test.v")
		if err != nil {
			cliutil.Redln(err)
			os.Exit(1)
		}
		disableLogs, err := cmd.Flags().GetBool("disableLogs")
		if err != nil {
			cliutil.Redln(err)
			os.Exit(1)
		}
		failfast, err := cmd.Flags().GetBool("failfast")
		if err != nil {
			cliutil.Redln(err)
			os.Exit(1)
		}
		parallel, err := cmd.Flags().GetUint("parallel")
		if err != nil {
			cliutil.Redln(err)
			os.Exit(1)
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

func loadEnvFromGivenPath(envPath string) {
	var err error
	if len(envPath) > 0 {
		if !fsutil.IsFile(envPath) {
			cliutil.Redln(fmt.Sprintf("no env file found at given path %s", envPath))
			os.Exit(1)
		}
		err = godotenv.Load(envPath)
		if err != nil {
			cliutil.Redln(err)
			os.Exit(1)
		}
		return
	}
	if fsutil.IsFile(".env") {
		err = godotenv.Load()
		if err != nil {
			cliutil.Redln(err)
			os.Exit(1)
		}
	}
}
