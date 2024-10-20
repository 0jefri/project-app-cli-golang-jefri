package domain

// struck untuk project atau properti
type Project struct {
	ID         int
	Nama       string
	Panjang    float64
	Lebar      float64
	HektarArea float64
	Tanaman    string
	Bibit      int
}

// membuat method atau function yang dibutuhkan
type ProjectRepository interface {
	SaveProject(project Project) error
	GetAllProjects() ([]Project, error)
	UpdateProject(project Project) error
	DeleteProject(id int) error
	FindProjectByID(id int) (Project, error)
}
