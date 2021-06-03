package cmd

/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"
	"os"

	"github.com/luisoliveira/heimdall/vault"
	"github.com/spf13/cobra"
)

var getSecrets = &cobra.Command{
	Use:   "get-secrets",
	Short: "get env vars of a project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		//envs := []string
		env := vault.GetEnv()

		projectName, _ := cmd.Flags().GetString("name")

		if projectName == "" {
			fmt.Println("usage : heimdall get-secrets `nome do projeto`")
			os.Exit(1)
		}

		token := "token"

		vault.KVSecrets(env, projectName, token)

	},
}

func init() {

	rootCmd.AddCommand(getSecrets)
	getSecrets.Flags().StringP("name", "n", "", "Application name")

}

var rootCmd = &cobra.Command{
	Use:   "heimdall",
	Short: "get access to the king's vault",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage : heimdall get-secrets `nome do projeto`")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
