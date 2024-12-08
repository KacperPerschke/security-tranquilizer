package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/KacperPerschke/security-tranquilizer/common"
	"github.com/KacperPerschke/security-tranquilizer/encoder"
)

var encodeCmd = &cobra.Command{
	Use:   "encode [flags] FILE_TO_BE_ENCODED",
	Short: "Your friendly encoder",
	Long:  "Encodes the contents of the given file as described in the Readme.\nThe name of a file to be encoded is given as an argument to the call.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkArgsCount(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		outFileName, inList, err := procCmdArgs(cmd, args)
		if err != nil {
			fmt.Printf("\nThere was an with command line args: %q\n\n", err.Error())
			os.Exit(1)
		}
		if err := encoder.EncodeFileToPNG(outFileName, inList); err != nil {
			fmt.Printf("\nThere was an error while encoding: %q\n\n", err.Error())
			os.Exit(1)
		}
	},
}

const (
	optOutFname        = "output"
	optOutFnameShort   = "o"
	timeFmtForFilename = "2006-01-02_15:04:05"
)

var outFileName string

func init() {
	rootCmd.AddCommand(encodeCmd)
	addOutputFlag(encodeCmd)
	encodeCmd.Flags().StringVarP(
		&outFileName,
		optOutFname,
		optOutFnameShort,
		fmt.Sprintf(
			"st_%s.png",
			time.Now().Format(timeFmtForFilename),
		),
		"name of output file",
	)
}

func procCmdArgs(cmd *cobra.Command, args []string) (string, []common.FileInfo, error) {
	emptyIL := []common.FileInfo{}
	oFN, err := cmd.Flags().GetString(optOutFname)
	if err != nil {
		return "", emptyIL, fmt.Errorf(
			"the ‘%s’ argument could not be extracted",
			optOutFname,
		)
	}
	if err := common.FSCanCreate(oFN); err != nil {
		return "", emptyIL, err
	}
	il, err := procIL(args)
	if err != nil {
		return "", emptyIL, err
	}
	return oFN, il, err
}

func procIL(al []string) ([]common.FileInfo, error) {
	intermediateIL := []common.FileInfo{}
	for _, el := range al {
		elPrep, err := common.FSPrepElI(el)
		if err != nil {
			return []common.FileInfo{}, err
		}
		intermediateIL = append(
			intermediateIL,
			elPrep,
		)
	}
	il := []common.FileInfo{}
	for _, prepEl := range intermediateIL {
		if prepEl.IsFileOrSymlink() {
			il = append(il, prepEl)
			continue
		}
		return []common.FileInfo{}, fmt.Errorf(
			"‘%s’ is not a file nor symlink",
			prepEl.Path,
		)
	}
	return il, nil

}
