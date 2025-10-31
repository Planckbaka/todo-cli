package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Planckbaka/todo-cli/internal/storage"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "根据任务名字查询任务",
	Long:  `查询任务，输入关于任务的标题的字可进行模糊查询，得到最多5条可能的任务`,
	Run: func(cmd *cobra.Command, args []string) {
		result, _ := pterm.DefaultInteractiveTextInput.WithDefaultText(pterm.Bold.Sprint("任务关键词")).Show()

		todos, err := storage.QueryTodoData(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		tableData := pterm.TableData{
			{"ID", "Title", "Description", "Priority", "DueDate", "Completed"},
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

			tableData = append(tableData, []string{
				strconv.Itoa(int(t.ID)), t.Title, t.Description, strings.TrimSpace(priority), dueDate, strconv.FormatBool(t.Completed)})
		}

		err = pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
