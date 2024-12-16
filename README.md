# Git Automator

Git Automator is a simple command-line tool built with Go (Golang) that automates common Git tasks such as adding changes, committing them, and pushing to a remote repository. It is designed to simplify the process of managing Git repositories for developers.

## Features

- Check if Git is installed
- Add all changes (`git add .`)
- Commit changes with a custom message (`git commit -m "message"`)
- Push changes to a remote repository (`git push`)
- Automated options to Add, Commit, and Push in a single command
- Ability to pull the latest changes before pushing
- Option to switch branches

## Installation

1. Clone this repository to your local machine:
    ```bash
    git clone https://github.com/Kayleexx/git-automator.git
    ```

2. Navigate to the project directory:
    ```bash
    cd git-automator
    ```

3. Install Go (Golang) if you haven't already:
    - Download Go from the official website: [https://golang.org/dl/](https://golang.org/dl/)
    - Follow the installation instructions for your operating system.

4. Build the project:
    ```bash
    go build main.go
    ```

5. Run the tool:
    ```bash
    ./main
    ```

## Usage

After running the program, you will be prompted to choose an action:

1. **Add all changes**: Adds all the modified files to the staging area.
2. **Commit changes**: Allows you to commit the changes with a custom message.
3. **Push to remote**: Pushes the committed changes to the remote repository.
4. **Add, Commit, and Push (automated)**: Automates the entire process in one step.
5. **Switch branches**: Allows you to switch between branches.
6. **Pull latest changes before push**: Pulls the latest changes from the remote repository before pushing your changes.

## Requirements

- Go 1.16 or higher
- Git installed and configured on your system
- A Git repository initialized on your local machine

