package webshot

import (
	"os/exec"
)

/*
ImageProcessing does the optical character recognition(OCR)

This method takes in two parameters which are: filePath and the name the processed image is to be saved as.
*/
func ImageProcessing(filePath, saveAs string) error {
	cmd := "tesseract"

	args := []string{filePath, saveAs}

	_, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return err
	}

	return nil
}
