# Crude (but cryptographic secure) IOTA Wallet Seed Generator

<p align="center">
    <img src="https://image.ibb.co/ccGQjv/Bildschirmfoto_20170811175141_2022x943.png" alt="example screenshot" />
</p>

# CAUTION

**!!!READ THIS!!!** Thanks to the amazing people in the iota slack - I discovered that the seed generator was biased towards some characters!!!
That means that the generated seeds were not as secure as they could be. You should download the new version of iota-seedgen and move your balance from the old seed to a new generated one!

# About

Since there is a shitload of "secure seed generators" out there, that use insecure random number generators like `Math.random()` in javascript or `Get-Random` in powershell I wanted to write a **SECURE** seed generator. [This guy's](https://www.reddit.com/r/Iota/comments/6vh8mv/urgent_all_of_my_iota_was_stolen_today_and_here/) iota have been stolen because he generated the seed in an insecure way.

I did not want to create a javascript in-browser seed generator because a hacked webserver can serve you backdoored javascript code that would generate insecure seeds or send the generated seeds back home.

That's why I made a **downloadable** version of a seed generator, that uses cryptographical secure operating system random number sources ([details here](https://golang.org/pkg/crypto/rand/#pkg-variables)) and is reasonably simple to use.
The best way to generate secure seeds is to download the generator, transfer it to an air-gapped computer and generate the tokens there.

# Download (Windows, Mac, Linux)

If you **TRUST ME** and just want to get the seed real fast download the programm
* [from github](https://github.com/erdii/iota-seedgen/releases/latest)
* [from my server](https://blog.werise.de/files/iota-seedgen/iota-seedgen_v0.1.0.zip)
* ([signature](https://blog.werise.de/files/iota-seedgen/iota-seedgen_v0.1.0.zip.asc) - [my keybase profile](https://keybase.io/erdii))

Read [how to verify a pgp signature written by gnupg.org](https://www.gnupg.org/gph/en/manual/x135.html#AEN160) or [how to verify a pgp signature written by torproject.org](https://www.torproject.org/docs/verifying-signatures.html.en#maincol) to verify the signature. my key id is **6682993C**


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

### Checksums
* `iota-seedgen_v0.1.0.zip`
    * md5: `ccc1ef56dfc1925a12a689de8421c51f`
    * sha1: `8983d3d991f6394b3846f54f8b2fa7bb8e2699c9`
    * sha256: `168d0db874d18e2dace0dbabdf6974be90a9cd88b65618795518c412e51f3e7b`
    * sha512: `f41b473fd38d44b10f847d4a89713bd2bce86ff74622efd133aaca38744eb0f8f4858fdf5a63f1e2a627ec7da874226799bcaa83bca37a911db2d7779156791a`
* `iota-seedgen_v0.0.2.zip`
    * md5: `b0cacd4bcff47ae415ed971678316889`
    * sha1: `b406c385716d151f5e5dc9b345ad70e01d7b7c42`
    * sha256: `9e028a052ab610f5992f10efe8649ca1f54b3bedd072adf7fa742f6a01640d16`
    * sha512: `19a822210d23de23f5a466d06200b7cbd8c749222bced6da07b9d6ca2f909750104888407e32f419aaeeed234934da2603372b8731eaac71d99de1ac0fb359a0`

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
