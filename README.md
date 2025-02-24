# Estatespace Command Line Interface (CLI)

Welcome to the **Estatespace** command line interface. This tool provides a set of commands to streamline various tasks related to Git and Jira.

## Installation


To install the `es-cli` tool, follow these steps:

1. **Locate the Executable**: Find the `es` executable in the `prog/` directory.
2. **Create a Directory**: Create a new directory named `system_tools`.
3. **Add to PATH**:
    - Add the path of the `system_tools` directory to your system's PATH environment variable.
4. **Source Your Shell Configuration File**: Source your shell configuration file (e.g., `.zshrc`, `.bashrc`) to apply the changes.
5. **Move the Executable**: Move the `es` executable into the `system_tools` directory.
6. **Test Installation**: Verify that the installation was successful by running:
   ```bash
   es --help
   ```
   
Alternatively, you can clone this repo and build it locally with `go build -o [name it whatever]`

For more detailed installation instructions or troubleshooting, please refer to our [GitHub repository](https://github.com/nerdingout/es_cli).

## Usage

The `es` CLI tool is designed to be intuitive and easy to use. Here are some common commands you can run:

### Basic Commands

- **Help**: Get help on any command by appending `--help`.
  ```bash
  es --help
  ```

- **Completion Script**: Generate the autocompletion script for your shell.
  ```bash
  es completion [shell]
  ```

### Git Commands

- **gac (git add and commit)**: Adds all files in the current directory to the staging area and commits them with a message starting with the ES ticket number.
  ```bash
  es gac "Commit message"
  ```

- **gcl (git delete local branches matching ticket pattern)**: Deletes local git branches that match a specific Jira ticket pattern.
  ```bash
  es gcl <ticket-pattern>
  ```

### Jira Commands

- **jbl (Jira backlog for specific user)**: Lists the backlog items assigned to a specific user in Jira.
  ```bash
  es jbl -u <username>
  ```

- **jtp (Jira ticket for current branch)**: Opens the corresponding Jira ticket for the current git branch.
  ```bash
  es jtp
  ```

## Flags

The `es` CLI tool supports a few flags:

- `-h, --help`: Display help information about any command.
- `-t, --toggle`: Toggle a specific feature (details depend on implementation).

For more detailed usage and examples of each command, please refer to the [GitHub repository](https://github.com/geekcentric/es-cli).

## Contributing

We welcome contributions from the community. If you find an issue or have ideas for new features, feel free to open a pull request.

To contribute:

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).