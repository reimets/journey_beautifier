# Itinerary Processor: README.md

## Overview

The Itinerary Processor is a Go application designed to process flight itineraries from a text file, replace airport codes with their full names, format dates and times, and output the processed information to another text file. It utilizes airport data from a CSV file to lookup airport names by their IATA or ICAO codes.

## Features

- **Airport Code Replacement:** Converts IATA and ICAO airport codes in the itinerary to their full names.
- **Date and Time Formatting:** Identifies and formats date and time strings into a more readable format.
- **Input and Output Handling:** Reads itineraries from an input text file and writes the processed itineraries to an output text file.
- **Error Checking:** Checks for common errors such as missing input files or incorrect command line arguments.
- **File Handling:** Reads from and writes to text files, and processes CSV files for airport data.

## Requirements

Go (Golang) installed on your system.
A CSV file containing airport data with the following columns: Name, Country, Municipality, ICAO, IATA, Coordinates.
An input text file containing flight itineraries with airport codes, dates, and times.

## Usage

Ensure you have the required CSV file with airport data and the input text file with flight itineraries.
Run the program using the Go command line:

```php
go run . <input_file.txt> <output_file.txt> <airport_data.csv>
```

#### Replace:
<input_file.txt> with the path to your input text file
<output_file.txt> with the path where you want the processed itineraries to be saved (file will be created automatically), 
and 
<airport_data.csv> with the path to your CSV file containing airport data.

## Example Command
```bash
go run . ./input.txt ./output.txt ./airport-lookup.csv
```
This command processes itineraries from input.txt, looks up airport data in airport-lookup.csv, and writes the processed itineraries to output.txt.

## Error Handling
The program includes error handling for various scenarios, such as missing files, empty input files, and unrecognized airport codes. If any issues are encountered, the program will output an error message and halt execution without overwriting any existing output file.

## Functions
- **loadAirportData(filePath string):** Loads airport data from a CSV file into a map.
- **processItinerary(inputFile \*os.File, airports map[string]Airport):** Processes the itinerary from the input file.
- **replaceAirportCodes(line string, airports map[string]Airport):** Replaces airport codes in the text with corresponding airport names.
- **formatDateAndTime(line string):** Formats date and time strings in the text.
- **checkingErrors(inputTxt, outputTxt, airportLookupCsv string):** Checks for errors in input arguments.
- **displayTheUsage():** Displays usage instructions for the program.

## Contributing
To contribute to the Itinerary Processor, please fork the repository, make your changes, and submit a pull request. Contributions are welcome to improve the application, extend its capabilities, or fix any issues.

## License
Reigo Reimets

## Authors
Reigo Reimets
