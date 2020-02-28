# About
A simple cli tool for converting subtitle files into csv files
# Instalation
`go install github.com/DarylSerrano/subtitles2csv`
# Usage
```
$ subtitles2csv --infile "subtitles.srt"

Saved into:  /home/daryl/go/src/subtitles2csv/out.csv
```
You can also set the outputh path:
```
$ subtitles2csv --infile "subtitles.srt" --outpath "../"

Saved into:  /home/daryl/go/src/out.csv
```
# License
[MIT](https://github.com/DarylSerrano/subtitles2csv/blob/master/LICENSE)