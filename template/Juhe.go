package template

type JuheInfo struct {
	Error           bool     `json:"error"`
	Errmsg          string   `json:"errmsg"`
	ConsumedFpoint  int      `json:"consumed_fpoint"`
	RequiredFpoints int      `json:"required_fpoints"`
	Size            int      `json:"size"`
	Distinct        Distinct `json:"distinct"`
	Aggs            Aggs     `json:"aggs"`
	LastUpdateTime  string   `json:"lastupdatetime"`
}

type Distinct struct {
	IP     int `json:"ip"`
	Server int `json:"server"`
}

type Aggs struct {
	AssetType []AssetType `json:"asset_type"`
	Countries []Country   `json:"countries"`
	Server    []Server    `json:"server"`
}

type AssetType struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

type Country struct {
	Code     string   `json:"code"`
	Count    int      `json:"count"`
	Name     string   `json:"name"`
	NameCode string   `json:"name_code"`
	Regions  []Region `json:"regions"`
}

type Region struct {
	Code  string `json:"code"`
	Count int    `json:"count"`
	Name  string `json:"name"`
}

type Server struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

type Data struct {
	Errmsg          string   `json:"errmsg"`
	Error           bool     `json:"error"`
	ConsumedFpoint  int      `json:"consumed_fpoint"`
	RequiredFpoints int      `json:"required_fpoints"`
	Size            int      `json:"size"`
	Page            int      `json:"page"`
	Mode            string   `json:"mode"`
	Query           string   `json:"query"`
	Results         []string `json:"results"`
}
