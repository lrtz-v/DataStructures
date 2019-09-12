
class Solution(object):

    def findSubsequences(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        subs = {()}
        for num in nums:
            subs |= {sub + (num,)
                    for sub in subs
                    if not sub or sub[-1] <= num}
            print(subs)
        return [sub for sub in subs if len(sub) >= 2]
        
if __name__ == '__main__':
    print(Solution().findSubsequences([4, 6, 7, 7]))
