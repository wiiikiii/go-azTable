package manipulateAzTable

type JsonStruct struct {
	Status         string        `json:"status,omitempty"`
	ErrorText      string        `json:"errorText,omitempty"`
	Meta           Meta          `json:"meta,omitempty"`
	Configurations Configuration `json:"configurations,omitempty"`
}

type Meta struct {
	Key        string `json:"key,omitempty"`
	Name       string `json:"name,omitempty"`
	LastUpdate string `json:"lastUpdate,omitempty"`
}

type Configuration struct {
	Name        string `json:"name,omitempty"`
	LuUpdate    string `json:"luUpdate,omitempty"`
	LuProcessed string `json:"luProcessed,omitempty"`
	Fields      Fields `json:"fields,omitempty"`
}

type Fields struct {
	Tier         Tier         `json:"tier,omitempty"`
	Location     Location     `json:"location,omitempty"`
	Usercount    Usercount    `json:"usercount,omitempty"`
	Maintenance  Maintenance  `json:"maintenance,omitempty"`
	Environment  Environment  `json:"environment,omitempty"`
	Backup       Backup       `json:"backup,omitempty"`
	Recovery     Recovery     `json:"recovery,omitempty"`
	Applications Applications `json:"applications,omitempty"`
	Hostpools    Hostpools    `json:"hostpools,omitempty"`
}

type Tier struct {
	Values       string   `json:"values,omitempty"`
	SelectValues []string `json:"selectValues,omitempty"`
}

type Hostpools struct {
	Values string `json:"values,omitempty"`
}

type Location struct {
	Values       string   `json:"values,omitempty"`
	SelectValues []string `json:"selectValues,omitempty"`
}

type Usercount struct {
	Values       string   `json:"values,omitempty"`
	SelectValues []string `json:"selectValues,omitempty"`
}

type Maintenance struct {
	Values       string   `json:"values,omitempty"`
	SelectValues []string `json:"selectValues,omitempty"`
}

type Environment struct {
	Values       string   `json:"values,omitempty"`
	SelectValues []string `json:"selectValues,omitempty"`
}

type Backup struct {
	Values       string   `json:"values,omitempty"`
	SelectValues []string `json:"selectValues,omitempty"`
}

type Recovery struct {
	Values       string   `json:"values,omitempty"`
	SelectValues []string `json:"selectValues,omitempty"`
}

type Applications struct {
	Values       []string `json:"values,omitempty"`
	SelectValues []string `json:"selectValues,omitempty"`
}
