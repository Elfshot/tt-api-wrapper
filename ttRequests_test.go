// Tests fail without ENV variables set. TYCOON_KEY_PRIVATE
package ttRequests_test

import (
	"testing"

	api "github.com/Elfshot/tt-api-wrapper"
)

func TestMainInit(t *testing.T) {
	api.Init()
}
func TestSingleAnyServRequest(t *testing.T) {
	api.Init()
	data, err := api.Get_DataAdv(59504)

	if data == nil {
		t.Error("Expected data to be non-nil")
	}
	if err != nil {
		t.Error(err)
	}
}
func TestSingleSpecServRequest(t *testing.T) {
	api.Init()
	data, err := api.Get_WidgetPlayers(1)

	if data == nil {
		t.Error("Expected data to be non-nil")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestSinglePlayersTotal(t *testing.T) {
	api.Init()
	data, err := api.GetTotalPlayers()
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error("Expected data to be non-nil")
	}
}
