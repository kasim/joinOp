package joinOp

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-joinOp")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	input1Nm := context.GetInput("input1").(string)
	input2Nm := context.GetInput("input2").(string)

	output1, ok := data.GetGlobalScope().GetAttr(input1Nm)
	if !ok {
		errorMsg := fmt.Sprintf("Attribute not defined: '%s'", input1Nm)
		log.Error(errorMsg)
		return false, activity.NewError(errorMsg, "", nil)
	}
	output2, ok := data.GetGlobalScope().GetAttr(input2Nm)
	if !ok {
		errorMsg := fmt.Sprintf("Attribute not defined: '%s'", input2Nm)
		log.Error(errorMsg)
		return false, activity.NewError(errorMsg, "", nil)
	}

	output := (output1.Value().(bool) && output2.Value().(bool))
	context.SetOutput("output", output)
	return true, nil
}
