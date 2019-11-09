# Leetcode Problem Solving in Go language

# 100DaysOfCode

## Day5

### Leetcode Problemset
#### 961.N-RepeatedElementinSize2NArray
There are 3 solutions
- 1.Using sort 
time-> O(nlogn) space O(1)
```
func repeatedNTimes(A []int) int {
	// sol 1
	sort.Ints(A)
	l:=len(A)/2
	if A[l-1]==A[l]{
		return A[l]
	} else {
		if A[l-1]==A[l-2]{
			return A[l-1]
		} else if A[l]==A[l+1] {
			return A[l]
		}
	}
	return 0
	
}

```
- 2.Using map
time-> o(n) space-> O(n)
```
func repeatedNTimes(A []int) int {
	
	m:=make(map[int]int)
	for i:=0;i<len(A);i++{
		if m[A[i]]>0{
			return A[i]
		} else {
			m[A[i]]++
		}
	}
	return 0
	
}
```
- 3. Using probabilistic intution
time-> O(n) space-> O(n)
```
func repeatedNTimes(A []int) int {
	for i:=1;i<4;i++{
		for j:=0;j<len(A)-i;j++{
			if A[j]==A[i+j]{
				return A[j]
			}
		}
	}
	return 0
}


```

#### 905.SortArrayByParity 
There are two solutions 
- 1. appending odd array after even array
time-> O(n) space-> O(n)
```
odd:=[]int{}
	even:=[]int{}
	for i:=0;i<len(A);i++{
		if A[i]%2==0{
			even= append(even, A[i])
		} else{
			odd=append(odd,A[i])
		}
	}
	return append(even,odd...)
```
- 2. negating even numbers ,then sort the array and then again negate even numbers
time->O(nlogn) space O(1)
```
for i:=0;i<len(A);i++{
			if A[i]%2==0{
				A[i]=-A[i]
			}
		}
		sort.Ints(A)
	for i:=0;i<len(A);i++{
		if A[i]<0{
			A[i]=-A[i]
		} else {
			break
		}
	}
	return A

```
### 1207.UniqueNumberofOccurrences
time->o(nlogn) space->O(n)

```
func uniqueOccurrences(arr []int) bool {
	m:=make(map[int]int)
	for i:=0;i<len(arr);i++{
		m[arr[i]]++
	}
	arr=[]int{}
	for _,v:=range m{
		arr=append(arr,v)
	}
	sort.Ints(arr)

	set:=[]int{}
	set=append(set,arr[0])
	for i:=1;i<len(arr);i++{
		if arr[i]!=set[len(set)-1]{
			set=append(set,arr[i])
		}
	}
	return (len(arr)==len(set))
}

```

#### 832. FlippinganImage
time->(n^2) space->O(1)
```
func flipAndInvertImage(A [][]int) [][]int {
	for i:=0;i<len(A);i++{
		l:=len(A[i])
		for j:=0;j<(l+1)/2;j++{
			A[i][j],A[i][l-j-1]=A[i][l-j-1]^1,A[i][j]^1
		}
	}
	return A
}
```

