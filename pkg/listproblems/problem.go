package listproblems

// Problem represents a LeetCode problem
type Problem struct {
	ID         string `json:"id" csv:"ID"`
	Title      string `json:"title" csv:"Title"`
	URL        string `json:"url" csv:"URL"`
	IsPremium  bool   `json:"isPremium" csv:"Is Premium"`
	Acceptance string `json:"acceptance" csv:"Acceptance %"`
	Difficulty string `json:"difficulty" csv:"Difficulty"`
	Frequency  string `json:"frequency" csv:"Frequency %"`
}
