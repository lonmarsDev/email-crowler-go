package models

type Content struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Verified_email bool   `json:"verified_email"`
	Picture        string `json:"picture"`
}

type Message_content struct {
	Messages           []Message `json:"messages"`
	ResultSizeEstimate int       `json:"resultSizeEstimate"`
}

type Message struct {
	Id       string `json:"id"`
	ThreadId string `json:"threadId"`
}

type Emails struct {
	ID           string    `json:"id`
	ThreadID     string    `json:"threadId"`
	LabelIds     []Label   `json:"labelIds"`
	Snippet      string    `json:"snippet"`
	Payload      *Payloads `json:"payload"`
	SizeEstimate int       `json:"sizeEstimate"`
	HistoryID    string    `json:"historyId"`
	InternalDate string    `json:"internalDate"`
}

type Payloads struct {
	PartID   string        `json:"partId"`
	MimeTYPE string        `json:"mimeType"`
	FileNAME string        `json:"filename"`
	HEADER   []Headers     `json:"headers"`
	Body     *Payload_body `json:"body"`
	Parts    []Parts_type  `json:"parts"`
}
type Payload_body struct {
	Siz int `json:"size"`
}

type Parts_type struct {
	PartID   string      `json:"partId"`
	MimeTYPE string      `json:"mimeType"`
	FileNAME string      `json:"filename"`
	HEADER   []Headers   `json:"headers"`
	Body     *Parts_body `json:"body"`
}

type Parts_body struct {
	SizeInt     int    `json:"size"`
	Data_string string `json:"data"`
}

type Label struct {
	Name string `json:"IMPORTANT"`
}
type Headers struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Email_Extract struct {
	ZGF0YQ string
}
