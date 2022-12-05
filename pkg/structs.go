package manipulateAzTable

type T struct {
	Status    int    `json:"status"`
	ErrorText string `json:"errorText"`
	Meta      struct {
		Key        string `json:"key"`
		Name       string `json:"name"`
		LastUpdate string `json:"lastUpdate"`
	} `json:"meta"`
	Configurations []struct {
		Name        string `json:"name"`
		LuUpdate    string `json:"luUpdate"`
		LuProcessed string `json:"luProcessed"`
		Fields      struct {
			Tier struct {
				Label        string   `json:"label"`
				Values       string   `json:"values"`
				SelectValues []string `json:"selectValues"`
			} `json:"tier"`
			Location struct {
				Label        string   `json:"label"`
				Values       string   `json:"values"`
				SelectValues []string `json:"selectValues"`
			} `json:"location"`
			Group struct {
				Label        string   `json:"label"`
				Values       string   `json:"values"`
				SelectValues []string `json:"selectValues"`
			} `json:"group"`
			Maintenance struct {
				Label        string   `json:"label"`
				Values       string   `json:"values"`
				SelectValues []string `json:"selectValues"`
			} `json:"maintenance"`
			Environment struct {
				Label        string        `json:"label"`
				Values       bool          `json:"values"`
				SelectValues []interface{} `json:"selectValues"`
			} `json:"environment"`
			Backup struct {
				Label        string        `json:"label"`
				Values       bool          `json:"values"`
				SelectValues []interface{} `json:"selectValues"`
			} `json:"backup"`
			Recovery struct {
				Label        string        `json:"label"`
				Values       bool          `json:"values"`
				SelectValues []interface{} `json:"selectValues"`
			} `json:"recovery"`
			Applications struct {
				Label        string   `json:"label"`
				Values       []string `json:"values"`
				SelectValues []string `json:"selectValues"`
			} `json:"applications"`
		} `json:"fields"`
	} `json:"configurations"`
}
