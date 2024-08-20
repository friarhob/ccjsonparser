# ccjsonparser

`ccjsonparser` is a command-line tool for validating JSON data.

This is a Go study project based on the [Coding Challenges](https://codingchallenges.fyi) exercises, particularly [this one](https://codingchallenges.fyi/challenges/challenge-json-parser).

## Development Status

This is still a work in progress. So far, we have:

Validate | Done
:--|:-:
`{}` | ✓
String keys and string values | ✓
Boolean and null values | ✓
Numeric values | 
Object and array values | 
Lists of JSONs | 


## Features

- **Validate JSON**: Validate JSON from file or standard input.

## Installation

You can build and install `ccjsonparser` from source using Go. Make sure you have Go installed on your system. 

1. Clone the repository:

   ```bash
   git clone https://github.com/username/ccjsonparser.git
   cd ccjsonparser
   ```

1. Build the executable:
   ```bash
   go build -o ccjsonparser ./cmd/ccjsonparser
   ```

1. (Optional) Move the binary to a directory in your PATH for easy access:
   ```bash
   mv ccjsonparser /usr/local/bin/
   ```

## Usage
   ```bash
   ccjsonparser [filepath]
   ```

### Examples
1. **Validate JSON from a file:**
   ```bash
   ccjsonparser example.json
   ```

1. **Validate JSON from standard input:**
   ```bash
   echo '{"name":"John","age":30}' | ccjsonparser
   ```

## Exit Codes

Code | Description
:-:|---
0 | Valid JSON
1 | Invalid JSON
2 | Error reading from file/stdin

## Test files

Test files (in `testdata/` folder) came from [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-json-parser) and [JSON.org](http://www.json.org/JSON_checker/test.zip).

## Future Ideas

- **Format JSON**: Beautify and format JSON data for better readability.
- **Extract Data**: Retrieve specific fields or values from JSON objects.
- **Filter Data**: Apply filters to JSON data to extract relevant information.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.