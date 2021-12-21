#! /bin/bash

echo "Compiling for Linux..."
archs=("amd64" "arm64" "386")
for arch in "${archs[@]}"
do
    GOOS=linux GOARCH=${arch} go build -o "build/linux/gitgud_linux_${arch}"
done

echo "Compiling for MacOS..."
archs=("amd64" "arm64")
for arch in "${archs[@]}"
do
    GOOS=darwin GOARCH=${arch} go build -o "build/mac/gitgud_macos_${arch}"
done

echo "Compiling for Windows..."
archs=("amd64" "arm64" "386")
for arch in "${archs[@]}"
do
    GOOS=windows GOARCH=${arch} go build -o "build/windows/gitgud_windows_${arch}.exe"
done