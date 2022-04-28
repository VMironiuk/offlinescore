package offlinescore

type Squad struct {
	Id          string   `json:"ID"`
	Name        string   `json:"Nm"`
	Image       string   `json:"Img"`
	CountryName string   `json:"CoNm"`
	Players     []Player `json:"Ps"`
}

type Player struct {
	Id             string `json:"Pid"`
	FirstName      string `json:"Fn,omitempty"`
	LastName       string `json:"Ln,omitempty"`
	ShirtName      string `json:"Shnm,omitempty"`
	ShortName      string `json:"Snm,omitempty"`
	CountryName    string `json:"CoNm"`
	RegionImage    string `json:"RegImg"`
	ShirtNumber    int    `json:"Snu,omitempty"`
	ActualPosition int    `json:"PosA"`
	Position       int    `json:"Pos"`
	ReasonType     int    `json:"Rt,omitempty"`
	ReturnInfo     string `json:"RtonS,omitempty"`
	ReasonInfo     string `json:"Rs,omitempty"`
	ReturnDate     uint64 `json:"Rton,omitempty"`
	PositionString string `json:"Pon,omitempty"`
}
