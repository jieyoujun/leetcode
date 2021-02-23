package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	inDegreeTable := make([]int, numCourses)
	ncan := 0
	learnQueue := make([]int, 0)

	for _, v := range prerequisites {
		// [1, 0]: 学1先学0, 即 0->1
		edges[v[1]] = append(edges[v[1]], v[0])
		inDegreeTable[v[0]]++
	}

	for v, outDegree := range inDegreeTable {
		if outDegree == 0 {
			// 从入度为0（不需要基础）的开始学
			learnQueue = append(learnQueue, v)
			inDegreeTable[v] = -1
		}
	}

	for len(learnQueue) > 0 {
		ncan++
		course := learnQueue[0]
		learnQueue = learnQueue[1:]
		for _, nextCourse := range edges[course] {
			if inDegreeTable[nextCourse] == 1 {
				learnQueue = append(learnQueue, nextCourse)
				inDegreeTable[nextCourse] = -1
			} else {
				inDegreeTable[nextCourse]--
			}
		}
	}

	return ncan == numCourses
}
