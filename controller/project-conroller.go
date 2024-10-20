package controller

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lumos-industry/usecase"
)

// ini buat struct dimana struck berisi usecase yang dibuat sebelumya dan kita kasi pointer
type ProjectController struct {
	ProjectUsecase *usecase.ProjectUsecase
}

// fungsi untuk buat projectnya
func (c *ProjectController) CreateProject() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Project: ")
	var id int
	fmt.Scan(&id)

	fmt.Print("Masukkan Nama Project: ")
	projectName, _ := reader.ReadString('\n')
	projectName = projectName[:len(projectName)-1]

	// Mendapatkan input kecepatan dari drone, panjang, dan lebar
	kecepatan, _ := usecase.GetInput("Masukan Kecepatan Drone (m/s): ")
	panjangWaktu, _ := usecase.GetInput("Masukan waktu terbang untuk mengukur panjang(seconds): ")
	lebarWaktu, _ := usecase.GetInput("Masukan waktu terbang untuk mengukur lebar(seconds): ")

	// Buat project melalui usecase
	project, err := c.ProjectUsecase.CreateProject(id, projectName, kecepatan, panjangWaktu, lebarWaktu)
	if err != nil {
		fmt.Println("Terjadi kesalahan membuat project", err)
		return
	}

	fmt.Printf("Selamat, Project '%s' sukses dibuat!\n", project.Nama)
	fmt.Printf("Ukuran lahan: %.2f meters x %.2f meters (%.2f hectares)\n", project.Panjang, project.Lebar, project.HektarArea)
}

// fungsi untuk menampilkan semua project yang sudah dibuat
func (c *ProjectController) ShowProjects() {
	projects, err := c.ProjectUsecase.ShowProjects()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(projects) == 0 {
		fmt.Println("Tidak ada proyek yang tersedia.")
		return
	}

	fmt.Printf("%-5s %-20s %-10s %-10s %-10s\n", "ID", "Nama Project", "Panjang (m)", "Lebar (m)", "Area (hectar)")
	for _, project := range projects {
		fmt.Printf("%-5d %-20s %-10.2f %-10.2f %-10.2f\n", project.ID, project.Nama, project.Panjang, project.Lebar, project.HektarArea)
	}
}

func (c *ProjectController) Update() {
	var projectID int
	fmt.Print("Masukan ID project untuk di edit: ")
	fmt.Scanln(&projectID)

	project, err := c.ProjectUsecase.FindProjectByID(projectID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Edit Project:", project.Nama)
	fmt.Printf("Tanaman : %s, Bibit: %d\n", project.Tanaman, project.Bibit)

	// Input pilihan jenis tanaman
	fmt.Println("Pilih Tanaman: ")
	fmt.Println("1. Kelapa Sawit")
	fmt.Println("2. Pohon Akasia")
	fmt.Print("Masukan Pilihan: ")

	var pilihTanaman int
	fmt.Scanln(&pilihTanaman)

	var tanaman string
	switch pilihTanaman {
	case 1:
		tanaman = "Kelapa Sawit"
	case 2:
		tanaman = "Pohon Akasia"
	default:
		fmt.Println("Gagal memilih.")
		return
	}

	err = c.ProjectUsecase.UpdateProject(projectID, tanaman)
	if err != nil {
		fmt.Println("Error updating project:", err)
	}
}

func (c *ProjectController) ShowUpdatedProjects() {
	projects, err := c.ProjectUsecase.ShowProjects()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(projects) == 0 {
		fmt.Println("Project tidak tersedia.")
		return
	}

	// Cek apakah ada proyek yang sudah di-update
	hasUpdatedProjects := false
	fmt.Printf("%-5s %-20s %-10s %-10s %-10s %-15s %-10s\n", "ID", "Nama Project", "Panjang (m)", "Lebar (m)", "Area (hectar)", "Tanaman", "Bibit")
	for _, project := range projects {
		if project.Tanaman != "" {
			hasUpdatedProjects = true
			fmt.Printf("%-5d %-20s %-10.2f %-10.2f %-10.2f %-15s %-10d\n", project.ID, project.Nama, project.Panjang, project.Lebar, project.HektarArea, project.Tanaman, project.Bibit)
		}
	}

	if !hasUpdatedProjects {
		fmt.Println("Belum ada proyek yang di-update.")
	}
}

func (p *ProjectController) DeleteProject() {
	var projectID int
	fmt.Print("Enter project ID to delete: ")
	fmt.Scanln(&projectID)

	err := p.ProjectUsecase.DeleteProject(projectID)
	if err != nil {
		fmt.Println("Gagal Menghapus Project:", err)
	}
}

func NewProjectController(p *usecase.ProjectUsecase) *ProjectController {
	return &ProjectController{
		ProjectUsecase: p,
	}
}
