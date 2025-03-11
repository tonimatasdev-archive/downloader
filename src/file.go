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
	startTime := time.Now()

	go updateSizePerSecond(file)

	for {
		size := GetSize(file)
		formattedSize := FormattedLength(ParseStr(size))
		percentage := FormattedSizePercentage(size, fileSize)
		elapsedTime := time.Now().Sub(startTime)

		fmt.Print(fmt.Sprintf("\rDownloaded: %s - %s%% - %s - %.2fs          ", formattedSize, percentage, sizePerSecond, elapsedTime.Seconds()))

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
