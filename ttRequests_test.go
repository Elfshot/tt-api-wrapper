package ttRequests

import (
	"testing"
)

// Tests
func TestAdd(t *testing.T) {
	Init()
	data, err := Get_DataAdv(59504)
	players, err := Get_Players()
	fxplayers, err := Get_FxPlayers()
	sotd, err := Get_Sotd()

	if data == nil || players == nil || fxplayers == nil || sotd == nil {
		t.Error("One of the requests failed (empty response)")
	}
	if err != nil {
		t.Error(err)
	}

}
