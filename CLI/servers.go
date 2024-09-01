package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Create and run the Vanilla server
func vanilla() {
	// Set the eula to true
	fmt.Println("Setting \"eula=true\"")
	eulaTrue := []byte("eula=true")
	err := os.WriteFile("eula.txt", eulaTrue, 0644)
	if err != nil {
		panic(err)
	}

	// Run the minecraft server
	fmt.Println("Generating Files....")
	cmd := exec.Command("java", "-Xmx1024M", "-Xms1024M", "-jar", "server.jar", "nogui")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
}

func fabric() {
	// Set the eula to true
	fmt.Println("Setting \"eula=true\"")
	eulaTrue := []byte("eula=true")
	err := os.WriteFile("eula.txt", eulaTrue, 0644)
	if err != nil {
		panic(err)
	}

	// Install and start the server
	fmt.Println("Generating Files....")
	cmd := exec.Command("java", "-jar", "server.jar")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

}

func forge() {
	// Set the eula to true
	fmt.Println("Setting \"eula=true\"")
	eulaTrue := []byte("eula=true")
	err := os.WriteFile("eula.txt", eulaTrue, 0644)
	if err != nil {
		panic(err)
	}

	// Install the server
	fmt.Println("Generating Files....")
	install := exec.Command("java", "-jar", "server.jar", "--installServer")

	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	install.Stdin = os.Stdin
	err = install.Start()
	if err != nil {
		panic(err)
	}

	// Wait for the installer to finish before trying to run the server
	err = install.Wait()
	if err != nil {
		panic(err)
	}

	// Run the server
	fmt.Println("Starting server....")
	// The forge installer automatically
	cmd := exec.Command("sh", "run.sh")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

}
