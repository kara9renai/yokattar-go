package object

type (
	Status struct {
		ID int `json:"id"`

		// AccountID of the status
		AccountID AccountID `json:"-" db:"account_id"`

		Account Account `json:"account"`

		Content string `json:"content"`

		FavoriteCount int64 `json:"favorite_count" db:"favorite_count"`

		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
