package cmd

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Long:  "Create a new project using ESBuild-TSC (test directory)",
	Short: "Create a new project using ESBuild-TSC",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateProject(args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func CreateProject(projectName string) {
	if _, err := os.Open(projectName); os.IsNotExist(err) {
		os.Mkdir(projectName, 0755)
	}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := os.WriteFile(fmt.Sprintf("%s/index.ts", projectName), []byte(`type Hello = "Hello World!" | "Hello";
const printMsg = (msg: Hello) => console.log(msg);
printMsg("Hello World!");`), 0755)
		if err != nil {
			log.Fatal(err)
		}
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := os.WriteFile(fmt.Sprintf("%s/esbuild-tsc.json", projectName), []byte(`{
  "$schema": "https://raw.githubusercontent.com/jeremiasbots/esbuild-tsc/refs/heads/main/esbuild-tsc-schema.json",
  "file": "index.ts",
  "engine": "bun",
  "tsconfig": "./tsconfig.json"
}`), 0755)
		if err != nil {
			log.Fatal(err)
		}
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := os.WriteFile(fmt.Sprintf("%s/tsconfig.json", projectName), []byte(`{
  "compilerOptions": {
	"target": "es2016",
	"module": "commonjs",
	"esModuleInterop": true,
	"forceConsistentCasingInFileNames": true,
	"strict": true,
	"skipLibCheck": true
  }
}`), 0755)
		if err != nil {
			log.Fatal(err)
		}
	}(wg)
	wg.Wait()
}
