package tt_requests_models

type UserData struct {
	User_id   uint32
	Data_type string
	Code      string
	Data      data
}

type data struct {
	Plate_override     string            `json:"plate_override"`
	Gaptitudes         gaptitudes        `json:"gaptitudes"`
	Gaptitudes_v       gaptitudesv       `json:"gaptitudes_v"`
	Thirst             float32           `json:"thirst"`
	Licenses           map[string]uint64 `json:"licenses"`
	Chat_title         string            `json:"chat_title"`
	Chat_prefix        string            `json:"chat_prefix"`
	Health             float32           `json:"health"`
	AcceptingPlayerEms bool              `json:"AcceptingPlayerEms"`
	Position           position          `json:"position"`
	Inventory          map[string]item   `json:"inventory"`
	Vehicle            vehicle           `json:"vehicle"`
	Groups             map[string]bool   `json:"groups"`
	Hunger             float32           `json:"hunger"`
	Ironman            bool              `json:"ironman"`
	Customization      customization     `json:"customization"`
}

type gaptitudes struct {
	Casino   casino   `json:"casino"`
	Business business `json:"business"`
	Hunting  hunting  `json:"hunting"`
	Ems      ems      `json:"ems"`
	Physical physical `json:"physical"`
	Farming  farming2 `json:"farming"`
	Piloting piloting `json:"piloting"`
	Train    train    `json:"train"`
	Player   player   `json:"player"`
	Trucking trucking `json:"trucking"`
	Police   police   `json:"police"`
}

type gaptitudesv struct {
	Casino   casino_v   `json:"casino"`
	Business business_v `json:"business"`
	Hunting  hunting_v  `json:"hunting"`
	Ems      ems_v      `json:"ems"`
	Physical physical_v `json:"physical"`
	Farming  farming2_v `json:"farming"`
	Piloting piloting_v `json:"piloting"`
	Train    train_v    `json:"train"`
	Player   player_v   `json:"player"`
	Trucking trucking_v `json:"trucking"`
	Police   *police_v  `json:"police"`
}

type customization struct {
	Num0      []int32   `json:"0"`
	Num1      []int32   `json:"1"`
	Num2      []int32   `json:"2"`
	Num3      []int32   `json:"3"`
	Num4      []int32   `json:"4"`
	Num5      []int32   `json:"5"`
	Num6      []int32   `json:"6"`
	Num7      []int32   `json:"7"`
	Num8      []int32   `json:"8"`
	Num9      []int32   `json:"9"`
	Num10     []int32   `json:"10"`
	Num11     []int32   `json:"11"`
	Num12     []int32   `json:"12"`
	Num13     []int32   `json:"13"`
	Num14     []int32   `json:"14"`
	Num15     []int32   `json:"15"`
	Num16     []int32   `json:"16"`
	Num17     []int32   `json:"17"`
	Num18     []int32   `json:"18"`
	Num19     []int32   `json:"19"`
	Num20     []int32   `json:"20"`
	H12       []float32 `json:"h12"`
	H7        []float32 `json:"h7"`
	H5        []float32 `json:"h5"`
	P7        []int16   `json:"p7"`
	P5        []int16   `json:"p5"`
	H9        []float32 `json:"h9"`
	P1        []int16   `json:"p1"`
	H1        []float32 `json:"h1"`
	H3        []float32 `json:"h3"`
	Modelhash int64     `json:"modelhash"`
	P3        []int16   `json:"p3"`
	P2        []int16   `json:"p2"`
	H8        []float32 `json:"h8"`
	H10       []float32 `json:"h10"`
	H2        []float32 `json:"h2"`
	P10       []int16   `json:"p10"`
	H6        []float32 `json:"h6"`
	P9        []int16   `json:"p9"`
	H4        []float32 `json:"h4"`
	P8        []int16   `json:"p8"`
	H0        []float32 `json:"h0"`
	P6        []int16   `json:"p6"`
	P4        []int16   `json:"p4"`
	H11       []float32 `json:"h11"`
	P0        []int16   `json:"p0"`
}

type vehicle struct {
	Has_trailer   bool   `json:"has_trailer"`
	Vehicle_model int32  `json:"vehicle_model"`
	Vehicle_label string `json:"vehicle_label"`
	Vehicle_type  string `json:"vehicle_type"`
	Vehicle_class int16  `json:"vehicle_class"`
	Vehicle_name  string `json:"vehicle_name"`
	Trailer       string `json:"trailer"`
	// owned_vehicles []any idk what this is
	// Vehicle_spawn string it both int and string??
}

type item struct {
	Amount uint64  `json:"amount"`
	Weight float32 `json:"weight"`
	Name   string  `json:"name"`
}

type position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type farming2 struct {
	Fishing float32 `json:"fishing"`
	Mining  float32 `json:"mining"`
	Farming float32 `json:"farming"`
}

type trucking struct {
	Trucking float32 `json:"trucking"`
	Garbage  float32 `json:"garbage"`
	Mechanic float32 `json:"mechanic"`
	Postop   float32 `json:"postop"`
}

type ems struct {
	Ems  float32 `json:"ems"`
	Fire float32 `json:"fire"`
}

type player struct {
	Racing float32 `json:"racing"`
	Player float32 `json:"player"`
}

type piloting struct {
	Piloting float32 `json:"piloting"`
	Cargos   float32 `json:"cargos"`
	Heli     float32 `json:"heli"`
}

type farming struct {
	Fishing float32 `json:"fishing"`
	Mining  float32 `json:"mining"`
	Animals float32 `json:"animals"`
	Farming float32 `json:"farming"`
}

type physical struct {
	Strength float32 `json:"strength"`
}

type train struct {
	Bus   float32 `json:"bus"`
	Train float32 `json:"train"`
}

type hunting struct {
	Skill float32 `json:"skill"`
}

type business struct {
	Business float32 `json:"business"`
}

type casino struct {
	Casino float32 `json:"casino"`
}

type police struct {
	Police float32 `json:"police"`
}

type farming2_v struct {
	Fishing float64 `json:"fishing"`
	Mining  float64 `json:"mining"`
	Farming float64 `json:"farming"`
}

type trucking_v struct {
	Trucking float64 `json:"trucking"`
	Garbage  float64 `json:"garbage"`
	Mechanic float64 `json:"mechanic"`
	Postop   float64 `json:"postop"`
}

type ems_v struct {
	Ems  float64 `json:"ems"`
	Fire float64 `json:"fire"`
}

type player_v struct {
	Racing float64 `json:"racing"`
	Player float64 `json:"player"`
}

type piloting_v struct {
	Piloting float64 `json:"piloting"`
	Cargos   float64 `json:"cargos"`
	Heli     float64 `json:"heli"`
}

type farming_v struct {
	Fishing float64 `json:"fishing"`
	Mining  float64 `json:"mining"`
	Animals float64 `json:"animals"`
	Farming float64 `json:"farming"`
}

type physical_v struct {
	Strength float64 `json:"strength"`
}

type train_v struct {
	Bus   float64 `json:"bus"`
	Train float64 `json:"train"`
}

type hunting_v struct {
	Skill float64 `json:"skill"`
}

type business_v struct {
	Business float64 `json:"business"`
}

type casino_v struct {
	Casino float64 `json:"casino"`
}

type police_v struct {
	Police float64 `json:"police"`
}
