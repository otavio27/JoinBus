package middlewares

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/otavio27/JoinBus-APP/back-end/contracts"
	"github.com/otavio27/JoinBus-APP/back-end/domain"
)

// HandleError is responsible for converting domain errors to HTTP errors
// simplifying error handling overall.
type ErrorHandler struct {
	logger contracts.LogProvider
}

func NewErrorHandler(logger contracts.LogProvider) ErrorHandler {
	return ErrorHandler{
		logger: logger,
	}
}

func (e ErrorHandler) Middleware(c *fiber.Ctx) error {
	err := c.Next()
	if err == nil {
		return nil
	}

	req := c.Request()
	status, body := handleDomainErrAsHTTP(
		c.Context(),
		e.logger,
		err,
		string(req.Header.Method()),
		string(req.RequestURI()),
	)
	c.Status(status).Send(body)
	return nil
}

func handleDomainErrAsHTTP(ctx context.Context, logger contracts.LogProvider, err error, method string, path string) (status int, responseBody []byte) {
	domainErr := domain.AsDomainErr(err)

	response := map[string]any{
		"code":  domainErr.Code,
		"title": domainErr.Title,
	}

	switch domainErr.Code {
	case "InternalErr":
		status = 500

		data := contracts.LogBody{
			"route": method + ": " + path,
		}
		for k, v := range domainErr.Data {
			data[k] = v
		}
		logger.Error(ctx, "request-error", data)

	case "BadRequest":
		status = 400
		for k, v := range domainErr.Data {
			response[k] = v
		}

	case "NotFoundErr":
		status = 404
		for k, v := range domainErr.Data {
			response[k] = v
		}
	}

	responseBody, _ = json.Marshal(response)
	return status, responseBody
}
