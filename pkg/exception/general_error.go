package exception

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"runtime"
	"strings"
)

type GeneralError interface {
	error
	WithError(err error) GeneralError
	WithMessageKey(key string) GeneralError
	WithMessage(msg string) GeneralError
	WithErrorCode(errorCode ErrorCode) GeneralError
	WithHttpStatusCode(statusCode int) GeneralError
	WithTrace(trace string) GeneralError
	WithPayload(payload []byte) GeneralError
	WithExtraInfo(extra map[string]interface{}) GeneralError
	WithParams(params []interface{}) GeneralError
	GetMessage() string
	GetError() error
	GetMessageKey() string
	GetTrace() string
	GetDevTrace() string
	GetErrorCode() ErrorCode
	GetHttpStatusCode() int
	GetPayload() []byte
	GetExtraInfo() map[string]interface{}
	StringDetail() string
	StringDetailSimple() string
	GetParams() []interface{}
	ReloadDevTrace(skip int) GeneralError
	Clone() GeneralError
	Unwrap() error
}

func New() GeneralError {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	s := &gError{DevTrace: fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)}
	return s
}

func Is(e error) (GeneralError, bool) {
	var se GeneralError
	if errors.As(e, &se) {
		return se, true
	}
	return nil, false
}

func Equals(ea error, eb error) bool {
	a, oka := Is(ea)
	if !oka {
		return false
	}

	b, okb := Is(eb)
	if !okb {
		return false
	}

	ae, ok := a.(gError)
	if !ok {
		return false
	}

	be, ok := b.(gError)
	if !ok {
		return false
	}

	//if strings.Compare(ae.GetMessageKey(), be.GetMessageKey()) != 0 {
	//	return false
	//}

	if ae.ErrorCode != be.ErrorCode {
		return false
	}

	return true
}

type gError struct {
	ErrorCode       ErrorCode              `json:"error_code"`
	HttpStatusCode  int                    `json:"http_status_code"`
	PayloadValue    []byte                 `json:"payload_value"`
	TextFormatValue []string               `json:"text_format_value"`
	Message         string                 `json:"message"`
	MessageKey      string                 `json:"message_key"`
	Trace           string                 `json:"trace"`
	ExtraInfo       map[string]interface{} `json:"extra_info"`
	DevTrace        string                 `json:"dev_trace"`
	Params          []interface{}          `json:"params"`
	Err             error                  `json:"-"`
}

func (e gError) Clone() GeneralError {
	se := New().WithMessageKey(e.GetMessageKey()).WithMessage(e.GetMessage()).
		WithTrace(e.GetTrace()).WithExtraInfo(e.GetExtraInfo()).WithErrorCode(e.GetErrorCode()).
		WithHttpStatusCode(e.GetHttpStatusCode()).WithPayload(e.GetPayload()).WithParams(e.GetParams())
	return se
}

func (e gError) ReloadDevTrace(skip int) GeneralError {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	e.DevTrace = fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
	return e
}

func (e gError) StringDetailSimple() string {
	s := e.StringDetail()
	s = strings.ReplaceAll(s, `"`, "")
	s = strings.ReplaceAll(s, `\n`, "")
	return s
}

func (e gError) GetParams() []interface{} {
	return e.Params
}

func (e gError) WithParams(params []interface{}) GeneralError {
	e.Params = params
	return e
}

func (e gError) WithExtraInfo(extra map[string]interface{}) GeneralError {
	e.ExtraInfo = extra
	return e
}

func (e gError) GetExtraInfo() map[string]interface{} {
	return e.ExtraInfo
}

func (e gError) GetError() error {
	return e.Err
}

func (e gError) GetMessage() string {
	return e.Message
}

func (e gError) GetMessageKey() string {

	return e.MessageKey
}

func (e gError) GetTrace() string {
	return e.Trace
}

func (e gError) GetDevTrace() string {
	return e.DevTrace
}

func (e gError) GetErrorCode() ErrorCode {
	return e.ErrorCode
}

func (e gError) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e gError) GetPayload() []byte {
	return e.PayloadValue
}

func (e gError) WithError(err error) GeneralError {
	e.Err = err
	return e
}

func (e gError) WithMessage(msg string) GeneralError {
	e.Message = msg
	return e
}

func (e gError) WithErrorCode(errorCode ErrorCode) GeneralError {

	e.ErrorCode = errorCode
	return e
}

func (e gError) WithHttpStatusCode(httpStatusCode int) GeneralError {
	e.HttpStatusCode = httpStatusCode
	return e
}

func (e gError) WithMessageKey(key string) GeneralError {
	e.MessageKey = key
	return e
}

func (e gError) WithTrace(trace string) GeneralError {
	e.Trace = trace
	return e
}

func (e gError) WithPayload(payload []byte) GeneralError {
	e.PayloadValue = payload
	return e
}

func (e gError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	} else if e.GetTrace() != "" {
		return e.GetTrace()
	} else {
		return e.GetMessage()
	}
}

func (e gError) Unwrap() error {
	return e.Err
}

func (e gError) StringDetail() string {
	marshal, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("status_code: %d , message: %v, trace: %s", e.ErrorCode, e.Message, e.Trace)
	}
	return string(marshal)
}

func ConvertDBError(err error) GeneralError {
	if err == nil {
		return nil
	}

	// GORM error mapping
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrRecordNotFound.WithError(err)
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return ErrInvalidTransaction.WithError(err)
	case errors.Is(err, gorm.ErrNotImplemented):
		return ErrNotImplemented.WithError(err)
	case errors.Is(err, gorm.ErrUnsupportedRelation), errors.Is(err, gorm.ErrPrimaryKeyRequired),
		errors.Is(err, gorm.ErrModelValueRequired), errors.Is(err, gorm.ErrModelAccessibleFieldsRequired),
		errors.Is(err, gorm.ErrSubQueryRequired), errors.Is(err, gorm.ErrInvalidData),
		errors.Is(err, gorm.ErrUnsupportedDriver), errors.Is(err, gorm.ErrRegistered),
		errors.Is(err, gorm.ErrInvalidField), errors.Is(err, gorm.ErrEmptySlice),
		errors.Is(err, gorm.ErrDryRunModeUnsupported), errors.Is(err, gorm.ErrInvalidDB),
		errors.Is(err, gorm.ErrInvalidValue), errors.Is(err, gorm.ErrCheckConstraintViolated),
		errors.Is(err, gorm.ErrInvalidValueOfLength), errors.Is(err, gorm.ErrPreloadNotAllowed),
		errors.Is(err, gorm.ErrMissingWhereClause):
		return ErrInvalidArgument.WithError(err)
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return ErrDuplicateRecord.WithError(err)
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return ErrForeignKeyViolation.WithError(err)
	}

	// MySQL error mapping
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1062: // ER_DUP_ENTRY
			return ErrDuplicateRecord.WithError(err)
		case 1048: // ER_BAD_NULL_ERROR
			return ErrNullViolation.WithError(err)
		case 1216, 1452: // Foreign key constraint fails
			return ErrForeignKeyViolation.WithError(err)
		case 1049: // ER_BAD_DB_ERROR
			return ErrDBInternal.WithError(err)
		case 1364: // ER_NO_DEFAULT_FOR_FIELD
			return ErrNullViolation.WithError(err)
		case 1146, 1051: // Table not found
			return ErrNotfound.WithError(err)
		case 1064: // SQL syntax error
			return ErrInvalidArgument.WithError(err)
		case 2002, 2003, 2013: // Connection issues
			return ErrDBInternal.WithError(err)
		default:
			return ErrDBInternal.WithError(err)
		}
	}

	// Fallback for unknown errors
	return ErrDBInternal.WithError(err)
}
