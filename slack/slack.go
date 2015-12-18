package slack

import "fmt"

const (
	ResponseTypeInChannel = "in_channel"
	ResponseTypeEphemeral = "ephemeral"
)

func Pre(text string) string {
	return fmt.Sprintf("```%s```", text)
}
