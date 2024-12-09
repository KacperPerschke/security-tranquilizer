package archiver

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

const conjecturalDirMode = 0755
const implicitDirMode = 0755

func UnpackFrom(gzipStream io.Reader) error {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf(
				"the next element in the archive cannot be read with *tar.Reader.Next(): %w",
				err,
			)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, implicitDirMode); err != nil {
				return fmt.Errorf(
					"failed to create a directory ‘%s’: %w",
					header.Name,
					err,
				)
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				return fmt.Errorf(
					"failed to create a file ‘%s’: %w",
					header.Name,
					err,
				)
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf(
					"failed to populate file ‘%s’ with data from archive: %w",
					header.Name,
					err,
				)
			}
			outFile.Close()

		default:
			return fmt.Errorf(
				"consecutive archive element ‘%s’ is of the nonprocessable type ‘%s’",
				header.Name,
				header.Typeflag,
			)
		}
	}
	return nil
}
