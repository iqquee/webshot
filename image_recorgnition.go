package puppet

import (
	"os/exec"
)

// ImageProcessing does the optical character recognition(OCR)
func ImageProcessing(filePath, saveAs string) error {
	cmd := "tesseract"

	args := []string{filePath, saveAs}

	_, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return err
	}

	return nil
}
