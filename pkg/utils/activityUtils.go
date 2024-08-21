package utils

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"net/http"
	"strings"
)

var DashStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#DC143C")).
	Margin(0, 0, 0, 4)

var ContentStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFD700"))

var InfoStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#32CD32"))

var ErrorStyle = lipgloss.NewStyle().
	Bold(true).
	Underline(true).
	Foreground(lipgloss.Color("#FF0000"))

func DisplayInfo(username string, page int) {
	data, err := fetchData(username, page)
	if err != nil {
		fmt.Println(ErrorStyle.Render(err.Error()))
	}
	for _, item := range data {
		muxData(item)
	}
}

func fetchData(username string, page int) ([]map[string]any, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events?page=%d", username, page))
	if err != nil {
		fmt.Println(ErrorStyle.Render(err.Error()))
	}
	if response.StatusCode != 200 {
		if response.StatusCode == 404 {
			return nil, fmt.Errorf("username not found")
		} else {
			return nil, fmt.Errorf("error fetching data. status code: %d", response.StatusCode)
		}
	}
	var result []map[string]any
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println(ErrorStyle.Render(err.Error()))
	}
	return result, nil
}

func muxData(data map[string]any) {
	switch data["type"] {
	case "CommitCommentEvent":
		commitCommentEvent(data)
	case "CreateEvent":
		createEvent(data)
	case "DeleteEvent":
		deleteEvent(data)
	case "ForkEvent":
		forkEvent(data)
	case "IssueCommentEvent":
		issueCommentEvent(data)
	case "IssuesEvent":
		issuesEvent(data)
	case "MemberEvent":
		memberEvent(data)
	case "PublicEvent":
		publicEvent(data)
	case "PullRequestEvent":
		pullRequestEvent(data)
	case "PullRequestReviewEvent":
		pullRequestReviewEvent(data)
	case "PullRequestReviewCommentEvent":
		pullRequestReviewCommentEvent(data)
	case "PullRequestReviewThreadEvent":
		pullRequestReviewThreadEvent(data)
	case "PushEvent":
		pushEvent(data)
	case "ReleaseEvent":
		releaseEvent(data)
	case "SponsorshipEvent":
		sponsorshipEvent(data)
	case "WatchEvent":
		watchEvent(data)
	}
}

func commitCommentEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	comment := payload["comment"].(map[string]any)
	body := comment["body"].(string)
	action := payload["action"].(string)
	commitID := comment["commit_id"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	time := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("%s a comment on %s saying \"%s\" in %s at %s!",
			strings.ToUpper(string(action[0]))+action[1:], commitID, body, repoName, time))
	fmt.Println(result)
}

func createEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	createType := payload["ref_type"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	time := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("Created a %s in %s at %s!",
			createType, repoName, time))
	fmt.Println(result)
}

func deleteEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	createType := payload["ref_type"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	time := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("Deleted a %s in %s at %s!",
			createType, repoName, time))
	fmt.Println(result)
}

func forkEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	fork := payload["forkee"].(map[string]any)
	repository := data["repository"].(map[string]any)
	forkRepoName := fork["name"].(string)
	repoName := repository["name"].(string)
	time := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("Forked a repo from %s into %s at %s!",
			forkRepoName, repoName, time))
	fmt.Println(result)
}

func issueCommentEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	comment := payload["comment"].(map[string]any)
	body := comment["body"].(string)
	issue := payload["issue"].(map[string]any)
	url := issue["html_url"].(string)
	action := payload["action"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	time := data["created_at"].(string)

	var result string
	switch action {
	case "edited":
		prev := payload["changes"].(map[string]any)["body"].(map[string]any)["from"].(string)
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a comment on %s saying \"%s\" to\"%s\" in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, body, prev, repoName, time))
	default:
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a comment on %s saying \"%s\" in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, body, repoName, time))
	}
	fmt.Println(result)
}

func issuesEvent(data map[string]any) string {
	// Handle IssuesEvent
	return "Handled IssuesEvent"
}

func memberEvent(data map[string]any) string {
	// Handle MemberEvent
	return "Handled MemberEvent"
}

func publicEvent(data map[string]any) string {
	// Handle PublicEvent
	return "Handled PublicEvent"
}

func pullRequestEvent(data map[string]any) string {
	// Handle PullRequestEvent
	return "Handled PullRequestEvent"
}

func pullRequestReviewEvent(data map[string]any) string {
	// Handle PullRequestReviewEvent
	return "Handled PullRequestReviewEvent"
}

func pullRequestReviewCommentEvent(data map[string]any) string {
	// Handle PullRequestReviewCommentEvent
	return "Handled PullRequestReviewCommentEvent"
}

func pullRequestReviewThreadEvent(data map[string]any) string {
	// Handle PullRequestReviewThreadEvent
	return "Handled PullRequestReviewThreadEvent"
}

func pushEvent(data map[string]any) string {
	// Handle PushEvent
	return "Handled PushEvent"
}

func releaseEvent(data map[string]any) string {
	// Handle ReleaseEvent
	return "Handled ReleaseEvent"
}

func sponsorshipEvent(data map[string]any) string {
	// Handle SponsorshipEvent
	return "Handled SponsorshipEvent"
}

func watchEvent(data map[string]any) string {
	// Handle WatchEvent
	return "Handled WatchEvent"
}
