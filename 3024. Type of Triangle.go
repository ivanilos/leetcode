func triangleType(nums []int) string {
    slices.Sort(nums)

    if nums[0] + nums[1] <= nums[2] {
        return "none"
    } else if nums[0] == nums[1] && nums[1] == nums[2] {
        return "equilateral"
    } else if nums[0] != nums[1] && nums[1] != nums[2] {
        return "scalene"
    } else {
        return "isosceles"
    }
}
