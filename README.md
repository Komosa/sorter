# sorter



## Usage

### struct literal

```go
package main

import (
	"fmt"
	"sort"

	"github.com/mattn/sorter"
)

func main() {
	s := []string{
		"2",
		"3",
		"1",
	}

	sort.Sort(&sorter.Wrapper{
		Length: len(s), // or just '3'
		LessFunc: func(i, j int) bool {
			return s[i] < s[j]
		},
		SwapFunc: func(i, j int) {
			s[i], s[j] = s[j], s[i]
		},
	})

	fmt.Println(s)
}
```

### NewWrapper

```go
package main

import (
	"fmt"
	"sort"

	"github.com/mattn/sorter"
)

func main() {
	s := []string{
		"2",
		"3",
		"1",
	}

	sort.Sort(sorter.NewWrapper(
		len(s),
		func(i, j int) bool {
			return s[i] < s[j]
		},
		func(i, j int) {
			s[i], s[j] = s[j], s[i]
		},
	))

	fmt.Println(s)
}
```

### NewWrapperWith

```go
package main

import (
	"fmt"
	"sort"

	"github.com/mattn/sorter"
)

func main() {
	s := []string{
		"2",
		"3",
		"1",
	}

	sort.Sort(sorter.NewWrapperWith(
		s,
		func(i, j int) bool {
			return s[i] < s[j]
		},
	))

	fmt.Println(s)
}
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
