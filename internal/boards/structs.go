package boards

type Options struct {
	Folder string
}

type Board struct {
	Name    string
	Vendor  string
	Core    string
	HasWifi bool `json:"has_wifi"`
}

type BoardWrapper struct {
	Boards   []Board
	Metadata Metadata `json:"_metadata"`
}

type Metadata struct {
	Totals Totals
	Errors Errors
}

type Totals struct {
	Vendors     int
	Boards      int
	WifiEnabled int `json:"wifi_enabled"`
}

type Errors struct {
	HasErrors      bool     `json:"has_errors"`
	FileReadErrors int      `json:"file_read_errors,omitempty"`
	Files          []string `json:"files,omitempty"`
}
