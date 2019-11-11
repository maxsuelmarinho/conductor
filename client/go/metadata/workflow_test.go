package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSubordinateDTO(t *testing.T) {
	var testCases = []struct {
		name        string
		workflowDef WorkflowDef
		shouldFail  bool
	}{
		{
			name: "should return that the workflowdef is valid when all fields are filled correctly",
			workflowDef: WorkflowDef{
				Name:    "test_workflow_def",
				Version: 1,
				Tasks: []WorkflowTask{
					WorkflowTask{},
				},
				SchemaVersion: 2,
			},
			shouldFail: false,
		},
		{
			name: "should return that the workflowdef is invalid when workflow name contain semi-colon",
			workflowDef: WorkflowDef{
				Name:    "test:workflow:def",
				Version: 1,
				Tasks: []WorkflowTask{
					WorkflowTask{},
				},
				SchemaVersion: 2,
			},
			shouldFail: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := tc.workflowDef.Validate()

			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
