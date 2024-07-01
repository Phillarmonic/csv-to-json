


# CSV to JSON Converter

This project is a simple command-line tool written in Go that converts CSV files to JSON format.  


**How to Use**
Download the appropriate binary for your operating system and architecture from the releases page. The binaries are named according to the OS and architecture, for example:  


csv-to-json-linux-amd64
csv-to-json-linux-arm64


After downloading, you can run the binary from the command line. Here's how to use it:  


To convert a CSV file to a JSON file, run the binary with the CSV file name as an argument. For example, if your CSV file is named 'data.csv', you would run:  <pre>./csv-to-json-linux-amd64 data.csv </pre> This will convert the 'data.csv' file into a JSON file named 'data.json' by default.  
If you want to output the result to stdout instead of a file, use the '-stdout' flag:  <pre>./csv-to-json-linux-amd64 -stdout data.csv </pre>
If you want to specify a different output file name, use the '-output' flag:  <pre>./csv-to-json-linux-amd64 -output output.json data.csv </pre> This will convert the 'data.csv' file into a JSON file named 'output.json'.  
Please note that you need to replace csv-to-json-linux-amd64 with the name of the binary you downloaded.
