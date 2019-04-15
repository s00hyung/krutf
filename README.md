# krutf
**EUC-KR <-> UTF-8**

Simple character encoding converter between EUC-KR and UTF-8 in golang. 



## Install
    go get github.com/soohl/krutf

## Usage

* Convert byte slices from EUC-KR to UTF-8 encoding and vice versa. `b` is a byte slices ([]byte) to convert and `c` is the expecting encoding (`EUCKR` or `UTF8`).
```
convertedByte, err := krutf.ConvertByte(b, c)
```

* Convert string from EUC-KR to UTF-8 encoding and vice versa. `s` is a string to convert and `c` is the expecting encoding (`EUCKR` or `UTF8`).

```
convertedString, err := krutf.ConvertString(s, c)
```

## Example
```
package main

import (
	"fmt"

	"github.com/soohl/krutf"
)

func main() {
	// Encoded in UTF
	utfString := "안녕하세요"
	fmt.Printf("%X\n", utfString)

	// Convert to EUC-KR
	krString, _ := krutf.ConvertString(utfString, krutf.EUCKR)
	fmt.Printf("%X\n", krString)
}
```

