package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func fetchData(username string, page int) (map[string]any, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events?page=%d", username, page))
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		if response.StatusCode == 404 {
			return nil, fmt.Errorf("username not found")
		} else {
			return nil, fmt.Errorf("error fetching data. status code: %d", response.StatusCode)
		}
	}
	var result map[string]any
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func muxData(data map[string]any) string {
	switch data["type"] {
	case "CommitCommentEvent":
		break
	case "CreateEvent":
		break
	case "DeleteEvent":
		break
	case "ForkEvent":
		break
	case "IssueCommentEvent":
		break
	case "IssuesEvent":
		break
	case "MemberEvent":
		break
	case "PublicEvent":
		break
	case "PullRequestEvent":
		break
	case "PullRequestReviewEvent":
		break
	case "PullRequestReviewCommentEvent":
		break
	case "PullRequestReviewThreadEvent":
		break
	case "PushEvent":
		break
	case "ReleaseEvent":
		break
	case "SponsorshipEvent":
		break
	case "WatchEvent":
		break
	}
	return "error multiplexing github event"
}

func commitCommentEvent(data map[string]any) string {
	// Handle CommitCommentEvent
	return "Handled CommitCommentEvent"
}

func createEvent(data map[string]any) string {
	// Handle CreateEvent
	return "Handled CreateEvent"
}

func deleteEvent(data map[string]any) string {
	// Handle DeleteEvent
	return "Handled DeleteEvent"
}

func forkEvent(data map[string]any) string {
	// Handle ForkEvent
	return "Handled ForkEvent"
}

func issueCommentEvent(data map[string]any) string {
	// Handle IssueCommentEvent
	return "Handled IssueCommentEvent"
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
