#!/bin/bash 

#::ARTR::title=Sample task with file transfer
#::ARTR::result-type=file 
#::ARTR::file-name=./testing-tasks/sample-file.txt

lorem_ipsum="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

echo "$lorem_ipsum" > ./testing-tasks/sample-file.txt
