package shortestPathBinaryMatrix

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n == 0 || grid[0][0] == 1 || grid[n - 1][n - 1] == 1 {
		return -1
	}

	queue := [][]int{}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		tmp := []bool{}
		for j := 0; j < n; j++ {
			tmp = append(tmp, false)
		}
		visited[i] = tmp
	}

	queue = append(queue, []int{0, 0})
	visited[0][0] = true
	step := 1

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[i]

			if point[0] == n - 1 && point[1] == n - 1 {
				return step
			}

			for j := point[0]-1; j <= point[0]+1; j++ {
				for k := point[1]-1; k <= point[1]+1; k++ {
					if j >= 0 && j < n && k >= 0 && k < n && grid[j][k] == 0 && !visited[j][k] {
						queue = append(queue, []int{j, k})
						visited[j][k] = true
					}
				}
			}
		}
		step++
		queue = queue[size:]
	}

	return -1
}




/*
int BFS(Node start, Node target) { // 计算从起点 start 到终点 target 的最近距离
    Queue<Node> q; // 核心数据结构
    Set<Node> visited; // 避免走回头路

    q.offer(start); // 将起点加入队列
    visited.add(start);
    int step = 0; // 记录扩散的步数

    while (q not empty) {
        int sz = q.size();
        for (int i = 0; i < sz; i++) { //将当前队列中的所有节点向四周扩散
            Node cur = q.poll();
            if (cur is target) // 划重点：这里判断是否到达终点
				return step;
				
            for (Node x : cur.adj())  // 将 cur 的相邻节点加入队列
                if (x not in visited) {
                    q.offer(x);
                    visited.add(x);
                }
        }

        step++;  // 划重点：更新步数在这里
    }
}

*/
