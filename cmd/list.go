/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Planckbaka/todo-cli/internal/storage"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		todos, err := storage.QueryTodoData()
		if err != nil {
			fmt.Println(err)
			return
		}
		tableData1 := pterm.TableData{
			{"ID", "Title", "Description", "Priority", "DueDate"},
		}
		for _, t := range todos {
			// 确保优先级字段有固定宽度，日期格式统一
			priority := fmt.Sprintf("%-8s", t.Priority) // 左对齐，最小8字符宽度
			dueDate := t.DueDate
			// 如果日期格式不完整，尝试标准化（这里假设都是YYYY-M-D格式）
			if len(dueDate) < 10 {
				// 简单的日期格式修正，实际项目中应该用time包处理
				parts := strings.Split(dueDate, "-")
				if len(parts) == 3 {
					year, month, day := parts[0], parts[1], parts[2]
					if len(month) == 1 {
						month = "0" + month
					}
					if len(day) == 1 {
						day = "0" + day
					}
					dueDate = fmt.Sprintf("%s-%s-%s", year, month, day)
				}
			}
			
			tableData1 = append(tableData1, []string{
				strconv.Itoa(int(t.ID)), t.Title, t.Description, strings.TrimSpace(priority), dueDate})
		}

		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData1).Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
