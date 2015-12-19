package responders

import (
	"fmt"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/nmonterroso/cowsay"
	"github.com/nmonterroso/cowsay-slackapp/models"
	"github.com/nmonterroso/cowsay-slackapp/restapi/operations"
	"github.com/nmonterroso/cowsay-slackapp/slack"
	"net/http"
)

func CowsayResponder(params operations.CowsayParams) middleware.Responder {
	cow, err := cowsay.ParseArgs(params.Text)

	if err != nil {
		return &operations.CowsayDefault{
			&models.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
	}

	message, err := cow.Speak()

	if err != nil {
		message = err.Error()
		code := http.StatusInternalServerError

		if err == cowsay.CowNotFound {
			message = fmt.Sprintf("%s not found", cow.Options.Cow)
			code = http.StatusNotFound
		}

		return &operations.CowsayDefault{
			&models.Error{
				Code:    int32(code),
				Message: message,
			},
		}
	}

	responseType := slack.ResponseTypeInChannel
	if cow.Options.List || cow.IsHelper {
		responseType = slack.ResponseTypeEphemeral
	}

	return &operations.CowsayOK{
		Payload: &models.SLACKResponse{
			Text:         slack.Pre(message),
			ResponseType: responseType,
		},
	}
}
