package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/t1utiu/javm/config"
)

func useJdk(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a JDK version")
		return
	}
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	version := args[0]
	jdkPath := filepath.Join(config.JdkDir, "jdk-"+version)
	if _, err := os.Stat(jdkPath); os.IsNotExist(err) {
		fmt.Printf("JDK version %s not found\n", version)
		return
	}
	cmdSetEnv := exec.Command("setx", config.EnvVar, jdkPath)
	if err := cmdSetEnv.Run(); err != nil {
		fmt.Printf("Error setting user environment variable: %v\n", err)
	}
	cmdSetEnv = exec.Command("setx", config.EnvVar, jdkPath, "/M")
	if err := cmdSetEnv.Run(); err != nil {
		fmt.Printf("Error setting system environment variable: %v\n", err)
	}
}

func listJdks(cmd *cobra.Command, args []string) {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := os.ReadDir(config.JdkDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		if match, _ := filepath.Match("jdk-*", file.Name()); match {
			fmt.Println(file.Name())
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "jdkvm",
	Short: "JDK Version Manager",
}

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Use a specific JDK version",
	Run:   useJdk,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed JDK versions",
	Run:   listJdks,
}

func Init() {
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(listCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
