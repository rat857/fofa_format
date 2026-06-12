package search

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"fofa_format/osDo"
	"fofa_format/template"
	"net/url"
	"strings"
	"time"

	"github.com/fatih/color"
)

const defaultAPIBase = "https://fofa.info"

var apiBase = defaultAPIBase

func SetAPIBase(base string) {
	base = strings.TrimSpace(base)
	base = strings.TrimRight(base, "/")
	if base == "" {
		apiBase = defaultAPIBase
		return
	}
	if !strings.HasPrefix(base, "http://") && !strings.HasPrefix(base, "https://") {
		base = "https://" + base
	}
	apiBase = base
}

func GetAPIBase() string {
	if apiBase == "" {
		return defaultAPIBase
	}
	return apiBase
}

func searchAllURL() string {
	return GetAPIBase() + "/api/v1/search/all"
}

func statsURL() string {
	return GetAPIBase() + "/api/v1/search/stats"
}

// FetchTargets runs a single FOFA search and returns nuclei-ready targets.
func FetchTargets(email, key, query string, size int) ([]string, error) {
	if strings.TrimSpace(query) == "" {
		return nil, fmt.Errorf("fofa query is empty")
	}
	if size < 1 {
		size = 1
	}
	if size > 10000 {
		size = 10000
	}

	color.Yellow("开始查询具体资产...")
	color.Green("%s", query)

	fields := "host"
	qbase64 := base64.StdEncoding.EncodeToString([]byte(query))
	params := url.Values{}
	if strings.TrimSpace(email) != "" {
		params.Set("email", email)
	}
	params.Set("key", key)
	params.Set("qbase64", qbase64)
	params.Set("fields", fields)
	params.Set("size", fmt.Sprintf("%d", size))

	requestURL := searchAllURL() + "?" + params.Encode()
	var lastErr error

	for i := 0; i < 5; i++ {
		payload, err := requestSearchPayload(requestURL)
		if err != nil {
			lastErr = err
			if strings.Contains(err.Error(), "速度") {
				color.Red(err.Error())
				time.Sleep(time.Duration(3+i*2) * time.Second)
				continue
			}
			return nil, err
		}
		if payload.Error {
			lastErr = fmt.Errorf("%s", payload.Errmsg)
			if strings.Contains(payload.Errmsg, "速度") {
				color.Red(payload.Errmsg)
				time.Sleep(time.Duration(3+i*2) * time.Second)
				continue
			}
			return nil, lastErr
		}
		results, err := parseSearchResults(payload.Results)
		if err != nil {
			return nil, err
		}
		formatted := osDo.Format(results)
		color.Green("FOFA 原始结果 %d 条，格式化去重后 %d 条", len(results), len(formatted))
		return formatted, nil
	}
	return nil, lastErr
}

type searchPayload struct {
	Error   bool   `json:"error"`
	Errmsg  string `json:"errmsg"`
	Results any    `json:"results"`
}

func requestSearchPayload(requestURL string) (searchPayload, error) {
	var payload searchPayload
	var lastErr error
	for i := 0; i < 5; i++ {
		raw := template.RestyStruct[searchPayload](requestURL)
		if !raw.Error {
			return raw, nil
		}
		lastErr = fmt.Errorf("%s", raw.Errmsg)
		if strings.Contains(raw.Errmsg, "速度") {
			time.Sleep(time.Duration(3+i*2) * time.Second)
			continue
		}
		return raw, lastErr
	}
	return payload, lastErr
}

// FetchTargetsSmart reuses the batching strategy from fofa_format for large result sets.
func FetchTargetsSmart(email, key, query string, maxTotal int) ([]string, error) {
	if maxTotal <= 0 {
		maxTotal = 10000
	}
	if maxTotal <= 10000 {
		targets, err := FetchTargets(email, key, query, maxTotal)
		if err != nil {
			return nil, err
		}
		if len(targets) > maxTotal {
			return targets[:maxTotal], nil
		}
		return targets, nil
	}

	base64AllList, _ := GetAllBase64(email, key)
	if len(base64AllList) == 0 {
		return FetchTargets(email, key, query, maxTotal)
	}

	seen := make(map[string]struct{})
	var all []string
	for _, qbase64 := range base64AllList {
		if len(all) >= maxTotal {
			break
		}
		remain := maxTotal - len(all)
		if remain > 10000 {
			remain = 10000
		}
		params := url.Values{}
		if strings.TrimSpace(email) != "" {
			params.Set("email", email)
		}
		params.Set("key", key)
		params.Set("qbase64", qbase64)
		params.Set("fields", "host")
		params.Set("size", fmt.Sprintf("%d", remain))

		requestURL := searchAllURL() + "?" + params.Encode()
		data := template.RestyStruct[template.Data](requestURL)
		if data.Error {
			return nil, fmt.Errorf("%s", data.Errmsg)
		}
		results, err := parseSearchResults(data.Results)
		if err != nil {
			return nil, err
		}
		for _, target := range osDo.Format(results) {
			if _, ok := seen[target]; ok {
				continue
			}
			seen[target] = struct{}{}
			all = append(all, target)
			if len(all) >= maxTotal {
				break
			}
		}
		time.Sleep(3 * time.Second)
	}
	if len(all) == 0 {
		return nil, fmt.Errorf("no targets returned from fofa_format smart search")
	}
	return all, nil
}

func parseSearchResults(raw any) ([][]string, error) {
	switch v := raw.(type) {
	case nil:
		return nil, nil
	case [][]string:
		return v, nil
	case []any:
		if len(v) == 0 {
			return nil, nil
		}
		if _, ok := v[0].([]any); ok {
			out := make([][]string, 0, len(v))
			for _, item := range v {
				rowAny, ok := item.([]any)
				if !ok {
					continue
				}
				row := make([]string, len(rowAny))
				for i, cell := range rowAny {
					row[i] = fmt.Sprint(cell)
				}
				out = append(out, row)
			}
			return out, nil
		}
		out := make([][]string, 0, len(v))
		for _, item := range v {
			out = append(out, []string{fmt.Sprint(item)})
		}
		return out, nil
	case []string:
		out := make([][]string, 0, len(v))
		for _, item := range v {
			out = append(out, []string{item})
		}
		return out, nil
	default:
		body, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("marshal fofa results: %w", err)
		}
		var matrix [][]string
		if err := json.Unmarshal(body, &matrix); err == nil && len(matrix) > 0 {
			return matrix, nil
		}
		var flat []string
		if err := json.Unmarshal(body, &flat); err == nil {
			out := make([][]string, 0, len(flat))
			for _, item := range flat {
				out = append(out, []string{item})
			}
			return out, nil
		}
		return nil, fmt.Errorf("unknown fofa results format")
	}
}
