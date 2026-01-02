# First Layer Strategies

 These ciphers could be manipulated in unexpected ways. For example, one may be obfuscated by encoding binary data that was then shifted, rotated, or similar.

 ## Binary Analysis

 > **Note:** Everything in this section has a first step of decoding the data to its binary representation based on its encoding. Everything should be attempted on the decoded binary of both the normal and reversed encoded string.

 ## Transformations

 | Transformation |
 | --- |
 | **Special transformation**: bitstream reversal. All below transformations should be attempted on the forward *and* reversed bitstream. |
 | Unshifted binary representation. |
 | The entire binary sequence is shifted 1-7 times in unison (rather than per byte) |
 | Bytes are shifted m times. |
 | Bytes are rotated m times. |

 ## Post-transformation analysis

 | Technique |
 | --- |
 | n-gram frequency analysis for varying size n |
 | longest repeating sequence |
 | most repeated n-gram |
 | n-gram index of coincidence |