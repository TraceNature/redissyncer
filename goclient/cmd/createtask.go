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
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// createtaskCmd represents the createtask command
var createtaskCmd = &cobra.Command{
	Use:   "createtask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//var dat []map[string]interface{}
		fmt.Println("createtask called")
		execfile, _ := cmd.Flags().GetString("execfile")
		if execfile == "" {
			log.Println("please put create file use '-e' parameter")
			os.Exit(1)
		}
		jsonFile, err := os.Open(execfile)
		defer jsonFile.Close()

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)
		//json.Unmarshal(byteValue, &dat)
		//for k, v := range dat {
		//	result, _ := json.Marshal(v)
		//	execcreatetask(viper.GetViper().GetString("server"), string(v))
		//}

		execcreatetask(viper.GetViper().GetString("server"), string(byteValue))
		os.Exit(0)

	},
}

func init() {
	rootCmd.AddCommand(createtaskCmd)
	createtaskCmd.Flags().StringP("execfile", "e", "", "json file path to create a task")
}

func execcreatetask(server string, createjson string) {

	url := server + "/api/v1/createtask"
	log.Println(createjson)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(createjson))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var dat map[string]interface{}
	json.Unmarshal(body, &dat)
	result, _ := json.MarshalIndent(dat, "", " ")
	fmt.Println(string(result))
}
