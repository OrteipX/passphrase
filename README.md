# Password Generator

A simple command-line tool written in Go for generating random passwords with customizable character sets.

## Features

- Generate passwords of any specified length.
- Customize your password to include:
  - Uppercase letters (`A-Z`)
  - Lowercase letters (`a-z`)
  - Numbers (`0-9`)
  - Special characters (`!#$%&()*+-.:;=?@[]^{}`
- Quick option to include a mix of all character types.
- Easy-to-use command-line interface.

## Installation

### Prerequisites

- Go 1.13 or later installed on your system. If you don't have Go installed, you can download it from the [official website](https://golang.org/dl/).

### Build from Source

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/password-generator.git
   cd password-generator
   ```
   
2. **Build the application**:
   ```bash
   go build -o passgen
   ```
   
3. **Move the executable to your PATH (optional)**:
   - Unix/Linux/macOS
   ```bash
   sudo mv passgen /usr/local/bin/
   ```

### Usage

```bash
passgen -l [length] [options]
```
- Options
  ```bash
  -l: (Required) Specify the length of the password (e.g., -l 16).
  -a: Include both uppercase and lowercase alphabetic characters (A-Za-z).
  -au: Include uppercase alphabetic characters only (A-Z).
  -al: Include lowercase alphabetic characters only (a-z).
  -n: Include numeric characters (0-9).
  -s: Include special characters (!#$%&()*+-.:;=?@[]^{}).
  -x: Include a mix of all character types (equivalent to -au -al -n -s).
  ```

### Examples

1. Generate a 12-character password with uppercase letters, lowercase letters, and numbers:
```bash
passgen -l 12 -a -n
```

2. Generate an 8-character password with lowercase letters and special characters:
```bash
passgen -l 12 -a -n
```

3. Generate a 16-character password with a mix of all character types:
```bash
passgen -l 16 -x
```

### Help
```bash
passgen -h
```
