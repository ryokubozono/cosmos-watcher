package watcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetProposals(status string) (*ProposalResponse, error) {
	req, err := http.NewRequest(http.MethodGet, mustGetenv("COSMOS_URL")+"?proposal_status="+status, nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 200 OK 以外の場合はエラーメッセージを表示して終了
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Response Body を読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// JSONを構造体にエンコード
	var res ProposalResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
