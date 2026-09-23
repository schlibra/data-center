package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test command",
	Run: func(cmd *cobra.Command, args []string) {
		count := 0
		for {
			client := resty.New()
			resp, err := client.
				R().
				SetAuthScheme("Bearer").
				SetAuthToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJ1c2VybmFtZSI6InNjaGxpYnJhIiwiZW5hYmxlIjoxLCJhZG1pbiI6MSwidG9rZW5faWQiOiI3YzYxZDM0ZjBhY2ZlMTY4MjhjNGU1YTI5MzZhMTM5ZWNlMDUyMzVhYWE4Njg0MmI1YWQ1YmVmOThkOWM5Zjk0IiwiaXNzIjoiZnJwLWF1dGgiLCJleHAiOjE3OTAyMzEzMjYsIm5iZiI6MTc5MDE0NDkyNiwiaWF0IjoxNzkwMTQ0OTI2fQ.wr8BX3hQN4IglokCeY2VADGoRRHp9ZHDevcJB9OEx0A").
				Get("http://127.0.0.1:49180/api/frp/token")
			if err != nil {
				fmt.Printf("运行次数：%d，错误：%s\n", count, err)
			}
			count++
			fmt.Printf("运行次数：%d，响应消息：%s\n", count, resp.String()[0:15])
		}
	},
}
