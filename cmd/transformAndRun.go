package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
)

type ESBuildTscJSON struct {
	Schema        string `json:"$schema"`
	File          string `json:"file"`
	TSConfig      string `json:"tsconfig"`
	ESBuildEngine string `json:"engine"`
}

var transformAndRunCmd = &cobra.Command{
	Use:     "tar",
	Aliases: []string{"transformAndRun"},
	Short:   "Transform a file using esbuild and run it",
	Long:    "Transform a file using esbuild and run it - using esbuild-tsc.json",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := os.ReadFile("esbuild-tsc.json")
		var data ESBuildTscJSON
		err = json.Unmarshal(result, &data)
		if err != nil {
			log.Fatal(err)
		}
		TransformAndRun(data.File, data.ESBuildEngine, data.TSConfig)
	},
}

func init() {
	rootCmd.AddCommand(transformAndRunCmd)
}

func TransformAndRun(file string, engine string, tsconfig string) {
	if _, err := os.Open("dist"); os.IsNotExist(err) {
		os.Mkdir("dist", 0755)
	}
	result := api.Build(api.BuildOptions{
		EntryPoints:  []string{file},
		Bundle:       true,
		Outfile:      "dist/index.js",
		MinifySyntax: true,
		Tsconfig:     tsconfig,
		Write:        true,
	})
	if len(result.Errors) != 0 {
		for _, err := range result.Errors {
			fmt.Println(err.Text)
		}
		os.Exit(1)
	}
	if engine == "deno" {
		out, err := exec.Command(engine, "run", "dist/index.js").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	} else {
		out, err := exec.Command(engine, "dist/index.js").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}
}
