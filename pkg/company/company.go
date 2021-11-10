package company

type Company struct {
	INN       int    `json:"INN,omitempty"`
	KPP       int    `json:"KPP,omitempty"`
	Name      string `json:"name,omitempty"`
	OwnerName string `json:"ownerName,omitempty"`
}
