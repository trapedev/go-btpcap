package lib

import "os"

func ReadFileAsByteArray(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, err
}