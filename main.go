/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package main

import (
	"fmt"
	bitcmd "github.com/chriswalz/bit/cmd"
	"log"
	"os"
)

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func main() {
	log.Println("Start")
	// defer needed to handle funkyness with CTRL + C & go-prompt
	defer bitcmd.HandleExit()
	if !bitcmd.IsGitRepo() {
		fmt.Println("fatal: not a git repository (or any of the parent directories): .git")
		return
	}
	log.Println("pre bit")
	argsWithoutProg := os.Args[1:]
	bitcliCmds := []string{"save", "sync", "version", "help", "info", "release"}
	if len(argsWithoutProg) == 0 || find(bitcliCmds, argsWithoutProg[0]) {
		log.Println("bit cli")
		bitcli()
	} else {
		log.Println("git cli start")
		completerSuggestionMap, _ := bitcmd.CreateSuggestionMap(bitcmd.ShellCmd)
		log.Println("post suggestion")
		yes := bitcmd.GitCommandsEnhancementPrompt(argsWithoutProg, completerSuggestionMap)
		log.Println("post enhance?")
		if yes {
			log.Println("git enhancement command used")
			return
		}
		log.Println("git command")
		bitcmd.RunGitCommandWithArgs(argsWithoutProg)
	}
}

func bitcli() {
	bitcmd.Execute()
}
