package main

import (
	"fmt"
	"os"

	"github.com/lumos-industry/controller"
	"github.com/lumos-industry/infrastructure"
	"github.com/lumos-industry/usecase"
)

func main() {
	// Inisialisasi repository, usecase, dan controller
	projectRepo := infrastructure.NewInMemoryProjectRepo()
	projectUsecase := usecase.NewProjectUsecase(projectRepo)
	projectController := controller.NewProjectController(projectUsecase)

	for {
		fmt.Println("\nPT. Lumos Industry")
		fmt.Println("1. Create Project")
		fmt.Println("2. Show Projects")
		fmt.Println("3. Update Project")
		fmt.Println("4. Delete Project")
		fmt.Println("5. Show Updated Projects")
		fmt.Println("6. Quit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			projectController.CreateProject()
		case 2:
			projectController.ShowProjects()
		case 3:
			projectController.Update()
		case 4:
			projectController.DeleteProject()
		case 5:
			projectController.ShowUpdatedProjects()
		case 6:
			fmt.Println("Exiting the program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}
