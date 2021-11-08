package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(n int) bool {
	return n > 3
}
func firstBadVersion(n int) int {
	l, r := 1, n
	if !isBadVersion(n) {
		return n + 1
	}
	for l < r {
		m := (l + r) / 2
		if isBadVersion(m+1) && !isBadVersion(m) {
			return m + 1
		}
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return 1
}
