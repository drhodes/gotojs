#!/usr/bin/env bash

echo "got: js-beauty ?"


command -v js-beautify >/dev/null 2>&1 || { 
	echo "---> nope."
	echo
	echo "Can't find js-beautify"
	echo "it can be found here: "
	echo "https://github.com/einars/js-beautify"
	echo 
	echo "or install" 
	echo "pip install jsbeautifier"

	echo >&2 "I require js-beatify but it's not installed.  Aborting."; exit 1; 
}

echo "---> yep."
