# CEP Information Retrieval

This Go source code is designed to retrieve address information based on a Brazilian CEP (Postal Code) using two different APIs: `cdn.apicep.com` and `viacep.com.br`. It demonstrates the usage of goroutines, channels, HTTP requests, and context management.

## How to Use

1. Make sure you have Go installed on your system.
2. Copy and paste the provided code into a `.go` file (e.g., `cep_retrieval.go`).
3. Open a terminal and navigate to the directory containing the file.
4. Run the following command to execute the code:
   
   ```
   go run main.go
   ```

5. You will be prompted to enter a CEP (Brazilian Postal Code).
6. Provide the CEP without dashes (e.g., `12345678`).
7. The program will retrieve address information using both APIs and display the results.

## Code Explanation

1. The program starts by creating two channels, `c1` and `c2`, to communicate with the goroutines responsible for API calls and data retrieval.

2. The user is prompted to input a CEP. If the CEP provided doesn't contain a dash, the program inserts a dash in the appropriate position to match the expected format.

3. If the length of the CEP is greater than 9, the program considers it an invalid input and terminates.

4. The program sets up a context with a timeout of 5 seconds to ensure that the API calls don't hang indefinitely.

5. Two goroutines are started concurrently:
   - The first goroutine performs an HTTP GET request to the `cdn.apicep.com` API to retrieve address information using the provided CEP.
   - The second goroutine performs a similar HTTP GET request to the `viacep.com.br` API.

6. The results from both APIs are sent to their respective channels (`c1` and `c2`) upon successful retrieval.

7. The `select` statement is used to wait for results from either channel. It will print the data received from the first API that responds, or it will print "timeout" if no response is received within 1 second.

## Libraries Used

- `context`: Used to manage context and timeouts for the HTTP requests.
- `fmt`: Used for formatted printing.
- `io`: Used for reading the response body from HTTP requests.
- `net/http`: Used for making HTTP requests.

## Note

Keep in mind that APIs can change or become unavailable over time. The code's functionality may be affected if the APIs mentioned in the code are no longer accessible or if their endpoints change.

**Disclaimer:** This code is provided as-is, and I recommend checking the APIs' documentation or terms of use before using this code in a production environment or making multiple requests in a short period to avoid potential issues.
