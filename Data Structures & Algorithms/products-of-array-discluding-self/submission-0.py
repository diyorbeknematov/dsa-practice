class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix = [1] * len(nums)
        suffix = [1] * len(nums)

        for i in range(0, len(nums)):
            if i > 0:
                prefix[i] = prefix[i-1]*nums[i-1]

        for i in range(len(nums)-1, -1, -1):
            if i < len(nums)-1:
                suffix[i] = suffix[i+1]*nums[i+1]

        
        for i in range(0, len(nums)):
            prefix[i] = prefix[i]*suffix[i]
        
        return prefix


        

# 1 2 4 6 
# 1 1 1 1 

# 1 1 2 8 

#  1  1 2 8 
# 48 24 6 1 
