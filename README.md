# KILO - GO

## Description

This is my attempt to recreate the `KILO` editor in `Go`.

Reference to this project: [Kilo Editor](https://viewsourcecode.org/snaptoken/kilo/index.html)

We will go through all the steps, from entering raw mode to editing text with syntax highlighting.

## Installation

This project currently will work only in `Linux`, if you are using `Windows` or `MacOS` you will have to make some changes in the code, feel free to create a pull request.

To install this project run the following command:

```bash
git clone https://github.com/alcb1310/kilo-go.git
cd gokilo
go build
./kilo-go
```

## Tutorial

The process of creating this project is being uploaded to [Dev.to](https://dev.to/alcb1310) in a series of parts:

- [Setup](https://dev.to/alcb1310/create-a-text-editor-with-go-setup-58ej)
- [Raw Mode](https://dev.to/alcb1310/create-a-text-editor-with-go-enter-raw-mode-5g2n)
- [Welcome Screen](https://dev.to/alcb1310/create-a-text-editor-with-go-welcome-screen-4hkm)
- [Moving the Cursor](https://dev.to/alcb1310/create-a-text-editor-in-go-moving-the-cursor-2bnk)
- [Text Viewer](https://dev.to/alcb1310/create-a-text-editor-in-go-text-viewer-4akp)
- [Status Bar](https://dev.to/alcb1310/create-a-text-editor-in-go-status-bar-2047)
- [Text Editor](https://dev.to/alcb1310/create-a-text-editor-in-go-a-text-editor-1m83)
- [Search](https://dev.to/alcb1310/create-a-text-editor-in-go-search-13l7)

## Deguging

To compile in debug mode:

```bash
sudo sysctl -w kernel.yama.ptrace_scope=0
go build -gcflags="-N -l" -o kilo
```

After it is built you run the program and then you can debug it with a debuger

## License

This project is distributed under the [GNU GPL V3 license](https://github.com/alcb1310/kilo-go/blob/main/LICENSE).

## Author

- [Andrés Court](https://x.com/alcb1310)
