RUN_NAME="go_web_demo"

run:
	sh build.sh
	./output/$(RUN_NAME)

build:
	sh build.sh

clean:
	rm -rf output