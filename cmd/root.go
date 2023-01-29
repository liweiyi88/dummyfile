package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/liweiyi88/dummyfile/dummyfile"
	"github.com/spf13/cobra"
)

var size string

func init() {
	rootCmd.Flags().StringVarP(&size, "size", "s", "", "the file size")
}

var rootCmd = &cobra.Command{
	Use:   "/path/to/dummy.txt",
	Short: "Create a fixed-size dummy file with random contents",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create dummy file: %v", err)
		}

		defer func() {
			err := file.Close()
			if err != nil {
				log.Printf("failed to close file: %s", file.Name())
			}
		}()

		if byteSize, err := strconv.ParseInt(size, 10, 64); err == nil {
			return dummyfile.Create(file, byteSize*dummyfile.B)
		}

		suffixes := dummyfile.GetSizeSuffixes()
		for _, suffix := range suffixes {
			filesize := strings.ToUpper(size)

			if !strings.HasSuffix(filesize, suffix) {
				continue
			}

			numberPart := strings.TrimSpace(strings.TrimSuffix(filesize, suffix))
			byteSize, err := strconv.ParseInt(numberPart, 10, 64)

			if err != nil {
				return fmt.Errorf("%s doesn't have valid number", filesize)
			}

			return dummyfile.Create(file, byteSize*dummyfile.SizeMap[suffix])

		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
