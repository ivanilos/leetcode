func canVisitAllRooms(rooms [][]int) bool {
    visited := make([]bool, len(rooms))
    
    result := DFS(0, rooms, visited)

    return result == len(rooms)
}

func DFS(room int, keys [][]int, visited []bool) int {
    visited[room] = true

    totalVisited := 1
    for _, nextRoom := range keys[room] {
        if !visited[nextRoom] {
            totalVisited += DFS(nextRoom, keys, visited)
        }
    }

    return totalVisited
}
