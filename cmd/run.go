package cmd

import (
	"fmt"
	"github.com/abhimanyu003/probe/runner"
	"github.com/gookit/goutil/fsutil"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
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
		fsutil.IsFile("")
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		envPath, err := cmd.Flags().GetString("env-file")
		if err != nil {
			panic(err)
		}
		loadEnvFromGivenPath(envPath)

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

func loadEnvFromGivenPath(envPath string) {
	var err error
	if len(envPath) > 0 {
		if !fsutil.IsFile(envPath) {
			panic(fmt.Sprintf("no env file found at given path %s", envPath))
		}
		err = godotenv.Load(envPath)
		if err != nil {
			panic(err)
		}
		return
	}
	if fsutil.IsFile(".env") {
		err = godotenv.Load()
		if err != nil {
			panic(err)
		}
	}
}
