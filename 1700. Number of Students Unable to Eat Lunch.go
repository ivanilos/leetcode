func countStudents(students []int, sandwiches []int) int {
    nextSandwich := 0
    notEatenInARow := 0
    for len(students) > 0 && notEatenInARow < len(students) {
        if students[0] == sandwiches[nextSandwich] {
            nextSandwich++
            notEatenInARow = 0
        } else {
            students = append(students, students[0])
            notEatenInARow++
        }
        students = students[1:len(students)]
    }
    return len(students)
    
}
