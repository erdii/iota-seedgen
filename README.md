# Crude (but cryptographic secure) IOTA Wallet Seed Generator

<p align="center">
    <img src="https://image.ibb.co/ccGQjv/Bildschirmfoto_20170811175141_2022x943.png" alt="example screenshot" />
</p>

# About

Since there is a shitload of "secure seed generators" out there, that use insecure random number generators like `Math.random()` in javascript or `Get-Random` in powershell I wanted to write a **SECURE** seed generator. [This guy's](https://www.reddit.com/r/Iota/comments/6vh8mv/urgent_all_of_my_iota_was_stolen_today_and_here/) iota have been stolen because he generated the seed in an insecure way.

I did not want to create a javascript in-browser seed generator because a hacked webserver can serve you backdoored javascript code that would generate insecure seeds or send the generated seeds back home.

That's why I made a **downloadable** version of a seed generator, that uses cryptographical secure operating system random number sources ([details here](https://golang.org/pkg/crypto/rand/#pkg-variables)) and is reasonably simple to use.

# Download (Windows, Mac, Linux)

If you **TRUST ME** and just want to get the seed real fast download the programm [here](https://blog.werise.de/files/iota-seedgen.zip) ([signature](https://blog.werise.de/files/iota-seedgen.zip.asc) - [my keybase profile](https://keybase.io/erdii)).
Read [how to verify a pgp signature written by gnupg.org](https://www.gnupg.org/gph/en/manual/x135.html#AEN160) or [how to verify a pgp signature written by torproject.org](https://www.torproject.org/docs/verifying-signatures.html.en#maincol) to verify the signature.


# Usage

* windows:
  * double click on the executable for your cpu (32 or 64 bit)
  * select one generated seed with your mouse and copy it via right-clicking
* linux + mac:
  * open a terminal
  * cd into the folder where the seedgen executables reside
  * mac 64bit: `./iota-seedgen_mac-64bit`
  * mac 32bit: `./iota-seedgen_mac-32bit`
  * linux 64bit: `./iota-seedgen_linux-64bit`
  * linux 32bit: `./iota-seedgen_linux-32bit`

### Checksums for `https://blog.werise.de/files/iota-seedgen.zip`

* md5: `cd0367666a2b6b221b4bbe741aa2bff1`
* sha1: `cd2ae82bb62a0078c3723f68c615514b51506f32`
* sha256: `ff6b5b91526686295edaa62f168b951b7962fcfb5e07ad9fd656b97f3e634f1c`
* sha512: `ba2ed19c006b38415fa8c887b0776c87e87b5ef40a362506f366b078b84ffbc7ac6556d40b9cda3b790ecd43ca064a6e26a07317fa1a4994f42dbe56b3fdbf06`


# Building it yourself (advanced users)

* requirements: go, keybase
* build with `make all`
* create a signed release with `make release` (this needs keybase installed)
* execute the binary for your operating system + cpu architecture
    * either by opening a terminal/cmd, navigation to the extracted zip folder and executing 
        * windows 32bit: `iota-seedgen_win-32bit.exe`
        * windows 64bit: `iota-seedgen_win-64bit.exe`
        * linux 32bit: `iota-seedgen_linux-32bit.exe`
        * linux 64bit: `iota-seedgen_linux-64bit.exe`
        * mac 32bit: `iota-seedgen_mac-32bit.exe`
        * mac 64bit: `iota-seedgen_mac-64bit.exe`
    * or by doubleclicking on the executable file


# License

Copyright (c) 2017 erdii <erdiicodes@gmail.com>  
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:  
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.  
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
