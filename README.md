### Due to limitations of webview I have moved this effort to go-astilectron.  If changes are made and these blocking issues are resolved, I may return to this method.

vectorworks-app-cleaner

Running this software will remove files that may or may not be associated with Vectorworks products.  Please use at your own risk.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.


required dlls: `webview.dll` `WebView2Loader.dll`

Build: `go build -ldflags="-H windowsgui" -o webview-example.exe`

Windows ENV needs this to run 
CheckNetIsolation.exe LoopbackExempt -a -n="Microsoft.Win32WebViewHost_cw5n1h2txyewy"
