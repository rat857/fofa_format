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
	Errmsg          string     `json:"errmsg"`
	Error           bool       `json:"error"`
	ConsumedFpoint  int        `json:"consumed_fpoint"`
	RequiredFpoints int        `json:"required_fpoints"`
	Size            int        `json:"size"`
	Page            int        `json:"page"`
	Mode            string     `json:"mode"`
	Query           string     `json:"query"`
	Results         [][]string `json:"results"`
}

type ResultExcel struct {
	Host           string `excel:"host"`
	Protocol       string `excel:"protocol"`
	CountryName    string `excel:"country_Name"`
	Region         string `excel:"region"`
	Domain         string `excel:"domain"`
	OS             string `excel:"os"`
	Server         string `excel:"server"`
	Title          string `excel:"title"`
	Lastupdatetime string `excel:"lastupdatetime"`
	Cname          string `excel:"cname"`
}
