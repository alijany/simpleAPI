package devices

type Device struct { // TODO remove tag ?
	Id          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}

type DeviceModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Note string `json:"note"`
}
