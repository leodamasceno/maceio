package bitbucket

import (
    "fmt"
    "os"
    "log"
    "net/http"
    "os/exec"
    "io/ioutil"
    "gopkg.in/yaml.v3"
    "bytes"
    "strings"
    "strconv"
    "encoding/json"
    "github.com/go-git/go-git/v5"
    . "github.com/go-git/go-git/v5/_examples"
    "github.com/go-git/go-git/v5/plumbing"
    git_http "github.com/go-git/go-git/v5/plumbing/transport/http"
)

type YAMLConfig struct {
    Tests []tests `yaml:"tests"`

}

type tests struct {
    Name     string `yaml:"name"`
    Cmd      string `yaml:"cmd"`
}

type CommitResponseData struct {
	Size       int  `json:"size"`
	Limit      int  `json:"limit"`
	IsLastPage bool `json:"isLastPage"`
	Values     []struct {
		ID          int    `json:"id"`
		Version     int    `json:"version"`
		Title       string `json:"title"`
		State       string `json:"state"`
		Open        bool   `json:"open"`
		Closed      bool   `json:"closed"`
		CreatedDate int64  `json:"createdDate"`
		UpdatedDate int64  `json:"updatedDate"`
		ClosedDate  int64  `json:"closedDate"`
		FromRef     struct {
			ID           string `json:"id"`
			DisplayID    string `json:"displayId"`
			LatestCommit string `json:"latestCommit"`
			Repository   struct {
				Slug          string `json:"slug"`
				ID            int    `json:"id"`
				Name          string `json:"name"`
				HierarchyID   string `json:"hierarchyId"`
				ScmID         string `json:"scmId"`
				State         string `json:"state"`
				StatusMessage string `json:"statusMessage"`
				Forkable      bool   `json:"forkable"`
				Project       struct {
					Key         string `json:"key"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Public      bool   `json:"public"`
					Type        string `json:"type"`
					Links       struct {
						Self []struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"links"`
				} `json:"project"`
				Public bool `json:"public"`
				Links  struct {
					Clone []struct {
						Href string `json:"href"`
						Name string `json:"name"`
					} `json:"clone"`
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"repository"`
		} `json:"fromRef"`
		ToRef struct {
			ID           string `json:"id"`
			DisplayID    string `json:"displayId"`
			LatestCommit string `json:"latestCommit"`
			Repository   struct {
				Slug          string `json:"slug"`
				ID            int    `json:"id"`
				Name          string `json:"name"`
				HierarchyID   string `json:"hierarchyId"`
				ScmID         string `json:"scmId"`
				State         string `json:"state"`
				StatusMessage string `json:"statusMessage"`
				Forkable      bool   `json:"forkable"`
				Project       struct {
					Key         string `json:"key"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Public      bool   `json:"public"`
					Type        string `json:"type"`
					Links       struct {
						Self []struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"links"`
				} `json:"project"`
				Public bool `json:"public"`
				Links  struct {
					Clone []struct {
						Href string `json:"href"`
						Name string `json:"name"`
					} `json:"clone"`
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"repository"`
		} `json:"toRef"`
		Locked bool `json:"locked"`
		Author struct {
			User struct {
				Name         string `json:"name"`
				EmailAddress string `json:"emailAddress"`
				ID           int    `json:"id"`
				DisplayName  string `json:"displayName"`
				Active       bool   `json:"active"`
				Slug         string `json:"slug"`
				Type         string `json:"type"`
				Links        struct {
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"user"`
			Role     string `json:"role"`
			Approved bool   `json:"approved"`
			Status   string `json:"status"`
		} `json:"author"`
		Reviewers    []interface{} `json:"reviewers"`
		Participants []interface{} `json:"participants"`
		Properties   struct {
			ResolvedTaskCount int `json:"resolvedTaskCount"`
			OpenTaskCount     int `json:"openTaskCount"`
		} `json:"properties"`
		Links struct {
			Self []struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"values"`
	Start int `json:"start"`
}

type PRResponseData struct {
	EventKey string `json:"eventKey"`
	Date     string `json:"date"`
	Actor    struct {
		Name         string `json:"name"`
		EmailAddress string `json:"emailAddress"`
		ID           int    `json:"id"`
		DisplayName  string `json:"displayName"`
		Active       bool   `json:"active"`
		Slug         string `json:"slug"`
		Type         string `json:"type"`
		Links        struct {
			Self []struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"actor"`
	PullRequest struct {
		ID          int    `json:"id"`
		Version     int    `json:"version"`
		Title       string `json:"title"`
		State       string `json:"state"`
		Open        bool   `json:"open"`
		Closed      bool   `json:"closed"`
		CreatedDate int64  `json:"createdDate"`
		UpdatedDate int64  `json:"updatedDate"`
		FromRef     struct {
			ID           string `json:"id"`
			DisplayID    string `json:"displayId"`
			LatestCommit string `json:"latestCommit"`
			Repository   struct {
				Slug          string `json:"slug"`
				ID            int    `json:"id"`
				Name          string `json:"name"`
				HierarchyID   string `json:"hierarchyId"`
				ScmID         string `json:"scmId"`
				State         string `json:"state"`
				StatusMessage string `json:"statusMessage"`
				Forkable      bool   `json:"forkable"`
				Project       struct {
					Key         string `json:"key"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Public      bool   `json:"public"`
					Type        string `json:"type"`
					Links       struct {
						Self []struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"links"`
				} `json:"project"`
				Public bool `json:"public"`
				Links  struct {
					Clone []struct {
						Href string `json:"href"`
						Name string `json:"name"`
					} `json:"clone"`
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"repository"`
		} `json:"fromRef"`
		ToRef struct {
			ID           string `json:"id"`
			DisplayID    string `json:"displayId"`
			LatestCommit string `json:"latestCommit"`
			Repository   struct {
				Slug          string `json:"slug"`
				ID            int    `json:"id"`
				Name          string `json:"name"`
				HierarchyID   string `json:"hierarchyId"`
				ScmID         string `json:"scmId"`
				State         string `json:"state"`
				StatusMessage string `json:"statusMessage"`
				Forkable      bool   `json:"forkable"`
				Project       struct {
					Key         string `json:"key"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Public      bool   `json:"public"`
					Type        string `json:"type"`
					Links       struct {
						Self []struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"links"`
				} `json:"project"`
				Public bool `json:"public"`
				Links  struct {
					Clone []struct {
						Href string `json:"href"`
						Name string `json:"name"`
					} `json:"clone"`
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"repository"`
		} `json:"toRef"`
		Locked bool `json:"locked"`
		Author struct {
			User struct {
				Name         string `json:"name"`
				EmailAddress string `json:"emailAddress"`
				ID           int    `json:"id"`
				DisplayName  string `json:"displayName"`
				Active       bool   `json:"active"`
				Slug         string `json:"slug"`
				Type         string `json:"type"`
				Links        struct {
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"user"`
			Role     string `json:"role"`
			Approved bool   `json:"approved"`
			Status   string `json:"status"`
		} `json:"author"`
		Reviewers    []interface{} `json:"reviewers"`
		Participants []interface{} `json:"participants"`
		Links        struct {
			Self []struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"pullRequest"`
}

type PushResponseData struct {
	EventKey string `json:"eventKey"`
	Date     string `json:"date"`
	Actor    struct {
		Name         string `json:"name"`
		EmailAddress string `json:"emailAddress"`
		ID           int    `json:"id"`
		DisplayName  string `json:"displayName"`
		Active       bool   `json:"active"`
		Slug         string `json:"slug"`
		Type         string `json:"type"`
		Links        struct {
			Self []struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"actor"`
	PullRequest struct {
		ID          int    `json:"id"`
		Version     int    `json:"version"`
		Title       string `json:"title"`
		State       string `json:"state"`
		Open        bool   `json:"open"`
		Closed      bool   `json:"closed"`
		CreatedDate int64  `json:"createdDate"`
		UpdatedDate int64  `json:"updatedDate"`
		FromRef     struct {
			ID           string `json:"id"`
			DisplayID    string `json:"displayId"`
			LatestCommit string `json:"latestCommit"`
			Repository   struct {
				Slug          string `json:"slug"`
				ID            int    `json:"id"`
				Name          string `json:"name"`
				HierarchyID   string `json:"hierarchyId"`
				ScmID         string `json:"scmId"`
				State         string `json:"state"`
				StatusMessage string `json:"statusMessage"`
				Forkable      bool   `json:"forkable"`
				Project       struct {
					Key         string `json:"key"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Public      bool   `json:"public"`
					Type        string `json:"type"`
					Links       struct {
						Self []struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"links"`
				} `json:"project"`
				Public bool `json:"public"`
				Links  struct {
					Clone []struct {
						Href string `json:"href"`
						Name string `json:"name"`
					} `json:"clone"`
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"repository"`
		} `json:"fromRef"`
		ToRef struct {
			ID           string `json:"id"`
			DisplayID    string `json:"displayId"`
			LatestCommit string `json:"latestCommit"`
			Repository   struct {
				Slug          string `json:"slug"`
				ID            int    `json:"id"`
				Name          string `json:"name"`
				HierarchyID   string `json:"hierarchyId"`
				ScmID         string `json:"scmId"`
				State         string `json:"state"`
				StatusMessage string `json:"statusMessage"`
				Forkable      bool   `json:"forkable"`
				Project       struct {
					Key         string `json:"key"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Public      bool   `json:"public"`
					Type        string `json:"type"`
					Links       struct {
						Self []struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"links"`
				} `json:"project"`
				Public bool `json:"public"`
				Links  struct {
					Clone []struct {
						Href string `json:"href"`
						Name string `json:"name"`
					} `json:"clone"`
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"repository"`
		} `json:"toRef"`
		Locked bool `json:"locked"`
		Author struct {
			User struct {
				Name         string `json:"name"`
				EmailAddress string `json:"emailAddress"`
				ID           int    `json:"id"`
				DisplayName  string `json:"displayName"`
				Active       bool   `json:"active"`
				Slug         string `json:"slug"`
				Type         string `json:"type"`
				Links        struct {
					Self []struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"links"`
			} `json:"user"`
			Role     string `json:"role"`
			Approved bool   `json:"approved"`
			Status   string `json:"status"`
		} `json:"author"`
		Reviewers    []interface{} `json:"reviewers"`
		Participants []interface{} `json:"participants"`
		Links        struct {
			Self []struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"pullRequest"`
	PreviousFromHash string `json:"previousFromHash"`
}

func CloneGitRepository(branch string, repo_url string) string {

    user := os.Getenv("GIT_USER")
    token := os.Getenv("GIT_TOKEN")

    local_dir := "repos/" + strings.ReplaceAll(branch, "/", "-")

    r, err := git.PlainClone(local_dir, false, &git.CloneOptions{
        Auth: &git_http.BasicAuth{
            Username: user, // yes, this can be anything except an empty string
            Password: token,
        },
        URL:           repo_url,
        ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
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

func RunCommand(branch string, command string) string {

    os.Setenv("TF_CLI_ARGS", "-no-color")

    cmd := exec.Command("/bin/sh", "-c", command)
    cmd.Dir = "repos/" + strings.ReplaceAll(branch, "/", "-")
    cmd_output, err := cmd.CombinedOutput()
    output := string(cmd_output)
    if err != nil {
        return output
    }

    return output

}

func GetPRId(git_host string, endpoint string, token string) int {

    var map_data CommitResponseData

    req, err := http.NewRequest("GET", git_host + endpoint, nil)
    req.Header.Set("Authorization", "Bearer " + token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("[ERROR] Failed to retrieve pull request ID: %v\n", err)
        return 0
    }

    //defer resp.Body.Close()

    body, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        fmt.Printf("[ERROR] Failed to process pull: %v\n", err2)
    }

    err = json.Unmarshal([]byte(body), &map_data)
    if err != nil {
        fmt.Printf("[ERROR] Failed to process pull request response: %v\n", err)
        return 0
    }

    return map_data.Values[0].ID

}

func CreateComment(cmd_output string, commit_id string, slug string, project string) {

    token := os.Getenv("GIT_TOKEN")
    git_host := os.Getenv("GIT_HOST")
    msg := "```"+cmd_output+"```"
    comment := strings.Replace(msg,"\n","\\n",-1)
    //git_post_data := []byte(`{"text": "` + msg +`"}`)
    git_post_data := []byte(`{"text": "`+comment+`"}`)

    pr_id := GetPRId(git_host, "/rest/api/1.0/projects/" + project + "/repos/" + slug + "/commits/" + commit_id + "/pull-requests", token)

    if pr_id != 0 {
        req, err := http.NewRequest("POST", git_host + "/rest/api/1.0/projects/" + project + "/repos/" + slug + "/pull-requests/" + strconv.Itoa(pr_id) + "/comments", bytes.NewBuffer(git_post_data))
        req.Header.Set("Authorization", "Bearer " + token)
        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            fmt.Printf("[ERROR] Failed to add comment to pull request: %v\n", err)
        }

        defer resp.Body.Close()
    }


}

func UpdateBuildStatus(status string, commit_id string, url string) bool {

    token := os.Getenv("GIT_TOKEN")
    git_host := os.Getenv("GIT_HOST")
    git_post_data := []byte(`{"state": "` + status +`",
                              "key": "Maceio",
                              "url": "` + url + `"}`)

    req, err := http.NewRequest("POST", git_host + "/rest/build-status/1.0/commits/" + commit_id, bytes.NewBuffer(git_post_data))
    req.Header.Set("Authorization", "Bearer " + token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return false
    }

    defer resp.Body.Close()

    return true

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

func EventHandler(action string, payload []byte) {

    var branch string
    var repo_url string
    var slug string
    var project string
    var url string
    var local_dir string

    switch action {
    case "pr:opened":
        var map_data PRResponseData
        err := json.Unmarshal(payload, &map_data)
        if err != nil {
            log.Fatal("[ERROR] Could not decode payload")
        }

        branch = map_data.PullRequest.FromRef.DisplayID
        repo_url = map_data.PullRequest.FromRef.Repository.Links.Clone[0].Href
        commit_id := CloneGitRepository(branch, repo_url)
        slug = map_data.PullRequest.FromRef.Repository.Slug
        project = map_data.PullRequest.FromRef.Repository.Project.Key
        url = strings.Replace(map_data.PullRequest.FromRef.Repository.Links.Self[0].Href, "browse", "builds", -1)
        local_dir = "repos/" + strings.ReplaceAll(branch, "/", "-")

        status := UpdateBuildStatus("INPROGRESS", commit_id, url)
        if status != true {
            fmt.Printf("[ERROR] Failed to update build status: %v\n", err)
        } else {
            config := ReadConfigFile(branch)
            for _, e := range config.Tests {
                cmd_output := RunCommand(branch, e.Cmd)
                if cmd_output == "" {
                    UpdateBuildStatus("FAILED", commit_id, url)
                } else {
                    UpdateBuildStatus("SUCCESSFUL", commit_id, url)
                }

                output := string(cmd_output)
                log.Printf(output)
                CreateComment(output, commit_id, slug, project)
            }
        }

    case "pr:from_ref_updated":
        var map_data PushResponseData
        err := json.Unmarshal(payload, &map_data)
        if err != nil {
            log.Fatal("[ERROR] Could not decode payload")
        }

        branch = map_data.PullRequest.FromRef.DisplayID
        repo_url = map_data.PullRequest.FromRef.Repository.Links.Clone[0].Href
        commit_id := CloneGitRepository(branch, repo_url)
        slug = map_data.PullRequest.FromRef.Repository.Slug
        project = map_data.PullRequest.FromRef.Repository.Project.Key
        url = strings.Replace(map_data.PullRequest.FromRef.Repository.Links.Self[0].Href, "browse", "builds", -1)
        local_dir = "repos/" + strings.ReplaceAll(branch, "/", "-")

        status := UpdateBuildStatus("INPROGRESS", commit_id, url)
        if status != true {
            fmt.Printf("[ERROR] Failed to update build status: %v\n", err)
        } else {
            config := ReadConfigFile(branch)
            for _, e := range config.Tests {
                cmd_output := RunCommand(branch, e.Cmd)
                if cmd_output == "" {
                    UpdateBuildStatus("FAILED", commit_id, url)
                } else {
                    UpdateBuildStatus("SUCCESSFUL", commit_id, url)
                }

                output := string(cmd_output)
                CreateComment(output, commit_id, slug, project)
            }
        }

    }

    os.RemoveAll(local_dir)

}
