package main

import "fmt"

type Job struct {
	name                string
	annualSalaryInEuros uint64
}

type Person struct {
	parents []*Person
	name    string
	age     int8
	job     *Job
}

// define a method for the person struct
func (person Person) greet() {
	employment := "I am unemployed."

	if person.job != nil {
		employment = fmt.Sprintf("I am a %s and my annual salary is %d euros.", person.job.name, person.job.annualSalaryInEuros)
	}
	fmt.Printf("Hi, my name is %s and I'm %d years old. %s\n", person.name, person.age, employment)
}

// if we want to mutate an entity, use a pointer instead
func (person *Person) newJob(job Job) {
	person.job = &job
}

func doStuffWithPerson() {
	john := Person{
		name: "John",
		age:  26,
	}

	sally := Person{
		name: "Sally",
		age:  28,
	}

	bob := Person{
		parents: []*Person{&john, &sally},
		name:    "Bob",
		age:     10,
		job: &Job{
			name:                "Software Engineer",
			annualSalaryInEuros: 140000,
		},
	}

	bob.greet()
	john.newJob(Job{
		name:                "Data Scientist",
		annualSalaryInEuros: 200000,
	})
	john.greet()
}
