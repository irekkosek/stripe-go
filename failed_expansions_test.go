package stripe

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailedExpansionTopLevelWithNull(t *testing.T) {
	fixture, _ := os.ReadFile("subscription.json")

	var result map[string]interface{}
	json.Unmarshal([]byte(fixture), &result)

	result["latest_invoice"] = nil

	newJSON, _ := json.Marshal(result)

	var s Subscription
	json.Unmarshal(newJSON, &s)

	assert.Nil(t, s.LatestInvoice)
}

func TestFailedExpansionTopLevelWithId(t *testing.T) {
	fixture, _ := os.ReadFile("subscription.json")

	var result map[string]interface{}
	json.Unmarshal([]byte(fixture), &result)

	result["latest_invoice"] = "il_xyz"

	newJSON, _ := json.Marshal(result)

	var s Subscription
	json.Unmarshal(newJSON, &s)

	assert.Equal(t, "il_xyz", s.LatestInvoice.ID)
}

func TestFailedExpansionTopLevelWithEmptyObject(t *testing.T) {
	fixture, _ := os.ReadFile("subscription.json")

	var result map[string]interface{}
	json.Unmarshal([]byte(fixture), &result)

	result["latest_invoice"] = make(map[string]interface{})

	newJSON, _ := json.Marshal(result)

	var s Subscription
	json.Unmarshal(newJSON, &s)

	assert.Empty(t, s.LatestInvoice.ID)
}

func TestFailedExpansionArrayWithNull(t *testing.T) {
	fixture, _ := os.ReadFile("subscription.json")

	var result map[string]interface{}
	json.Unmarshal([]byte(fixture), &result)

	var items = result["items"].(map[string]interface{})
	var array = items["data"].([]interface{})

	array = append(array, nil)
	var expectedLength = len(array)
	items["data"] = array

	newJSON, _ := json.Marshal(result)

	var s Subscription
	json.Unmarshal(newJSON, &s)

	assert.Equal(t, expectedLength, len(s.Items.Data))
	assert.Nil(t, s.Items.Data[expectedLength-1])
}

func TestFailedExpansionArrayWithId(t *testing.T) {
	fixture, _ := os.ReadFile("subscription.json")

	var result map[string]interface{}
	json.Unmarshal([]byte(fixture), &result)

	var items = result["items"].(map[string]interface{})
	var array = items["data"].([]interface{})

	array = append(array, "si_xyz")
	items["data"] = array

	newJSON, _ := json.Marshal(result)

	var s Subscription
	err := json.Unmarshal(newJSON, &s)

	assert.NotNil(t, err)
	assert.IsType(t, &json.UnmarshalTypeError{}, err)
}

func TestFailedExpansionArrayWithEmptyObject(t *testing.T) {
	fixture, _ := os.ReadFile("subscription.json")

	var result map[string]interface{}
	json.Unmarshal([]byte(fixture), &result)

	var items = result["items"].(map[string]interface{})
	var array = items["data"].([]interface{})

	array = append(array, make(map[string]interface{}))
	var expectedLength = len(array)
	items["data"] = array

	newJSON, _ := json.Marshal(result)

	var s Subscription
	json.Unmarshal(newJSON, &s)

	assert.Equal(t, expectedLength, len(s.Items.Data))
	assert.NotNil(t, s.Items.Data[expectedLength-1])
	assert.IsType(t, &SubscriptionItem{}, s.Items.Data[expectedLength-1])
	assert.Empty(t, s.Items.Data[expectedLength-1].ID)
}
