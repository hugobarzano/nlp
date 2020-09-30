 setup:
	chmod +x bin/*
build:
	chmod +x bin/_buildImage.sh
	bin/_buildImage.sh
test:
	chmod +x bin/_test.sh
	bin/_test.sh
run:
	chmod +x bin/_run.sh
	bin/_run.sh
