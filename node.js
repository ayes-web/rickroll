var t=(
process.platform == 'win32'? 'start':
process.platform == 'darwin'? 'open':
'xdg-open');
require('child_process').exec(t + " https://files.catbox.moe/ok50ut.mp4");
