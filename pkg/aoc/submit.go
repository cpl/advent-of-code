package aoc

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type SubmitResponse byte

func (response SubmitResponse) String() string {
	switch response {
	case SubmitResponseUndefined:
		return "undefined"
	case SubmitResponseError:
		return "error"
	case SubmitResponseSuccess:
		return "success"
	case SubmitResponseWrong:
		return "wrong"
	case SubmitResponseTooHigh:
		return "too_high"
	case SubmitResponseTooLow:
		return "too_low"
	}

	return "unknown"
}

func (response SubmitResponse) Emoji() string {
	switch response {
	case SubmitResponseUndefined:
		return "❓"
	case SubmitResponseError:
		return "🛑"
	case SubmitResponseSuccess:
		return "⭐️"
	case SubmitResponseWrong:
		return "😕"
	case SubmitResponseTooHigh:
		return "📈"
	case SubmitResponseTooLow:
		return "📉"
	}

	return "⁉️"
}

const (
	SubmitResponseUndefined SubmitResponse = iota
	SubmitResponseError
	SubmitResponseSuccess
	SubmitResponseWrong
	SubmitResponseTooHigh
	SubmitResponseTooLow
)

var submitTicker *time.Ticker

func init() {
	submitTicker = time.NewTicker(time.Second)
}

func Submit(year, day, part int, answer any) (SubmitResponse, error) {
	if submitSuccessExists(year, day, part) {
		return SubmitResponseSuccess, nil
	}

	response, err := submitInner(year, day, part, answer, 1)
	if err != nil {
		return response, err
	}

	if response == SubmitResponseSuccess {
		submitSuccessWrite(year, day, part, answer)
	}

	return response, nil
}

func submitInner(year, day, part int, answer any, attempt int) (SubmitResponse, error) {
	if attempt == 3 {
		return SubmitResponseError, fmt.Errorf("too many attempts")
	}

	var answerStr string
	switch v := answer.(type) {
	case string:
		answerStr = v
	case int:
		answerStr = strconv.Itoa(v)
	default:
		panic("invalid answer type")
	}

	body := strings.NewReader("level=" + strconv.Itoa(part) + "&answer=" + answerStr)

	r, err := newRequest("POST", fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day), body)
	if err != nil {
		return SubmitResponseError, fmt.Errorf("creating puzzle download request: %w", err)
	}

	<-submitTicker.C

	resp, err := httpClient.Do(r)
	if err != nil {
		return SubmitResponseError, fmt.Errorf("submitting answer (%d/%d) (%v): %w", year, day, answer, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return SubmitResponseError, fmt.Errorf("cannot submit answer (%d/%d) (%v): %w", year, day, answer, err)
	}

	datas := string(data)

	if idx := strings.Index(datas, "you have to wait after submitting an answer before trying again. You have "); idx != -1 {
		idx2 := strings.Index(datas[idx:], "s")
		if idx2 == -1 {
			return SubmitResponseError, fmt.Errorf("bad response: %s", datas)
		}

		wait, err := time.ParseDuration(datas[idx : idx2+1])
		if err != nil {
			return SubmitResponseError, fmt.Errorf("cannot parse wait duration: %w", err)
		}

		time.Sleep(wait + time.Second)

		return submitInner(year, day, part, answerStr, attempt+1)
	}

	if idx := strings.Index(datas, "your answer is too high"); idx != -1 {
		return SubmitResponseTooHigh, nil
	}

	if idx := strings.Index(datas, "your answer is too low"); idx != -1 {
		return SubmitResponseTooLow, nil
	}

	if idx := strings.Index(datas, "not the right answer"); idx != -1 {
		return SubmitResponseWrong, nil
	}

	return SubmitResponseSuccess, nil
}
