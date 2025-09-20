type Packet struct {
    source int
    destination int
    timestamp int
}

type Router struct {
    memoryLimit int
    packets []Packet
    packetsMap map[Packet]bool
    destinationToTimestampMap map[int][]int
}


func Constructor(memoryLimit int) Router {
    return Router {
        memoryLimit,
        []Packet{},
        map[Packet]bool{},
        map[int][]int{},
    }
}


func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
    newPacket := Packet{source, destination, timestamp}

    if _, ok := this.packetsMap[newPacket]; ok {
        return false
    }

    if len(this.packets) >= this.memoryLimit {
        droppedPacket := this.packets[0]
        this.packets = this.packets[1:]

        delete(this.packetsMap, droppedPacket)

        this.destinationToTimestampMap[droppedPacket.destination] = this.destinationToTimestampMap[droppedPacket.destination][1:]
    }

    this.packets = append(this.packets, newPacket)
    this.packetsMap[newPacket] = true
    this.destinationToTimestampMap[newPacket.destination] = append(this.destinationToTimestampMap[newPacket.destination], newPacket.timestamp)

    return true
}


func (this *Router) ForwardPacket() []int {
    if len(this.packets) == 0 {
        return []int{}
    }

    forwarded := this.packets[0]
    this.packets = this.packets[1:]

    delete(this.packetsMap, forwarded)

    this.destinationToTimestampMap[forwarded.destination] = this.destinationToTimestampMap[forwarded.destination][1:]

    return []int{forwarded.source, forwarded.destination, forwarded.timestamp}
}


func (this *Router) GetCount(destination int, startTime int, endTime int) int {
    return countLowerEqual(this.destinationToTimestampMap[destination], endTime) - countLowerEqual(this.destinationToTimestampMap[destination], startTime - 1)
}

func countLowerEqual(v []int, upperBound int) int {
    low := 0
    high := len(v) - 1
    bestIdx := -1

    for low <= high {
        mid := (low + high) / 2

        if v[mid] <= upperBound {
            bestIdx = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return bestIdx + 1
}


/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
