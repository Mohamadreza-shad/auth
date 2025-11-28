package http

import (
	"encoding/json"
	"errors"
	"github.com/Mohamadreza-shad/auth/pkg/exception"
	"github.com/Mohamadreza-shad/auth/pkg/i18n"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ContentTypeMessage     = "Content-Type"
	ApplicationJsonMessage = "application/json"
)

type ResponseFailure struct {
	Error  ErrorDetail `json:"error"`
	Status int         `json:"status"`
}

type ErrorDetail struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type SuccessResponse struct {
	Result   interface{} `json:"result"`
	MetaData interface{} `json:"meta_data"`
	Status   int64       `json:"status"`
}

func MakeSuccessResponse(w http.ResponseWriter, data SuccessResponse) {
	jData, err := json.Marshal(&data)
	if err != nil {
		MakeErrorResponseWithCode(w, http.StatusInternalServerError, int(exception.ErrCodeSomethingWentWrong), "something went wrong")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set(ContentTypeMessage, ApplicationJsonMessage)
	_, _ = w.Write(jData)
}

func MakeErrorLocalizedResponse(c *gin.Context, err error, localize i18n.I18n, lang string) {
	httpStatus := 500
	errCode := 500
	message := "SOMETHING_WENT_WRONG"
	var locErr error

	var gerr exception.GeneralError
	if errors.As(err, &gerr) {
		message, locErr = localize.GetLocalizedMessage(gerr.GetMessageKey(), lang)
		if locErr != nil {
			message = gerr.Error()
		}
		errCode = int(gerr.GetErrorCode())
		httpStatus = gerr.GetHttpStatusCode()
	}

	responseJson := ResponseFailure{
		Error: ErrorDetail{
			Msg:  message,
			Code: errCode,
		},
		Status: httpStatus,
	}

	c.JSON(httpStatus, responseJson)
}

func MakeErrorResponseWithCode(w http.ResponseWriter, st, c int, message string) {
	responseJson := ResponseFailure{
		Status: st,
		Error: ErrorDetail{
			Msg:  message,
			Code: c,
		},
	}
	jData, _ := json.Marshal(&responseJson)
	w.WriteHeader(st)
	w.Header().Set(ContentTypeMessage, ApplicationJsonMessage)
	_, _ = w.Write(jData)

}
