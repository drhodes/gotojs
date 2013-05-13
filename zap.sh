test -e .zapsafe
if [[ $? -ne 0 ]]
then
	echo "You're trying to run zap.sh in a directory without .zapsafe"
	echo "that could be dangerous."
	exit 1
fi

find | grep -v "\.git" | grep ~$
trash `find | grep -v "\.git" | grep ~$` 2> /dev/null
trash ./gotojs 2> /dev/null

for pkg in $(ls ./test/pass) 
do
	echo cleaning ./test/pass/$pkg
	go clean ./test/pass/$pkg
	rm -f ./test/pass/$pkg/main.test
done;


for pkg in $(ls ./test/run) 
do
	echo cleaning ./test/run/$pkg
	cd ./test/run/$pkg && make clean && cd - > /dev/null
done;

rm -f ./bin/gotojs

