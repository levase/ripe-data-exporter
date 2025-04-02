package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"ripe-data-exporter/internal/api"
	"ripe-data-exporter/internal/storage"
	"ripe-data-exporter/pkg/types"
)

type Config struct {
	APIURL      string
	OutputFile  string
	HTTPTimeout time.Duration
}

func main() {
	asn, err := getUserInput()
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	config := Config{
		APIURL:      fmt.Sprintf("https://stat.ripe.net/data/announced-prefixes/data.json?resource=%s", asn),
		OutputFile:  fmt.Sprintf("as%s_response.json", asn),
		HTTPTimeout: 30 * time.Second,
	}

	if err := run(config); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func getUserInput() (string, error) {
	fmt.Print("Input AS number (e.g., 12389 or AS12389): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if err := types.ValidateASN(input); err != nil {
		return "", fmt.Errorf("invalid AS number: %w", err)
	}

	re := regexp.MustCompile(`^(AS)?(\d+)$`)
	matches := re.FindStringSubmatch(input)
	return matches[2], nil
}

func run(config Config) error {
	response, err := api.FetchRIPEData(config.APIURL, config.HTTPTimeout)
	if err != nil {
		return err
	}

	if err := storage.SaveToFile(response, config.OutputFile); err != nil {
		return err
	}

	log.Printf("Data successfully saved to %s", config.OutputFile)
	return nil
}
