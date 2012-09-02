#!/usr/bin/env bash

echo "got: js-beauty ?"

js-beautify > /dev/null 2> /dev/null

if [[ $? -ne 0 ]]
then
	echo "---> nope."
	echo
	echo "Can't find js-beautify"
	echo "it can be found here: "
	echo "https://github.com/einars/js-beautify"
	echo 
	echo "or clone it:"
	echo "git clone https://github.com/einars/js-beautify.git"
	exit 1
else
	echo "---> yep."
fi
