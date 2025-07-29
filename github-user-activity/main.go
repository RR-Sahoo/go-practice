package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload map[string]interface{} `json:"payload"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected a username.")
		return
	}
	username := os.Args[1]
	apiUrl := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("GitHub API returned non-200 status:", resp.Status)
		return
	}

	var events []Event
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			commits := event.Payload["commits"].([]interface{})
			fmt.Printf("🔼 Pushed %d commits to %s\n", len(commits), event.Repo.Name)
		case "IssuesEvent":
			action := event.Payload["action"].(string)
			fmt.Printf("🐛 %s an issue in %s\n", capitalize(action), event.Repo.Name)
		case "IssueCommentEvent":
			action := event.Payload["action"].(string)
			fmt.Printf("💬 %s a comment on an issue in %s\n", capitalize(action), event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("⭐ Starred %s\n", event.Repo.Name)
		case "PullRequestEvent":
			action := event.Payload["action"].(string)
			fmt.Printf("🔃 %s a pull request in %s\n", capitalize(action), event.Repo.Name)
		default:

		}
	}

}
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}
