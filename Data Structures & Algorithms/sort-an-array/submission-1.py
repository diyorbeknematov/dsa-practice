class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        return quick_sort(nums)

def merge_sort(nums: List(int)):
    if len(nums) <= 1:
        return nums
    
    mid = len(nums) // 2 

    left = merge_sort(nums[:mid])
    right = merge_sort(nums[mid:])

    return merge(left, right)

def merge(right: List[int], left: List[int]) -> List[int]:
    i, j = 0, 0
    res = []

    while i < len(right) and j < len(left):
        if right[i] < left[j]:
            res.append(right[i])
            i += 1
        else:
            res.append(left[j])
            j += 1

    while i < len(right):
        res.append(right[i])
        i += 1

    while j < len(left):
        res.append(left[j])
        j += 1

    return res


def quick_sort(nums):
    if len(nums) <= 1:
        return nums

    pivot = nums[len(nums) // 2]

    left = []
    right = []
    equal = []

    for num in nums:
        if num < pivot:
            left.append(num)
        elif num > pivot:
            right.append(num)
        else:
            equal.append(num)

    left = quick_sort(left)
    right = quick_sort(right)

    return left + equal + right

