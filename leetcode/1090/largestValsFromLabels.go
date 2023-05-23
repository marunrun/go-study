package main

import "sort"

func main() {

}

/**
根据题目描述，我们需要从 nnn 个元素的集合中选出一个子集，子集元素个数不超过 numWantednumWantednumWanted，且子集中最多有相同标签的 useLimituseLimituseLimit 项，使得子集的值之和最大。因此，我们应该贪心地选择集合中值较大的元素，同时记录每个标签出现的次数，当某个标签出现的次数达到 useLimituseLimituseLimit 时，我们就不能再选择该标签对应的元素了。

具体地，我们先将集合中的元素按照值从大到小进行排序，然后从前往后遍历排序后的元素。在遍历的过程中，我们使用一个哈希表 cntcntcnt 记录每个标签出现的次数，如果某个标签出现的次数达到了 useLimituseLimituseLimit，那么我们就跳过该元素，否则我们就将该元素的值加到最终的答案中，并将该标签出现的次数加 111。同时，我们用一个变量 numnumnum 记录当前子集中的元素个数，当 numnumnum 达到 numWantednumWantednumWanted 时，我们就可以结束遍历了。

遍历结束后，我们就得到了最大的分数。

作者：ylb
链接：https://leetcode.cn/problems/largest-values-from-labels/solutions/2280034/python3javacgotypescript-yi-ti-yi-jie-ta-evaq/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) (ans int) {
	n := len(values)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{values[i], labels[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})
	cnt := map[int]int{}
	for i, num := 0, 0; i < n && num < numWanted; i++ {
		v, l := pairs[i][0], pairs[i][1]
		if cnt[l] < useLimit {
			cnt[l]++
			num++
			ans += v
		}
	}
	return ans
}
