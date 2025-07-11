package aoc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func DownloadInput(year, problem string) (io.ReadCloser, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, problem)
	req, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("cookie", os.Getenv("AOC_SESSION_COOKIE"))

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("couldn't request input %w", err)
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("error from request %s : status %s", url, resp.Status)
	}

	return resp.Body, nil
}

func SaveInput(reader io.ReadCloser, outputFile string) error {
	defer reader.Close()

	path := path.Dir(outputFile)

	os.MkdirAll(path, 0777)

	f, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return fmt.Errorf("couldn't open input file to save %w", err)
	}
	defer f.Close()

	br := bufio.NewReader(reader)

	bw := bufio.NewWriter(f)

	_, err = io.Copy(bw, br)

	if err != nil {
		return fmt.Errorf("couldn't write to input file %w", err)
	}

	return nil
}

func GetInput(year, problem string) (string, error) {
	filename := fmt.Sprintf("../inputs/%s/%s.txt", year, problem)
	if _, err := os.Stat(filename); err == nil {
		return read(filename)
	}

	reader, err := DownloadInput(year, problem)
	if err != nil {
		return "", fmt.Errorf("couldn't download input %w", err)
	}

	if err := SaveInput(reader, filename); err != nil {
		return "", fmt.Errorf("couldn't save input %w", err)
	}
	return read(filename)
}

func read(filename string) (string, error) {
	b, err := os.ReadFile(filename)

	if err != nil {
		return "", fmt.Errorf("failed to read saved input %w", err)
	}

	return string(b), nil
}
