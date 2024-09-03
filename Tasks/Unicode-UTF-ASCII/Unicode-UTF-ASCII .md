#### ASCII
**ASCII** or **A**merican **S**tandard **C**ode for **I**nformation **I**nterchange is a character encoding for electronic communication. ASCII codes represent text in computers, telecommunications equipment and other devices.
#### Unicode
Unicode is an international encoding standard. That is used with different languages and scripts, by which each letter, digit or symbol is assigned a unique numeric value that applies across different platforms and programs.
#### UTF
**U**nicode **T**ransformation **F**ormat, short **UTF**, is designed to encode characters from all the world's writing systems in a standardized way.
#### Non printable / control characters
**Control charcters** or **non-printing characters** (**NPC**) is a code point in a character set that does not represent a written character or symbol. They are used as in-band signaling to cause effects other than the addition of a symbol to the tet. All other characters are mainly graphic characters, also known as printing characters (or printable characters), except for space characters. In the ASCII standard there are 33 control characters such as code 7, BEL, which rings a terminal bell.

##### Examples
| ASCII | Written<br>(Code) | Function                                                                                                               |
| ----- | ----------------- | ---------------------------------------------------------------------------------------------------------------------- |
| 0x00  | \0                | mark the end of a string                                                                                               |
| 0x07  | \a                | may cause the device to emit a warning<br>(Bell or beep sound or screen flashing)                                      |
| 0x08  | \b                | overprint previous character                                                                                           |
| 0x09  | \t                | moves the printing position right to <br>the next tab stop                                                             |
| 0x0A  | \n                | moves the print head down one line, <br>or to the left edgeand down.<br>(often used as an end of line marker)          |
| 0x0B  | \v                | vertical tabulation                                                                                                    |
| 0x0C  | \f                | to cause a printer to eject paper to the top of <br>the next page, or a video terminal to clear the screen.            |
| 0x0D  | \r                | moves the printing position to the start of the line, <br>allowing overprinting. (often used as an end of line marker) |
| 0x1A  |                   | acts as an end of file for the Windows text-mode file i/o                                                              |
| 0x1B  | \e                | introduces an escape sequence                                                                                          |
#### Different formats of UTF
There are 3 different formats of UTF there are UTF-8, UTF-16 and UTF-32.
Each character will then be represented either as a sequence of one to four 8-bit bytes (UTF-8), one or two 16-bit code units (UTF-16) or a single 32-bit code unit (UTF32).
#### Why everyone uses UTF-8
Because UTF-8 is the one that results in fewer internationalization issues than any alternate text encoding.
#### How emoji's work
Emoji's are built by a sequence of Unicode code points. The most basic emojis are single Unicode code points.
Emoji's look different on different operating systems because the emoji standard defines what code point sequences stands for what emoji's, but it doesnt dictate how they should look. Every operating system vendor creates their own images for these emoji characters and puts them all into the system font as glyphs.
#### Zero Width Joiner
The zero-width joiner is a non-printing character used to join two or more characters together in sequence to create a new emoji.
#### How colors work for emoji's
Unicode 8.0 added support for skin tone variations on human emojis. This is supported via a set of five modifiers based on the [Fizpatrick scale](https://en.wikipedia.org/wiki/Fitzpatrick_scale) for measuring human skin color. These modifiers are also Unicode code points, starting at `U+1F3FB` and ending at `U+1F3FF`

#### URL encoding
Here is an example URL and a version of this URL split in its parts.
https://maxmuster:geheim@www.example.com:8080/index.html?p1=A&p2=B#resource

| Schema | User      | Password | Host            | Port | Path        | Query     | Fragment |
| ------ | --------- | -------- | --------------- | ---- | ----------- | --------- | -------- |
| https  | maxmuster | geheim   | www.example.com | 8080 | /index.html | p1=A&p2=B | resource |
#### Reserved Characters Substitution
These characters have a use in the unix system

| Character  | Use                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| ---------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| *          | wildcard for anything. There are no limits on things like the number of characters or what characters (aside from so called special characters and metacharacters)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| ?          | wildcard for a single character. There are no limits on what characters.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [ ]        | delimiter for a range of single character wild cards. For example if you only wanted to list images from a numeric series where the files of interest were inthe range 140 thru 179 you could use `$ ls -l myImages_001[4567]?.dm3` or `$ ls -l myImages_001[4-7]?.dm3`. The delimiter only replaces this one character nothing else.                                                                                                                                                                                                                                                                                                                                                                       |
| \| (pipe)  | this (\|) symbol is called **pipe**. Its used for pipelineing commands, meaning that the output of one command becomes the input to the next command<br>`$ cmd1 \| cmd2 \| cmd3 \| cmd4` etc.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| ;          | the semi colon (:) is a "command separator" and it allows you to string a series of commands onto a single command line. This is quite different from the pipe in that the commands strung together do not depend on each other for inputs as an example, when you are actively working in one area of the computer, but want to move someplace else and look at how much has been accomplished in that area, you could put together a series of commands that would do this and return to the starting location:<br>`$ cd $HOME/ProjectX; ls -ltr; cd -`                                                                                                                                                   |
| &          | the ampersand (&) cases a command to run in the shell's "background". Meaning that the user continues to have access to the interactive shell.<br>For example, if you start a graphical text editor from the command line without using the &, the editor will run but your terminal wil be inactive until you exit the editor. Instead the editor should be started and run in the background using this syntax:<br>`$ nedit myTextFile.txt &`                                                                                                                                                                                                                                                             |
| >,>> and < | These symbols are called redirection operators and tthey allow the user to manipulate where stdin, stdout, stderr (names for the linux standard input/output channels) end up. For example the stdout output from a command can be directed into anew file using `$ cmd [options] [target] > newFile`<br>or into an existing file using `$ cmd [options] [targer] >> oldFile`<br>In this second example, it doesn't matter if oldFile exists or not: the ">>" redirection command wil create the file if necessary but append to an existing file. If the ">" redirection command is used, the file will always be new and an existing file with the same name will be over-written (and lost)              |
| " "        | a pair of double quotes (the " symbol encloses an expression where you want to see everything exactly as it is written except for any parts that involve a \$, a back quote(´) or a backslash(\\))<br>Example command:<br>`$ echo "My home is ${HOME}"`<br>Output:<br>My home is /N/u/myUserName/Quarry                                                                                                                                                                                                                                                                                                                                                                                                     |
| ' '        | a pair of single quotes (the ' symbol) encloses an expression where you want to see everything exactly the way it is written. In this case the command:<br>`$ echo 'My home is ${HOME}'` outputs `My home is ${HOME}`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `          | back quotes (the \` symbol) causes something called command substitution which means that a command placed between a pair of back quotes is executed and the output of the command replaces everything from the first \` to the second \`. Back quotes are protected by single quotes (command substition does not occur)<br>Example:<br>command: $ echo "Today is \` date \`"<br>outputs: My home is Thu May 16 19:47:27 EDT 201<br>command: $ echo 'Today is \` date \`' outputs: My home is \` date \`                                                                                                                                                                                                   |
| \          | the \\ (Backslash) has a number of uses, only some of which will be described here. One major use is when quoting the \$, \` and " characters. In this contect the character which follows \\ is taken literally. This means it will be shown as a normal character and doesnt have any special use.<br>Example:<br>command: $ echo "Today is \` date \`"<br>outputs: My home is Thu May 16 19:47:27 EDT 201<br>command: $ echo 'Today is \\\` date \\\`' outputs: My home is \\\` date \\\`<br>                                                                                                                                                                                                            |
| !          | the ! (exclamation mark, often reffered to as a bang) has several special uses when used as a logical operator. The ! means "not".<br>Example:<br>`$ ls -l myImages_001[!4-7]?.dm3`<br><br>the ! (exclamation mark) also can get the last information that is needed for the new command from the old command.<br>Example<br>`!ec`<br>outputs the last command that began with "ec".<br><br>Lastly there is #! (pronounced "hash-bang" or "shebang"). It is a special symbol pair that is used as the first line of a shell script to invoke a paricular shell.<br>For example, if you want to run the bash shell. Start the script using #!/bin/bash                                                       |
| #          | the # symbol is often called "hash or "pound" and the use of hash-bang (#!) was described above. The # is most commonly used to signify a comment in programming and scripting languages.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| /          | the / (forward slash) symbol is used to delineate and separate directory names when describing a location (e.g. cd $HOME/ProjectX/Images/2024_09_02)<br>when used alone (e.g. cd /), the / symbol indicates "the root of the filesystem" (i.e., the place where the full pathname to all files and directories starts)                                                                                                                                                                                                                                                                                                                                                                                      |
| ~          | the most common use and meaning for the ~ symbol (the "tilde") is simply "your home directory" and (for example) cd ~ is one of many ways to move from anywhere to your home. You will also sometimes see the ~ appended to the end of file name where the file has been edited using the emacs text editor. The appended ~ signifies that a file is the backup copy of an edited file. Another use for the ~ is in conjunction with a user's name on the computer system (e.g. ~userName). In this case ~userName indicates the home directory for a user called userName. Since the IU computer clusters will not allow you to look at another user's home area, this is not so useful on those machines. |
#### Escape / Unescape URLS with url package
When URLs escape the URL Query will be [percent-encoded](https://en.wikipedia.org/wiki/Percent-encoding). And when the URL Unescapes it will be [percent-decoded](https://en.wikipedia.org/wiki/Percent-encoding).
[ASCII Encoding Reference](https://www.w3schools.com/tags/ref_urlencode.ASP)

**Example:** Escape
```
package main

import (
	"fmt"
	"net/url"
)

func main() {
	query := url.QueryEscape("my/cool+blog&about,stuff")
	fmt.Println(query)

} 
```
---
```
Output:
my%2Fcool%2Bblog%26about%2Cstuff
```

**Example:** Unescape
```
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	escapedQuery := "my%2Fcool%2Bblog%26about%2Cstuff"
	query, err := url.QueryUnescape(escapedQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(query)

}
```
---
```
Output:
my/cool+blog&about,stuff
```
