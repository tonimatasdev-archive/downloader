package src

import (
	"fmt"
	"os"
	"time"
)

func CreateFile(name string) *os.File {
	file, err := os.Create("./" + name)

	if err != nil {
		panic(err)
	}

	return file
}

func Downloaded(fileSize string, file *os.File) {
	var oldSize int64
	startTime := time.Now()

	for {
		size := GetSize(file)
		formattedSize := FormattedLength(ParseStr(size))
		percentage := FormattedSizePercentage(size, fileSize)
		sizePerSecond := FormattedSizePerSecond(size, oldSize, 50)
		elapsedTime := time.Now().Sub(startTime)

		fmt.Print(fmt.Sprintf("\rDownloaded: %s - %s%% - %s - %.2fs          ", formattedSize, percentage, sizePerSecond, elapsedTime.Seconds()))
		oldSize = size

		time.Sleep(time.Millisecond * 50)
	}
}

func GetSize(file *os.File) int64 {
	stat, err := file.Stat()

	if err != nil {
		panic(err)
	}

	return stat.Size()
}
