# File Converter CLI (Image and Text)

This is a command-line application written in Go that allows you to convert image and text files between supported formats. It supports converting JPEG, JPG, and PNG images, as well as text-based formats like `.txt`, `.odt`, and `.docx` (implementation required).

---

## 🚀 Features

- ✅ Convert image files between `.jpg`, `.jpeg`, and `.png`
- ✅ Convert text files between formats (e.g., `.txt`, `.odt`, `.docx`)
- ✅ Easy-to-use CLI interface
- ❗ Extensible with your own `ConvertImg` and `ConvertTxt` implementations

---

## 🧩 Project Structure

```
|   .gitignore
|   cmd.go
|   converter.go
|   go.mod
|   go.sum
|   main.go
|   README.md
|   
\---txtFiles
        docxHandler.go
        odtHandler.go
        txtHandler.go
```

---

This file is the entry point and handles parsing CLI flags, validating inputs, and calling the correct conversion logic.

---

## ⚙️ Usage

Build the app:

```bash
go build -o file-converter
```

## Convert Images
```bash
./file-converter -img input.jpg:output.png
```

## Convert Text Files
```bash
./file-converter -txt input.txt:output.odt
```