package archiver

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"

	"github.com/KacperPerschke/security-tranquilizer/common"
)

func PackToArchive(fList []common.FileInfo) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	err := populateArchive(fList, buf)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func addToArchive(tw *tar.Writer, el common.FileInfo) error {
	header, err := tar.FileInfoHeader(
		el.Info,
		el.Info.Name(), // FileInfoHeader only takes the basename. Read †).
	)
	if err != nil {
		return err
	}

	// †) In order to save the directory structure in the archive
	// we have to give the whole path to the file.
	header.Name = el.Path

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	fh, err := os.Open(el.Path)
	if err != nil {
		return err
	}
	defer fh.Close()
	_, err = io.Copy(tw, fh)
	if err != nil {
		return err
	}

	return nil
}

func populateArchive(els []common.FileInfo, b io.Writer) error {
	gw := gzip.NewWriter(b)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, el := range els {
		err := addToArchive(tw, el)
		if err != nil {
			return err
		}
	}

	return nil
}
