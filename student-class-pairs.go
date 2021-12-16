package main

import "fmt"

func main() {
	student_course_pairs1 := [][]string{
		[]string{"58","Linear Algebra"},
		[]string{"94","Art History"},
		[]string{"94","Operating Systems"},
		[]string{"17","Software Design"},
		[]string{"58","Mechanics"},
		[]string{"58","Economics"},
		[]string{"17","Linear Algebra"},
		[]string{"17","Political Science"},
		[]string{"94","Economics"},
		[]string{"25","Economics"},
		[]string{"58","Software Design"},
	}

	fmt.Println(findPairs(student_course_pairs1))

}

func findPairs(studentCoursePairs [][]string) []string {
	ans := []string{}
	m := make(map[string]map[string]bool) // key: students val: classes
	students := []string{}

	// create the map and the unique list of students
	for i := 0; i < len(studentCoursePairs); i++ {
		student := studentCoursePairs[i][0]
		class := studentCoursePairs[i][1]

		_, present := m[student]
		if present {
			classes := m[student]
			classes[class] = true

			m[student] = classes
		} else {
			m[student] = map[string]bool{class:true}
			students = append(students, student)
		}
	}

	fmt.Println(students)
	fmt.Println(m)

	// check who shares the same classes
	for i := 0; i < len(students)-1; i++ {
		student := students[i]
		classes := m[student]

		for j := i+1; j < len(students); j++ {
			student2 := students[j]
			classes2 := m[student2]

			pairStr := student + ", " + student2 + " : "
			ans = append(ans, pairStr)

			for key, _ := range classes {
				_, present := classes2[key]
				if present {
					ans[len(ans)-1] += key + " "
				}
			}

		}
	}
	return ans
}