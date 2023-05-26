package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatal(err)
	}

	return !info.IsDir()
}

func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func SliceContains[M comparable](s []M, value M) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func SplitToString(a []uint, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(b, sep)
}

func StringSlice2UintSlice(a []string) ([]uint, error) {
	b := make([]uint, len(a))
	for i, v := range a {
		aux, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return b, err
		}
		b[i] = uint(aux)
	}
	return b, nil
}

func ReadTail(filePath string, limit int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	fileSize := fileInfo.Size()

	offset := fileSize - int64(limit)
	if offset < 0 {
		offset = 0
	}

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, limit)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func RemoveSliceItem[T comparable](slice []T, itemToRemove T) []T {
	indexToRemove := -1

	for i, item := range slice {
		if item == itemToRemove {
			indexToRemove = i
			break
		}
	}
	if indexToRemove != -1 {
		slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	}

	return slice
}
