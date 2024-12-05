# sanktldex
# Domain and Subdomain Extraction Tool

This Go tool allows you to extract base domains or subdomains from a list of URLs. It supports single-domain extraction using the `-f` flag or multiple domains using the `-fL` flag, which takes a file containing domains.

## Features

- Extract base domains or subdomains.
- Support for concurrent processing with `-t` flag for controlling the number of threads (default is 8).
- Extract subdomains only for specific domains (single domain or a list of domains).
- Input can be provided via standard input (stdin) or from a file.

## Flags

- `-t <num>`: Number of threads to utilize (default is 8).
- `-s`: Dump subdomains instead of base domains.
- `-f <domain>`: Extract subdomains for a single domain (e.g., `example.com`).
- `-fL <file>`: Extract subdomains for domains listed in a file. Each line in the file should contain one domain.

## Example Usage

### Extract subdomains for a single domain:

```bash
cat urls.txt | sanktldex -f example.com -s
```

## Example Usage From github :
```
git clone https://github.com/5hank4r/sanktldex.git
cd sanktldex
go build sanktldex.go
sudo mv sanktldex /usr/local/bin/
```
