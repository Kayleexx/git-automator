package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Welcome to Git Automator!")

	// Check if Git is installed
	if _, err := exec.LookPath("git"); err != nil {
		fmt.Println("Error: Git is not installed or not in your PATH.")
		return
	}

	fmt.Println("Checking Git Status...")
	runGitCommand("status")

	// Display options
	fmt.Println("Choose an action:")
	fmt.Println("1. Add all changes")
	fmt.Println("2. Commit changes")
	fmt.Println("3. Push to remote")
	fmt.Println("4. Add, Commit, and Push (automated)")
	fmt.Println("5. Switch branches")
	fmt.Println("6. Pull latest changes before push")

	// Take user input for action
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter choice (1/2/3/4): ")
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "1":
		runGitCommand("add", ".")
	case "2":
		fmt.Print("Enter commit message: ")
		scanner.Scan()
		commitMessage := scanner.Text()
		runGitCommand("commit", "-m", commitMessage)
	case "3":
		runGitCommand("push")
	case "4":
		runGitCommand("add", ".")
		fmt.Print("Enter commit message (or press Enter to use default): ")
		scanner.Scan()
		commitMessage := scanner.Text()
		if commitMessage == "" {
			commitMessage = "Auto commit"
		}
		runGitCommand("commit", "-m", commitMessage)
		runGitCommand("push")
	case "5":
		fmt.Println("Fetching branches...")
		runGitCommand("branch")
		fmt.Print("Enter the branch to switch to: ")
		scanner.Scan()
		branchName := scanner.Text()
		runGitCommand("checkout", branchName)
	case "6":
		runGitCommand("pull")
		runGitCommand("push")
	default:
		fmt.Println("Invalid choice. Exiting.")
	}
}

func runGitCommand(args ...string) {
	// Dynamically execute the Git command
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Print the Git command being executed
	fmt.Printf("Running: git %s\n", strings.Join(args, " "))

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Command executed successfully!")
}
