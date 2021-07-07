package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "os/exec"
    "strings"
    "io/ioutil"
    "gopkg.in/yaml.v3"
    "github.com/google/go-github/github"
    "github.com/go-git/go-git/v5"
    . "github.com/go-git/go-git/v5/_examples"
    "github.com/go-git/go-git/v5/plumbing"
    git_http "github.com/go-git/go-git/v5/plumbing/transport/http"
    "golang.org/x/net/context"
    "golang.org/x/oauth2"
)

type YAMLConfig struct {
    Tests []tests `yaml:"tests"`

}

type tests struct {
    Name     string `yaml:"name"`
    Cmd      string `yaml:"cmd"`
}

func addCommentToGithub(pr_number int, body string) {

    var err error

    msg := "```\n" + body
    // Construct github HTTP client
    token := os.Getenv("GIT_TOKEN")
    ts := oauth2.StaticTokenSource(&oauth2.Token{
        AccessToken: token,
    })
    tc := oauth2.NewClient(oauth2.NoContext, ts)
    client := github.NewClient(tc)
    _, _, err = client.Issues.CreateComment(
        context.Background(),
        "leodamasceno",
        "bazer-test",
        pr_number,
        &github.IssueComment{Body: &msg},
    )
    if err != nil {
        log.Fatalf("[ERROR] Failed to create comment: %s", err)
    } else {

    }

}

func updatePRStatusCheck(token string, stat string, desc string, commit string) {

    ctx := context.Background()
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(ctx, ts)

    client := github.NewClient(tc)

    _, _, err := client.Users.Get(ctx, "")
    if err != nil {
        fmt.Printf("\nerror: %v\n", err)
        return
    }

    name := "bazer"
    status := stat
    description := desc
    context_git := "Maceio"
    creator := github.User{
        Name:  &name,
    }

    repoStatus := &github.RepoStatus{
        State:       &status,
        Description: &description,
        Context:     &context_git,
        Creator:     &creator,
    }

    createStatus, _, err := client.Repositories.CreateStatus(context.Background(), "leodamasceno", "bazer-test", commit, repoStatus)
    if err != nil {
        log.Fatal(createStatus)
    }

}

func runCommand(token string, name string, command string, branch string, commit_id string, pr_number int) {

    os.Setenv("TF_CLI_ARGS", "-no-color")
    cmd := exec.Command("/bin/sh", "-c", command)
    cmd.Dir = "repos/" + strings.ReplaceAll(branch, "/", "-")
    cmd_output, err := cmd.CombinedOutput()
    if err != nil {
        updatePRStatusCheck(token, "error", name, commit_id)
    } else {
        updatePRStatusCheck(token, "success", name, commit_id)
    }
    body := string(cmd_output)
    addCommentToGithub(pr_number, body)

}

func readConfigFile (branch string) YAMLConfig {

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

func cloneGithubRepository(token string, local_dir string, branch string, repo_url string) string {
    // Clone the given repository to the given directory
    Info("Cloning repository: %s. Branch: %s", repo_url, branch)

    r, err := git.PlainClone(local_dir, false, &git.CloneOptions{
        Auth: &git_http.BasicAuth{
            Username: "git", // yes, this can be anything except an empty string
            Password: token,
        },
        URL:           repo_url,
        ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
        //Progress:      os.Stdout,
    })
    CheckIfError(err)

    // Retrieving the branch being pointed by HEAD
    ref, err := r.Head()
    CheckIfError(err)
    // ... retrieving the commit object
    commit, err := r.CommitObject(ref.Hash())
    CheckIfError(err)
    commit_id := commit.ID().String()

    return commit_id

}


func handleWebhook(w http.ResponseWriter, r *http.Request) {

    // Get Github Token and Secret from env variable
    token := os.Getenv("GIT_TOKEN")
    secret := os.Getenv("GIT_SECRET")

	payload, err := github.ValidatePayload(r, []byte(secret))
	if err != nil {
		log.Printf("[ERROR] Error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("[ERROR] Could not parse webhook: err=%s\n", err)
		return
	}

	switch e := event.(type) {
	case *github.PullRequestEvent:
        pr_number := *e.Number
        branch := *e.PullRequest.Head.Ref
        repo_url := *e.Repo.CloneURL
        local_dir := "repos/" + strings.ReplaceAll(branch, "/", "-")
        commit_id := cloneGithubRepository(token, local_dir, branch, repo_url)
        config := readConfigFile(branch)

        for _, e := range config.Tests {
            runCommand(token, e.Name, e.Cmd, branch, commit_id, pr_number)
        }

        // Delete local repo clone
        os.RemoveAll(local_dir)
	}
}

func main() {
    log.Println("Maceio started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
