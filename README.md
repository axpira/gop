<div id="top"></div>

[![Go Reference](https://pkg.go.dev/badge/github.com/axpira/gop.svg)](https://pkg.go.dev/github.com/axpira/gop)
[![Go Report Card](https://goreportcard.com/badge/github.com/axpira/gop)](https://goreportcard.com/report/github.com/axpira/gop)

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">

<h3 align="center">GOP</h3>

  <p align="center">
    Golang Patterns is commom used interfaces to use in your projects
    <br />
    <a href="https://github.com/axpira/gop"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/axpira/gop/issues">Report Bug</a>
    ·
    <a href="https://github.com/axpira/gop/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project


Using this project you can just send log (for example) 
and don't think about how implementation you'll use.
This can be done any time, and can be changed any time too.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- GETTING STARTED -->

## Installation

```go
go get -u github.com/axpira/gop/log
```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

The implementation will format the output, but you can change anytime implementation.

### Choose a implementation

For now:
- [goplogjson](https://github.com/axpira/goplogjson)
- [adapter](https://github.com/axpira/goplogadapter)


### Simple Log
```go
package main

import (
	"github.com/axpira/gop/log"
)

func main() {
	log.Info("Hello World")
}
```

### Add fields to log
```go
package main

import (
	"github.com/axpira/gop/log"
)

func main() {
	log.Inf(log.
		Str("str_field", "hello").
		Int("int_field", 42).
		Msg("Hello World"),
	)
}
```

### Leveled logging
```go
package main

import (
	"errors"

	"github.com/axpira/gop/log"
)

func main() {
	log.Error("Hello World", errors.New("unknown error"))
}
```

### Other Leveled logging
```go
package main

import (
	"errors"

	"github.com/axpira/gop/log"
)

func main() {
	log.Log(
		log.ErrorLevel,
		log.Msg("Hello World").Err(errors.New("unknown error")),
	)
```

### Sub logger
```go
package main

import (
	"errors"

	"github.com/axpira/gop/log"
)

func main() {
	l := log.With(log.Str("mykey", "value"))
	l.Trace("Hello World")
```

<p align="right">(<a href="#top">back to top</a>)</p>




<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Thiago Ferreira - thiagogbferreira@gmail.com

Project Link: [https://github.com/axpira/gop](https://github.com/axpira/gop)

<p align="right">(<a href="#top">back to top</a>)</p>




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/axpira/gop.svg?style=for-the-badge
[contributors-url]: https://github.com/axpira/gop/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/axpira/gop.svg?style=for-the-badge
[forks-url]: https://github.com/axpira/gop/network/members
[stars-shield]: https://img.shields.io/github/stars/axpira/gop.svg?style=for-the-badge
[stars-url]: https://github.com/axpira/gop/stargazers
[issues-shield]: https://img.shields.io/github/issues/axpira/gop.svg?style=for-the-badge
[issues-url]: https://github.com/axpira/gop/issues
[license-shield]: https://img.shields.io/github/license/axpira/gop.svg?style=for-the-badge
[license-url]: https://github.com/axpira/gop/blob/master/LICENSE.txt
