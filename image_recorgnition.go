package puppet

import gocr "github.com/otiai10/gosseract"

// ImageProcessing does the optical character recognition(OCR)
func ImageProcessing(filePath string) (string, error) {
	client := gocr.NewClient()
	defer client.Close()

	client.SetImage(filePath)
	text, err := client.Text()
	if err != nil {
		return "", err
	}

	return text, nil
}
