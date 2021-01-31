package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Scraper scrapes for a key given a path
type Scraper interface {
	Scrape() (*Report, error)
}

// DefaultScraper is a generic implementation of Scraper
type DefaultScraper struct {
	key  string
	path string
}

// NewDefaultScraper create a generic scraper for key at path
func NewDefaultScraper(key, path string) (*DefaultScraper, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w "+path, ErrNotExist)
	}
	absPath, _ := filepath.Abs(path)
	return &DefaultScraper{
		key:  key,
		path: absPath,
	}, nil
}

// Scrape path and generate report
func (s *DefaultScraper) Scrape() (*Report, error) {
	var report Report
	err := filepath.Walk(s.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			//TODO continue on error
			return err
		}

		if info.Mode().IsRegular() {
			f, err := os.Open(path)
			if err != nil {
				//TODO continue on error
				return err
			}

			scanner := bufio.NewScanner(f)
			var line int
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), s.key) {
					// absPath, _ := filepath.Abs(path)
					report = append(report, &ReportEntry{
						Line: line,
						Path: path,
					})
				}
				line++
			}
		}
		return nil

	})

	if err != nil {
		return nil, fmt.Errorf("error walking path %w", err)
	}
	return &report, nil
}
