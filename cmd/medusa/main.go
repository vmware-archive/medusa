package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"io/ioutil"
	"strings"
	"path/filepath"

	"github.com/BurntSushi/toml"	
)

type MedusaConfig struct {
	Org string
	ApiKey string
}

func main() {
	//The init command
	initCommandPtr := flag.Bool("init", false, "Set/initialize the Medusa configs e.g. org and API key")

	//The repos command
	reposCommand := flag.NewFlagSet("repos", flag.ExitOnError)
	repoTypePtr := reposCommand.String("type", "all", "Repo type all|private|public, defaults to all")
	reposVerbosePtr := reposCommand.Bool("verbose", false, "Verbose mode i.e. full detais")
	reposCsvPtr := reposCommand.Bool("csv", false, "Report results in CSV format")

	//The repo command
	repoCommand := flag.NewFlagSet("repo", flag.ExitOnError)
	repoNamePtr := repoCommand.String("name", "", "The repo's name (required)")
	repoVerbosePtr := repoCommand.Bool("verbose", false, "Verbose mode i.e. full detais")
	repoCsvPtr := repoCommand.Bool("csv", false, "Report results in CSV format")

	//TODO
	//setApiKey := flag.NewFlagSet("api_key", flag.ExitOnError)
	//infoCommand := flag.NewFlagSet("info", flag.ExitOnError)
	//medusa users --filters -2fa -verbose -csv
	//medusa user <user_name> -verbose -teams -repos -csv
	//medusa teams|groups -verbose -users -repos -csv
	//medusa team|group <team_group_name> --filters
	//medusa collaborators --filters
	//medusa collaborator <collborator_name> --filters
	
	flag.Parse()
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	homeDir := os.Getenv("HOME")
	dot_medusa := filepath.Join(homeDir, ".medusa")
	config := loadConfig(dot_medusa)
	
	switch os.Args[1] {
	case "init":
		setConfig(dot_medusa)
	case "repos":
		reposCommand.Parse(os.Args[2:])
		repos(&config, repoTypePtr, reposVerbosePtr, reposCsvPtr)
	case "repo":
		repoCommand.Parse(os.Args[2:])
		if *repoNamePtr == "" {
			repoCommand.PrintDefaults()
			os.Exit(1)
		}
		repo(&config, repoNamePtr, repoVerbosePtr, repoCsvPtr)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}	
}

func setConfig(confFilePath string) (MedusaConfig) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of your GitHub organization: ")
	org, _ := reader.ReadString('\n')
	org = strings.TrimSpace(org)
	fmt.Print("Copy/paste your GitHub API key: ")
	apiKey, _ := reader.ReadString('\n')
	apiKey = strings.TrimSpace(apiKey)
	confData := []byte(fmt.Sprintf("Org=\"%s\"\nApiKey=\"%s\"\n", org, apiKey))
	err := ioutil.WriteFile(confFilePath, confData, 0644)
	if err != nil {
		panic(err)
	}
	return MedusaConfig{org, apiKey}
}

func loadConfig(confFilePath string) (MedusaConfig) {
	var config MedusaConfig
	confExists, _ := confFileExists(confFilePath)
	if confExists {
		if _, err := toml.DecodeFile(confFilePath, &config); err != nil {
			fmt.Println(err)
			return config
		}
	} else {
		config = setConfig(confFilePath)
	}
	return config
}

func confFileExists(confFilePath string) (bool, error){
	exists := true
	var existsError error
	if _, err := os.Stat(confFilePath); err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			existsError = err
		}
	}
	return exists, existsError
}

/*func init(org string){
}*/

func repos(config *MedusaConfig, repoType *string, verbose *bool, csv *bool){
	//TODO
}	

func repo(config *MedusaConfig, repoName *string, verbose *bool, csv *bool){
	//TODO
}	

func users(){
	fmt.Println("users")
}	

func user(){
	fmt.Println("user")
}	

func teams(){
	fmt.Println("teams")
}	

func team(){
	fmt.Println("team")
}	

func groups(){
	fmt.Println("groups")
}	

func group(){
	fmt.Println("group")
}	

func collaborator(){
	fmt.Println("collaborator")
}	

func collaborators(){
	fmt.Println("collaborators")
}	




