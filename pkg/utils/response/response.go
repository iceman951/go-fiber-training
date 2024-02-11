package response

import (
	"calcal/pkg/utils/constant"
	"calcal/pkg/utils/library"

	"strings"

	"github.com/baac-tech/zlogwrap"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func SuccessResponse(c *fiber.Ctx, msg string, data interface{}) error { // 200
	statusCode := fiber.StatusOK
	res := ResponseMessageStatusData{
		Status:    fasthttp.StatusMessage(statusCode),
		Message:   msg,
		Data:      data,
		Timestamp: library.GetTimestamp(),
	}

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "SuccessResponse",
		Context:     c,
	})
	logger.SetField("event", map[string]interface{}{
		"name":        constant.EVENT_TYPE_RESPONSE,
		"status_code": statusCode,
		"response":    res,
	}).Debug()
	return c.Status(statusCode).JSON(res)
}

func CreatedResponse(c *fiber.Ctx) error { // 201
	statusCode := fiber.StatusCreated
	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "CreatedResponse",
		Context:     c,
	})
	logger.SetField("event", map[string]interface{}{
		"name":        constant.EVENT_TYPE_RESPONSE,
		"status_code": statusCode,
	}).Debug()
	return c.SendStatus(statusCode)
}

func NoContentResponse(c *fiber.Ctx) error { // 204
	statusCode := fiber.StatusNoContent
	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "NoContentResponse",
		Context:     c,
	})
	logger.SetField("event", map[string]interface{}{
		"name":        constant.EVENT_TYPE_RESPONSE,
		"status_code": statusCode,
	}).Debug()
	return c.SendStatus(statusCode)
}

func BadRequestResponse(c *fiber.Ctx, msg ...string) error { // 400
	statusCode := fiber.StatusBadRequest
	if len(msg) == 0 {
		msg = append(msg, fasthttp.StatusMessage(statusCode))
	}
	res := ResponseMessageStatus{
		Status:  fiber.ErrBadRequest.Message,
		Message: strings.Join(msg, " "),
	}

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "BadRequestResponse",
		Context:     c,
	})
	logger.SetField("event", map[string]interface{}{
		"name":        constant.EVENT_TYPE_RESPONSE,
		"status_code": statusCode,
		"response":    res,
	}).Debug()
	return c.Status(statusCode).JSON(res)
}

func UnauthorizedResponse(c *fiber.Ctx, msg ...string) error { // 401
	statusCode := fiber.StatusUnauthorized
	if len(msg) == 0 {
		msg = append(msg, fasthttp.StatusMessage(statusCode))
	}
	res := ResponseMessageStatus{
		Status:  fiber.ErrUnauthorized.Message,
		Message: strings.Join(msg, " "),
	}

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "UnauthorizedResponse",
		Context:     c,
	})
	logger.SetField("event", map[string]interface{}{
		"name":        constant.EVENT_TYPE_RESPONSE,
		"status_code": statusCode,
		"response":    res,
	}).Debug()
	return c.Status(statusCode).JSON(res)
}

func ForbiddenResponse(c *fiber.Ctx, msg ...string) error { // 403
	statusCode := fiber.StatusForbidden
	if len(msg) == 0 {
		msg = append(msg, fasthttp.StatusMessage(statusCode))
	}
	res := ResponseMessageStatus{
		Status:  fiber.ErrForbidden.Message,
		Message: strings.Join(msg, " "),
	}

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "ForbiddenResponse",
		Context:     c,
	})
	logger.SetField("event", map[string]interface{}{
		"name":        constant.EVENT_TYPE_RESPONSE,
		"status_code": statusCode,
		"response":    res,
	}).Debug()
	return c.Status(statusCode).JSON(res)
}

func NotFoundResponse(c *fiber.Ctx, msg ...string) error { // 404
	statusCode := fiber.StatusNotFound
	if len(msg) == 0 {
		msg = append(msg, fasthttp.StatusMessage(statusCode))
	}
	res := ResponseMessageStatus{
		Status:  fiber.ErrNotFound.Message,
		Message: strings.Join(msg, " "),
	}

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "NotFoundResponse",
		Context:     c,
	})
	logger.SetField("event", map[string]interface{}{
		"name":        constant.EVENT_TYPE_RESPONSE,
		"status_code": statusCode,
		"response":    res,
	}).Debug()
	return c.Status(statusCode).JSON(res)
}

func InternalServerError(c *fiber.Ctx, msg ...string) error { // 500
	statusCode := fiber.StatusInternalServerError
	if len(msg) == 0 {
		msg = append(msg, fasthttp.StatusMessage(statusCode))
	}
	res := ResponseMessageStatus{
		Status:  fiber.ErrInternalServerError.Message,
		Message: strings.Join(msg, " "),
	}

	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "InternalServerError",
		Context:     c,
	})
	logger.SetField("event", map[string]interface{}{
		"name":        constant.EVENT_TYPE_RESPONSE,
		"status_code": statusCode,
		"response":    res,
	}).Debug()
	return c.Status(statusCode).JSON(res)
}
