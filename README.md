# grep-cli-golang
golang bootcamp https://playbook.one2n.in/go-bootcamp/go-projects/grep-in-go

clone the repo
```
git clone https://github.com/biswaone/grep-cli-golang.git
```
```
cd grep-cli-golang
make
```
To run 
```
cd bin
./go-grep search_string ../sample.txt
```
## Search a directory for a string
```
./go-grep test ../test-directory
../test-directory/file1.txt this is the first test file in this directory
../test-directory/file2.txt hi a test line is present here 
../test-directory/file2.txt this is another test line 
../test-directory/inner/inner_file.txt this file contains a test line
```
## Flags
-i: case insensitive search 
```
./go-grep -i foo
bar
barbazfoo
Foobar
food
^D
barbazfoo
Foobar
food
```

-o: output to file
```
./go-grep search_string ../sample.txt -o outfile.txt
cat outfile.txt
I found the search_string in the file.
Another line also contains the search_string
```
-c: count match
```
./go-grep search_string -c  ../sample.txt
2
```