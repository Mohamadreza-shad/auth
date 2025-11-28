package exception

const (
	OkayErrMsgKey               string = "msg.ok.status"
	BadRequestErrMsgKey         string = "msg.error.bad_request"
	UnauthorizedErrMsgKey       string = "msg.error.unauthorized"
	NotFoundErrMsgKey           string = "msg.error.not_found"
	TooManyRequestsErrMsgKey    string = "msg.error.too_many_requests"
	SomethingWentWrongErrMsgKey string = "msg.error.Something_went_wrong"
	NotImplementedErrMsgKey     string = "msg.error.not_implemented"

	UnknownErrMsgKey                string = "msg.error.unknown"
	InvalidArgumentErrMsgKey        string = "msg.error.invalid_argument"
	SerializationFailureErrMsgKey   string = "msg.error.serialization_failure"
	DeserializationFailureErrMsgKey string = "msg.error.deserialization_failure"
	AlreadyExistsErrMsgKey          string = "msg.error.already_exists"

	DBForeignKeyViolationErrMsgKey string = "msg.error.db_foreign_key_violation"
	DBNullViolationErrMsgKey       string = "msg.error.db_null_violation"
	DBInternalErrMsgKey            string = "msg.error.db_internal"
	DBInvalidTransactionErrMsgKey  string = "msg.error.db_invalid_transaction"
	DBDuplicateRecordErrMsgKey     string = "msg.error.db_duplicate_record"
	DBRecordNotFoundErrMsgKey      string = "msg.error.db_record_not_found"

	ActionFailedErrMsgKey        string = "msg.error.action_failed"
	InsufficientBalanceErrMsgKey string = "msg.error.insufficient_balance"

	LimitationExceededErrMsgKey string = "msg.error.limitation_exceeded"
	InvalidAmountErrMsgKey      string = "msg.error.invalid_amount"

	MissingTranslationErrMsgKey  string = "msg.error.missing_translation_key"
	MissingI18nResourceErrMsgKey string = "msg.error.missing_i18n_resource"
	LoadingI18nResourceErrMsgKey string = "msg.error.loading_i18n_resource"
	TimeoutErrMsgKey             string = "msg.error.timeout"
	InvalidConfigErrMsgKey       string = "msg.error.invalid_config"

	DuplicateRequestErrMsgKey string = "msg.error.duplicate_request"

	KafkaNewPublisherSubscriberMsgKey string = "msg.error.kafka_new_subscriber_publisher"
)
