package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"path/filepath"

	//"github.com/BurntSushi/toml"	
)

type medusaConfig struct {
	Org string
	EncryptedApiKey string
}

func main() {
	//The repos command
	reposCommand := flag.NewFlagSet("repos", flag.ExitOnError)
	reposTypePtr := reposCommand.String("type", "all", "Repo type all|private|public, defaults to all")
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
	
	switch os.Args[1] {
	case "init":
		//prompt for the org and API key here
		fmt.Println(confFileExists())
		org := readInput()
		fmt.Println(org)
	case "repos":
		reposCommand.Parse(os.Args[2:])
		fmt.Println(*reposTypePtr)
		fmt.Println(*reposVerbosePtr)
		fmt.Println(*reposCsvPtr)		
	case "repo":
		repoCommand.Parse(os.Args[2:])
		if *repoNamePtr == "" {
			repoCommand.PrintDefaults()
			os.Exit(1)
		}
		fmt.Println(*repoNamePtr)
		fmt.Println(*repoVerbosePtr)
		fmt.Println(*repoCsvPtr)		
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}	
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of your GitHub organization: ")
	org, _ := reader.ReadString('\n')
	fmt.Print(org)
	return org
}

	


func loadConfig(){
}

func confFileExists() (bool, error){
	homeDir := os.Getenv("HOME")
	dot_medusa := filepath.Join(homeDir, ".medusa")
	exists := true
	var existsError error
	if _, err := os.Stat(dot_medusa); err != nil {
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

func repos(){
	fmt.Println("repos")
}	

func repo(){
	fmt.Println("repo")
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




