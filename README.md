# HttpFromAndroid

## build

### in windows

```bash
set GOOS=linux&&set GOARCH=arm64&&go build
```

### in linux / macos

```bash
GOOS=linux GOARCH=arm64 go build
```

## usage

```bash
# PUSH
adb push HttpFromAndroid /data/local/tmp

# PERMISSION
adb shell chmod 777 /data/local/tmp/HttpFromAndroid

# USAGE
adb shell
cd /data/local/tmp
HttpFromAndroid -method GET -url http://www.baidu.com
```
