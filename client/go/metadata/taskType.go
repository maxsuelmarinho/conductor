package metadata

// TaskType represents each of the possible enumeration values.
type TaskType int

// Constants representing each of the possible enumeration values.
const (
	Simple TaskType = iota + 1
	Dynamic
	ForkJoin
	ForkJoinDynamic
	Decision
	Join
	DoWhile
	SubWorkflow
	Event
	Wait
	UserDefined
	HTTP
	Lambda
	ExclusiveJoin
	Terminate
	KafkaPublish
)

var taskTypeNames = []string{
	"SIMPLE",
	"DYNAMIC",
	"FORK_JOIN",
	"FORK_JOIN_DYNAMIC",
	"DECISION",
	"JOIN",
	"DO_WHILE",
	"SUB_WORKFLOW",
	"EVENT",
	"WAIT",
	"USER_DEFINED",
	"HTTP",
	"LAMBDA",
	"EXCLUSIVE_JOIN",
	"TERMINATE",
	"KAFKA_PUBLISH",
}

var taskTypeNameToEnum = map[string]TaskType{
	"SIMPLE":            Simple,
	"DYNAMIC":           Dynamic,
	"FORK_JOIN":         ForkJoin,
	"FORK_JOIN_DYNAMIC": ForkJoinDynamic,
	"DECISION":          Decision,
	"JOIN":              Join,
	"DO_WHILE":          DoWhile,
	"SUB_WORKFLOW":      SubWorkflow,
	"EVENT":             Event,
	"WAIT":              Wait,
	"USER_DEFINED":      UserDefined,
	"HTTP":              HTTP,
	"LAMBDA":            Lambda,
	"EXCLUSIVE_JOIN":    ExclusiveJoin,
	"TERMINATE":         Terminate,
	"KAFKA_PUBLISH":     KafkaPublish,
}

// Name returns the string representation of the enum
func (t TaskType) Name() string {
	return toString(&t, taskTypeNames)
}

// Ordinal returns the ordinal representation of the enum
func (t TaskType) Ordinal() int {
	return int(t)
}

// MarshalJSON marshals the enum as a quoted json string
func (t TaskType) MarshalJSON() ([]byte, error) {
	return marshalJSON(&t)
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (t *TaskType) UnmarshalJSON(b []byte) error {
	enumName, err := unmarshalJSON(b)
	if err != nil {
		return err
	}

	*t = taskTypeNameToEnum[enumName]
	return nil
}
