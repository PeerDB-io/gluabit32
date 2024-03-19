# bit32 for gopher-lua

Implements Lua 5.2 [bit32](https://www.lua.org/manual/5.2/manual.html#6.7) for [gopher-lua](https://github.com/yuin/gopher-lua). To use, call
```go
import (
	"github.com/PeerDB-io/gluabit32"
)

// add so that `local bit32 = require("bit32")` works
L.PreloadModule("bit32", gluabit32.Loader)

// or add to global env
L.Push(ls.NewFunction(gluabit32.Loader))
L.Call(0, 1)
L.Env.RawSetString("bit32", L.Get(-1))
L.Pop(1)
```
