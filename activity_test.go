package joinOp

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/stretchr/testify/assert"
	cache "github.com/patrickmn/go-cache"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("input1", "b1")
	tc.SetInput("input2", "avgTemp")
	dt, _ := data.ToTypeEnum("bool")
	c := cache.New(time.Duration(1)*time.Minute, time.Duration(1)*time.Minute)
	data.GetGlobalScope().AddAttr("GlobalCache", data.TypeAny, c)
	set(c, "b1", "true")
	data.GetGlobalScope().AddAttr("b1", dt, true)
	data.GetGlobalScope().AddAttr("avgTemp", dt, true)
	act.Eval(tc)

	output := tc.GetOutput("output")

	assert.Equal(t, output, true)
	//check result attr
}

func set(c *cache.Cache, key string, value string) {
		c.Set(key, value, cache.DefaultExpiration)
}
