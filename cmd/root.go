package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yuisofull/leetgo/pkg/listcompanies"
	"github.com/yuisofull/leetgo/pkg/listproblems"
	"github.com/yuisofull/leetgo/pkg/storage/api"
)

var (
	listCompanies bool
	companyName   string
	allDetails    bool

	byFrequency  bool
	byDifficulty string
	byAcceptance bool
	limit        int
	isNotPremium bool
)

var RootCmd = &cobra.Command{
	Use:   "leetgo",
	Short: "LeetGo is a CLI tool for LeetCode",
	Long:  "LeetGo is a CLI tool for LeetCode that provides information about problems asked by companies in interviews with multiple operations.",
	Example: `  # List all available companies
  leetgo -l

  # List problems by a specific company
  leetgo -c google

  # List problems by a specific company with filters
  leetgo -c google --frequency --limit 10 --non-premium

  # Show detailed problem information
  leetgo -c google -a`,
	Run: func(cmd *cobra.Command, args []string) {
		if listCompanies {
			repo := apistore.NewStorage()
			svc := listcompanies.NewService(repo)
			companies, err := svc.GetCompanies()
			if err != nil {
				fmt.Println("Error fetching companies:", err)
				return
			}
			fmt.Println("Available Companies:")
			for _, company := range companies {
				fmt.Printf("%s\n", company.Name)
			}
			return
		}

		if companyName != "" {
			repo := apistore.NewStorage()
			svc := listproblems.NewService(repo)

			filter := listproblems.Filter{
				ByFrequency:  byFrequency,
				ByDifficulty: byDifficulty,
				ByAcceptance: byAcceptance,
				Limit:        limit,
				IsNotPremium: isNotPremium,
			}

			problems, err := svc.GetProblemsFromCompany(companyName, filter)
			if err != nil {
				fmt.Printf("Error fetching problems for company '%s': %v\n", companyName, err)
				return
			}

			fmt.Printf("Problems for company '%s':\n", companyName)
			for _, problem := range problems {
				if allDetails {
					fmt.Printf("ID: %s\nTitle: %s\nURL: %s\nDifficulty: %s\nFrequency: %.2f%%\nAcceptance: %.2f%%\nPremium: %t\n\n",
						problem.ID, problem.Title, problem.URL, problem.Difficulty, problem.Frequency, problem.Acceptance, problem.IsPremium)
				} else {
					fmt.Printf("%s | Frequency: %.2f%%\n", problem.URL, problem.Frequency)
				}
			}
			return
		}

		fmt.Println("Welcome to LeetGo!")

		if err := cmd.Help(); err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.Flags().BoolVarP(&listCompanies, "list-companies", "l", false, "List all companies available in LeetCode data")
	RootCmd.Flags().StringVarP(&companyName, "company", "c", "", "Specify the company to list problems for")

	RootCmd.Flags().BoolVarP(&byFrequency, "frequency", "f", false, "Sort problems by frequency")
	RootCmd.Flags().StringVar(&byDifficulty, "difficulty", "", "Get problems by difficulty(easy, medium, hard)")
	RootCmd.Flags().BoolVar(&byAcceptance, "acceptance", false, "Sort problems by acceptance rate")
	RootCmd.Flags().IntVar(&limit, "limit", 0, "Limit the number of problems returned (0 for no limit)")
	RootCmd.Flags().BoolVar(&isNotPremium, "non-premium", false, "Exclude premium problems")

	RootCmd.Flags().BoolVarP(&allDetails, "all-details", "a", false, "Show all details for each problem")
}

func Execute() error {
	return RootCmd.Execute()
}
