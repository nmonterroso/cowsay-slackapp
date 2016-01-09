package responders

import (
	"fmt"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/nmonterroso/cowsay"
	"github.com/nmonterroso/cowsay-slackapp/models"
	"github.com/nmonterroso/cowsay-slackapp/restapi/operations/defaultop"
	"github.com/nmonterroso/cowsay-slackapp/slack"
	"net/http"
)

func CowsayResponder(params defaultop.CowsayParams) middleware.Responder {
	cow, err := cowsay.ParseArgs(params.Text)

	if cow.IsHelper {
		cow.Options.Random = true
	}

	if err != nil {
		return defaultop.NewCowsayDefault(http.StatusInternalServerError).WithPayload(
			&models.Error{
				Message: err.Error(),
			})
	}

	message, err := cow.Speak()

	if err != nil {
		message = err.Error()
		code := http.StatusInternalServerError

		if err == cowsay.CowNotFound {
			message = fmt.Sprintf("%s not found", cow.Options.Animal)
			code = http.StatusNotFound
		}

		return defaultop.NewCowsayDefault(code).WithPayload(
			&models.Error{
				Message: message,
			})
	}

	responseType := slack.ResponseTypeInChannel
	if cow.Options.List || cow.IsHelper {
		responseType = slack.ResponseTypeEphemeral
	}

	return defaultop.NewCowsayOK().WithPayload(
		&models.SLACKResponse{
			Text:         slack.Pre(message),
			ResponseType: responseType,
		})
}
