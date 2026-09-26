class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        # quick_sort(nums, 0, len(nums)-1)

        l, r = 0, len(nums)-1
        i = 0

        while i <= r:
            if nums[i] == 0:
                nums[i], nums[l] = nums[l], nums[i]
                l += 1
            elif nums[i] == 2:
                nums[i], nums[r] = nums[r], nums[i]
                r -=1 
                i -=1 
            
            i += 1

def quick_sort(nums: List[int], low: int, high: int):
    if low >= high:
        return 
    
    pivot = nums[high]
    i = low 
    for j in range(low, high):
        if nums[j] < pivot:
            nums[i], nums[j] = nums[j], nums[i]
            i += 1
    
    nums[i], nums[high] = nums[high], nums[i]

    quick_sort(nums, low, i - 1)
    quick_sort(nums, i + 1, high)

# 2 0 1 1 0 
# 