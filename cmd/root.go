package cmd

import (
	"fmt"
	"os"

	"github.com/nullsploit01/cc-wc/cmd/utils"
	"github.com/spf13/cobra"
)

var countBytes bool
var countLines bool
var countWords bool
var countCharacters bool

var rootCmd = &cobra.Command{
	Use:   "cc-wc [flags] [file]",
	Short: "Counts the number of bytes or lines in a file",
	Long: `cc-wc is a simplified clone of the Unix wc (word count) tool, focusing primarily on counting bytes and lines.
				It allows you to quickly determine the size of a file in bytes or count the number of lines. 
				This tool supports counting bytes with the -c flag and lines with the -l flag. More functionalities might be added in future versions.

				Example:
					cc-wc -c filename.txt  # Counts the bytes in filename.txt
					cc-wc -l filename.txt  # Counts the lines in filename.txt
					cc-wc -w filename.txt  # Counts the words in filename.txt
					cc-wc -m filename.txt  # Counts the characters in filename.txt
					`,

	Run: func(cmd *cobra.Command, args []string) {
		var inputFile *os.File

		if len(args) < 1 {
			fileInfo, _ := os.Stdin.Stat()
			if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
				inputFile = os.Stdin
			} else {
				cmd.PrintErr("Error: A file name is required as an argument or data must be piped.\n")
				cmd.Usage()
				return
			}
		}

		if inputFile == nil {
			file, err := os.Open(args[0])
			if err != nil {
				cmd.PrintErrf("Error reading file: %v\n", err)
				return
			}

			inputFile = file
		}

		if countBytes {
			cmd.Printf("Bytes: %d\n", utils.ByteCount(inputFile))
		} else if countLines {
			cmd.Printf("Lines: %d\n", utils.LineCount(inputFile))
		} else if countWords {
			cmd.Printf("Words: %d\n", utils.WordCount(inputFile))
		} else if countCharacters {
			cmd.Printf("Characters: %d\n", utils.CharacterCount(inputFile))
		} else {
			cmd.Printf("%d %d %d %s\n", utils.LineCount(inputFile), utils.WordCount(inputFile), utils.ByteCount(inputFile), inputFile.Name())
		}

		defer inputFile.Close()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Command execution error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&countBytes, "count-bytes", "c", false, "Count the number of bytes in the specified file")
	rootCmd.Flags().BoolVarP(&countLines, "count-lines", "l", false, "Count the number of lines in the specified file")
	rootCmd.Flags().BoolVarP(&countWords, "count-words", "w", false, "Count the number of words in the specified file")
	rootCmd.Flags().BoolVarP(&countCharacters, "count-characters", "m", false, "Count the number of characters in the specified file")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
