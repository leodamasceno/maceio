package main

import (
    "fmt"
    "log"
    //"io/ioutil"
    "net/http"
    "strings"
    //"encoding/json"
    "io/ioutil"
    //bitbucket "./bitbucket"
    //"github.com/ktrysmt/go-bitbucket"
)

type MapData struct {
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
	Repository struct {
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
	Changes []struct {
		Ref struct {
			ID        string `json:"id"`
			DisplayID string `json:"displayId"`
			Type      string `json:"type"`
		} `json:"ref"`
		RefID    string `json:"refId"`
		FromHash string `json:"fromHash"`
		ToHash   string `json:"toHash"`
		Type     string `json:"type"`
	} `json:"changes"`
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {

    //var map_data MapData

    if strings.Contains(r.UserAgent(), "GitHub") {
        fmt.Println("Github")
    } else if strings.Contains(r.UserAgent(), "Bitbucket") {
        fmt.Println("Bitbucket")
        //bitbucket.CreateComment("final test",w,r)
        //bitbucket.UpdateBuildStatus()
        //err := json.NewDecoder(r.Body).Decode(&map_data)
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        action := r.Header.Get("X-Event-Key")

        if strings.Contains(action, "pr:opened") {


        } else if strings.Contains(action, "pr:from_ref_updated") {
            fmt.Println(string(body))
            //repo_url := map_data.Repository.Links.Clone[0].Href
            //branch := map_data.Changes[0].Ref.DisplayID
            //fmt.Printf("Branch: %v - url: %+v", branch, repo_url)
            fmt.Println("DONE")
        } else {
            fmt.Println("WRONG ACTION")
        }


    }
}

func main() {
    log.Println("Maceio started")
    http.HandleFunc("/webhook", handleWebhook)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
