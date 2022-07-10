package ttRequests_models

type FxPlayers []struct {
	Endpoint    string   `json:"endpoint"`
	ID          uint32   `json:"id"`
	Identifiers []string `json:"identifiers"`
	Name        string   `json:"name"`
	Ping        uint16   `json:"ping"`
}
