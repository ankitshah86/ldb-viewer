# LevelDB Viewer
> **Warning: This program is work in progress.** For the most part it should work as expected, it may break in certain situations.

## ABOUT
This program can be used to view levelDB database in the browser. You must have the latest version of GoLang installed on your machine. 
   

## How to Run 
To run the program, you will need to supply the absolute path of the existing LevelDB directory as an argument. If the dbpath argument is not supplied, A test database (testdb) will be created and will be shown in the browser.
    
For windows 

 ``` go build && ./ldb-viewer.exe ```

 ``` go build && ./ldb-viewer.exe --dbpath=Absolute/Path/To/The/LevelDB/Folder ```

For linux

 ``` go build && ./ldb-viewer --dbpath=Absolute/Path/To/The/LevelDB/Folder ```



## TODO
- [ ] Add Tests
- [ ] Add support for integers
- [ ] Check database validity
- [ ] Add Caching
- [ ] Add support for csv export
