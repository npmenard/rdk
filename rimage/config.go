package rimage

// AttrConfig is exported to be used as an attribute map for all camera types.
type AttrConfig struct {
	Color              string      `json:"color"`
	Depth              string      `json:"depth"`
	Host               string      `json:"host"`
	Port               int         `json:"port"`
	Aligned            bool        `json:"aligned"`
	Debug              bool        `json:"debug"`
	Stream             string      `json:"stream"`
	Num                string      `json:"num"`
	Args               string      `json:"args"`
	IntrinsicExtrinsic interface{} `json:"intrinsic_extrinsic"`
	Homography         interface{} `json:"homography"`
	Warp               interface{} `json:"warp"`
	Intrinsic          interface{} `json:"intrinsic"`
}