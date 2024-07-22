package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	var cmd_string string = "az aks list --output json"
	var args []string = strings.Fields(cmd_string)
	var cmd *exec.Cmd = exec.Command(args[0], args[1:]...)
  var output []byte
  var err error
  var data []map[string]string
  var resource_group string
  var name string

	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(output, &data)

	for _, item := range data {
		resource_group = item["resourceGroup"]
		name = item["name"]
		cmd_string = "az aks get-credentials --name %s --resource-group %s"
		cmd_string = fmt.Sprintf(cmd_string, string(name), string(resource_group))
		args = strings.Fields(cmd_string)
		cmd = exec.Command(args[0], args[1:]...)
		output, err = cmd.CombinedOutput()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(output))
	}
}
