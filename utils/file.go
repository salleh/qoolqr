/*
 * Project: QoolQR
 * Filename: /utils/file.go
 * Created Date: Friday September 13th 2024 11:01:07 +0800
 * Author: Sallehuddin Abdul Latif (salleh@sallehuddin.com)
 * Company: Sallehuddin Industries
 * --------------------------------------
 * Last Modified: Friday September 13th 2024 11:40:31 +0800
 * Modified By: Sallehuddin Abdul Latif (salleh@sallehuddin.com)
 * --------------------------------------
 * Copyright (c) 2024 Sallehuddin Industries
 */

package utils

import (
	"image"
	"os"
)

func GetImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
