package types

type Response struct {
	DataCallName string `json:"data_call_name"`
	Time         string `json:"time"`
	Data         struct {
		Resources string   `json:"resource"`
		Prefixes  []Prefix `json:"prefixes"`
	} `json:"data"`
}

type Prefix struct {
	Prefix    string     `json:"prefix"`
	Timelines []Timeline `json:"timelines"`
}

type Timeline struct {
	StartTime string `json:"starttime"`
	EndTime   string `json:"endtime"`
}
