<!-- ABOUT THE PROJECT -->
## About The Project
This is a project to test out parallel request with Golang. With some basic go features.



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites
Before you begin you must have Go installed and configured properly for your computer. Please see [https://golang.org/doc/install](https://golang.org/doc/install)

### Build from the source 

1. Clone the repo
   ```sh
   git clone https://github.com/nixon1333/golang-multi-routine.git
   ```
2. Build the binary and run
   build
   ```sh
   go build -o gohttp
   ```
   run
   ```sh
   ./gohttp -parallel 3 google.com facebook.com yahoo.com
   ```
or Run the go file
   ```sh 
   go run main.go -parallel 3 google.com facebook.com yahoo.com
   ``` 

<!-- LICENSE -->
## License

Distributed under the GNU License. See `LICENSE` for more information.


<!-- CONTACT -->
## Contact

Ashraful Islam Nixon - [@nixon13333](https://twitter.com/nixon1333)

Project Link: [https://github.com/nixon1333/golang-multi-routine](https://github.com/nixon1333/golang-multi-routine)


