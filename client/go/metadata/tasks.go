package metadata

type TimeoutPolicy int

const (
	Retry TimeoutPolicy = iota + 1
	TimeoutWF
	AlertOnly
)

var timeoutPolicyNames = []string{
	"RETRY",
	"TIMEOUT_WF",
	"ALERT_ONLY",
}

var timeoutPolicyNameToEnum = map[string]TimeoutPolicy{
	"RETRY":      Retry,
	"TIMEOUT_WF": TimeoutWF,
	"ALERT_ONLY": AlertOnly,
}

// Name returns the string representation of the enum
func (t TimeoutPolicy) Name() string {
	return toString(&t, timeoutPolicyNames)
}

// Ordinal returns the ordinal representation of the enum
func (t TimeoutPolicy) Ordinal() int {
	return int(t)
}

// MarshalJSON marshals the enum as a quoted json string
func (t TimeoutPolicy) MarshalJSON() ([]byte, error) {
	return marshalJSON(&t)
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (t *TimeoutPolicy) UnmarshalJSON(b []byte) error {
	enumName, err := unmarshalJSON(b)
	if err != nil {
		return err
	}

	*t = timeoutPolicyNameToEnum[enumName]
	return nil
}

type RetryLogic int

const (
	Fixed RetryLogic = iota + 1
	ExponentialBackoff
)

var retryLogicNames = []string{
	"FIXED",
	"EXPONENTIAL_BACKOFF",
}

var retryLogicNameToEnum = map[string]RetryLogic{
	"FIXED":               Fixed,
	"EXPONENTIAL_BACKOFF": ExponentialBackoff,
}

// Name returns the string representation of the enum
func (r RetryLogic) Name() string {
	return toString(&r, retryLogicNames)
}

// Ordinal returns the ordinal representation of the enum
func (r RetryLogic) Ordinal() int {
	return int(r)
}

// MarshalJSON marshals the enum as a quoted json string
func (r RetryLogic) MarshalJSON() ([]byte, error) {
	return marshalJSON(&r)
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (r *RetryLogic) UnmarshalJSON(b []byte) error {
	enumName, err := unmarshalJSON(b)
	if err != nil {
		return err
	}

	*r = retryLogicNameToEnum[enumName]
	return nil
}

type TaskDef struct {
	Name                        string                 `json:"name" validate:"required,gt=0"`
	Description                 string                 `json:"description"`
	RetryCount                  int                    `json:"retryCount" validate:"gte=0"`
	TimeoutSeconds              int                    `json:"timeoutSeconds"`
	InputKeys                   []string               `json:"inputKeys"`
	OutputKeys                  []string               `json:"outputKeys"`
	TimeoutPolicy               TimeoutPolicy          `json:"timeoutPolicy"`                           // default TIME_OUT_WF
	RetryLogic                  RetryLogic             `json:"retryLogic"`                              // default FIXED
	RetryDelaySeconds           int                    `json:"retryDelaySeconds"`                       // default 60
	ResponseTimeoutSeconds      int                    `json:"responseTimeoutSeconds" validate:"min=1"` // default 60*60 = one hour
	ConcurrentExecLimit         int                    `json:"concurrentExecLimit"`
	InputTemplate               map[string]interface{} `json:"inputTemplate"`
	RateLimitPerFrequency       int                    `json:"rateLimitPerFrequency"`
	RateLimitFrequencyInSeconds int                    `json:"rateLimitFrequencyInSeconds"`
	IsolationGroupID            string                 `json:"isolationGroupId"`
	ExecutionNameSpace          string                 `json:"executionNameSpace"`
}
