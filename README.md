GoFM
=====
GoFM is a cross-platform real-time audio streaming server written in Go. It allows you to stream MP3 audio files from a specified directory over HTTP and synchronize playback across multiple clients.

[![Go Version](https://img.shields.io/badge/Go-v1.16-blue)](https://golang.org/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/pxgo/go-fm)](https://goreportcard.com/report/github.com/pxgo/go-fm)
[![Downloads](https://img.shields.io/github/downloads/pxgo/go-fm/total)](https://github.com/pxgo/go-fm/releases)
[![References](https://img.shields.io/github/forks/pxgo/go-fm?label=references)](https://github.com/pxgo/go-fm/network/members)
[![License](https://img.shields.io/github/license/pxgo/go-fm)](https://github.com/pxgo/go-fm/blob/main/LICENSE)

<div style="margin:1.5rem 0;border:1px solid #e9e9e9;border-radius:6px;overflow:hidden;">
  <iframe title="fm.stackstore.net" src="https://fm.stackstore.net" style="width:100%;height:350px;border:0;display:block;" loading="lazy"></iframe>
</div>

<p style="text-align:center;"><a href="https://fm.stackstore.net" target="_blank" rel="noopener">fm.stackstore.net</a></p>

## Usage

To use GoFM, download the latest release from the [Releases page](https://github.com/pxgo/go-fm/releases) and run the server with the following command:

```
./GoFM -d /path/to/your/music/directory
```

By default, GoFM listens on port 8090. You can customize the server's behavior using the following command-line options:

```
-p int
    Specifies the server port number (default 8090).
-host string
    Specifies the server host address (default "0.0.0.0").
-r
    Enables random playback mode.
-debug
    Enables debug mode for the server.
-d string
    Specifies the directory to play audio files from (default "/path/to/your/music/directory").
-h
    Shows help information.
-n string
    Specifies the name of the FM (default "GoFM").
```

For example, to change the server's port number to 8080, use the -p option followed by the desired port number, like this:
```
./GoFM -d /path/to/your/music/directory -p 8080
```

You can find more information about GoFM on [STACKSTORE](https://stackstore.net/GoFM).

## License

GoFM is released under the [MIT License](https://github.com/pxgo/go-fm/blob/main/LICENSE). Feel free to use, modify, and distribute the software. Contributions are welcome!
