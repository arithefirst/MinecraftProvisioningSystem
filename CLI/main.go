package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {

	// Check if the command was formatted correctly
	if len(os.Args) != 3 {
		fmt.Println("Usage: MCPS [platform] [version]")
		fmt.Println("EX: MCPS vanilla 1.12.2")
		return
	}

	// Check if the platform is valid (Is Forge, Fabric, or Vanilla)
	if strings.ToLower(os.Args[1]) != "vanilla" && strings.ToLower(os.Args[1]) != "forge" && strings.ToLower(os.Args[1]) != "fabric" {
		panic("\"" + os.Args[1] + "\"" + " is not a valid platform.\nPlease choose from Vanilla, Forge, or Fabric.")
	}

	validVersions := []string{"1.12", "1.12.1", "1.12.2", "1.13", "1.13.1", "1.14",
		"1.14.1", "1.14.2", "1.14.3", "1.14.4", "1.15", "1.15.1", "1.15.2", "1.16",
		"1.16.1", "1.16.2", "1.16.3", "1.16.4", "1.16.5", "1.17", "1.17.1", "1.18",
		"1.18.1", "1.18.2", "1.19", "1.19.1", "1.19.2", "1.19.3", "1.19.4", "1.20",
		"1.20.1", "1.20.2", "1.20.3", "1.20.4", "1.20.5", "1.20.6", "1.21", "1.21.1"}

	if !slices.Contains(validVersions, os.Args[2]) {
		panic("\"" + os.Args[2] + "\"" + " is not a valid Minecraft version.\nAll Minecraft versions above 1.12 are supported.\n(See https://minecraft.fandom.com/wiki/Java_Edition_version_history for more)")
	}

	// Download server.jar for requested version
	fmt.Printf("Downloading server.jar for %v %v\n", os.Args[1], os.Args[2])
	retriveServerJar(os.Args[1], os.Args[2])

	// Run the different setup procedures depending on the server type
	if strings.ToLower(os.Args[1]) == "vanilla" {
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

	} else if strings.ToLower(os.Args[1]) == "fabric" {
		invalidFabric := []string{"1.12", "1.12.1", "1.12.2", "1.13", "1.13.1"}
		if slices.Contains(invalidFabric, os.Args[2]) {
			panic("Err: Fabric only supports 1.14+")
		} else {
			// Fabric code here
		}

	} else if strings.ToLower(os.Args[1]) == "forge" {
		// Set the eula to true
		fmt.Println("Setting \"eula=true\"")
		eulaTrue := []byte("eula=true")
		err := os.WriteFile("eula.txt", eulaTrue, 0644)
		if err != nil {
			panic(err)
		}

		// Run the minecraft server
		fmt.Println("Generating Files....")
		cmd := exec.Command("java", "-jar", "server.jar", "--installServer")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err = cmd.Start()
		if err != nil {
			panic(err)
		}
	}

}
