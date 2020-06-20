.PHONY: build-android
build-android:
	# Add the path of your NDK_TOOLCHAIN here
	export NDK_TOOLCHAIN=~/build-deps/ndk-toolchain
	export CC=$NDK_TOOLCHAIN/bin/arm-linux-androideabi-gcc
	export GO111MODULES=on
	export GOOS=android
	export GOARCH=arm
	export GOARM=7
	export CGO_ENABLED=1
	go build -v -o ./bin/chutkulabot-android -x main.go

.PHONY: podman-build
podman-build:
	podman image build -t chutkulabot .

.PHONY: docker-build
docker-build:
	docker image build -f Containerfile -t chutkulabot .

.PHONY: build
build:
	go build -o ./bin/chutkulabot .
