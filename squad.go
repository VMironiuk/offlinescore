package offlinescore

type Squad struct {
	Id          string   `json:"ID"`
	Name        string   `json:"Nm"`
	Image       string   `json:"Img"`
	CountryName string   `json:"CoNm"`
	Players     []Player `json:"Ps"`
}

type Player struct {
	Id                  string `json:"Pid"`
	FirstName           string `json:"Fn"`
	LastName            string `json:"Ln"`
	ShortName           string `json:"Shnm"`
	CountryName         string `json:"CoNm"`
	Residence           string `json:"Rcd"`
	ShirtNumber         int    `json:"Snu,omitempty"`
	AlternativePosition int    `json:"PosA"`
	Position            int    `json:"Pos"`
	ReasonType          int    `json:"Rt,omitempty"`
	ReturnInfo          string `json:"RtonS,omitempty"`
	ReasonInfo          string `json:"Rs,omitempty"`
}
