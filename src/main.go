package main

import (
    "log"
    "os"
    "net/http"
    "strings"
    "io/ioutil"
    "github.com/google/go-github/github"
    bitbucket "maceio/bitbucket"
    //github_local "maceio/github"
    "reflect"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {


    if strings.Contains(r.UserAgent(), "Bitbucket") {

        action := r.Header.Get("X-Event-Key")
        signature := r.Header.Get("X-Hub-Signature")

        if action != "diagnostics:ping" {

            payload, err := ioutil.ReadAll(r.Body)
            if err != nil {
                log.Fatal("[ERROR] The payload is invalid: %v", err.Error())
            }

            err_payload := bitbucket.ValidatePayload(payload, signature)
            if err_payload != true {
                log.Fatal("[ERROR] The payload could not be validated: %v", err.Error())
            }

            bitbucket.EventHandler(action, payload)
        }
    } else if strings.Contains(r.UserAgent(), "GitHub") {

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
        log.Fatal(reflect.TypeOf(event))
        //github_local.EventHandler(event)
    }
}

func main() {
    log.Println("Maceio started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
