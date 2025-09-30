package delivery

type (
	CallbackData struct {
		Inter *Interval `json:"in,omitempty"`
		A     *Answer   `json:"a,omitempty"`
	}

	Interval struct {
		Val string `json:"v,omitempty"`
	}

	Answer struct {
		AnswerNum int   `json:"a,omitempty"`
		CaseID    int64 `json:"c,omitempty"`
	}
)
