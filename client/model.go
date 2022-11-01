package client

type Response struct {
	Info Info        `json:"info,omitempty"`
	Data []Character `json:"results,omitempty"`
}

type Info struct {
	Count int `json:"count"`
	Pages int `json:"pages"`
}

type Character struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Image   string `json:"image"`
}
