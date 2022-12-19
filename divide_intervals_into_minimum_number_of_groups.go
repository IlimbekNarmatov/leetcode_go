type holder struct {
	a *[][]int
}

func minGroups(intervals [][]int) int {
	r := make([]holder, 0)
	for _, v := range intervals {
		i := index(&r, &v)
		if i < 0 {
			a := make([][]int, 1)
			a[0] = v
			r = append(r, holder{a: &a})
			continue
		}
		a := append(*r[i].a, v)
		r[i].a = &a
	}
	return len(r)
}

func index(a *[]holder, b *[]int) int {
	for i, h := range *a {
		t := false
		for _, aa := range *h.a {
			if isIntersect(&aa, b) {
				t = true
				break
			}
		}
		if !t {
			return i
		}
	}
	return -1
}

func isIntersect(a *[]int, b *[]int) bool {
	if (*a)[0] > (*b)[0] {
		return isIntersect(b, a)
	}
	if (*a)[0] == (*b)[0] {
		return true
	} else if (*a)[1] >= (*b)[0] && (*a)[1] <= (*b)[1] {
		return true
	}
	return false
}
