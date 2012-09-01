// package fmt
var init = function  () {
for (var i  =  0; i<nByte; i++){
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '0';
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  ' ';
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
f.wid  =  0;
f.widPresent  =  false;
f.prec  =  0;
f.precPresent  =  false;
f.minus  =  false;
f.plus  =  false;
f.sharp  =  false;
f.space  =  false;
f.unicode  =  false;
f.uniQuote  =  false;
f.zero  =  false;
}
fmt.prototype.init = function (buf) {
f.buf  =  buf;
f.clearflags();
}
fmt.prototype.computePadding = function (width) {
var left  =  unhandled Expr in func ShowExpr: *ast.UnaryExpr;
var w  =  f.wid;
if (w<0) {
var left  =  false;
var w  =  unhandled Expr in func ShowExpr: *ast.UnaryExpr;
}
var w  -=  width;
if (w>0) {
if (left&&f.zero) {
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
f.buf.Write(unhandled Expr in func ShowExpr: *ast.SliceExpr);
var n  -=  m;
}
}
fmt.prototype.pad = function (b) {
if (unhandled Expr in func ShowExpr: *ast.UnaryExpr||f.wid==0) {
f.buf.Write(b);
return ;
}
AssignStmt missing case: {[padding left right] 2126 := [0x18761dc0]};
if (left>0) {
f.writePadding(left, padding);
}
f.buf.Write(b);
if (right>0) {
f.writePadding(right, padding);
}
}
fmt.prototype.padString = function (s) {
if (unhandled Expr in func ShowExpr: *ast.UnaryExpr||f.wid==0) {
f.buf.WriteString(s);
return ;
}
AssignStmt missing case: {[padding left right] 2500 := [0x18769220]};
if (left>0) {
f.writePadding(left, padding);
}
f.buf.WriteString(s);
if (right>0) {
f.writePadding(right, padding);
}
}
var putint = function  (buf, base, val, digits) {
var i  =  len(buf)-1;
for (; val>=base; ){
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  unhandled Expr in func ShowExpr: *ast.IndexExpr;
i--;
var val  /=  base;
}
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  unhandled Expr in func ShowExpr: *ast.IndexExpr;
return i-1;
}
fmt.prototype.fmt_boolean = function (v) {
if (v) {
f.pad(trueBytes);
}else{
f.pad(falseBytes);
}
}
fmt.prototype.integer = function (a, base, signedness, digits) {
if (f.precPresent&&f.prec==0&&a==0) {
return ;
}
unhandled Stmt in func ShowStmt: *ast.DeclStmt
var negative  =  signedness==signed&&a<0;
if (negative) {
var a  =  unhandled Expr in func ShowExpr: *ast.UnaryExpr;
}
var prec  =  0;
if (f.precPresent) {
var prec  =  f.prec;
f.zero  =  false;
}elseif (f.zero&&f.widPresent&&unhandled Expr in func ShowExpr: *ast.UnaryExpr&&f.wid>0) {
var prec  =  f.wid;
if (negative||f.plus||f.space) {
prec--;
}
}
var i  =  len(f.intbuf);
var ua  =  uint64(a);
for (; ua>=base; ){
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  unhandled Expr in func ShowExpr: *ast.IndexExpr;
var ua  /=  base;
}
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  unhandled Expr in func ShowExpr: *ast.IndexExpr;
for (; i>0&&prec>nByte-i; ){
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '0';
}
if (f.sharp) {
switch (base) {
case 8: if (unhandled Expr in func ShowExpr: *ast.IndexExpr!='0') {
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '0';
}break;
case 16: i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  'x'+unhandled Expr in func ShowExpr: *ast.IndexExpr-'a';
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '0';break;
}
}
if (f.unicode) {
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '+';
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  'U';
}
if (negative) {
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '-';
}elseif (f.plus) {
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '+';
}elseif (f.space) {
i--;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  ' ';
}
if (f.unicode&&f.uniQuote&&a>=0&&a<=utf8.MaxRune&&strconv.IsPrint(rune(a))) {
var runeWidth  =  utf8.RuneLen(rune(a));
var width  =  1+1+runeWidth+1;
copy(unhandled Expr in func ShowExpr: *ast.SliceExpr, unhandled Expr in func ShowExpr: *ast.SliceExpr);
var i  -=  width;
var j  =  len(buf)-width;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  ' ';
j++;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '\'';
j++;
utf8.EncodeRune(unhandled Expr in func ShowExpr: *ast.SliceExpr, rune(a));
var j  +=  runeWidth;
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '\'';
}
f.pad(unhandled Expr in func ShowExpr: *ast.SliceExpr);
}
fmt.prototype.truncate = function (s) {
if (f.precPresent&&f.prec<utf8.RuneCountInString(s)) {
var n  =  f.prec;
for (var i in s) { 
 var %!s(<nil>) = s[i];
if (n==0) {
var s  =  unhandled Expr in func ShowExpr: *ast.SliceExpr;
break;
}
n--;
}
}
return s;
}
fmt.prototype.fmt_s = function (s) {
var s  =  f.truncate(s);
f.padString(s);
}
fmt.prototype.fmt_sx = function (s, digits) {
unhandled Stmt in func ShowStmt: *ast.DeclStmt
for (var i  =  0; i<len(s); i++){
if (i>0&&f.space) {
var b  =  append(b, ' ');
}
var v  =  unhandled Expr in func ShowExpr: *ast.IndexExpr;
var b  =  append(b, unhandled Expr in func ShowExpr: *ast.IndexExpr, unhandled Expr in func ShowExpr: *ast.IndexExpr);
}
f.pad(b);
}
fmt.prototype.fmt_q = function (s) {
var s  =  f.truncate(s);
unhandled Stmt in func ShowStmt: *ast.DeclStmt
if (f.sharp&&strconv.CanBackquote(s)) {
var quoted  =  "`"+s+"`";
}else{
if (f.plus) {
var quoted  =  strconv.QuoteToASCII(s);
}else{
var quoted  =  strconv.Quote(s);
}
}
f.padString(quoted);
}
fmt.prototype.fmt_qc = function (c) {
unhandled Stmt in func ShowStmt: *ast.DeclStmt
if (f.plus) {
var quoted  =  strconv.AppendQuoteRuneToASCII(unhandled Expr in func ShowExpr: *ast.SliceExpr, rune(c));
}else{
var quoted  =  strconv.AppendQuoteRune(unhandled Expr in func ShowExpr: *ast.SliceExpr, rune(c));
}
f.pad(quoted);
}
var doPrec = function  (f, def) {
if (f.precPresent) {
return f.prec;
}
return def;
}
fmt.prototype.formatFloat = function (v, verb, prec, n) {
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  ' ';
var slice  =  strconv.AppendFloat(unhandled Expr in func ShowExpr: *ast.SliceExpr, v, verb, prec, n);
switch (unhandled Expr in func ShowExpr: *ast.IndexExpr) {
case '-', '+': var slice  =  unhandled Expr in func ShowExpr: *ast.SliceExpr;break;
case : if (f.plus) {
unhandled Expr in func ShowExpr: *ast.IndexExpr  =  '+';
}elseif (f.space) {

}else{
var slice  =  unhandled Expr in func ShowExpr: *ast.SliceExpr;
}break;
}
f.pad(slice);
}
fmt.prototype.fmt_e64 = function (v) {
f.formatFloat(v, 'e', doPrec(f, 6), 64);
}
fmt.prototype.fmt_E64 = function (v) {
f.formatFloat(v, 'E', doPrec(f, 6), 64);
}
fmt.prototype.fmt_f64 = function (v) {
f.formatFloat(v, 'f', doPrec(f, 6), 64);
}
fmt.prototype.fmt_g64 = function (v) {
f.formatFloat(v, 'g', doPrec(f, unhandled Expr in func ShowExpr: *ast.UnaryExpr), 64);
}
fmt.prototype.fmt_G64 = function (v) {
f.formatFloat(v, 'G', doPrec(f, unhandled Expr in func ShowExpr: *ast.UnaryExpr), 64);
}
fmt.prototype.fmt_fb64 = function (v) {
f.formatFloat(v, 'b', 0, 64);
}
fmt.prototype.fmt_e32 = function (v) {
f.formatFloat(float64(v), 'e', doPrec(f, 6), 32);
}
fmt.prototype.fmt_E32 = function (v) {
f.formatFloat(float64(v), 'E', doPrec(f, 6), 32);
}
fmt.prototype.fmt_f32 = function (v) {
f.formatFloat(float64(v), 'f', doPrec(f, 6), 32);
}
fmt.prototype.fmt_g32 = function (v) {
f.formatFloat(float64(v), 'g', doPrec(f, unhandled Expr in func ShowExpr: *ast.UnaryExpr), 32);
}
fmt.prototype.fmt_G32 = function (v) {
f.formatFloat(float64(v), 'G', doPrec(f, unhandled Expr in func ShowExpr: *ast.UnaryExpr), 32);
}
fmt.prototype.fmt_fb32 = function (v) {
f.formatFloat(float64(v), 'b', 0, 32);
}
fmt.prototype.fmt_c64 = function (v, verb) {
f.buf.WriteByte('(');
var r  =  real(v);
for (var i  =  0; ; i++){
switch (verb) {
case 'e': f.fmt_e32(r);break;
case 'E': f.fmt_E32(r);break;
case 'f': f.fmt_f32(r);break;
case 'g': f.fmt_g32(r);break;
case 'G': f.fmt_G32(r);break;
}
if (i!=0) {
break;
}
f.plus  =  true;
var r  =  imag(v);
}
f.buf.Write(irparenBytes);
}
fmt.prototype.fmt_c128 = function (v, verb) {
f.buf.WriteByte('(');
var r  =  real(v);
for (var i  =  0; ; i++){
switch (verb) {
case 'e': f.fmt_e64(r);break;
case 'E': f.fmt_E64(r);break;
case 'f': f.fmt_f64(r);break;
case 'g': f.fmt_g64(r);break;
case 'G': f.fmt_G64(r);break;
}
if (i!=0) {
break;
}
f.plus  =  true;
var r  =  imag(v);
}
f.buf.Write(irparenBytes);
}
main();
