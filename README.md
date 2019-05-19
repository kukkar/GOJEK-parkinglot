# GOJEK-parkinglot


1. Since there is not formal try catch mechanism in golang except for panic, recover an Error Wrapper was written which reads the output of function, write the error to stdout and returns it.
2. For solving the problem, the data structure used for representing parking lot is a mix of Heap and HashMaps to solve all of the commands in optimal manner as possible, assuming space is not an issue. Further details have been added in the comments alongside as to why each struct was used and why.
3. No external testing library has been used as golang provides internal "testing" library. All test cases have been written using it.
4. The build process is written in the form a very simple Makefile which formats the code, and then builds and creates a binary according to the system on which it is being built, assuming go is already installed and GOPATH is alo set already and that make process is being called when the current working directory is the path of the root of the project.
5. the parking_lot shell script has been written which first makes the project, then checks if the build or tests failed.


## Run
```
$ ./parking_lot
test
....
```
OR
```
$ ./parking_lot abc.txt
output....
```