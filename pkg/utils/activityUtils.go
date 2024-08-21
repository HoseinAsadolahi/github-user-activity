package utils

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"net/http"
	"strings"
	"time"
)

var DashStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#DC143C")).
	Margin(0, 0, 0, 4)

var IdStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#DC143C")).
	Margin(0, 0, 0, 8)

var ContentStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFD700"))

var InfoStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#32CD32"))

var ErrorStyle = lipgloss.NewStyle().
	Bold(true).
	Italic(true).
	Foreground(lipgloss.Color("#FF0000"))

func DisplayInfo(username string, page int) {
	data, err := fetchData(username, page)
	if err != nil {
		if strings.Contains(err.Error(), "deadline") {
			fmt.Println(ErrorStyle.Render("Timeout! Please check your connection!"))
		} else {
			fmt.Println(ErrorStyle.Render(err.Error()))
		}
		return
	}
	for _, item := range data {
		muxData(item)
	}
}

func fetchData(username string, page int) ([]map[string]any, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Get(fmt.Sprintf("https://api.github.com/users/%s/events?page=%d",
		username, page))
	if err != nil {
		return nil, err
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
	createdAt := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("%s a comment on %s saying \"%s\" in %s at %s!",
			strings.ToUpper(string(action[0]))+action[1:], commitID, body, repoName, createdAt))
	fmt.Println(result)
}

func createEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	createType := payload["ref_type"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	createdAt := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("Created a %s in %s at %s!",
			createType, repoName, createdAt))
	fmt.Println(result)
}

func deleteEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	createType := payload["ref_type"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	createdAt := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("Deleted a %s in %s at %s!",
			createType, repoName, createdAt))
	fmt.Println(result)
}

func forkEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	fork := payload["forkee"].(map[string]any)
	repository := data["repo"].(map[string]any)
	forkRepoName := fork["name"].(string)
	repoName := repository["name"].(string)
	createdAt := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("Forked a repo from %s into %s at %s!",
			forkRepoName, repoName, createdAt))
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
	createdAt := data["created_at"].(string)

	var result string
	switch action {
	case "edited":
		updatedAt := data["updated_at"].(string)
		prev := payload["changes"].(map[string]any)["body"].(map[string]any)["from"].(string)
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a comment on %s saying \"%s\" to \"%s\" in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, body, prev, repoName, updatedAt))
	default:
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a comment on %s saying \"%s\" in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, body, repoName, createdAt))
	}
	fmt.Println(result)
}

func issuesEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	issue := payload["issue"].(map[string]any)
	url := issue["html_url"].(string)
	title := issue["title"].(string)
	body := issue["body"].(string)
	action := payload["action"].(string)
	createdAt := data["created_at"].(string)

	var result string
	switch action {
	case "edited":
		prevBody := payload["changes"].(map[string]any)["body"].(map[string]any)["from"].(string)
		prevTitle := payload["changes"].(map[string]any)["title"].(map[string]any)["from"].(string)
		updatedAt := data["updated_at"].(string)
		result = DashStyle.Render("-") + ContentStyle.Render(fmt.Sprintf(
			"%s an issue at %s"+IfThenElse(prevTitle == title, fmt.Sprintf(" with \"%s\" as title",
				title), fmt.Sprintf(" and changed its title from %s into %s",
				prevTitle, title)).(string)+IfThenElse(prevBody == body, " ",
				" and updated the description ").(string)+"at %s",
			strings.ToUpper(string(action[0]))+action[1:], url, updatedAt))
	default:
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf(
				"%s an issue at %s with \"%s\" as title at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, title, createdAt))
	}
	fmt.Println(result)
}

func memberEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	member := payload["member"].(map[string]any)
	memberName := member["name"].(string)
	action := payload["action"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	createdAt := data["created_at"].(string)

	var result string
	switch action {
	case "edited":
		role := payload["changes"].(map[string]any)["permission"].(map[string]any)["to"]
		prevRole := payload["changes"].(map[string]any)["old_permission"].(map[string]any)["from"].(string)
		updatedAt := data["updated_at"].(string)
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf(
				"%s member: github.com/%s "+
					IfThenElse(role != nil, fmt.Sprintf("from %s to %s ", prevRole, role.(string)),
						fmt.Sprintf("as %s ", role.(string))).(string)+"in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], memberName, repoName, updatedAt))
	case "removed":
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf(
				"%s member: github.com/%s in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], memberName, repoName, createdAt))
	default:
		role := payload["changes"].(map[string]any)["role_name"].(map[string]any)["to"].(string)
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf(
				"%s member: github.com/%s as %s in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], memberName, role, repoName, createdAt))
	}
	fmt.Println(result)
}

func publicEvent(data map[string]any) {
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	createdAt := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("Changed visability of %s from private into public at %s!",
			repoName, createdAt))
	fmt.Println(result)
}

func pullRequestEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	pullReq := payload["pull_request"].(map[string]any)
	url := pullReq["url"].(string)
	title := pullReq["title"].(string)
	action := payload["action"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	updatedAt := pullReq["updated_at"].(string)

	var result string
	switch action {
	case "edited":
		prevTitle := payload["changes"].(map[string]any)["title"].(map[string]any)["from"].(string)
		result = DashStyle.Render("-") + ContentStyle.Render(fmt.Sprintf(
			"%s a pull request:%s "+IfThenElse(prevTitle == title, fmt.Sprintf("with \"%s\" as title",
				title), fmt.Sprintf(" and changed its title from %s into %s",
				prevTitle, title)).(string)+"in %s at %s",
			strings.ToUpper(string(action[0]))+action[1:], url, repoName, updatedAt))
	case "dequeued":
		reason := payload["reason"].(string)
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a pull request: %s with \"%s\" as title in %s because %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, title, repoName, reason, updatedAt))
	default:
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a pull request: %s with \"%s\" as title in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, title, repoName, updatedAt))
	}
	fmt.Println(result)
}

func pullRequestReviewEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	body := payload["review"].(map[string]any)["body"].(string)
	pullReq := payload["pull_request"].(map[string]any)
	url := pullReq["url"].(string)
	action := payload["action"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	createdAt := data["created_at"].(string)

	var result string
	switch action {
	case "edited":
		updatedAt := data["updated_at"].(string)
		prev := payload["changes"].(map[string]any)["body"].(map[string]any)["from"].(string)
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a comment on %s saying \"%s\" to \"%s\" in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, body, prev, repoName, updatedAt))
	default:
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a review on %s saying \"%s\" in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, body, repoName, createdAt))
	}
	fmt.Println(result)
}

func pullRequestReviewCommentEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	comment := payload["comment"].(map[string]any)
	body := comment["body"].(string)
	pullReq := payload["pull_request"].(map[string]any)
	url := pullReq["url"].(string)
	action := payload["action"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	createdAt := data["created_at"].(string)

	var result string
	switch action {
	case "edited":
		prev := payload["changes"].(map[string]any)["body"].(map[string]any)["from"].(string)
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a comment on %s saying \"%s\" to \"%s\" in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, body, prev, repoName, createdAt))
	default:
		result = DashStyle.Render("-") +
			ContentStyle.Render(fmt.Sprintf("%s a comment on %s saying \"%s\" in %s at %s!",
				strings.ToUpper(string(action[0]))+action[1:], url, body, repoName, createdAt))
	}
	fmt.Println(result)
}

func pullRequestReviewThreadEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	pullReq := payload["pull_request"].(map[string]any)
	url := pullReq["url"].(string)
	action := payload["action"].(string)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	updatedAt := data["updated_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("Marked pull request:%s as %s in %s at %s!",
			url, action, repoName, updatedAt))
	fmt.Println(result)
}

func pushEvent(data map[string]any) {
	payload := data["payload"].(map[string]any)
	commits := payload["commits"].([]any)
	repository := data["repo"].(map[string]any)
	repoName := repository["name"].(string)
	createdAt := data["created_at"].(string)

	result := DashStyle.Render("-") +
		ContentStyle.Render(fmt.Sprintf("Pushed %d commits in %s at %s shown below:",
			len(commits), repoName, createdAt))
	fmt.Println(result)
	for _, commit := range commits {
		converted := commit.(map[string]any)
		message := converted["message"].(string)
		sha := converted["sha"].(string)
		author := converted["author"].(map[string]any)["email"].(string)
		fmt.Println(IdStyle.Render(sha[0:7]+": ") + ContentStyle.Render(message+" By "+author))
	}
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

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
