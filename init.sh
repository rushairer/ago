#!/bin/bash
echo "Please enter a new package name:"
read name
grep -rl "PACKAGENAME" | xargs -I {} sed -i "" "s|PACKAGENAME|${name}|g" {}