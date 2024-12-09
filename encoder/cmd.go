package encoder

import (
	"fmt"
	"image/png"
	"os"

	"github.com/KacperPerschke/security-tranquilizer/archiver"
	"github.com/KacperPerschke/security-tranquilizer/common"
	"github.com/KacperPerschke/security-tranquilizer/img"
)

func EncodeToPNG(of string, il []common.FileInfo) error {
	b, err := archiver.PackToArchive(il)
	if err != nil {
		return err // Should we include more information?
	}
	wErr := os.WriteFile("check.tar.gz", b, 0644)
	if wErr != nil {
		return err
	}

	img, err := img.PackToImg(b)
	if err != nil {
		return err
	}

	oHandle, err := os.OpenFile(string(of), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return fmt.Errorf("Problem during attempt to open file '%s' for writing: %w", of, err)
	}

	errPNG := png.Encode(oHandle, img)
	if errPNG != nil {
		return fmt.Errorf("Problem during attempt to write image to file '%s': %w", of, errPNG)
	}

	errClose := oHandle.Close()
	if err != nil {
		return fmt.Errorf("Problem during releasing handle to file '%s' after write: %w", of, errClose)
	}

	return nil
}
