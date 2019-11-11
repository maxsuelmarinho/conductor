package metadata

// WorkflowDef holds the workflow definitions
type WorkflowDef struct {
	Name                          string                 `json:"name" validate:"required,gt=0,excludes=:"`
	Description                   string                 `json:"description"`
	Version                       int                    `json:"version" validate:"required,min=1"` // TODO: default 1
	Tasks                         []WorkflowTask         `json:"tasks" validate:"required,gt=0"`
	InputParameters               []string               `json:"inputParameters"`
	OutputParameters              map[string]interface{} `json:"outputParameters"`
	FailureWorkflow               string                 `json:"failureWorkflow"`
	SchemaVersion                 int                    `json:"schemaVersion" validate:"required,min=2,max=2"` // TODO: default 2
	Restartable                   bool                   `json:"restartable"`                                   // TODO: default true
	WorkflowStatusListenerEnabled bool                   `json:"workflowStatusListenerEnabled"`                 // TODO: default false
}

// Validate defines an implementation of the InputValidation interface
func (w WorkflowDef) Validate() error {
	return govalidator.Struct(w)
}

// WorkflowTask is the task definition definied as part of the WorkflowDef. The tasks definied in the Workflow definition are saved as part of WorkflowDef.Tasks
type WorkflowTask struct {
	Name                           string                    `json:"name" validate:"required,gt=0"`
	TaskReferenceName              string                    `json:"taskReferenceName" validate:"required,gt=0"`
	Description                    string                    `json:"description"`
	InputParameters                map[string]interface{}    `json:"inputParameters"`
	Type                           TaskType                  `json:"type"`
	DynamicTaskNameParam           string                    `json:"dynamicTaskNameParam"`
	CaseValueParam                 string                    `json:"caseValueParam"`
	CaseExpression                 string                    `json:"caseExpression"`
	ScriptExpression               string                    `json:"scriptExpression"`
	DecisionCases                  map[string][]WorkflowTask `json:"decisionCases"`
	DynamicForkJoinTasksParam      string                    `json:"dynamicForkJoinTasksParam"`
	DynamicForkTasksParam          string                    `json:"dynamicForkTasksParam"`
	DynamicForkTasksInputParamName string                    `json:"dynamicForkTasksInputParamName"`
	DefaultCase                    []WorkflowTask            `json:"defaultCase"`
	ForkTasks                      [][]WorkflowTask          `json:"forkTasks"`
	StartDelay                     int                       `json:"startDelay" validate:"gte=0"`
	SubWorkflowParam               SubWorkflowParams         `json:"subWorkflowParam"`
	JoinOn                         []string                  `json:"joinOn"`
	Sink                           string                    `json:"sink"`
	Optional                       bool                      `json:"optional"`
	TaskDefinition                 TaskDef                   `json:"taskDefinition"`
	RateLimited                    bool                      `json:"rateLimited"`
	DefaultExclusiveJoinTask       []string                  `json:"defaultExclusiveJoinTask"`
	AsyncComplete                  bool                      `json:"asyncComplete"`
	LoopCondition                  string                    `json:"loopCondition"`
	LoopOver                       []WorkflowTask            `json:"loopOver"`
}

// Validate defines an implementation of the InputValidation interface
func (w WorkflowTask) Validate() error {
	return govalidator.Struct(w)
}

type SubWorkflowParams struct {
	Name         string            `json:"name" validate:"required,gt=0"`
	Version      int               `json:"version"`
	TaskToDomain map[string]string `json:"taskToDomain"`
}

// Validate defines an implementation of the InputValidation interface
func (s SubWorkflowParams) Validate() error {
	return govalidator.Struct(s)
}
