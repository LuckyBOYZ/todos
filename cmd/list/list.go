package list

import (
	"github.com/LuckyBOYZ/todos/cmd/db"
	"github.com/mergestat/timediff"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List your tasks",
	Long: `This command lists all your tasks.
Depends on the --all flag you can list all tasks or only the ones that are not done.`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		renderTasks(all)
	},
}

func init() {
	Cmd.Flags().BoolP("all", "a", false, "all tasks")
}

func renderTasks(all bool) {
	todos := db.GetTodos(all)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "description", "done", "created at"})
	for _, v := range todos {
		table.Append([]string{
			v[0],
			v[1],
			v[2],
			getTimeDifferenceFromEpochString(v[3]),
		})
	}
	table.Render()
}

func getTimeDifferenceFromEpochString(timestamp string) string {
	epoch, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return "Cannot parse time from CSV"
	}
	t := time.Unix(epoch, 0)
	return timediff.TimeDiff(t)
}
