package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"strconv"

	jira "github.com/andygrunwald/go-jira"
	notify "github.com/TheCreeper/go-notify"
	"github.com/joho/godotenv"
)

type jiraConfig struct {
	host     string
	username string
	password string
	jql	 string
	timeout  string
}

type jiraIssue struct {
	*jira.Issue
}

type notification struct {
	body     string
	title    string
	subtitle string
	link     string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jc := jiraConfig{host: os.Getenv("JIRA_HOST"), username: os.Getenv("JIRA_USERNAME"), password: os.Getenv("JIRA_PASSWORD"), jql: os.Getenv("JIRA_JQL"), timeout: os.Getenv("TIMEOUT") }

	jiraClient, err := jira.NewClient(nil, jc.host)
	if err != nil {
		panic(err)
	}
	res, err := jiraClient.Authentication.AcquireSessionCookie(jc.username, jc.password)
	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}

	jql := jc.jql

	for {
		issues, _, _ := jiraClient.Issue.Search(jql, nil)
		i, err := strconv.ParseInt(jc.timeout, 10, 32)
                if err != nil {
                    panic(err)
                }

		if len(issues) > 0 {
			for _, issue := range issues {

				n := newNotificationFromIssue(issue)
				fmt.Println(issue.Key, "has some changes ->", jiraIssue{&issue}.getLink(), "at", timeNow())

				go func(n notification) {
				    sendNotification(n)
				    time.Sleep( time.Duration(i) * time.Second)
				}(n)
			}
		}
		time.Sleep(( time.Duration(i) * time.Minute) - (1 * time.Second))
	}
}

func sendNotification(n notification) {
	ntf := notify.NewNotification(n.title, n.link + "\n\n" + n.subtitle)
	ntf.AppIcon =  "assets/jiraicon.png"

	if _, err := ntf.Show(); err != nil {
	    return
	}

}

func newNotificationFromIssue(i jira.Issue) notification {
	return notification{
		body:     "You have new activity",
		title:    i.Key,
		subtitle: i.Fields.Summary,
		link:     jiraIssue{&i}.getLink(),
	}
}

func (ji jiraIssue) getLink() string {
	return os.Getenv("JIRA_HOST") + "browse/" + ji.Key
}

func timeNow() string {
	t := time.Now()
	return t.Format("15:04:05")
}
