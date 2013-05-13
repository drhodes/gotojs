var lib = require('../../lib.js');

const  nByte = 65;
const  ldigits = "0123456789abcdef";
const  udigits = "0123456789ABCDEF";
const  signed = true;
const  unsigned = false;
var  padZeroBytes = make(array-type-show, nByte);
var  padSpaceBytes = make(array-type-show, nByte);
var  newline = ['\n'];
var init = function  () {
for (var i  =  0; i<nByte; i++){
var padZeroBytes[i]  =  '0';
var padSpaceBytes[i]  =  ' ';
}
}
function fmt(intbuf, buf, wid, prec, widPresent, precPresent, minus, plus, sharp, space, unicode, uniQuote, zero){this.intbuf=intbuf;
this.buf=buf;
this.wid=wid;
this.prec=prec;
this.widPresent=widPresent;
this.precPresent=precPresent;
this.minus=minus;
this.plus=plus;
this.sharp=sharp;
this.space=space;
this.unicode=unicode;
this.uniQuote=uniQuote;
this.zero=zero;}
fmt.prototype.clearflags = function () {
this.wid  =  0;
this.widPresent  =  false;
this.prec  =  0;
this.precPresent  =  false;
this.minus  =  false;
this.plus  =  false;
this.sharp  =  false;
this.space  =  false;
this.unicode  =  false;
this.uniQuote  =  false;
this.zero  =  false;
} 
fmt.prototype.init = function (buf) {
this.buf  =  buf;
this.clearflags();
} 
fmt.prototype.computePadding = function (width) {
var left  =  !this.minus;
var w  =  this.wid;
if (w<0) {
var left  =  false;
var w  =  -w;
}
var w  -=  width;
if (w>0) {
if (left&&this.zero) {
return [padZeroBytes,w,0];
}
if (left) {
return [padSpaceBytes,w,0];
}else{
return [padSpaceBytes,0,w];
}
}
return ;
} 
fmt.prototype.writePadding = function (n, padding) {
for (; n>0; ){
var m  =  n;
if (m>nByte) {
var m  =  nByte;
}
this.buf.Write(padding.slice(0,m));
var n  -=  m;
}
} 
fmt.prototype.pad = function (b) {
if (!this.widPresent||this.wid==0) {
this.buf.Write(b);
return ;
}
AssignStmt missing case: {[padding left right] 2126 := [0xc20008c700]};
if (left>0) {
this.writePadding(left, padding);
}
this.buf.Write(b);
if (right>0) {
this.writePadding(right, padding);
}
} 
fmt.prototype.padString = function (s) {
if (!this.widPresent||this.wid==0) {
this.buf.WriteString(s);
return ;
}
AssignStmt missing case: {[padding left right] 2500 := [0xc20008cc80]};
if (left>0) {
this.writePadding(left, padding);
}
this.buf.WriteString(s);
if (right>0) {
this.writePadding(right, padding);
}
} 
var putint = function  (buf, base, val, digits) {
var i  =  len(buf)-1;
for (; val>=base; ){
var buf[i]  =  digits[val%base];
i--;
var val  /=  base;
}
var buf[i]  =  digits[val];
return i-1;
}
var  trueBytes = array-type-show("true");
var  falseBytes = array-type-show("false");
fmt.prototype.fmt_boolean = function (v) {
if (v) {
this.pad(trueBytes);
}else{
this.pad(falseBytes);
}
} 
fmt.prototype.integer = function (a, base, signedness, digits) {
if (this.precPresent&&this.prec==0&&a==0) {
return ;
}
DeclStmt not implemented
var negative  =  signedness==signed&&a<0;
if (negative) {
var a  =  -a;
}
var prec  =  0;
if (this.precPresent) {
var prec  =  this.prec;
this.zero  =  false;
}elseif (this.zero&&this.widPresent&&!this.minus&&this.wid>0) {
var prec  =  this.wid;
if (negative||this.plus||this.space) {
prec--;
}
}
var i  =  len(this.intbuf);
var ua  =  uint64(a);
for (; ua>=base; ){
i--;
var buf[i]  =  digits[ua%base];
var ua  /=  base;
}
i--;
var buf[i]  =  digits[ua];
for (; i>0&&prec>nByte-i; ){
i--;
var buf[i]  =  '0';
}
if (this.sharp) {
switch (base) {
case 8: if (buf[i]!='0') {
i--;
var buf[i]  =  '0';
}break;
case 16: i--;
var buf[i]  =  'x'+digits[10]-'a';
i--;
var buf[i]  =  '0';break;
}
}
if (this.unicode) {
i--;
var buf[i]  =  '+';
i--;
var buf[i]  =  'U';
}
if (negative) {
i--;
var buf[i]  =  '-';
}elseif (this.plus) {
i--;
var buf[i]  =  '+';
}elseif (this.space) {
i--;
var buf[i]  =  ' ';
}
if (this.unicode&&this.uniQuote&&a>=0&&a<=utf8.MaxRune&&strconv.IsPrint(rune(a))) {
var runeWidth  =  utf8.RuneLen(rune(a));
var width  =  1+1+runeWidth+1;
copy(buf.slice(i-width,buf.length), buf.slice(i,buf.length));
var i  -=  width;
var j  =  len(buf)-width;
var buf[j]  =  ' ';
j++;
var buf[j]  =  '\'';
j++;
utf8.EncodeRune(buf.slice(j,buf.length), rune(a));
var j  +=  runeWidth;
var buf[j]  =  '\'';
}
this.pad(buf.slice(i,buf.length));
} 
var array-type-show buf = this.intbuf.slice(0,this.intbuf.length);
fmt.prototype.truncate = function (s) {
if (this.precPresent&&this.prec<utf8.RuneCountInString(s)) {
var n  =  this.prec;
for (var i in s) { 
 var %!s(<nil>) = s[i];
if (n==0) {
var s  =  s.slice(0,i);
break;
}
n--;
}
}
return s;
} 
fmt.prototype.fmt_s = function (s) {
var s  =  this.truncate(s);
this.padString(s);
} 
fmt.prototype.fmt_sx = function (s, digits) {
DeclStmt not implemented
for (var i  =  0; i<len(s); i++){
if (i>0&&this.space) {
var b  =  lib.append(b, ' ');
}
var v  =  s[i];
var b  =  lib.append(b, digits[v>>4], digits[v&0xF]);
}
this.pad(b);
} 
