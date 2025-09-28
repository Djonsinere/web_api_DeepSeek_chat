
# AI API Chat Web (DeepSeek)

This is a simple web chat application built with Go, Gin, and DeepSeek AI API. It provides a web interface for chatting with an AI model via OpenRouter's DeepSeek endpoint.

## Project Structure

```
├── go.mod
├── go.sum
├── main.go              # Main server entrypoint
├── src/
│   ├── ai_cals.go       # DeepSeek API call logic
│   └── config.go        # API key config (see below)
├── templates/
│   └── index.html       # Web chat UI
└── README.md
```

## Setup

1. **Clone the repository:**
	```sh
	git clone https://github.com/Djonsinere/web_api_DeepSeek_chat.git
	cd web_api_DeepSeek_chat
	```

2. **Install dependencies:**
	```sh
	go mod tidy
	```

3. **Set your DeepSeek/OpenRouter API key:**
	- Create `src/config.go` (or edit it) and add:
	  ```go
	  package src
	  const Api_key = "your_openrouter_api_key"
	  ```

4. **Run the server:**
	```sh
	go run main.go
	```
	The server will start on [http://localhost:8080](http://localhost:8080)

## API

- **POST** `/api/chat`
  - Request JSON: `{ "message": "your message" }`
  - Response JSON: `{ "reply": "AI response" }`

## Notes
- Requires a valid OpenRouter/DeepSeek API key.
- The frontend is a simple HTML/JS page in `templates/index.html`.
- All AI logic is in `src/ai_cals.go`.

## License
MIT
