package cmd

import (
	"fmt"

	"github.com/Planckbaka/todo-cli/internal/storage"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "标记一条任务完成",
	Long:  `标记完成一个已知的任务。必须提供任务ID。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done called")
		result, _ := pterm.DefaultInteractiveTextInput.WithDefaultText(pterm.Bold.Sprint("请输入你已完成的任务ID")).Show()
		err := storage.DoneTodoData(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Enable debug messages in PTerm.
		pterm.EnableDebugMessages()
		// Print a success message with PTerm.
		pterm.Success.Println("task gets done successfully")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
