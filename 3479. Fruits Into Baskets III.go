const BUCKET_SIZE = 300

func numOfUnplacedFruits(fruits []int, baskets []int) int {
    bucketsMaxCap, buckets := getBucketsMaxCap(baskets)

    result := 0
    for _, fruit := range fruits {
        foundFit := false

        for i := 0; i < buckets; i++ {
            if fruit <= bucketsMaxCap[i] {
                foundFit = true
                removeFirstFitFromBucket(i, fruit, baskets)
                updateBucketMaxCap(i, baskets, bucketsMaxCap)
                break
            }
        }

        if !foundFit {
            result++
        }
    }

    return result
}

func getBucketsMaxCap(baskets []int) ([]int, int) {
    buckets := (len(baskets) + BUCKET_SIZE - 1) / BUCKET_SIZE
    bucketsMaxCap := make([]int, buckets)

    for i, cap := range baskets {
        bucketId := i / BUCKET_SIZE

        bucketsMaxCap[bucketId] = max(bucketsMaxCap[bucketId], cap)
    }

    return bucketsMaxCap, buckets
}

func removeFirstFitFromBucket(bucketId int, qt int, baskets []int) {
    startBucket := bucketId * BUCKET_SIZE

    for i := startBucket; i < (bucketId + 1) * BUCKET_SIZE && i < len(baskets); i++ {
        if baskets[i] >= qt {
            baskets[i] = 0
            return
        }
    } 
}

func updateBucketMaxCap(bucketId int, baskets []int, bucketsMaxCap []int) {
    bucketsMaxCap[bucketId] = 0
    startBucket := bucketId * BUCKET_SIZE

    for i := startBucket; i < (bucketId + 1) * BUCKET_SIZE && i < len(baskets); i++ {
        bucketsMaxCap[bucketId] = max(bucketsMaxCap[bucketId], baskets[i])
    }
}
