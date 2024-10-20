package usecase

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/lumos-industry/domain"
)

// membuat struck dimana struct ini menyimpan struct yang ada di domain
type ProjectUsecase struct {
	ProjectRepo domain.ProjectRepository
}

// membuat fungsi untuk create atau buat project.
func (p *ProjectUsecase) CreateProject(id int, nama string, kecepatan, panjangWaktu, lebarWaktu float64) (domain.Project, error) {

	// lakukan pengecekan, takutnya nanti ada yang input 0
	if nama == "" || kecepatan <= 0 || panjangWaktu <= 0 || lebarWaktu <= 0 {
		return domain.Project{}, errors.New("invalid input values")
	}

	// Cek apakah ID sudah ada
	exists, err := p.IsIDExists(id)
	if err != nil {
		fmt.Println("Error checking ID existence:", err)
		return domain.Project{}, err
	}
	if exists {
		fmt.Println("ID project sudah ada:", id)
		return domain.Project{}, errors.New("ID project sudah ada, gunakan ID yang berbeda")
	}

	// menghitung panjang, lebar , dan luas dalam bentuk hektar
	panjang := kecepatan * panjangWaktu
	lebar := kecepatan * lebarWaktu
	area := panjang * lebar
	hektar := area / 10000

	//ini untuk mengecek tipe data pakai reflect
	println("Mengecek Type Data:")
	println("Type of panjang:", reflect.TypeOf(panjang).String())
	println("Type of lebar:", reflect.TypeOf(lebar).String())
	println("Type of area dalam hektar:", reflect.TypeOf(hektar).String())

	// membuat project baru
	project := domain.Project{
		ID:         id,
		Nama:       nama,
		Panjang:    panjang,
		Lebar:      lebar,
		HektarArea: hektar,
	}

	//ini untuk menyimpan project
	err = p.ProjectRepo.SaveProject(project)
	if err != nil {
		return domain.Project{}, err
	}
	return project, nil
}

// Tambahkan fungsi ShowProjects
func (p *ProjectUsecase) ShowProjects() ([]domain.Project, error) {
	return p.ProjectRepo.GetAllProjects()
}

// fungsi untuk update
func (p *ProjectUsecase) UpdateProject(id int, tanaman string) error {
	//cari project berdasarkan id, ini yg pertaman dilakukan untuk update
	project, err := p.ProjectRepo.FindProjectByID(id)
	if err != nil {
		return errors.New("project tidak ditemukan")
	}

	var bibit int
	if tanaman == "Kelapa Sawit" {
		bibit = int(project.HektarArea * 10000 / 25) // ini untuk jarak 5m atau area perbibit = 5m^2
	} else if tanaman == "Pohon Akasia" {
		bibit = int(project.HektarArea * 10000 / 9)
	} else {
		return errors.New("salah memasukan pilihan")
	}

	// Updatae informasi tanaman dan bibit
	project.Tanaman = tanaman
	project.Bibit = bibit

	//menyimpan update
	err = p.ProjectRepo.UpdateProject(project)
	if err != nil {
		return err
	}

	fmt.Printf("Project '%s' di update dengan tanaman: %s, bibit yang dibutuhkan sebanyak: %d\n", project.Nama, tanaman, bibit)
	return nil
}

// fungsi untuk menemukan project by id
func (p *ProjectUsecase) FindProjectByID(id int) (domain.Project, error) {
	return p.ProjectRepo.FindProjectByID(id)
}

func (p *ProjectUsecase) DeleteProject(id int) error {
	err := p.ProjectRepo.DeleteProject(id)
	if err != nil {
		return err
	}
	fmt.Println("Sukses Menghapus Project.")
	return nil
}

// fungsi untuk mendapatkan input dari terminal
func GetInput(prompt string) (float64, error) {
	var input string
	print(prompt)
	fmt.Scanln(&input)

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("invalid input, Please Input Number")
	}

	return value, nil
}

// fungsi untuk memeriksa apakah ID sudah ada
func (p *ProjectUsecase) IsIDExists(id int) (bool, error) {
	_, err := p.ProjectRepo.FindProjectByID(id)
	if err != nil {
		if err.Error() == "project tidak ditemukan" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// disini saya membuat constructor fungsi, berguna untuk instance baru dari ProjectUsecase denga repo yang di inject ke dalamnya.
func NewProjectUsecase(repo domain.ProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{
		ProjectRepo: repo,
	}
}
