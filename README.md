# aksh (Shell version 1.0.0)

A custom shell implementation written in Go that provides basic command-line functionality with command history support.

## Features

- **Built-in Commands**: `exit`, `echo`, `type`, `cd`, `history`
- **External Command Execution**: Run any executable found in your PATH
- **Command History**: Persistent history stored in `~/.aksh_history`
- **Path Resolution**: Automatic binary lookup in system PATH

## Prerequisites

- Go 1.16 or higher
- Unix-like operating system (Linux, macOS, BSD)

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/jhaksh-24/aksh.git
cd aksh
```

### 2. Build the Shell

```bash
go build -o aksh
```

This creates an executable named `aksh` in the current directory.

### 3. (Optional) Install Globally

To use the shell from anywhere on your system:

```bash
# Move to a directory in your PATH
sudo mv aksh /usr/local/bin/

# Or add to your personal bin directory
mkdir -p ~/bin
mv aksh ~/bin/
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc  # or ~/.zshrc
source ~/.bashrc
```

## Usage

### Starting the Shell

```bash
./aksh
```

Or if installed globally:

```bash
aksh
```

### Built-in Commands

#### `cd` - Change Directory
```bash
cd /path/to/directory    # Change to specified directory
cd ~                      # Change to home directory
cd                        # Change to home directory (no args)
```

#### `echo` - Print Text
```bash
echo Hello World          # Outputs: Hello World
echo "Multiple words"     # Outputs: Multiple words
```

#### `type` - Show Command Type
```bash
type cd                   # Outputs: cd is a shell builtin
type ls                   # Outputs: ls is /usr/bin/ls
type nonexistent          # Outputs: nonexistent: not found
```

#### `exit` - Exit the Shell
```bash
exit 0                    # Exit with status code 0
exit                      # Exit with default status
```

### External Commands

Any executable in your PATH can be run:

```bash
ls -la
cat file.txt
grep "pattern" file.txt
python3 script.py
```

### Command History

All executed commands are automatically saved to `~/.aksh_history` and persist across sessions.

## Project Structure

```
aksh/
├── builtincmd/
│   ├── ChangeDirectory.go    # cd command implementation
│   ├── Echo.go                # echo command implementation
│   └── Type.go                # type command implementation
├── externalcmd/
│   └── Execute.go             # External command execution
├── history/
│   └── History.go             # Command history management
├── utils/
│   ├── InvalidCmd.go          # Error handling utilities
│   └── SearchBinInPath.go     # PATH search functionality
├── go.mod
└── main.go                    # Entry point
```

## Configuration

### History File Location

Command history is stored in:
```
~/.aksh_history
```

You can view or edit this file manually with any text editor.

## Uninstallation

### Remove the Binary

If installed globally:
```bash
sudo rm /usr/local/bin/aksh
# or
rm ~/bin/aksh
```

If only built locally:
```bash
cd /path/to/shell
rm aksh
```

### Remove History File (Optional)

```bash
rm ~/.aksh_history
```

### Remove Source Code

```bash
cd ..
rm -rf aksh/
```

## Troubleshooting

### Command Not Found

If external commands aren't working:
```bash
# Check your PATH
echo $PATH

# Verify the command exists
which ls
type ls
```

### Permission Denied

If you get permission errors:
```bash
# Make the binary executable
chmod +x aksh

# For global installation
sudo mv aksh /usr/local/bin/
```

### History File Issues

If history isn't saving:
```bash
# Check if file exists and is writable
ls -la ~/.aksh_history

# Recreate if needed
rm ~/.aksh_history
# Restart the shell
```

## Development

### Building from Source

```bash
go build -o aksh
```

### Running Tests (Will be added soon)

```bash
go test ./...
```

### Adding Dependencies

```bash
go get <package-name>
go mod tidy
```

## Limitations

- No command pipelines (`|`) support
- No input/output redirection (`>`, `<`)
- No background processes (`&`)
- No tab completion
- No command editing (arrow keys)
- Limited to Unix-like systems

## Contributing

Contributions are welcome! Feel free to:
- Report bugs
- Suggest features
- Submit pull requests

## License

This project is open source and available under the MIT License.

## Author

[@jhaksh-24](https://github.com/jhaksh-24)

## Acknowledgments

Built as a learning project to understand:
- Shell implementation fundamentals
- Go programming language
- Operating system interactions
- Process management
