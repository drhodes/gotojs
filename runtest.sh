HERE=`pwd`
for d in $(ls ./test); do
	cd ./test/$d;
	pwd
	go clean 
	go build -v
	cd $HERE
done;