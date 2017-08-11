# Crude (but cryptographic secure) IOTA Wallet Seed Generator

<center>
    ![example screenshot](https://image.ibb.co/ccGQjv/Bildschirmfoto_20170811175141_2022x943.png)
</center>

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
