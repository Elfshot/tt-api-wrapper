// Tests fail without ENV variables set. TYCOON_KEY_PRIVATE
package ttRequests

import (
	"testing"
)

func TestMainInit(t *testing.T) {
	checkInit()
}
func TestSingleAnyServRequest(t *testing.T) {
	checkInit()
	data, err := Get_DataAdv(59504)

	if data == nil {
		t.Error("Expected data to be non-nil")
	}
	if err != nil {
		t.Error(err)
	}
}
func TestSingleSpecServRequest(t *testing.T) {
	checkInit()
	data, err := Get_WidgetPlayers(1)

	if data == nil {
		t.Error("Expected data to be non-nil")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestSinglePlayersTotal(t *testing.T) {
	checkInit()
	data, err := GetTotalPlayers()
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error("Expected data to be non-nil")
	}
}
