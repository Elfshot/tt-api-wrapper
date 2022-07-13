package ttRequests_models

type BaseTotalPlayer struct {
	VrpId     uint32 `json:"vrpId"`
	Name      string `json:"name"`
	ServerId  uint8  `json:"serverId"`
	AvatarUrl string `json:"avatarUrl"`
	IsStaff   bool   `json:"isStaff"`
	IsDonor   bool   `json:"isDonor"`
	Job       string `json:"job"`
}
