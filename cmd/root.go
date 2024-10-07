package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cc-wc [flags] [file]",
	Short: "Counts the number of bytes in a file",
	Long: `cc-wc is a simplified clone of the Unix wc (word count) tool, focusing on byte counting. 
It can be used to quickly determine the size of a file in bytes. This tool currently supports only 
the counting of bytes, but may be extended in the future to include word, line, and character counts. 
For example:

To count the bytes in a file:
cc-wc -c filename.txt`,

	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("count-bytes", "c", false, "Count number of bytes in the specified file")
}
