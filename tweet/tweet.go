package tweet

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Get a random quote tweet from a directory with subfolders and lyric files
func GetTweet(directory string) (string, error) {
	var quotation, content string
	{
		lyrics, err := getFiles(directory)
		if err != nil {
			return "", err
		}
		if len(lyrics) < 1 {
			return "", errors.New("no files given")
		}
		lyric := lyrics[getRandomIndex(0, len(lyrics)-1)]

		lines, err := readLines(lyric)
		if err != nil {
			return "", err
		}
		if len(lines) < 2 {
			return "", errors.New("no lyrics given")
		}
		quotation = lines[0]
		for content == "" {
			content = lines[getRandomIndex(1, len(lines)-1)]
		}
	}

	return fmt.Sprintf("%s\n-\n%s", content, quotation), nil
}

func getRandomIndex(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getFiles(path string) ([]string, error) {
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, f := range fileInfo {
		if f.IsDir() {
			content, err := getFiles(path + "/" + f.Name())
			if err != nil {
				return nil, err
			}
			files = append(files, content...)
		} else {
			files = append(files, path+"/"+f.Name())
		}
	}
	return files, nil
}
