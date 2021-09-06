package github

import (
    /*
    "fmt"
    */
    "os"
    "log"
    /*
    "net/http"
    "os/exec"
    */
    "io/ioutil"
    "gopkg.in/yaml.v3"
    /*
    "bytes"
    */
    "strings"
    /*
    "strconv"
    "encoding/json"
    "github.com/go-git/go-git/v5"
    . "github.com/go-git/go-git/v5/_examples"
    "github.com/go-git/go-git/v5/plumbing"
    */
    "github.com/google/go-github/github"
)

type YAMLConfig struct {
    Tests []tests `yaml:"tests"`

}

type tests struct {
    Name     string `yaml:"name"`
    Cmd      string `yaml:"cmd"`
}

func ReadConfigFile (branch string) YAMLConfig {

    config := YAMLConfig{}

    // Read config file
    yamlFile, err := ioutil.ReadFile("repos/" + strings.ReplaceAll(branch, "/", "-") + "/maceio.yaml")
    if err != nil {
        log.Printf("[ERROR] Config file error: #%v ", err)
    }

    err = yaml.Unmarshal(yamlFile, &config)

    if err != nil {
        log.Printf("[ERROR] Config file error: #%v ", err)
    }

    return config

}

func EventHandler(event string) {

    var branch string
    var repo_url string
    var slug string
    var project string
    var url string
    var local_dir string

	switch e := event.(type) {
	case *github.PullRequestEvent:
        pr_number := *e.Number
        branch := *e.PullRequest.Head.Ref
        repo_url := *e.Repo.CloneURL
        org := strings.Split(*e.Repo.FullName, "/")[0]
        repo_name := strings.Split(*e.Repo.FullName, "/")[1]
        local_dir := "repos/" + strings.ReplaceAll(branch, "/", "-")
        commit_id := cloneGithubRepository(token, local_dir, branch, repo_url)
        config := readConfigFile(branch)

        for _, e := range config.Tests {
            runCommand(token, org, repo_name, e.Name, e.Cmd, branch, commit_id, pr_number)
        }

        // Delete local repo clone
        os.RemoveAll(local_dir)
	}

}
