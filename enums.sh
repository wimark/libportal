#!/bin/sh

for i in `find ./ -name "*.toml"`
do 
	cat $i | enum-generator -enable-json -enable-bson -package=libportal | gofmt > "${i%.*}".go
done
