package main

import (
	"encoding/json"
	// "flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type CompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
	Echo        bool    `json:"echo"`
}

type CompletionResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Missing OpenAI API key")
		os.Exit(1)
	}

	var prompt string
	if len(os.Args) > 1 {
		prompt = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Fprintln(os.Stderr, "Usage: ./myprogram.go <prompt>")
		os.Exit(1)
	}
	
	prompt = "Provide only the appropriate git commands for:" + prompt

	reqBody := CompletionRequest{
		Model:       "text-davinci-003",
		Prompt:      prompt,
		MaxTokens:   60,
		Temperature: 0.5,
		TopP:        1.0,
		Echo:        false,
	}
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", strings.NewReader(string(reqJSON)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error: %s\n", resp.Status)
		os.Exit(1)
	}

	var respBody CompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	text := respBody.Choices[0].Text

	codeRegex := regexp.MustCompile("`([^`]*)`")
	match := codeRegex.FindStringSubmatch(text)
	if match != nil {
		code := match[1]
		runWithConfirmation(code)
	} else {
		runWithConfirmation(text)
	}
}

func runWithConfirmation(command string) {
	fmt.Printf("Would you like to run the following command: %s [y/N] ", command)
	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.ToLower(confirmation) == "y" {
		cmd := exec.Command("sh", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Println("Command not executed.")
	}
}
