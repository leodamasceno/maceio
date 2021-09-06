package main

import (
    "fmt"
    //"crypto/hmac"
    //"crypto/sha256"
    //"encoding/hex"
    "github.com/go-git/go-git/v5"
    . "github.com/go-git/go-git/v5/_examples"
    //"github.com/go-git/go-git/v5/plumbing"
    git_http "github.com/go-git/go-git/v5/plumbing/transport/http"
)


func main() {

    r, err := git.PlainClone("repo", false, &git.CloneOptions{
        Auth: &git_http.BasicAuth{
            Username: "git", // yes, this can be anything except an empty string
            //Password: "NzY5MTk1ODI0MDQ1OlsmzYzcqmSyhYnSY1QIfT4q2vU0",
            Password: "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",
        },
        //URL:           "https://git.gartner.com/scm/awsdm/leonardo-webhook-test2.git",
        URL:           "https://github.com/leodamasceno/ping2me.io.git",
        //ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", "master")),
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
    fmt.Println(commit_id)
}
