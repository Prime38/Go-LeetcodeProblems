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

## Day4
### Leetcode Problemset
#### 771.JewelsandStones
time-> o(n) space->O(n)
```
func numJewelsInStones(J string, S string) int {
	ans:=0
	arr:=[128]bool{}
	for i:=0;i<len(J);i++{
		arr[int(J[i])-65]=true
	}
	for i:=0;i<len(S);i++{
		if arr[int(S[i])-65]==true{
			ans++
		}
	}
	return ans
}

```
#### 804.UniqueMorseCodeWords
time-> O(n^2) space->O(n)
```
func uniqueMorseRepresentations(words []string) int {
	morse:=[]string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	m:=make(map[string]bool)
	l1:=len(words)
	for i:=0;i<l1;i++{
		l2:=len(words[i])
		word:=""
		for j:=0;j<l2;j++ {
			word=word+morse[int(words[i][j])-97]

		}
		fmt.Println(word)
		m[word]=true
	}

	return len(m)
}

```



## Day3
### Leetcode Problemset
#### 14.Longest Common Prefix
Charecter by Charecter matching <br>
time->O(nm) space -> O(m) . m=length of prefix
```
if len(strs)==0{
		return ""
	}
	minlen:=len(strs[0])
	LowStr:=strs[0]
	for i:=0;i<len(strs);i++{
		if minlen> len(strs[i]){
			minlen=len(strs[i])
			LowStr=strs[i]
		}
	}
	ans := 0
	for i:=0;i<minlen;i++{
		match:=true
		for j:=0;j<len(strs);j++{
			if strs[j][i]!=LowStr[i]{
				match=false
			}
		}
		if match==true{
			ans = ans + 1
		} else {
			break
		}
	}
	return LowStr[:ans]
```
#### 169. Majority Element
There are two solutions
- 1. Using map <br> 
time->O(n) space -> O(n)
```
m:=make(map[int]int)
	max:=len(nums)/2
	var  ans int
	for i:=0;i<len(nums);i++{
		m[nums[i]]++
		if max<m[nums[i]]{
			max=m[nums[i]]
			ans=nums[i]
		}
	}
	return ans
```
- 2. Using sort <br> 
time->O(nlogn) space -> O(n)
```
l:=len(nums)
	sort.Ints(nums)
	return nums[l/2]
```
## Day2
### Leetcode Problemset
#### 1.Two Sum
Using map to find the sum of two numbers that is equal to target number.<br>
time->O(n) space -> O(n)
```
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, k}
		}
		m[v] = j
	}
	return nil
}

```

#### 136.Single Number<br>
There are 3 solutions for this problem.
- 1. using map to find double values .<br>
time->O(n) space -> O(n)
```
func singleNumber(nums []int) int {
    m:=make(map[int]int)
	min:=-199999999
	for k,v:=range nums{
		if m[v]!=min{
			if m[v]==0{
				m[v]=k+1
			} else {
				nums[m[v]-1]=min
				nums[k]=min
			}
	
		}
	}
	for i:=0;i<len(nums);i++{
		if nums[i]!=min{
			return nums[i]
		}
	}
	return 0
}


```
- 2. (total of all numbers in array) - 2*(total of double numbers in array).<br>
time->O(n) space -> O(n)
```
func singleNumber(nums []int) int {
    m:=make(map[int]bool)
	total:=0
	twices:=0
	for i:=0;i<len(nums);i++{
		total+=nums[i]
		if m[nums[i]]==true{
			twices+=2*nums[i]
		} else{
			m[nums[i]]=true
		}
	}
	return total-twices
}


```
- 3. using bitwise xor . <br>
time->O(n) ,space -> O(1)

```
func singleNumber(nums []int) int {
    ans:=0
	for i:=0;i<len(nums);i++{
		ans^=nums[i]
	}
	return ans
}

```

## Day1
### Leetcode ProblemSet
#### 1108.Defanging an IP Address
```
func defangIPaddr(address string) string {
	var ans string
	for i:=0;i< len(address);i++ {
		if address[i]=='.' {
			ans+="[.]"
		} else {
			ans+=string(address[i])
		}
	}
	return ans
}

```
#### 448.Find All Numbers Disappeared in an Array

There are 3 solutions for this problem
- 1. 440 ms 8.2 MB 
```
ans:=[]int{}
for i:=0;i<len(nums);i++{
		t:=int(math.Abs(float64(nums[i]))-1)
		//fmt.Println(t)
		nums[t]= int(-math.Abs(float64(nums[t])))
	}
	fmt.Println(nums)
	for i:=0;i<len(nums);i++{
		if nums[i]>0{
			f:=i+1
			ans=append(ans,f)
		}
	}
	return ans
```
- 2. 392 ms 8.3 MB 
```
ans:=[]int{}
	
	for i:=0;i<len(nums);i++{
		if nums[i]==0{
			continue
		}
		next:=nums[i]-1
		cur:=i
		for nums[next]>0{
	
			cur=next
			next=nums[next]-1
			nums[cur]=0
			fmt.Println(cur,next)
		}
		fmt.Println(nums)
	}
```
- 3. 396 ms 8.5 MB 
```
for i := 0; i < len(nums); i++ {
		fmt.Println(nums)
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			fmt.Println(nums)
		}
	}

	ans := make([]int, 0, len(nums))

	for i, n := range nums {
		if n != i+1 {
			ans = append(ans, i+1)
		}
	}

	return ans
```

