
## Features

- *Allows reading and writing of ID3 tags*
- *Works for both ID3v1 and ID3v2 (2.3, and in most cases 2.4)*
- *Extremely simple to use. Abstracts the complicated parts*

## Installation

```sh
go get github.com/kalbhor/mutago
```

## Usage

```go
import (
"github.com/kalbhor/mutago"
)
```

Note : The [examples](https://github.com/makebyte/mutago/tree/master/examples) folder contains implementations of all the methods below. It uses a sample mp3 file licensed under creative commons.

### Read methods

##### Open file
```go
f, err := mutago.Open("file.mp3")
if err != nil {
// Handle error
}
defer f.Close() // Always good to close after using
```

##### Provides a list of available tags
```go
tags := f.List() // List() Provides an array of tags in the file. 
```

##### Provides values of basic tags
```go
/*
For basic tags such as song title, album name and artist there are seperate methods. 
(So that you don't need to know the ID3 tag identifier)
*/
title, err := f.Title()
album, err := f.Album()
artist, err := f.Artist()
```

##### Provides values of other tags
```go
/*
For more advanced tags such as comments, year, lead singer, etc.
Use Get() method that fetches the ID3 tag's value, provided the tag identifier.
For a list of valid tags check id3.org.

Get() will return an error if the tag is invalid or doesn't exist in the file.
*/

text, err := f.Get("TXXX")
lead, err := f.Get("TPE1")

```

### Write methods
```go
/*
Yet to be built.

Will be added soon.
*/

```


## License

[![PyPI](https://img.shields.io/pypi/l/Django.svg)](https://github.com/makebyte/mutago/blob/master/LICENSE)

```

The BSD 2-Clause License
Copyright © 2022 Lakshay Kalbhor
All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

```
