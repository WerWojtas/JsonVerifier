# Json Verifier
Verifies if any Resource field in Json input contains single "*".
#### System requirements
- go1.22.1
#### Running verifier
Clone this repo
```
git clone https://github.com/WerWojtas/JsonVerifier
cd JsonVerifier
cd src
```
Place your json data in file yourJson
Then run command (in src directory):
```
go run structs.go verify.go  main.go
```
#### Running tests
Run command
```
go test
```
