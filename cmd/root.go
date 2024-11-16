package cmd

import (
	"fmt"
	"github.com/yuisofull/leetgo/pkg/listcompanies"
	"github.com/yuisofull/leetgo/pkg/listproblems"
	"github.com/yuisofull/leetgo/pkg/storage/api"

	"github.com/spf13/cobra"
)

var (
	listCompanies bool
	companyName   string
)

var RootCmd = &cobra.Command{
	Use:   "leetgo",
	Short: "LeetGo is a CLI tool for LeetCode",
	Long:  "LeetGo is a CLI tool for LeetCode to help you get tagged questions grouped by companies by frequency",
	Example: `  # List all available companies
  leetgo -l

  # List problems by a specific company
  leetgo -c google`,
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
			// Handle `-c <company>` flag to list problems
			repo := apistore.NewStorage()
			svc := listproblems.NewService(repo)
			problems, err := svc.GetProblemsFromCompany(companyName, listproblems.Filter{})
			if err != nil {
				fmt.Printf("Error fetching problems for company '%s': %v\n", companyName, err)
				return
			}

			fmt.Printf("Problems for company '%s':\n", companyName)
			for _, problem := range problems {
				fmt.Printf("- %s (ID: %s)\n", problem.Title, problem.ID)
			}
			return
		}

		// Default message if no flags are provided
		fmt.Println("Welcome to LeetGo! Use -l to list companies or -c <company> to list problems by company.")
	},
}

func init() {
	// Define flags for the root command
	RootCmd.Flags().BoolVarP(&listCompanies, "list-companies", "l", false, "List all companies available in LeetCode data")
	RootCmd.Flags().StringVarP(&companyName, "company", "c", "", "Specify the company to list problems for")
}

func Execute() error {
	return RootCmd.Execute()
}
