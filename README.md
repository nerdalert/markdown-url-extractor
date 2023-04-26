# Markdown URL Extractor
This command-line program extracts the URLs of all markdown files in a GitHub repository to the raw githubusercontent URL

## Usage
To use the program, you'll need to have Go installed on your system. If you don't already have it, you can download it from the official website: https://golang.org/dl/.

Once you have Go installed, you can clone the program's source code from GitHub:

```
git clone https://github.com/nerdalert/markdown-url-extractor.git
```

Change into the program's directory:

```
cd markdown-url-extractor
```

Then, build the program using the following command:

```
go build
```

This will create a binary executable file in the same directory as the source code.

You can then run the program with the following command:

```
./markdown-url-extractor --repo-url <repository URL>
```

Replace <repository URL> with the URL of the GitHub repository you want to extract markdown URLs from.

For example, to extract the URLs from the nexodus-io/nexodus repository, you would run the following command:

```
./markdown-url-extractor --repo-url https://github.com/nexodus-io/nexodus.git
```

The program will then clone the repository (if it doesn't already exist), extract the URLs of all markdown files in the repository, and print them to the console. If the repo is already cloned, it will use the existing directory.

## License
This program is licensed under the MIT License. See the LICENSE file for more information.
