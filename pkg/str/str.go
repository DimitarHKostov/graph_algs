package str

import (
	"math"
)

const (
	domain = "abcdefgihj"
)

func KnuthMorrisPratt(str, substr string) bool {
	n := len(str)
	m := len(substr)
	prefixTable := buildPrefixTable(substr)
	result := make([]int, 0)

	i, j := 0, 0

	for i < n {
		if substr[j] == str[i] {
			i++
			j++
		}

		if j == m {
			result = append(result, i-j)
			j = prefixTable[j-1]
		} else if i < n && substr[j] != str[i] {
			if j != 0 {
				j = prefixTable[j-1]
			} else {
				i++
			}
		}
	}

	return len(result) > 0
}

func buildPrefixTable(substr string) []int {
	m := len(substr)
	prefixTable := make([]int, m)

	j := 0
	i := 1

	for i < m {
		if substr[i] == substr[j] {
			j++
			prefixTable[i] = j
			i++
		} else {
			if j != 0 {
				j = prefixTable[j-1]
			} else {
				prefixTable[i] = 0
				i++
			}
		}
	}

	return prefixTable
}

func RabinKarp(str, substr string) bool {
	lenStr := int64(len(str))
	lenSubstr := int64(len(substr))

	if lenSubstr > lenStr {
		return false
	}

	lenDomain := int64(len(domain))

	var hashStr int64
	var targetHash int64
	var i int64

	for i = 0; i < lenSubstr; i++ {
		hashIStr := calculateHash(str[i], lenDomain, int64(lenSubstr-1-i))
		hashISubStr := calculateHash(substr[i], lenDomain, int64(lenSubstr-1-i))
		hashStr += hashIStr
		targetHash += hashISubStr
	}

	if hashStr == targetHash {
		return true
	}

	rightIndex := lenSubstr
	leftIndex := 0

	for rightIndex != lenStr {
		hashStr -= calculateHash(str[leftIndex], lenDomain, int64(lenSubstr-1))
		hashStr *= int64(lenDomain)
		hashStr += calculateHash(str[rightIndex], lenDomain, 0)

		leftIndex++
		rightIndex++

		if hashStr == targetHash {
			return true
		}
	}

	return false

	//either uncomment the for loop or the recursive solution
	//return checkRecursive(hashStr, targetHash, int64(leftIndex), rightIndex, lenStr, lenSubstr, lenDomain, str, substr)
}

func checkRecursive(hashStr, targetHash, leftIndex, rightIndex, lenStr, lenSubstr, lenDomain int64, str, subStr string) bool {
	hashStr -= calculateHash(str[leftIndex], lenDomain, int64(lenSubstr-1))
	hashStr *= int64(lenDomain)
	hashStr += calculateHash(str[rightIndex], lenDomain, 0)

	leftIndex++
	rightIndex++

	if hashStr == targetHash {
		return true
	}

	if rightIndex == lenStr {
		return false
	}

	return checkRecursive(hashStr, targetHash, leftIndex, rightIndex, lenStr, lenSubstr, lenDomain, str, subStr)
}

func calculateHash(c byte, constant, power int64) int64 {
	return int64(float64(reduce(c)) * math.Pow(float64(constant), float64(power)))
}

func reduce(b byte) byte {
	return b - 'a' + 1
}
