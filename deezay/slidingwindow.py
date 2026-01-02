
def maxSlidingWindow(nums, k):
    if not nums:
        return []
    
    result = []
    window = []
    
    for i in range(len(nums)):
        while window and window[0] < i - k + 1:
            window.pop(0)
            
        while window and nums[window[-1]] < nums[i]:
            window.pop()
            
        window.append(i)
        
        if i >= k - 1:
            result.append(nums[window[0]])
            
    return result

nums = [1,3,-1,-3,5,3,6,7]
k = 3
print(maxSlidingWindow(nums, k))