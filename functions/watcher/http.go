package watcher

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

var (
	PROPOSAL_STATUS_DEPOSIT_PERIOD = "1"
	PROPOSAL_STATUS_VOTING_PERIOD  = "2"
	PROPOSAL_STATUS_PASSED         = "3"
	PROPOSAL_STATUS_REPJECTED      = "4"
	PROPOSAL_STATUS_FAILED         = "5"
)

func init() {
	functions.HTTP("WatcherHTTP", WatcherHTTP)
}

// WatcherHTTP is an HTTP Cloud Function with a request parameter.
func WatcherHTTP(w http.ResponseWriter, r *http.Request) {
	internal := mustGetenv("INTERVAL_HOURS")
	intervalFloat, err := strconv.ParseFloat(internal, 64)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	params := []struct {
		Status string
	}{
		{Status: PROPOSAL_STATUS_DEPOSIT_PERIOD},
		{Status: PROPOSAL_STATUS_VOTING_PERIOD},
		// {Status: PROPOSAL_STATUS_PASSED}, // 件数が増え続けるためコメントアウト
		// {Status: PROPOSAL_STATUS_REPJECTED}, // 件数が増え続けるためコメントアウト
		// {Status: PROPOSAL_STATUS_FAILED}, // 件数が増え続けるためコメントアウト
	}
	for _, p := range params {
		res, err := GetProposals(p.Status)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var messages []string
		for _, proposal := range res.Proposals {
			var header string
			switch p.Status {
			case PROPOSAL_STATUS_DEPOSIT_PERIOD:
				header = "[Proposal in deposit period]"
				if time.Now().Sub(ParseDatetime(proposal.SubmitTime)).Hours() > intervalFloat {
					continue
				}
				messages = append(messages, fmt.Sprintf("%s\nProposalID: %s%s\nTitle: %s\nSubmitted: %s\n", header, mustGetenv("EXPLORER_URL"), proposal.ProposalID, proposal.Content.Title, proposal.SubmitTime))
			case PROPOSAL_STATUS_VOTING_PERIOD:
				header = "[Proposal in voting period]"
				if time.Now().Sub(ParseDatetime(proposal.VotingStartTime)).Hours() > intervalFloat {
					continue
				}
				messages = append(messages, fmt.Sprintf("%s\nProposalID: %s%s\nTitle: %s\nSubmitted: %s\nVote Period: %s ~ %s\n", header, mustGetenv("EXPLORER_URL"), proposal.ProposalID, proposal.Content.Title, proposal.SubmitTime, proposal.VotingStartTime, proposal.VotingEndTime))
				// 件数が増え続けるためコメントアウト
			// case PROPOSAL_STATUS_PASSED:
			// 	header = "[Proposal of passed]"
			// 	if time.Now().Sub(ParseDatetime(proposal.VotingEndTime)).Hours() > intervalFloat {
			// 		continue
			// 	}
			// case PROPOSAL_STATUS_REPJECTED:
			// 	header = "[Proposal of rejected]"
			// 	if time.Now().Sub(ParseDatetime(proposal.VotingEndTime)).Hours() > intervalFloat {
			// 		continue
			// 	}
			// case PROPOSAL_STATUS_FAILED:
			// 	header = "[Proposal of failed]"
			// 	if time.Now().Sub(ParseDatetime(proposal.VotingEndTime)).Hours() > intervalFloat {
			// 		continue
			// 	}
			default:
				header = "[Proposal]"
				continue
			}
		}
		for _, m := range messages {
			if err := Notify(m); err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

// parse string "2023-12-12T08:56:53.722526361Z" to time.Time
func ParseDatetime(datetime string) time.Time {
	t, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		panic(err)
	}
	return t
}
