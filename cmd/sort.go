package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/miromax42/go-sort/pkg/line"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// sortCmd represents the sort command
var sortCmd = &cobra.Command{ //nolint
	Use:   "sort",
	Short: "sort file",
	Long:  "This subcommand sort provided file\nexample: sort test.txt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sort called with with args=%v\n", args)

		if len(args) != 1 {
			fmt.Printf("wrong format")

			return
		}

		filename := args[0]

		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)

			return
		}
		defer file.Close()

		sc := bufio.NewScanner(file)

		lines, total, err := line.NewLines(sc)
		if err != nil {
			fmt.Printf("Error process file: %v\n", err)

			return
		}

		// generate filename
		filenameFields := strings.Split(filename, ".")
		filenameFields[0] += "(sorted)"
		newFilename := strings.Join(filenameFields, ".")

		err = printLines(newFilename, lines)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("scan file, lines=%v\n", total)
	},
}

func printLines(filePath string, lines []line.Line) error {
	data := line.ConvertStringArray(lines)

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("cant create file [%v]", filePath)
	}
	defer f.Close()

	table := tablewriter.NewWriter(f)
	table.SetHeader([]string{"Type", "Number", "Comment", "Registry", "Date"})
	table.SetFooter([]string{"", "", "", "Total", fmt.Sprintf("%d", len(lines))})
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
	})
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()

	return nil
}

func init() {
	RootCmd.AddCommand(sortCmd)
}
