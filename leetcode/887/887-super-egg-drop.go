/*
你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。

每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。

你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。

每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。

你的目标是确切地知道 F 的值是多少。

无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？



示例 1：

输入：K = 1, N = 2
输出：2
解释：
鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
如果它没碎，那么我们肯定知道 F = 2 。
因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
示例 2：

输入：K = 2, N = 6
输出：3
示例 3：

输入：K = 3, N = 14
输出：4


提示：

1 <= K <= 100
1 <= N <= 10000
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(superEggDrop2(2, 100))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

var dp = map[int]map[int]int{}

func superEggDrop(K int, N int) int {
	if K == 0 {
		return 0
	}
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	if K == 1 {
		return N
	}

	if dp[K] == nil {
		dp[K] = make(map[int]int)
	} else if dp[K][N] > 0 {
		return dp[K][N]
	}

	var result = math.MaxInt64

	// 线性搜索
	//for i := 1; i <= N; i++ {
	//	result = min(result, max(
	//		// 鸡蛋碎了
	//		superEggDrop(K-1, i-1),
	//		// 鸡蛋没碎
	//		superEggDrop(K, N-i),
	//	)+1)
	//}

	// 二分搜索
	lo := 1
	hi := N
	for lo <= hi {
		mid := (lo + hi) / 2
		broken := superEggDrop(K-1, mid-1)
		noBroken := superEggDrop(K, N-mid)
		result = min(result, max(broken, noBroken)+1)
		if broken < noBroken {
			lo = mid + 1
		} else if broken > noBroken {
			hi = mid - 1
		} else if broken == noBroken {
			break
		}
	}
	dp[K][N] = result
	return result
}

func superEggDrop2(K int, N int) int {
	if K == 0 || N == 0 {
		return 0
	}
	// 通过另一种状态方程
	// 对于状态方程back[k][m] = n, 表示k个鸡蛋,尝试m次，最多能够测量的楼层
	// 那么转移方程为back[k][m] = back[k][m-1] + back[k-1][m-1] + 1
	var back = make([][]int, K+1)
	for i := range back {
		back[i] = make([]int, N+1)
	}
	// 尝试m次
	m := 0
	for back[K][m] < N {
		m++
		for k := 1; k <= K; k++ {
			back[k][m] = back[k][m-1] + back[k-1][m-1] + 1
		}
	}
	return m
}
