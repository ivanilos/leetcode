type Person struct {
    name string
    height int
}

func sortPeople(names []string, heights []int) []string {
    people := []Person{}

    for i := 0; i < len(names); i++ {
        people = append(people, Person{names[i], heights[i]})
    }

    sort.Slice(people, func(a, b int) bool {
        return people[a].height > people[b].height
    })

    result := []string{}
    for i := 0; i < len(people); i++ {
        result = append(result, people[i].name)
    }
    return result
}
