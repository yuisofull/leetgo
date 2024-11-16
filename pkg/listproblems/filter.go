package listproblems

type Filter struct {
	ByFrequency  bool `json:"byFrequency"`
	ByDifficulty bool `json:"byDifficulty"`
	ByAcceptance bool `json:"byAcceptance"`
	Limit        int  `json:"limit"`
	IsPremium    bool `json:"isPremium"`
}

func (f *Filter) Process() {
	if f.Limit <= 0 {
		f.Limit = 50
	}
}
