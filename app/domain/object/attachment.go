package object

type (
	Attachment struct {
		ID         int64  `json:"id"`
		Type       string `json:"type"`
		URL        string `json:"url"`
		Desription string `json:"description"`
	}
)
