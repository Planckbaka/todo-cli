package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Planckbaka/todo-cli/internal/storage"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除任务",
	Long: `根据你提供的任务ID进行任务的删除，
若不知道任务ID多少可以通过list and query 指令来查询所要删除的任务ID`,
	Run: func(cmd *cobra.Command, args []string) {
		result, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("请输入你需要删除的任务的id").Show()
		deletedlist, err := storage.DeleteTodoData(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		pterm.Success.Println("You delete the task successfully")

		//make a table to show the deleted data
		tableData := pterm.TableData{
			{"ID", "Title", "Description", "Priority", "DueDate", "Completed"},
		}
		tableData = append(tableData, []string{strconv.Itoa(int(deletedlist.ID)), deletedlist.Title, deletedlist.Description, strings.TrimSpace(deletedlist.Priority), deletedlist.DueDate, strconv.FormatBool(deletedlist.Completed)})

		//Show deleted data
		err = pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
