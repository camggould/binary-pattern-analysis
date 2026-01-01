# Binary Pattern Analysis

## Background

I have been working on several decade-old ciphers for the past year or so and wanted to organize analysis into one place. These ciphers are all known to be some form of classical encryption (possibly multiple layers) rather than modern encryption, but contain additional layers of obfuscation via tricky data encoding techniques. This extra layer has stumped analysts for a decade now, and I would like to take a crack at seriously analyzing them to possibly find a way out.

## The data

The data for these ciphers as they are presented can be found in the `data` folder. I have used phonetic alphabet conventions for naming them, alpha through juliett. Some of the data has a clear first obfuscation layer, such as decoding the appearing text from ASCII decimal, which produces a hexadecimal sequence.

### Character encoding techniques

An important primer on the encoding techniques leveraged as obfuscation here. Several forms of data encodings are being leveraged across these ciphers: octal, decimal, hexadecimal, and base64 are all used or have been used on other complex ciphers created by the same author.

Octal, decimal, hexadecimal, and base64 can all be used to encode ASCII or Latin1 text from an [ASCII/extended ASCII table](https://www.schoolcoders.com/data-representation/characters/extended-ascii/). The ASCII table maps numbers to symbols/letters. Those numbers can be converted out of decimal to other bases. Binary, octal, and hex are more common but base64 is also possible. Base64 can also be used to encode text directly.

Hexadecimal and Base64 are commonly utilized for encoding raw binary data, as a way to compress its representation to fewer characters.