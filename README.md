# Natural Language Git CLI assistant

Interfacing with git sucks. With gitgpt you can use natural langauge instead of git commands to do what you want. 

Example:
```
$ ./gitgptv2 add .gitignore commit with msg adding ignore and push
Would you like to run the following command:

git add .gitignore
git commit -m "adding ignore"
git push [y/N] y
[main c70490c] adding ignore
 1 file changed, 3 insertions(+)
 create mode 100644 .gitignore
Enumerating objects: 4, done.
Counting objects: 100% (4/4), done.
Delta compression using up to 12 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (3/3), 312 bytes | 312.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/Hesse/gitgpt.git
   3a5828e..c70490c  main -> main
```



## Requirements

- You must have an OpenAI API key. This key should be set as an environment vairable OPENAI_API_KEY.

```
export OPENAI_API_KEY=<yourkey>
```

- You must have go installed in order to build the source OR you can download the pre-built binary

I have only tested this on Mac OS, however I'm pretty sure it'll work without issue on Linux as well. I can't say the same about Windows because I haven't tested it. 


## Build / Installation

### Build

1. Clone the repo
2. go build -o dist/gitgpt gitgpt.go
3. Add the file to your PATH
```
echo 'export PATH=$PATH:<path_to_repository/dist>' >> ~/.bash_profile

```


