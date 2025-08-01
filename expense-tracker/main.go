package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	ID          int
	Description string
	Amount      int
	Date        time.Time
}

var expenses []Expense

func loadExpenses() {
	data, err := os.Open("expenses.csv")
	if err != nil {
		expenses = []Expense{}
		return
	}
	defer data.Close()
	reader := csv.NewReader(data)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading expenses:", err)
		expenses = []Expense{}
		return
	}
	var loadedExpenses []Expense
	for _, record := range records {
		if len(record) < 4 {
			fmt.Println("Invalid record format:", record)
			continue
		}
		amount, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println("Error converting amount to int:", err)
			continue
		}
		id, err := strconv.Atoi(record[3])
		if err != nil {
			fmt.Println("Error converting ID to int:", err)
			continue
		}
		loadedExpenses = append(loadedExpenses, Expense{
			ID:          id,
			Description: record[0],
			Amount:      amount,
			Date: func() time.Time {
				t, err := time.Parse("2006-01-02", record[1])
				if err != nil {
					fmt.Println("Error parsing date:", err)
					return time.Now()
				}
				return t
			}(),
		})
	}
	expenses = loadedExpenses

}
func saveExpenses() {
	file, err := os.Create("expenses.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, exp := range expenses {
		record := []string{exp.Description, exp.Date.Format("2006-01-02"), strconv.Itoa(exp.Amount), strconv.Itoa(exp.ID)}
		writer.Write(record)
	}
}
func getNextID() int {
	maxID := 0
	for _, expense := range expenses {
		if expense.ID > maxID {
			maxID = expense.ID
		}
	}
	return maxID + 1
}
func exportCsv() {
	file, err := os.Create("expenses.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Description", "Date", "Amount", "ID"})
	for _, exp := range expenses {
		writer.Write([]string{exp.Description, exp.Date.Format("2006-01-02"), strconv.Itoa(exp.Amount), strconv.Itoa(exp.ID)})
	}
	writer.Flush()
	fmt.Println("Expenses exported to expenses.csv")
}
func main() {
	loadExpenses()
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add, view, or clear")
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCmd.String("description", "", "Description for expense")
		amount := addCmd.Int("amount", 0, "Amount for expense")
		addCmd.Parse(os.Args[2:])
		if *description == "" && *amount < 0 {
			fmt.Println("Description and amount are required")
			return
		}
		exp := Expense{
			ID:          getNextID(),
			Description: *description,
			Amount:      *amount,
			Date:        time.Now(),
		}
		expenses = append(expenses, exp)
		saveExpenses()
		fmt.Println("Expense added successfully")
	case "list":
		if len(expenses) == 0 {
			fmt.Println("No expenses found")
			return
		}
		for _, exp := range expenses {
			fmt.Printf("ID: %d, Description: %s, Amount: %d, Date: %s\n", exp.ID, exp.Description, exp.Amount, exp.Date.Format("2006-01-02"))
		}
	case "summary":
		summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summaryCmd.Int("month", 0, "Month to summarize expenses")
		summaryCmd.Parse(os.Args[2:])
		total := 0
		if *month < 1 || *month > 12 {
			fmt.Println("Invalid month. Please provide a month between 1 and 12.")
			return
		}

		if *month == 0 {
			for _, exp := range expenses {
				total += exp.Amount

			}
			fmt.Printf("Total expenses: %d\n", total)
		} else {
			for _, exp := range expenses {
				if exp.Date.Month() == time.Month(*month) {
					total += exp.Amount
				}
			}
			fmt.Printf("Total expenses for month %s: %d\n", time.Month(*month), total)
		}

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "ID of expense to delete")
		deleteCmd.Parse(os.Args[2:])
		for i, exp := range expenses {
			if exp.ID == *id {
				expenses = append(expenses[:i], expenses[i+1:]...)
				saveExpenses()
				fmt.Println("Expense deleted successfully")
				return
			}
		}
		fmt.Println("Expense not found")
	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		description := updateCmd.String("description", "", "Description for expense")
		amount := updateCmd.Int("amount", 0, "Amount for expense")
		id := updateCmd.Int("id", 0, "ID of expense to update")
		updateCmd.Parse(os.Args[2:])
		for i, exp := range expenses {
			if exp.ID == *id {
				if *description != "" {
					expenses[i].Description = *description
				}
				if *amount > 0 {
					expenses[i].Amount = *amount
				}
				saveExpenses()
				fmt.Println("Expense updated successfully")
				return
			}
		}
		fmt.Println("Expense not found")
	case "export":
		exportCsv()
	}

}
