package main

import "fmt"

type User struct {
	Name        string
	Email       string
	Designation string
}

type Story struct {
	Summary     string
	Description string
	Priority    string
	Assignee    User
	Reporter    User
}

type Project struct {
	Name    string
	Owner   User
	Stories []Story
}

func (p *Project) Clone() Project {
	tempProject := new(Project)
	tempProject.Name = p.Name
	tempProject.Owner = p.Owner
	tempProject.Stories = p.Stories

	return *tempProject
}

func main() {
	projectOne := createSampleProject()
	fmt.Println(projectOne)

	owner := User{
		Name:        "john",
		Email:       "john@test.com",
		Designation: "Program Manager",
	}
	projectTwo := projectOne.Clone()
	projectTwo.Owner = owner

	fmt.Println(projectTwo)
}

func createSampleProject() Project {
	owner := User{
		Name:        "dev",
		Email:       "dev@test.com",
		Designation: "Program Manager",
	}

	developerOne := User{
		Name:        "ram",
		Email:       "ram@test.com",
		Designation: "Backend Developer",
	}

	developerTwo := User{
		Name:        "raj",
		Email:       "raj@test.com",
		Designation: "Frontend Developer",
	}

	qa := User{
		Name:        "alex",
		Email:       "alex@test.com",
		Designation: "QA Engineer",
	}

	storyOne := Story{
		Summary:     "Product List API",
		Description: "List all active products",
		Priority:    "major",
		Assignee:    developerOne,
		Reporter:    qa,
	}

	storyTwo := Story{
		Summary:     "Product List Page",
		Description: "Design a responsive product list page",
		Priority:    "minor",
		Assignee:    developerTwo,
		Reporter:    qa,
	}

	projectOne := Project{
		Name:    "Project A",
		Owner:   owner,
		Stories: []Story{storyOne, storyTwo},
	}

	return projectOne
}
