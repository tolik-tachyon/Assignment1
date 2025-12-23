package Employee

import "fmt"

type Employee interface {
	CalculateSalary() float64
	CalculateBonus() float64
	GetWorkHours() int
}

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
}

func (f FullTime) CalculateSalary() float64 {
	return f.MonthlySalary
}

func (f FullTime) CalculateBonus() float64 {
	return f.MonthlySalary * f.BonusRate
}

func (f FullTime) GetWorkHours() int {
	return 160
}

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}

func (p PartTime) CalculateSalary() float64 {
	return p.HourlyRate * float64(p.HoursWorked)
}

func (p PartTime) CalculateBonus() float64 {
	if p.HoursWorked > 100 {
		return 100
	}
	return 0
}

func (p PartTime) GetWorkHours() int {
	return p.HoursWorked
}

type Contractor struct {
	ProjectRate       float64
	ProjectsCompleted int
}

func (c Contractor) CalculateSalary() float64 {
	return c.ProjectRate * float64(c.ProjectsCompleted)
}

func (c Contractor) CalculateBonus() float64 {
	return 0
}

func (c Contractor) GetWorkHours() int {
	return c.ProjectsCompleted * 40
}

type Intern struct {
	DailyRate  float64
	DaysWorked int
}

func (i Intern) CalculateSalary() float64 {
	return i.DailyRate * float64(i.DaysWorked)
}

func (i Intern) CalculateBonus() float64 {
	return 50
}

func (i Intern) GetWorkHours() int {
	return i.DaysWorked * 8
}

func DemoEmployee() {
	fmt.Println("\n--- EMPLOYEE DEMO ---")

	employees := []Employee{
		FullTime{3000, 0.1},
		PartTime{10, 120},
		Contractor{500, 3},
		Intern{20, 15},
	}

	for _, e := range employees {
		fmt.Println("Salary:", e.CalculateSalary(),
			"Bonus:", e.CalculateBonus(),
			"Hours:", e.GetWorkHours())
	}
}
