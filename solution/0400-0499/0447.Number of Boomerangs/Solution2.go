func numberOfBoomerangs(points [][]int) (ans int) {
	for _, p1 := range points {
		cnt := map[int]int{}
		for _, p2 := range points {
			d := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
			cnt[d]++
		}
		for _, x := range cnt {
			ans += x * (x - 1)
		}
	}
	return
}