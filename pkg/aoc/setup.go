package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/anaskhan96/soup"
)

func Setup(year string, day string) error {
	url := createURL(year, day)

	// Get the sample for the puzzle
	getSample(url)

	// Get the input for the puzzle
	getInput(url)

	// Check if solution exists and if so copy
	filename := fmt.Sprintf("%s/day%s.go", year, day)
	_, err := os.Open(filename)
	if err == nil {
		err := copyFile(filename, "pkg/aoc/solve.go")
		if err != nil {
			return err
		}
	}

	fmt.Printf("%s-%s AOC puzzle has been setup successfully 🎄🎁🎅\n", year, day)

	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func createURL(year string, day string) string {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s", year, day)
	return url
}

func getSample(url string) error {
	// Parse the HTML
	resp, err := soup.Get(url)
	if err != nil {
		return err
	}
	doc := soup.HTMLParse(resp)
	block := doc.Find("pre").Find("code").Text()

	// Write sample to file
	err = os.WriteFile("sample", []byte(block), 0644)
	if err != nil {
		return err
	}

	return nil
}

func getInput(url string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/input", url), nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", "github.com/benpops89/aoc-in-go")
	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", os.Getenv("AOC_SESSION")
	req.AddCookie(cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download file with status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Write the body to the destiniation file
	err = os.WriteFile("input", body, 0644)
	if err != nil {
		return err
	}
	return nil
}
