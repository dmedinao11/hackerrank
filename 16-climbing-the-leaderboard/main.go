package main

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	var result []int32
	rankSet := removeRepeated(ranked)

	for _, p := range player {
		result = append(result, getScoreRank(rankSet, p))
	}

	return result
}

func getScoreRank(rank []int32, p int32) int32 {
	return findPlayerRankIndex(rank, p, 0, len(rank)-1) + 1
}

func findPlayerRankIndex(rank []int32, p int32, start int, end int) int32 {
	if rank[start] == p {
		return int32(start)
	}

	middle := (end - start) / 2

	if middle != 0 {
		index := middle + start
		if rank[index] > p {
			return findPlayerRankIndex(rank, p, index, end)
		} else {
			return findPlayerRankIndex(rank, p, start, index)
		}
	}

	if p == rank[end] {
		return int32(end)
	} else if p > rank[start] {
		return int32(start)
	} else if p > rank[end] {
		return int32(start + 1)
	} else {
		return int32(end + 1)
	}

}

func removeRepeated(ranked []int32) []int32 {
	var set []int32
	var current int32

	for i, r := range ranked {
		if i == 0 {
			set = append(set, r)
			current = r
			continue
		}

		if current == r {
			continue
		}

		set = append(set, r)
		current = r
	}

	return set
}
