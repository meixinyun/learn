[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: random

- [Reference](#reference)
- [random integer](#random-integer)
- [random float](#random-float)
- [random duration](#random-duration)
- [random `bytes`](#random-bytes)
- [random select](#random-select)

[↑ top](#go-random)
<br><br><br><br><hr>


#### Reference

- [package `math/rand`](http://golang.org/pkg/math/rand/)

[↑ top](#go-random)
<br><br><br><br><hr>


#### random integer

[Code](http://play.golang.org/p/88gzcG-r4v):

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	fmt.Println(src.Int63()) // 8965630292270293660

	random := rand.New(src)
	fmt.Println(random.Int())      // 7742198863449996164
	fmt.Println(random.Int31())    // 1780122247
	fmt.Println(random.Int31n(3))  // 0
	fmt.Println(random.Int63())    // 838216768439018635
	fmt.Println(random.Int63n(10)) // 7
}

```

[↑ top](#go-random)
<br><br><br><br><hr>


#### random float

[Code](http://play.golang.org/p/AyWtUt-W7U):

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	fmt.Println(random.Float32())     // 0.7096111
	fmt.Println(random.Float64())     // 0.7267748269300062
	fmt.Println(random.ExpFloat64())  // 1.4478015992783408
	fmt.Println(random.NormFloat64()) // -1.7676830716730048
}

```

[↑ top](#go-random)
<br><br><br><br><hr>


#### random duration

[Code](http://play.golang.org/p/251bHlcW9S):

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println(duration(5*time.Second, 12*time.Second)) // 5.061306405s
}

func duration(min, max time.Duration) time.Duration {
	if min >= max {
		// return a random duration
		return 7*time.Second + 173*time.Microsecond
	}
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	adt := time.Duration(random.Int63n(int64(max - min)))
	return min + adt
}

```

[↑ top](#go-random)
<br><br><br><br><hr>


#### random `bytes`

```go
package main

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	b := make([]byte, 10)
	if _, err := crand.Read(b); err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// ��.���ms#

	fmt.Println(string(randBytes(10)))
	// IdPDZOxast

	fmt.Println(multiRandBytes(3, 5))
	// [[119 121 67] [114 70 70] [112 90 100] [74 85 77] [84 101 101]]
}

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func randBytes(bytesN int) []byte {
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, bytesN)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := bytesN-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

func multiRandBytes(bytesN, sliceN int) [][]byte {
	m := make(map[string]struct{})
	rs := [][]byte{}
	for len(rs) != sliceN {
		b := randBytes(bytesN)
		if _, ok := m[string(b)]; !ok {
			rs = append(rs, b)
			m[string(b)] = struct{}{}
		} else {
			continue
		}
	}
	return rs
}

```

[↑ top](#go-random)
<br><br><br><br><hr>


#### random select

```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// nr := rand.New(rand.NewSource(time.Now().UnixNano()))
	// fmt.Println(nr.Float32())

	// fmt.Println(rand.Float32())
}

func main() {
	ents := []weightEntry{
		{weight: 0.05},
		{weight: 0.8},
		{weight: 0.15},
	}
	wt := createWeightTable(ents)

	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
}

/*
{0.8 }
{0.8 }
{0.15 }
{0.8 }
{0.8 }
{0.8 }
{0.8 }
{0.8 }
*/

type weightEntry struct {
	weight float32
	name   string
}

type weightTable struct {
	entries       []weightEntry
	distributions []float32
}

func (wt weightTable) Len() int           { return len(wt.entries) }
func (wt weightTable) Swap(i, j int)      { wt.entries[i], wt.entries[j] = wt.entries[j], wt.entries[i] }
func (wt weightTable) Less(i, j int) bool { return wt.entries[i].weight < wt.entries[j].weight }

func createWeightTable(entries []weightEntry) *weightTable {
	wt := weightTable{entries: entries}
	sort.Sort(wt)
	var cw float32
	for _, entry := range wt.entries {
		cw += entry.weight
		wt.distributions = append(wt.distributions, cw)
	}
	return &wt
}

func (wt *weightTable) choose() weightEntry {
	entryN := len(wt.entries)
	lastWeight := wt.entries[len(wt.entries)-1].weight
	idx := sort.Search(entryN, func(i int) bool {
		// returns the smallest index, of which distribution is
		// greater than random weight value
		return wt.distributions[i] >= rand.Float32()*lastWeight
	})
	return wt.entries[idx]
}

```

```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// nr := rand.New(rand.NewSource(time.Now().UnixNano()))
	// fmt.Println(nr.Float32())

	// fmt.Println(rand.Float32())
}

func main() {
	ents := []weightEntry{
		{weight: 0.05},
		{weight: 0.8},
		{weight: 0.15},
	}
	wt := createWeightTable(ents)

	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
}

/*
{0.8 }
{0.8 }
{0.8 }
{0.15 }
{0.15 }
{0.8 }
{0.8 }
{0.8 }
*/

type weightEntry struct {
	weight float32
	name   string
}

type weightTable struct {
	entries    []weightEntry
	sumWeights float32
}

func (wt weightTable) Len() int           { return len(wt.entries) }
func (wt weightTable) Swap(i, j int)      { wt.entries[i], wt.entries[j] = wt.entries[j], wt.entries[i] }
func (wt weightTable) Less(i, j int) bool { return wt.entries[i].weight < wt.entries[j].weight }

func createWeightTable(entries []weightEntry) *weightTable {
	wt := weightTable{entries: entries}
	sort.Sort(wt)
	for _, entry := range wt.entries {
		wt.sumWeights += entry.weight
	}
	return &wt
}

func (wt *weightTable) choose() weightEntry {
	v := rand.Float32() * wt.sumWeights
	var sum float32
	var idx int
	for i := range wt.entries {
		sum += wt.entries[i].weight
		if sum >= v {
			idx = i
			break
		}
	}
	return wt.entries[idx]
}

```

[↑ top](#go-random)
<br><br><br><br><hr>
