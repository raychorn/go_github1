#!/bin/bash

if [ -z "$1" ]
  then
    echo "No argument supplied"
    exit 1
fi

find * -size +50M -type f -print >> .gitignore
git add -A
git commit -m "Added some code."
git push -u origin main
