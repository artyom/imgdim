imgdim command returns image dimensions in format "width,height". Created as a
lightweight replacement for `identify -format "%w,%h"` which can be too heavy
on big files (especially on big GIF files).

	Usage: imgdim filename.{jpeg,png,gif,bmp,tiff}

LICENSE: [MIT](http://opensource.org/licenses/MIT).
