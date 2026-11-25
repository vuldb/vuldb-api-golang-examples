# VulDB API – Go Examples

[![Language: Go](https://img.shields.io/badge/Language-Go-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: GPL-3.0](https://img.shields.io/badge/License-GPL--3.0-blue.svg)](LICENSE)
[![VulDB API](https://img.shields.io/badge/VulDB-API%20Documentation-663399)](https://vuldb.com/kb/api)

Go (Golang) example for retrieving vulnerability intelligence through the official [VulDB API](https://vuldb.com/kb/api).

[VulDB](https://vuldb.com/) provides structured information about security vulnerabilities, technical details, risk assessments and cyber threat intelligence (CTI). Its API enables developers and security teams to integrate vulnerability data into their own applications, vulnerability management platforms, SIEM and SOAR solutions, and automated security workflows.

This repository provides a standalone Go demonstration using only the Go standard library. It can be used to explore the API or as a starting point for a custom integration.

## Table of Contents

- [Features](#features)
- [Repository Structure](#repository-structure)
- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Usage Examples](#usage-examples)
- [API Request and Authentication](#api-request-and-authentication)
- [Response Handling](#response-handling)
- [API Credits and Rate Limits](#api-credits-and-rate-limits)
- [Troubleshooting](#troubleshooting)
- [Production Integration](#production-integration)
- [Contributing](#contributing)
- [License](#license)
- [Support and Resources](#support-and-resources)

## Features

- Retrieve the most recent vulnerability entries.
- Retrieve an individual vulnerability by its VulDB entry ID.
- Search for a vulnerability by its CVE identifier.
- Select basic or detailed vulnerability information.
- Authenticate using the `X-VulDB-ApiKey` HTTP header.
- Send HTTPS POST requests and display the API response in the terminal.
- Use native Go packages without external dependencies.

## Repository Structure

```text
vuldb-api-golang-examples/
├── VulDBApiDemoGo.go    # Standalone Go API demonstration
├── LICENSE             # GNU General Public License v3.0
└── README.md           # Repository documentation
```

## Requirements

- Go installed on your system (a currently supported release is recommended).
- A VulDB account with a personal API key and sufficient API credits.
- Internet access to `https://vuldb.com` over HTTPS.
- Git, if cloning the repository instead of downloading the source file.

The example uses the standard-library packages `bytes`, `fmt`, `io/ioutil` and `net/http`. No additional packages or Go module setup are required when running the source file directly.

## Getting Started

### 1. Download the repository

```bash
git clone https://github.com/vuldb/vuldb-api-golang-examples.git
cd vuldb-api-golang-examples
```

Alternatively, download the source file [VulDBApiDemoGo.go](VulDBApiDemoGo.go) directly from GitHub.

### 2. Obtain a VulDB API key

1. Visit [VulDB](https://vuldb.com/) and sign in or create an account.
2. Open your user profile to access your personal API key and check your available API credits.
3. Open `VulDBApiDemoGo.go` in a text editor.
4. Set the `personalAPIKey` variable to your personal API key:

   ```go
   personalAPIKey := "YOUR_PERSONAL_API_KEY"
   ```

> [!IMPORTANT]
> Treat your API key as a secret. Do not commit a populated key to Git, publish it in examples, or include it in logs or support requests. For automated integrations, use an environment variable or a secret management solution instead (see [Production Integration](#production-integration)).

### 3. Run the example

Verify that Go is available:

```bash
go version
```

Execute the demonstration:

```bash
go run VulDBApiDemoGo.go
```

With the default query settings, the program requests the five most recent vulnerability entries with basic information and prints the response to standard output.

Alternatively, compile and run a standalone executable on Linux or macOS:

```bash
go build -o vuldb-api-demo VulDBApiDemoGo.go
./vuldb-api-demo
```

On Windows (PowerShell):

```powershell
go build -o vuldb-api-demo.exe VulDBApiDemoGo.go
.\vuldb-api-demo.exe
```

## Configuration

All settings are defined inside the `main()` function in [VulDBApiDemoGo.go](VulDBApiDemoGo.go). The current example does not accept command-line arguments or automatically read environment variables.

| Variable | Default | Description |
| --- | --- | --- |
| `url` | `https://vuldb.com/api` | API endpoint used by the demonstration. |
| `personalAPIKey` | Empty string | Personal API key used for authentication. |
| `userAgent` | `VulDB API Advanced Go Demo Agent` | HTTP User-Agent identifying the client. |
| `recent` | `"5"` | Requested number of recent vulnerability entries. |
| `details` | `"0"` | `0` for basic data; `1` for detailed data. |
| `id` | Empty string | Numeric VulDB entry ID to retrieve. |
| `cve` | Empty string | CVE identifier to search for. |

### Query Selection

The program performs one query per execution. The following priority applies:

| Priority | Condition | Request type |
| --- | --- | --- |
| 1 | `cve` is not empty | `search` |
| 2 | `cve` is empty and `id` is not empty | `id` |
| 3 | Both `cve` and `id` are empty | `recent` |

The `details` setting applies to all three query types. If both `id` and `cve` are populated, the CVE search takes precedence. The `recent` value is ignored during an ID or CVE lookup.

## Usage Examples

Modify the existing variables in `VulDBApiDemoGo.go` as shown below, save the file and execute `go run VulDBApiDemoGo.go` after each change. The examples are independent alternatives; replace the existing query settings rather than adding duplicate variable declarations.

### Retrieve recent vulnerabilities

Request the five most recent entries with basic information:

```go
recent := "5"
details := "0"
id := ""
cve := ""
```

Request body:

```text
recent=5&details=0
```

### Retrieve detailed information about recent vulnerabilities

```go
recent := "10"
details := "1"
id := ""
cve := ""
```

Request body:

```text
recent=10&details=1
```

Detailed requests may consume more API credits than basic requests.

### Retrieve a vulnerability by VulDB ID

```go
recent := "5"
details := "1"
id := "290848"
cve := ""
```

Request body:

```text
id=290848&details=1
```

### Search for a CVE identifier

The following example searches for CVE-2024-1234:

```go
recent := "5"
details := "1"
id := ""
cve := "CVE-2024-1234"
```

Request body:

```text
search=CVE-2024-1234&details=1
```

The CVE identifier is mapped to the API's `search` parameter rather than sent as a separate `cve` parameter.

## API Request and Authentication

The demonstration uses the following HTTP request configuration:

| Property | Value |
| --- | --- |
| Method | `POST` |
| Endpoint | `https://vuldb.com/api` |
| Authentication header | `X-VulDB-ApiKey` |
| Content-Type | `application/x-www-form-urlencoded` |
| Default response format | JSON |

For example, the default request is equivalent to:

```http
POST /api HTTP/1.1
Host: vuldb.com
X-VulDB-ApiKey: YOUR_PERSONAL_API_KEY
User-Agent: VulDB API Advanced Go Demo Agent
Content-Type: application/x-www-form-urlencoded
Content-Length: 18

recent=5&details=0
```

The [official API documentation](https://vuldb.com/kb/api) also documents `https://vuldb.com/?api` as the API endpoint and supports additional query types, filters and response formats beyond those implemented in this example.

## Response Handling

The program reads the complete HTTP response body. If the HTTP status is `200 OK`, it prints the raw body prefixed with `Response:`. For other HTTP statuses, it prints the status code. Request creation, network and response reading errors are reported separately.

A typical JSON API response contains the following top-level sections:

| Section | Description |
| --- | --- |
| `response` | API status, result count and credit usage information. |
| `request` | Information about the processed query. |
| `result` | Matching vulnerability records. |

The demonstration does not parse JSON, extract individual fields, save responses to files, or implement pagination. Refer to the [API documentation](https://vuldb.com/kb/api) for response schemas and supported data fields.

> [!NOTE]
> The API status in `response.status` can differ from the HTTP status. A production integration should inspect both rather than treating every HTTP 200 response as a successful API query.

## API Credits and Rate Limits

API access and available query sizes depend on the VulDB account and license.

- Basic vulnerability queries (`details=0`) generally consume one API credit per request.
- Detailed vulnerability queries (`details=1`) generally consume one credit per returned entry, with a minimum of one credit per request.
- Monitor `response.consumption` and `response.remaining` to track usage.
- Credits are replenished according to a rolling 24-hour window. Avoid aggressive polling or immediate retries after a limit is reached.
- VulDB recommends no more than 30 requests per minute; observe any additional restrictions applicable to your account.

See [API Credits](https://vuldb.com/kb/apicredits) and the [API documentation](https://vuldb.com/kb/api) for details.

## Troubleshooting

The following are common VulDB API status values. Except for HTTP-specific errors such as 429, these may appear inside the JSON response rather than as the actual HTTP status.

| Status | Meaning | Recommended action |
| --- | --- | --- |
| `200` | Results returned | Process the response. |
| `204` | No matching entries | Check the identifier or search criteria. |
| `400` | Invalid query data | Verify the request parameters. |
| `401` | Authentication failed | Verify your personal API key and account access. |
| `402` | Subscription insufficient | Check your license and enabled features. |
| `403` | API allowance exhausted | Check available credits and wait for replenishment or upgrade. |
| `405` | Unsupported query type | Check the API documentation. |
| `409` | Suspected API abuse | Stop sending requests and review the account restrictions. |
| `429` (HTTP) | Excessive request frequency | Respect the suspension period and reduce request frequency. |

If the program reports a network error, verify DNS resolution, Internet access, proxy configuration and the availability of HTTPS connections to VulDB. Recent eligible API requests can also be inspected in your VulDB user profile.

## Production Integration

The supplied code is intentionally minimal. Before using it in an automated or production environment, consider the following improvements:

- Load the API key from a protected environment variable or secret manager.
- Encode request parameters with `net/url.Values` rather than manually concatenating arbitrary input.
- Define an HTTP client timeout and use request contexts for cancellation.
- Parse responses with `encoding/json` and validate the API status separately from the HTTP status.
- Replace the deprecated `ioutil.ReadAll` with `io.ReadAll` when targeting Go 1.16 or later.
- Handle rate limits, credit exhaustion and transient network errors with appropriate backoff; avoid blindly retrying authentication or subscription errors.
- Implement pagination or incremental synchronization where appropriate, and avoid logging credentials or unrestricted response data.

### Optional: Read the API key from an environment variable

The existing example requires a small source-code change to support this configuration. Add `"os"` to the import block and replace the existing `personalAPIKey` declaration with:

```go
personalAPIKey := os.Getenv("VULDB_API_KEY")
if personalAPIKey == "" {
    fmt.Fprintln(os.Stderr, "VULDB_API_KEY is not set")
    os.Exit(1)
}
```

On Bash, enter the key without echoing it or putting the literal key in shell history:

```bash
read -r -s -p 'VulDB API key: ' VULDB_API_KEY
printf '\n'
export VULDB_API_KEY
go run VulDBApiDemoGo.go
unset VULDB_API_KEY
```

For CI/CD, inject `VULDB_API_KEY` through the platform's protected secrets mechanism.

## Contributing

Bug reports, documentation improvements and suggestions for additional Go API examples are welcome.

- Open an issue in the [GitHub issue tracker](https://github.com/vuldb/vuldb-api-golang-examples/issues).
- Fork the repository and submit a pull request with your changes.
- Format Go source code with `gofmt` and never include real API keys or confidential data in contributions.

## License

The source code in this repository is distributed under the [GNU General Public License v3.0](LICENSE).

Access to the VulDB service and use of data obtained through its API are subject to the applicable [VulDB Terms of Use](https://vuldb.com/kb/terms) and your subscription agreement. The open-source license of the example code does not grant an unrestricted license to redistribute VulDB data.

## Support and Resources

For questions about API access, credits, licensing or custom integrations, contact the VulDB support team at [support@vuldb.com](mailto:support@vuldb.com).

| Resource | Link |
| --- | --- |
| VulDB website | [vuldb.com](https://vuldb.com/) |
| API documentation | [vuldb.com/kb/api](https://vuldb.com/kb/api) |
| API credits | [vuldb.com/kb/apicredits](https://vuldb.com/kb/apicredits) |
| Support documentation | [vuldb.com/kb/support](https://vuldb.com/kb/support) |
| Official GitHub organization | [github.com/vuldb](https://github.com/vuldb) |
| Repository issues | [GitHub Issues](https://github.com/vuldb/vuldb-api-golang-examples/issues) |
| Go documentation | [go.dev/doc](https://go.dev/doc/) |

---

Maintained by [VulDB](https://vuldb.com/).
