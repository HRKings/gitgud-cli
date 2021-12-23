#! /bin/bash
VERSION=$(git describe --long --tags stable | cut -d '-' -f 1)

echo "Removing old versions..."
rm -rf ./build

echo "Compiling for Linux..."
archs=("amd64" "arm64" "386")
for arch in "${archs[@]}"
do
    GOOS=linux GOARCH=${arch} go build -o "build/linux/gitgud"
    tar -C "build/linux" -czf "build/linux/gitgud_v${VERSION}_linux_${arch}.tar.gz" "gitgud"
    rm "build/linux/gitgud"
done

echo "Compiling for MacOS..."
archs=("amd64" "arm64")
for arch in "${archs[@]}"
do
    GOOS=darwin GOARCH=${arch} go build -o "build/macos/gitgud"
    tar -C "build/macos" -czf "build/macos/gitgud_v${VERSION}_macos_${arch}.tar.gz" "gitgud"
    rm "build/macos/gitgud"
done

echo "Compiling for Windows..."
archs=("amd64" "arm64" "386")
for arch in "${archs[@]}"
do
    GOOS=windows GOARCH=${arch} go build -o "build/windows/gitgud.exe"
    tar -C "build/windows" -czf "build/windows/gitgud_v${VERSION}_windows_${arch}.tar.gz" "gitgud.exe"
    rm "build/windows/gitgud.exe"
done