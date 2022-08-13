package object

type (
	Status struct {
		ID int `json:"id"`

		// AccountID of the status
		AccountID AccountID `json:"-" db:"account_id"`

		Account Account `json:"account,omitempty"`

		Content *string `json:"content,omitempty"`

		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
