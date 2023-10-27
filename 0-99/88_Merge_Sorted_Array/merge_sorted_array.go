package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int)  {
    firstPointer := m - 1
	secondPointer := n - 1
	thirdPointer := m + n - 1
	for secondPointer >= 0 {
		if firstPointer >= 0 && (nums1[firstPointer] > nums2[secondPointer]) {
			nums1[thirdPointer] = nums1[firstPointer]
			firstPointer--
		} else {
			nums1[thirdPointer] = nums2[secondPointer]
			secondPointer--
		} 
		thirdPointer--
	}
}