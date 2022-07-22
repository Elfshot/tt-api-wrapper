package ttRequests_models

// add json and bson omitempty tags to every struct
type UserData struct {
	User_id   uint32       `json:"user_id" bson:"user_id,omitempty"`
	Data_type string       `json:"data_type" bson:"data_type,omitempty"`
	Code      string       `json:"code" bson:"code,omitempty"`
	Data      UserDataData `json:"data" bson:"data,omitempty"`
}

type UserDataData struct {
	Plate_override     string            `json:"plate_override" bson:"plate_override,omitempty"`
	Gaptitudes         gaptitudes        `json:"gaptitudes" bson:"gaptitudes,omitempty"`
	Gaptitudes_v       gaptitudesv       `json:"gaptitudes_v" bson:"gaptitudes_v,omitempty"`
	Thirst             float32           `json:"thirst" bson:"thirst,omitempty"`
	Licenses           map[string]uint64 `json:"licenses" bson:"licenses"`
	Chat_title         string            `json:"chat_title" bson:"chat_title"`
	Chat_prefix        string            `json:"chat_prefix" bson:"chat_prefix"`
	Health             float32           `json:"health" bson:"health,omitempty"`
	AcceptingPlayerEms bool              `json:"AcceptingPlayerEms" bson:"AcceptingPlayerEms,omitempty"`
	Position           position          `json:"position" bson:"position,omitempty"`
	Inventory          map[string]item   `json:"inventory" bson:"inventory"`
	Vehicle            vehicle           `json:"vehicle" bson:"vehicle,omitempty"`
	Groups             map[string]bool   `json:"groups" bson:"groups"`
	Hunger             float32           `json:"hunger" bson:"hunger,omitempty"`
	Ironman            bool              `json:"ironman" bson:"ironman,omitempty"`
	Customization      customization     `json:"customization" bson:"customization,omitempty"`
}

type gaptitudes struct {
	Casino   casino   `json:"casino" bson:"casino,omitempty"`
	Business business `json:"business" bson:"business,omitempty"`
	Hunting  hunting  `json:"hunting" bson:"hunting,omitempty"`
	Ems      ems      `json:"ems" bson:"ems,omitempty"`
	Physical physical `json:"physical" bson:"physical,omitempty"`
	Farming  farming2 `json:"farming" bson:"farming,omitempty"`
	Piloting piloting `json:"piloting" bson:"piloting,omitempty"`
	Train    train    `json:"train" bson:"train,omitempty"`
	Player   player   `json:"player" bson:"player,omitempty"`
	Trucking trucking `json:"trucking" bson:"trucking,omitempty"`
	Police   police   `json:"police" bson:"police,omitempty"`
}

type gaptitudesv struct {
	Casino   casino_v   `json:"casino" bson:"casino,omitempty"`
	Business business_v `json:"business" bson:"business,omitempty"`
	Hunting  hunting_v  `json:"hunting" bson:"hunting,omitempty"`
	Ems      ems_v      `json:"ems" bson:"ems,omitempty"`
	Physical physical_v `json:"physical" bson:"physical,omitempty"`
	Farming  farming2_v `json:"farming" bson:"farming,omitempty"`
	Piloting piloting_v `json:"piloting" bson:"piloting,omitempty"`
	Train    train_v    `json:"train" bson:"train,omitempty"`
	Player   player_v   `json:"player" bson:"player,omitempty"`
	Trucking trucking_v `json:"trucking" bson:"trucking,omitempty"`
	Police   *police_v  `json:"police" bson:"police,omitempty"`
}

type customization struct {
	Num0      []int32   `json:"0" bson:"0,omitempty"`
	Num1      []int32   `json:"1" bson:"1,omitempty"`
	Num2      []int32   `json:"2" bson:"2,omitempty"`
	Num3      []int32   `json:"3" bson:"3,omitempty"`
	Num4      []int32   `json:"4" bson:"4,omitempty"`
	Num5      []int32   `json:"5" bson:"5,omitempty"`
	Num6      []int32   `json:"6" bson:"6,omitempty"`
	Num7      []int32   `json:"7" bson:"7,omitempty"`
	Num8      []int32   `json:"8" bson:"8,omitempty"`
	Num9      []int32   `json:"9" bson:"9,omitempty"`
	Num10     []int32   `json:"10" bson:"10,omitempty"`
	Num11     []int32   `json:"11" bson:"11,omitempty"`
	Num12     []int32   `json:"12" bson:"12,omitempty"`
	Num13     []int32   `json:"13" bson:"13,omitempty"`
	Num14     []int32   `json:"14" bson:"14,omitempty"`
	Num15     []int32   `json:"15" bson:"15,omitempty"`
	Num16     []int32   `json:"16" bson:"16,omitempty"`
	Num17     []int32   `json:"17" bson:"17,omitempty"`
	Num18     []int32   `json:"18" bson:"18,omitempty"`
	Num19     []int32   `json:"19" bson:"19,omitempty"`
	Num20     []int32   `json:"20" bson:"20,omitempty"`
	H12       []float32 `json:"h12" bson:"h12,omitempty"`
	H7        []float32 `json:"h7" bson:"h7,omitempty"`
	H5        []float32 `json:"h5" bson:"h5,omitempty"`
	P7        []int16   `json:"p7" bson:"p7,omitempty"`
	P5        []int16   `json:"p5" bson:"p5,omitempty"`
	H9        []float32 `json:"h9" bson:"h9,omitempty"`
	P1        []int16   `json:"p1" bson:"p1,omitempty"`
	H1        []float32 `json:"h1" bson:"h1,omitempty"`
	H3        []float32 `json:"h3" bson:"h3,omitempty"`
	Modelhash int64     `json:"modelhash" bson:"modelhash,omitempty"`
	P3        []int16   `json:"p3" bson:"p3,omitempty"`
	P2        []int16   `json:"p2" bson:"p2,omitempty"`
	H8        []float32 `json:"h8" bson:"h8,omitempty"`
	H10       []float32 `json:"h10" bson:"h10,omitempty"`
	H2        []float32 `json:"h2" bson:"h2,omitempty"`
	P10       []int16   `json:"p10" bson:"p10,omitempty"`
	H6        []float32 `json:"h6" bson:"h6,omitempty"`
	P9        []int16   `json:"p9" bson:"p9,omitempty"`
	H4        []float32 `json:"h4" bson:"h4,omitempty"`
	P8        []int16   `json:"p8" bson:"p8,omitempty"`
	H0        []float32 `json:"h0" bson:"h0,omitempty"`
	P6        []int16   `json:"p6" bson:"p6,omitempty"`
	P4        []int16   `json:"p4" bson:"p4,omitempty"`
	H11       []float32 `json:"h11" bson:"h11,omitempty"`
	P0        []int16   `json:"p0" bson:"p0,omitempty"`
}

type vehicle struct {
	// Has_trailer   bool `json:"has_trailer" bson:"has_trailer,omitempty"`
	Vehicle_model int32  `json:"vehicle_model" bson:"vehicle_model,omitempty"`
	Vehicle_label string `json:"vehicle_label" bson:"vehicle_label,omitempty"`
	Vehicle_type  string `json:"vehicle_type" bson:"vehicle_type,omitempty"`
	Vehicle_class int16  `json:"vehicle_class" bson:"vehicle_class,omitempty"`
	Vehicle_name  string `json:"vehicle_name" bson:"vehicle_name,omitempty"`
	Trailer       string `json:"trailer" bson:"trailer,omitempty"`
	// owned_vehicles []any idk what this is
	// Vehicle_spawn string it both int and string??
}

type item struct {
	Amount uint64  `json:"amount" bson:"amount,omitempty"`
	Weight float32 `json:"weight" bson:"weight,omitempty"`
	Name   string  `json:"name" bson:"name,omitempty"`
}

type position struct {
	X float64 `json:"x" bson:"x,omitempty"`
	Y float64 `json:"y" bson:"y,omitempty"`
	Z float64 `json:"z" bson:"z,omitempty"`
}

type farming2 struct {
	Fishing float32 `json:"fishing" bson:"fishing,omitempty"`
	Mining  float32 `json:"mining" bson:"mining,omitempty"`
	Farming float32 `json:"farming" bson:"farming,omitempty"`
}

type trucking struct {
	Trucking float32 `json:"trucking" bson:"trucking,omitempty"`
	Garbage  float32 `json:"garbage" bson:"garbage,omitempty"`
	Mechanic float32 `json:"mechanic" bson:"mechanic,omitempty"`
	Postop   float32 `json:"postop" bson:"postop,omitempty"`
}

type ems struct {
	Ems  float32 `json:"ems" bson:"ems,omitempty"`
	Fire float32 `json:"fire" bson:"fire,omitempty"`
}

type player struct {
	Racing float32 `json:"racing" bson:"racing,omitempty"`
	Player float32 `json:"player" bson:"player,omitempty"`
}

type piloting struct {
	Piloting float32 `json:"piloting" bson:"piloting,omitempty"`
	Cargos   float32 `json:"cargos" bson:"cargos,omitempty"`
	Heli     float32 `json:"heli" bson:"heli,omitempty"`
}

type farming struct {
	Fishing float32 `json:"fishing" bson:"fishing,omitempty"`
	Mining  float32 `json:"mining" bson:"mining,omitempty"`
	Animals float32 `json:"animals" bson:"animals,omitempty"`
	Farming float32 `json:"farming" bson:"farming,omitempty"`
}

type physical struct {
	Strength float32 `json:"strength" bson:"strength,omitempty"`
}

type train struct {
	Bus   float32 `json:"bus" bson:"bus,omitempty"`
	Train float32 `json:"train" bson:"train,omitempty"`
}

type hunting struct {
	Skill float32 `json:"skill" bson:"skill,omitempty"`
}

type business struct {
	Business float32 `json:"business" bson:"business,omitempty"`
}

type casino struct {
	Casino float32 `json:"casino" bson:"casino,omitempty"`
}

type police struct {
	Police float32 `json:"police" bson:"police,omitempty"`
}

type farming2_v struct {
	Fishing float64 `json:"fishing" bson:"fishing,omitempty"`
	Mining  float64 `json:"mining" bson:"mining,omitempty"`
	Farming float64 `json:"farming" bson:"farming,omitempty"`
}

type trucking_v struct {
	Trucking float64 `json:"trucking" bson:"trucking,omitempty"`
	Garbage  float64 `json:"garbage" bson:"garbage,omitempty"`
	Mechanic float64 `json:"mechanic" bson:"mechanic,omitempty"`
	Postop   float64 `json:"postop" bson:"postop,omitempty"`
}

type ems_v struct {
	Ems  float64 `json:"ems" bson:"ems,omitempty"`
	Fire float64 `json:"fire" bson:"fire,omitempty"`
}

type player_v struct {
	Racing float64 `json:"racing" bson:"racing,omitempty"`
	Player float64 `json:"player" bson:"player,omitempty"`
}

type piloting_v struct {
	Piloting float64 `json:"piloting" bson:"piloting,omitempty"`
	Cargos   float64 `json:"cargos" bson:"cargos,omitempty"`
	Heli     float64 `json:"heli" bson:"heli,omitempty"`
}

type farming_v struct {
	Fishing float64 `json:"fishing" bson:"fishing,omitempty"`
	Mining  float64 `json:"mining" bson:"mining,omitempty"`
	Animals float64 `json:"animals" bson:"animals,omitempty"`
	Farming float64 `json:"farming" bson:"farming,omitempty"`
}

type physical_v struct {
	Strength float64 `json:"strength" bson:"strength,omitempty"`
}

type train_v struct {
	Bus   float64 `json:"bus" bson:"bus,omitempty"`
	Train float64 `json:"train" bson:"train,omitempty"`
}

type hunting_v struct {
	Skill float64 `json:"skill" bson:"skill,omitempty"`
}

type business_v struct {
	Business float64 `json:"business" bson:"business,omitempty"`
}

type casino_v struct {
	Casino float64 `json:"casino" bson:"casino,omitempty"`
}

type police_v struct {
	Police float64 `json:"police" bson:"police,omitempty"`
}
