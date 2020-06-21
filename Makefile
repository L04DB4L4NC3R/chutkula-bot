.PHONY: build-android

# Note that in the .env, remove the +srv for it to run on aarch64 linux
build-android:
	# Add the path of your NDK_TOOLCHAIN here
	CC=~/build-deps/ndk-toolchain/bin/aarch64-linux-android30-clang \
	GO111MODULES=on \
	GOOS=android \
	GOARCH=arm64 \
	GOARM=7 \
	CGO_ENABLED=1 \
	go build -v -o ./bin/chutkulabot-android -x .

.PHONY: podman-build
podman-build:
	podman image build -t chutkulabot .

.PHONY: docker-build
docker-build:
	docker image build -f Containerfile -t chutkulabot .

.PHONY: build
build:
	go build -o ./bin/chutkulabot .
