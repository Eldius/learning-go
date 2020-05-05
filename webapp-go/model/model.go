/*
Package model is where entities and TOs live...
*/
package model

import (
	"time"
	"sort"
)

/*
Employee it's an employee
*/
type Employee struct {
	ID int64
	Name string
	Role string
	SalaryHistory []Salary
	Active bool
}

/*
Salary is the salary representation
*/
type Salary struct {
	Since time.Time
	Until time.Time
	Value float64
}
/*
EmployeeTO is the TO object to transfer
*/
type EmployeeTO struct {
	ID int64
	Name string
	Role string
	SalaryHistory []Salary
	LastSalary Salary
	Active bool
}

/*
ToTO parses entity to TO
*/
func (e *Employee)ToTO() EmployeeTO {
	return EmployeeTO{
		ID: e.ID,
		Name: e.Name,
		Role: e.Role,
		SalaryHistory: e.SalaryHistory,
		LastSalary: e.GetLastSalary(),
		Active: e.Active,
	}
}

/*
SortSalaryBy sorts servers by
*/
type SortSalaryBy func(s1, s2 *Salary) bool

// serverSorter joins a By function and a slice of servers to be sorted.
type salarySorter struct {
	salaryHistory []Salary
	by      func(s1, s2 *Salary) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *salarySorter) Len() int {
	return len(s.salaryHistory)
}

// Swap is part of sort.Interface.
func (s *salarySorter) Swap(i, j int) {
	s.salaryHistory[i], s.salaryHistory[j] = s.salaryHistory[j], s.salaryHistory[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *salarySorter) Less(i, j int) bool {
	return s.by(&s.salaryHistory[i], &s.salaryHistory[j])
}

/*
SortSalary sorts server
*/
func (by SortSalaryBy) SortSalary(salaryHistory []Salary) {
	ss := &salarySorter{
		salaryHistory: salaryHistory,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ss)
}

/*
GetLastSalary returns the last salary from history
*/
func (e *Employee)GetLastSalary() Salary {
	var sortBySinceDate = func(s1, s2 *Salary) bool {
		return s1.Since.IsZero() || !s2.Since.IsZero() || (s1.Since.Unix() > s2.Since.Unix())
	}
	SortSalaryBy(sortBySinceDate).SortSalary(e.SalaryHistory)
	return e.SalaryHistory[0]
}

/*
ListEmployees returns a list of employees
*/
func ListEmployees() []Employee {
	employees := make([]Employee,0)

	employees = append(employees, Employee{
		ID: 1,
		Name: "Fulano de Tal",
		Role: "Developer",
		Active: false,
		SalaryHistory: []Salary{
			Salary{
				Since: time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC),
				Until: time.Date(2018, time.August, 15, 0, 0, 0, 0, time.UTC),
				Value: 1000.0,
			},
		},
	})

	employees = append(employees, Employee{
		ID: 2,
		Name: "Ciclano da Silva",
		Role: "Developer",
		Active: true,
		SalaryHistory: []Salary{
			Salary{
				Since: time.Date(2017, time.January, 15, 0, 0, 0, 0, time.UTC),
				Until: time.Date(2018, time.August, 15, 0, 0, 0, 0, time.UTC),
				Value: 1000.0,
			},
			Salary{
				Since: time.Date(2018, time.August, 15, 0, 0, 0, 0, time.UTC),
				Until: time.Date(2019, time.August, 15, 0, 0, 0, 0, time.UTC),
				Value: 2000.0,
			},
			Salary{
				Since: time.Date(2019, time.August, 15, 0, 0, 0, 0, time.UTC),
				Value: 3000.0,
			},
		},
	})

	return employees
}
