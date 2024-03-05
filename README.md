# Go Data Structures

## Tree

```sh
go get github.com/adrian-lin-1-0-0/go-ds/tree@latest
```

- [Fenwick tree (Binary Index Tree)](./tree/bit.go)

  Time Complexity
  
  | Operation   |  Average | Worst case |
  |  ----       | ----     | ---        |
  | Update      | O(log n) | O(log n)   |
  | Query       | O(log n) | O(log n)   |
  | Range Query | O(log n) | O(log n)   |
  
  Space Complexity
  
  |  Average | Worst case |
  | ----     | ---        |
  | O(n)     |  O(n)      |

  ```go
  package main
  
  import (
	  "github.com/adrian-lin-1-0-0/go-ds/tree"
  )
  
  func main(){
    bit := tree.NewBIT[int](10)
    bit.Update(1, 1)
    bit.Update(2, 2)
    bit.Update(3, 3)
  
    bit.Query(3) // 6
    bit.RangeQuery(1,2) // 3
  }
  ```
