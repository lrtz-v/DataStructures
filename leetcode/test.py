from collections import defaultdict

class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        n = len(prices)

        start = 0
        end = 0
        first = 0
        second = 0

        points = list()

        for i in range(1, n):
            if prices[i] < prices[i-1] or i == n-1:
                if prices[i] < prices[i-1]:
                    end = i - 1
                else:
                    end = i
                points.append((start, end))
                # val = prices[end] - prices[start]
                # print("start: {}, end: {}, val: {}\n".format(start, end, val))
                # if val > first and val > second:
                #     first, second = val, first
                # elif val > second:
                #     second = val
                start = i
        
        print(points)

        # [(0, 2), (3, 5), (6, 8)]
        dp = defaultdict(dict)
        for i in [item[0] for item in points]:
            for j in [item[1] for item in points]:
                pass








        return first + second
        
if __name__ == "__main__":
    print(Solution().maxProfit([1, 2, 4, 2, 5, 7, 2, 4, 9, 0]))
