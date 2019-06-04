package main

import (
	"log"
	"reflect"
)

/* Finding the formula ---
k = number of rotations
n = array size
i = index of the rotated array if we try to do it in O(n)
x = index where to read from meaning rotated[i] = a[x]

Scenario 1 - [1, 2, 3, 4, 5] -> [5, 1, 2, 3, 4]
| k | n | i | x |
| 4 | 5 | 0 | 4 |
| 4 | 5 | 1 | 0 |
| 4 | 5 | 2 | 1 |
| 4 | 5 | 3 | 2 |
| 4 | 5 | 4 | 3 |

Scenario 2 - [1, 2, 3, 4, 5] -> [1, 2, 3, 4, 5]
| k | n | i | x |
| 5 | 5 | 0 | 0 |
| 5 | 5 | 1 | 1 |
| 5 | 5 | 2 | 2 |
| 5 | 5 | 3 | 3 |
| 5 | 5 | 4 | 4 |

f = ((k % n) + i) % n
*/
func rotate(a []int, k int) []int {
	l := len(a)
	mod := k % l
	rotated := make([]int, l)
	for i := 0; i < l; i++ {
		rotated[i] = a[(mod+i)%l]
	}
	return rotated
}

func main() {
	scenarios := []scenario{
		{
			a:        []int{1, 2, 3, 4, 5},
			k:        4,
			expected: []int{5, 1, 2, 3, 4},
		},
		{
			a:        []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			a:        []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{2, 3, 4, 5, 1},
		},
	}

	for i, s := range scenarios {
		actual := rotate(s.a, s.k)
		if !reflect.DeepEqual(actual, s.expected) {
			log.Printf("Scenario %d failed, got %v instead of %v", i+1, actual, s.expected)
		}
	}
}

type scenario struct {
	a        []int
	k        int
	expected []int
}
