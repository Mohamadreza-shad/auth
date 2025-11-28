package exception

import "net/http"

type ErrorCode int

const (
	// Internal & Application-Specific Error Codes (1000+)
	ErrCodeSomethingWentWrong     ErrorCode = 1000
	ErrCodeNotImplemented         ErrorCode = 1001
	ErrCodeBadRequest             ErrorCode = 1002
	ErrCodeUnauthorized           ErrorCode = 1003
	ErrCodeNotFound               ErrorCode = 1004
	ErrCodeTooManyRequests        ErrorCode = 1005
	ErrCodeUnknown                ErrorCode = 1006 // Unclassified error
	ErrCodeInvalidArgument        ErrorCode = 1007 // Argument validation failed
	ErrCodeSerializationFailure   ErrorCode = 1008 // Serialization or deserialization failure
	ErrCodeDeserializationFailure ErrorCode = 1009 // Serialization or deserialization failure
	ErrCodeAlreadyExists          ErrorCode = 1010 // Resource already exists
	ErrCodeActionFailed           ErrorCode = 1011 // Account wrapper error
	ErrCodeLimitationExceeded     ErrorCode = 1012 // Application/business limitation exceeded
	ErrCodeInvalidAmount          ErrorCode = 1013 // Invalid amount specified
	ErrCodeTimeout                ErrorCode = 1014
	ErrCodeMissingTranslationKey  ErrorCode = 1015 // Missing Translation Message Key
	ErrCodeMissingI18nResource    ErrorCode = 1016 //
	ErrCodeLoadingI18nResource    ErrorCode = 1017 //
	ErrCodeInvalidConfig          ErrorCode = 1018 //
	ErrCodeDuplicateRequest       ErrorCode = 1020 //
	ErrCodeDBForeignKeyViolation  ErrorCode = 1021 // DB foreign key violation
	ErrCodeDBNullViolation        ErrorCode = 1022 // DB null constraint violation
	ErrCodeDBInternal             ErrorCode = 1023 // Unknown database error
	ErrCodeDBRecordNotFound       ErrorCode = 1024 // Unknown database error
	ErrCodeDBInvalidTransaction   ErrorCode = 1025 // Unknown database error
	ErrCodeDBDuplicateRecord      ErrorCode = 1026 //
	ErrCodeKafkaNewPubSub         ErrorCode = 1027 //

	//
	ErrCodeInsufficientBalance ErrorCode = 2000 // Account wrapper error
)

// ErrorCodeToHTTPStatus maps custom error codes to HTTP status codes
var ErrorCodeToHTTPStatus = map[ErrorCode]int{
	ErrCodeSomethingWentWrong:     http.StatusInternalServerError, // 500
	ErrCodeNotImplemented:         http.StatusNotImplemented,      // 501
	ErrCodeBadRequest:             http.StatusBadRequest,          // 400
	ErrCodeUnauthorized:           http.StatusUnauthorized,        // 401
	ErrCodeNotFound:               http.StatusNotFound,            // 404
	ErrCodeTooManyRequests:        http.StatusTooManyRequests,     // 429
	ErrCodeUnknown:                http.StatusInternalServerError, // 500
	ErrCodeInvalidArgument:        http.StatusBadRequest,          // 400
	ErrCodeSerializationFailure:   http.StatusInternalServerError, // 500
	ErrCodeDeserializationFailure: http.StatusInternalServerError, // 500
	ErrCodeAlreadyExists:          http.StatusConflict,            // 409
	ErrCodeActionFailed:           http.StatusInternalServerError, // 500
	ErrCodeLimitationExceeded:     http.StatusForbidden,           // 403
	ErrCodeInvalidAmount:          http.StatusBadRequest,          // 400
	ErrCodeTimeout:                http.StatusGatewayTimeout,      // 504
	ErrCodeMissingTranslationKey:  http.StatusInternalServerError, // 500
	ErrCodeMissingI18nResource:    http.StatusInternalServerError, // 500
	ErrCodeLoadingI18nResource:    http.StatusInternalServerError, // 500
	ErrCodeInvalidConfig:          http.StatusInternalServerError, // 500
	ErrCodeInsufficientBalance:    http.StatusPaymentRequired,     // 402
	ErrCodeDuplicateRequest:       http.StatusConflict,            // 409
	ErrCodeDBForeignKeyViolation:  http.StatusInternalServerError, // 500
	ErrCodeDBNullViolation:        http.StatusBadRequest,          // 400
	ErrCodeDBInternal:             http.StatusInternalServerError, // 500
	ErrCodeDBRecordNotFound:       http.StatusNotFound,            // 404
	ErrCodeDBInvalidTransaction:   http.StatusBadRequest,          // 400
	ErrCodeDBDuplicateRecord:      http.StatusConflict,            // 409
}
