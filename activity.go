package joinOp

import (
	"fmt"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	cache "github.com/patrickmn/go-cache"
)

var log = logger.GetLogger("activity-joinOp")

const (
	CacheName = "GlobalCache"
)

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

	var c *cache.Cache
	val, ok := data.GetGlobalScope().GetAttr(CacheName)

	if !ok {
		log.Error("cache doesn't exist")
		return false, fmt.Errorf("cache doesn't exist")
	} else {
		c = val.Value().(*cache.Cache)
	}
	cacheVal, found := get(c, input1Nm)
	if !found {
		log.Infof("No cache entry was found for [%s]", input1Nm)
		errorMsg := fmt.Sprintf("Attribute not defined: '%s'", input1Nm)
		log.Error(errorMsg)
		return false, activity.NewError(errorMsg, "", nil)
	}
	output1, err := strconv.ParseBool(cacheVal.(string))
	if err != nil {
		errorMsg := fmt.Sprintf("Attribute '%s' is not in right content: '%s'", input1Nm, cacheVal)
		log.Error(errorMsg)
		return false, activity.NewError(errorMsg, "", nil)
	}
	output2, ok := data.GetGlobalScope().GetAttr(input2Nm)
	if !ok {
		errorMsg := fmt.Sprintf("Attribute not defined: '%s'", input2Nm)
		log.Error(errorMsg)
		return false, activity.NewError(errorMsg, "", nil)
	}

	output := (output1 && output2.Value().(bool))
	context.SetOutput("output", output)
	return true, nil
}

func get(c *cache.Cache, key string) (interface{}, bool) {
		return c.Get(key)

}
