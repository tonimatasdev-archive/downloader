package src

import (
	"fmt"
	"strconv"
	"time"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

func FormattedLength(lengthStr string) string {
	length := ParseInt(lengthStr)

	var sizeType string
	var size float64

	if length > KB && length < MB {
		size = float64(length) / KB
		sizeType = "KB"
	} else if length > MB && length < GB {
		size = float64(length) / MB
		sizeType = "MB"
	} else {
		size = float64(length) / GB
		sizeType = "GB"
	}

	return fmt.Sprintf("%.2f %s", size, sizeType)
}

func FormattedSizePercentage(size int64, maxSize string) string {
	maxSizeInt, err := strconv.ParseInt(maxSize, 10, 64)

	if err != nil {
		panic(err)
	}

	percentage := float64(size) / float64(maxSizeInt) * 100

	return fmt.Sprintf("%.2f", percentage)
}

func FormattedSizePerSecond(size int64, oldSize int64, millis int64) string {
	downloadedSize := size - oldSize
	timesToSecond := time.Second.Milliseconds() / millis
	downloadPerSecond := downloadedSize * timesToSecond

	return FormattedLength(ParseStr(downloadPerSecond)) + "/s"
}
