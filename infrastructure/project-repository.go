package infrastructure

import (
	"errors"

	"github.com/lumos-industry/domain"
)

// In-memory repository untuk menyimpan project
type InMemoryProjectRepo struct {
	projects []domain.Project
}

// DeleteProject implements domain.ProjectRepository.
func (repo *InMemoryProjectRepo) DeleteProject(id int) error {
	for i, project := range repo.projects {
		if project.ID == id {
			repo.projects = append(repo.projects[:i], repo.projects[i+1:]...)
			return nil
		}
	}
	return errors.New("project tidak ditemukan")
}

// FindProjectByID implements domain.ProjectRepository.
func (repo *InMemoryProjectRepo) FindProjectByID(id int) (domain.Project, error) {
	for _, project := range repo.projects {
		if project.ID == id {
			return project, nil
		}
	}
	return domain.Project{}, errors.New("project tidak ditemukan")
}

// UpdateProject implements domain.ProjectRepository.
func (repo *InMemoryProjectRepo) UpdateProject(updatedProject domain.Project) error {
	for i, project := range repo.projects {
		if project.ID == updatedProject.ID {
			repo.projects[i] = updatedProject
			return nil
		}
	}
	return errors.New("project tidak ditemukan")
}

func (repo *InMemoryProjectRepo) SaveProject(project domain.Project) error {
	// Cek apakah ID sudah ada
	for _, p := range repo.projects {
		if p.ID == project.ID {
			return errors.New("ID project sudah ada, gunakan ID yang berbeda")
		}
	}

	// Menyimpan project kedalam list
	repo.projects = append(repo.projects, project)
	return nil
}

func (repo *InMemoryProjectRepo) GetAllProjects() ([]domain.Project, error) {
	// Mengembalikan semua project
	if len(repo.projects) == 0 {
		return nil, errors.New("project tidak ditemukan")
	}
	return repo.projects, nil
}

func NewInMemoryProjectRepo() *InMemoryProjectRepo {
	return &InMemoryProjectRepo{
		projects: []domain.Project{},
	}
}
