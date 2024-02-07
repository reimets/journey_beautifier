package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"regexp"
	"time"
)

// Airport represents the structure for airport data from CSV
type Airport struct {
	Name         string
	Country      string
	Municipality string
	ICAO         string
	IATA         string
	Coordinates  string
}


// Main entry point for the program.
func main() {
    // Checks if the correct number of arguments are provided from the command line.
    if len(os.Args) != 4 {
        displayTheUsage()
        return
    }

    // Reads command line arguments.
	inputTxt := os.Args[1]
	outputTxt := os.Args[2]
	airportLookupCsv := os.Args[3]

    // Checks for errors such as input existence and usage of '-h' flag.
    if !checkingErrors(inputTxt, outputTxt, airportLookupCsv) {
        return
    }

    // Loads airport data from CSV file.
    airports, err := loadAirportData(airportLookupCsv)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Opens the input file for reading.
    inputFile, err := os.Open(inputTxt)
    if err != nil {
        fmt.Println("\033[1m\033[31m-------------------\033[0m\033[22m\n"+
            "\033[31m| Input not found |\033[0m\n"+
            "\033[1m\033[31m-------------------\033[0m\033[22m")
        return
    }
    defer inputFile.Close()

    allCodesFound, output, err := processItinerary(inputFile, airports)
    if err != nil {
        if err.Error() == "Input file is empty" {
            fmt.Println(err) // Outputs an error message in the terminal.
            return // Exits the program without creating or overwriting the output.txt file.
        }
        fmt.Println(err) // Outputs another error message in the terminal.
        return
    }

    // If not all codes were found, the program stops with an error message.
    if !allCodesFound {
        fmt.Println("\033[1m\033[31m----------------------------------------------------------\033[0m\033[22m\n"+
        "\033[31m| Error: Not all airport codes were found in the lookup. |\033[0m\n"+
        "\033[1m\033[31m----------------------------------------------------------\033[0m\033[22m")
        return // Program stops and does not overwrite output.txt file
    }    
    
    fmt.Println(output) // prints output info to terminal

    output = trimColor(output) // removes colors from text, because of need to print info to putput.txt file

    // Writes the processed data to the output file.
    err = os.WriteFile(outputTxt, []byte(output), 0644)
    if err != nil {
        fmt.Println("\033[31mError writing output:\033[0m", err)
        return
    }

    // Addition: Displays success message after successfully writing to the output file.
    fmt.Println("\033[1m\033[32m-------------------------------------\033[0m\033[22m\n"+
        "\033[1m\033[32m| Itinerary processed successfully. |\033[0m\033[22m\n"+
        "\033[1m\033[32m-------------------------------------\033[0m\033[22m")
}


// Processes travel details, replacing airport codes and formatting dates.
func processItinerary(inputFile *os.File, airports map[string]Airport) (bool, string, error) {
    scanner := bufio.NewScanner(inputFile)
    var output strings.Builder
    allCodesFound := true
    isFirstLine := true

    for scanner.Scan() {
        if isFirstLine {
            isFirstLine = false
            // Checks if the first line is empty
            if len(scanner.Text()) == 0 {
                continue // Continues to the next line if the first line is empty
            }
        }

        line := scanner.Text()
        line = cleanUpText(line) // Removes extra spaces and empty lines.
        processedLine, codesFound := replaceAirportCodes(line, airports)
        if !codesFound {
            allCodesFound = false
        }
        processedLine = formatDateAndTime(processedLine)
        output.WriteString(processedLine + "\n")
    }

    // Checks if no lines were read (file may be empty)
    if isFirstLine {
        return false, "", fmt.Errorf("\033[1m\033[31m-----------------------\033[0m\033[22m\n"+
            "\033[31m| Input file is empty |\033[0m\n"+
            "\033[1m\033[31m-----------------------\033[0m\033[22m")
    }

    if err := scanner.Err(); err != nil {
        return false, "", fmt.Errorf("\033[31mError reading input: %v \033[0m", err)
    }

    finalOutput := cleanUpText(output.String()) // Checks that the final result is cleaned up as well.
    return allCodesFound, finalOutput, nil
}


// removes colors from text, because of need to print info to putput.txt file
// without that func output.txt file lines were with strange marks
func trimColor(text string) string {
	// Defineerib regulaaaravaldis, mis sobitub ANSI värvikoodidega.
	colorCodeRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	// Asendab kõik leitud värvikoodid tühja stringiga.
	return colorCodeRegex.ReplaceAllString(text, "")
}


// Replaces airport codes with their full names.
func replaceAirportCodes(line string, airports map[string]Airport) (string, bool) {
    allCodesFound := true
    reIATA := regexp.MustCompile(`#\w{3}`)
    reICAO := regexp.MustCompile(`##\w{4}`)

	// Replaces ICAO and IATA codes with airport names. Order matters!
    line = reICAO.ReplaceAllStringFunc(line, func(code string) string {
        if airport, exists := airports[code[2:]]; exists {
            return "\033[1m\033[33m"+airport.Name+"\033[0m\033[22m"
        } else {
            allCodesFound = false
            return code // The code is left unchanged if not found.
        }
    })

    line = reIATA.ReplaceAllStringFunc(line, func(code string) string {
        if airport, exists := airports[code[1:]]; exists {
            return "\033[1m\033[33m"+airport.Name+"\033[0m\033[22m"
        } else {
            allCodesFound = false
            return code // The code is left unchanged if not found.
        }
    })
    return line, allCodesFound
}


// Loads airport data from a CSV file.
func loadAirportData(filePath string) (map[string]Airport, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("\033[1m\033[31m----------------------------\033[0m\033[22m\n"+
            "\033[31m| Airport lookup not found |\033[0m\n"+
            "\033[1m\033[31m----------------------------\033[0m\033[22m")
    }
    defer file.Close()

    reader := csv.NewReader(bufio.NewReader(file))
    records, err := reader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("\033[1m\033[31m--------------------------------\033[0m\033[22m\n"+
			"\033[31m| Error reading airport lookup |\033[0m\n"+
			"\033[1m\033[31m--------------------------------\033[0m\033[22m")
    }

    airports := make(map[string]Airport)
    for i, record := range records {
        if i == 0 { // Checks header
            for _, header := range record {
                if header == "" {
                    return nil, fmt.Errorf("\033[1m\033[31m----------------------------\033[0m\033[22m\n"+
                    "\033[31m| Airport lookup malformed |\033[0m\n"+
                    "\033[1m\033[31m----------------------------\033[0m\033[22m")
                }
            }
            continue
        }

        // Creates an airport data structure.
        airport := Airport{
            Name:         record[0],
            Country:      record[1],
            Municipality: record[2],
            ICAO:         record[3],
            IATA:         record[4],
            Coordinates:  record[5],
        }

        airports[airport.ICAO] = airport
        airports[airport.IATA] = airport

        // Checks data integrity.
        if len(record) != 6 || airport.IATA == "" || airport.ICAO == "" || airport.Name == "" || airport.Municipality == "" || airport.Country == "" || airport.Coordinates == "" {
            return nil, fmt.Errorf("\033[1m\033[31m----------------------------\033[0m\033[22m\n"+
                "\033[31m| Airport lookup malformed |\033[0m\n"+
                "\033[1m\033[31m----------------------------\033[0m\033[22m")
        }
    }
    return airports, nil
}


// Cleans up text by removing extra spaces and multiple consecutive newlines.
func cleanUpText(text string) string {
    // Removes extra spaces and tabs.
    spacePattern := regexp.MustCompile(`[ \t]+`)
    text = spacePattern.ReplaceAllString(text, " ")

    // Removes multiple consecutive newlines.
    newlinePattern := regexp.MustCompile(`\n\n+`)
    text = newlinePattern.ReplaceAllString(text, "\n\n")

    // Removes empty lines from the start and end of the text.
    text = strings.TrimSpace(text)

    return text
}


// formatDateAndTime identifies and formats date and time strings in the text.
// This function uses regular expressions to detect patterns of dates and times
// and converts them into a human-readable format. If a date and time pattern is detected
// in the text, it is replaced with the appropriate format. If no pattern is found or it is
// invalid, the original text is returned.
func formatDateAndTime(line string) string {
    // Checks if the line starts with the specified text and leaves it in the original color
    if strings.HasPrefix(line, "Your flight departs from") {
		return line // returns line without changing color (here)
	}
	// working with patterns using regular expressions
	regs := regexp.MustCompile(`\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}([Z+-]\d{2}:\d{2}|Z)?)\)`) 
	matches := regs.FindStringSubmatch(line)

	if len(matches) > 1 {
		parsedTime, err := time.Parse("2006-01-02T15:04Z07:00", matches[1])
		if err != nil {
			return "\033[31m"+line+"\033[0m" // Return original line if date/time is invalid (in red)
		}

		var parsedDate string
		// Formats the date or time according to the prefix or suffix.
		if strings.HasPrefix(line, "D") {
			parsedDate = "\033[1m\033[44m"+parsedTime.Format(" 02 Jan 2006 ")+"\033[0m\033[22m"
		} else if strings.HasPrefix(line, "T12") {
			parsedDate = "\033[32m"+parsedTime.Format("03:04PM (-07:00)")+"\033[0m"
		} else if strings.HasPrefix(line, "T24") {
			parsedDate = "\033[32m"+parsedTime.Format("15:04 (-07:00)")+"\033[0m"
		} else {
			return "\033[31m"+line+"\033[0m" // Return original line if wrong prefix
		}
		if strings.HasSuffix(line, "Z" ) {
			parsedDate += " (+00:00)"
		} 
		return parsedDate 
	} 
	return line // Returns original line if no pattern found
}


// Checks the validity of input, output, and airport lookup file as well as usage of '-h' flag.
// This function checks that the input and output files, as well as the airport lookup file,
// do not overlap and are correctly provided. It also handles the '-h' flag to display
// usage instructions if requested by the user. The function returns 'true' if all checks
// are successful, otherwise 'false', and the usage instructions are displayed.
func checkingErrors(inputTxt, outputTxt, airportLookupCsv string) bool {
	// checking errors
	helpFlag := flag.Bool("h", false, "Displays the usage")
	flag.Parse()

	if *helpFlag {
		displayTheUsage()
		return false
	}

	if inputTxt == outputTxt || inputTxt == airportLookupCsv || outputTxt == airportLookupCsv {
		displayTheUsage()
		return false
	}
	return true
}


// displayTheUsage displays the usage instructions for correctly running the program.
// This function outputs explanatory examples and guidelines that help users understand
// how to use the program, including the purpose and required format of each input parameter.
func displayTheUsage() {
	fmt.Println("\033[31mitinerary usage:\033[0m")
	fmt.Println("\033[41m go run . ./input.txt ./output.txt ./airport-lookup.csv \033[0m")
}
