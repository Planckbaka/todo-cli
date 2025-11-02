/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/Planckbaka/todo-cli/internal/models"
	"github.com/Planckbaka/todo-cli/internal/storage"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加一条新的任务",
	Long: `创建一条新的任务。必须提供标题（位置参数或 --title）。
可选字段：--desc, --priority (low|medium|high), --due YYYY-MM-DD, --tag tag1,tag2。
成功后会输出新任务 ID。`,
	Example: `./todo-cli add "写作业" --desc "写有机化学第四章和第五章作业" --due 2025-10-29 --tag homework`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: 缺少任务标题，例如: todo add \"写作业\"")
			return
		}
		title := args[0]
		// retrieved flag arguments
		desc, _ := cmd.Flags().GetString("desc")
		if desc == "" {
			result, _ := pterm.DefaultInteractiveTextInput.WithDefaultText(pterm.Bold.Sprint("请描述你的任务")).Show()
			desc = result

		}
		priority, _ := cmd.Flags().GetString("priority")
		if priority == "" {
			result, _ := pterm.DefaultInteractiveTextInput.WithDefaultText(pterm.Bold.Sprint("请描述你的任务的紧急度（low/medium/high)")).Show()
			priority = result
		}
		due, _ := cmd.Flags().GetString("due")
		if due == "" {
			result, _ := pterm.DefaultInteractiveTextInput.WithDefaultText(pterm.Bold.Sprint("请输入你的任务的deadline(YYYY-MM-DD)")).Show()
			due = result
		}

		tagStr, _ := cmd.Flags().GetString("tag")
		if tagStr == "" {
			result, _ := pterm.DefaultInteractiveTextInput.WithDefaultText(pterm.Bold.Sprint("请输入你的任务的标签，使用逗号隔开")).Show()
			tagStr = result
		}
		tagStr = strings.ReplaceAll(tagStr, "，", ",")

		// 3️⃣ 处理 tag 字符串为切片
		var tags []string
		if tagStr != "" {
			tags = strings.Split(tagStr, ",")
		}

		todo := models.Todo{
			Title:       title,
			Description: desc,
			Priority:    priority,
			DueDate:     due,
			Tags:        tags,
		}
		err := storage.SaveTodoData(&todo)
		if err != nil {
			fmt.Println(err)
		}

		// 4️⃣ 打印或保存任务（这里只是打印模拟）
		fmt.Println("✅ 新任务创建成功！")
		fmt.Printf("标题: %s\n", title)
		fmt.Printf("描述: %s\n", desc)
		fmt.Printf("优先级: %s\n", priority)
		fmt.Printf("截止日期: %s\n", due)
		fmt.Printf("标签: %v\n", tags)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	//定义flag
	addCmd.Flags().StringP("desc", "d", "", "任务描述（可选）")
	addCmd.Flags().StringP("priority", "p", "", "优先级，low|medium|high")
	addCmd.Flags().String("due", "", "截止日期，格式 YYYY-MM-DD 或自然语言（若支持）")
	addCmd.Flags().StringP("tag", "t", "", "逗号分隔的标签列表")
}
