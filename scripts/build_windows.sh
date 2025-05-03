#!/bin/bash 

outfile="./bin/artr.exe"

if $(GOOS=windows CGO_ENABLED=1 go build -installsuffix 'static' -o "$outfile" main.go); then 
  echo "build successfully"
  echo "$outfile"
fi
