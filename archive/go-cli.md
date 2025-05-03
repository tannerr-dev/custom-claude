
# can you help me write a golang CLI, that takes an argument, uses it as a prompt for a api call to claude, then returns the reponse to the console? thank you

**1745552944180**

---
# Writing a Golang CLI for Claude API Interaction

I'll help you create a Golang CLI application that takes a command-line argument, sends it as a prompt to Claude via the Anthropic API, and returns the response to the console.

## Prerequisites

1. Go installed on your system
2. An Anthropic API key (for Claude)

## Step-by-Step Implementation

### 1. Setup Project

```bash
mkdir claude-cli
cd claude-cli
go mod init claude-cli
```

### 2. Install Dependencies

```bash
go get github.com/spf13/cobra
go get github.com/spf13/viper
```

### 3. Create Main Application

Create a file named `main.go`:

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type AnthropicRequest struct {
	Model     string  `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int     `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AnthropicResponse struct {
	Content []ContentBlock `json:"content"`
	Error   struct {
		Message string `json:"message"`
	} `json:"error"`
}

type ContentBlock struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
}

func main() {
	var apiKey string
	var maxTokens int
	var model string

	rootCmd := &cobra.Command{
		Use:   "claude-cli [prompt]",
		Short: "Send a prompt to Claude and get a response",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			prompt := strings.Join(args, " ")
			response, err := askClaude(prompt, apiKey, maxTokens, model)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(response)
		},
	}

	rootCmd.Flags().StringVar(&apiKey, "api-key", "", "Anthropic API key (required)")
	rootCmd.Flags().IntVar(&maxTokens, "max-tokens", 1000, "Maximum number of tokens in the response")
	rootCmd.Flags().StringVar(&model, "model", "claude-3-opus-20240229", "Claude model to use")
	
	rootCmd.MarkFlagRequired("api-key")

	// Allow environment variable for API key
	if apiKey == "" {
		apiKey = os.Getenv("ANTHROPIC_API_KEY")
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func askClaude(prompt, apiKey string, maxTokens int, model string) (string, error) {
	if apiKey == "" {
		return "", fmt.Errorf("API key is required. Set it with --api-key or ANTHROPIC_API_KEY environment variable")
	}

	requestBody := AnthropicRequest{
		Model: model,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens: maxTokens,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var response AnthropicResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	if response.Error.Message != "" {
		return "", fmt.Errorf("API error: %s", response.Error.Message)
	}

	// Extract the text from the response
	var result string
	for _, block := range response.Content {
		if block.Type == "text" {
			result += block.Text
		}
	}

	return result, nil
}
```

### 4. Build the Application

```bash
go build -o claude-cli
```

### 5. Using the CLI

You can use the CLI in two ways:

**With command-line flag:**
```bash
./claude-cli --api-key="your-api-key-here" "What is the capital of France?"
```

**With environment variable:**
```bash
export ANTHROPIC_API_KEY="your-api-key-here"
./claude-cli "What is the capital of France?"
```

You can also customize additional parameters:
```bash
./claude-cli --max-tokens=2000 --model="claude-3-haiku-20240307" "Explain quantum computing"
```

## Additional Improvements You Could Make

1. **Add configuration file support** - Use Viper to load settings from a config file
2. **Implement stream mode** - Support streaming responses from Claude
3. **Add colorful output** - Use a package like `fatih/color` for nicer console output
4. **Support conversation context** - Store and send previous messages for continued conversations
5. **Add more parameters** - Support temperature, top_p, and other API parameters

This CLI provides a basic but functional way to interact with Claude from your command line. You can expand on it based on your specific needs.
---

