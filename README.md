# Golang crawler for wzry skin pictures

Use goroutine to concurrently and incrementall download
hero pictures from [wzry](pvp.qq.com).

## config

You can create a `config.ini` file in the same folder of the executable file,
with some customized configures. Such as

```ini
savepath = D:\test
```

Up till now it supports keywords

- `savepath`, where to save the pictures
