package ttRequests_models

// Dunno how I'd implement the mixed json array so...
//
//[string, uint32 uint32]
//
//[Name, Src, VrpId]
type Players struct {
	Players [][]interface{} `json:"players"`
}
