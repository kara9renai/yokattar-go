package object

type (
	Like struct {
		ID       int      `json:"id"`
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
