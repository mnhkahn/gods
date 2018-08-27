// Package xsort
// Reference: https://segmentfault.com/a/1190000005111605
package xsort

// MergeAndUints computes intersection of ids list.
func MergeAndUints(ids ...[]uint32) []uint32 {
	if len(ids) == 0 {
		return nil
	} else if len(ids) == 1 {
		return ids[0]
	}

	// Find out the shortest list.
	shortest := 0
	for i := 1; i < len(ids); i++ {
		if len(ids[i]) < len(ids[shortest]) {
			shortest = i
		}
	}

	res := make([]uint32, 0, len(ids[shortest]))
	// Cursors of all lists.
	js := make([]int, len(ids))
	// Intersection.
	for _, id := range ids[shortest] {
		exists := true
		for m := 0; m < len(ids); m++ {
			if m == shortest {
				continue
			}
			mIds := ids[m]
			j := js[m]
			if j >= len(mIds) {
				exists = false
				break
			} else if id < mIds[j] {
				exists = false
				break
			} else if id > mIds[j] {
				// If the value of the cursor is less than the current id value, move the cursor until the value indicated is as equal as possible to the current id value.
				jj := j
				for ; jj < len(mIds) && mIds[jj] < id; jj++ {
				}
				js[m] = jj

				// If the cursor is not equal to the current id after the end of the movement, this value is not in the intersection.
				if jj >= len(mIds) || mIds[jj] != id {
					exists = false
				}
			}
		}

		if exists {
			res = append(res, id)
		}
	}
	return res
}

// MergeOrUints merge sort ids list.
func MergeOrUints(ids ...[]uint32) []uint32 {
	if len(ids) == 0 {
		return nil
	} else if len(ids) == 1 {
		return ids[0]
	}

	res := make([]uint32, len(ids[0]))
	copy(res, ids[0])
	for i := 1; i < len(ids); i++ {
		res = MergeUInt(res, ids[i])
	}

	return res
}

// MergeUInt merge sort a & b.
func MergeUInt(a, b []uint32) []uint32 {
	i, j := 0, 0

	res := make([]uint32, 0, len(a)+len(b))
	for i <= len(a)-1 && j <= len(b)-1 {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else if a[i] > b[j] {
			res = append(res, b[j])
			j++
		} else { // a[i] == b[j]
			res = append(res, a[i])
			i++
			j++
		}
	}

	for i <= len(a)-1 {
		res = append(res, a[i])
		i++
	}

	for j <= len(b)-1 {
		res = append(res, b[j])
		j++
	}

	return res
}
