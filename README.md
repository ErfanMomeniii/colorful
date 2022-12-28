![image description](./assets/photo/logo.png)

# colorful

colorful is a lightweight package for rendering colorful text in terminal with Go you can use this instead of **fmt** package for having many features.


# Documentation
## Install
```bash
go get github.com/ErfanMomeniii/colorful
```   
Next, include it in your application:
```bash
import "github.com/ErfanMomeniii/colorful"
``` 
### Supported Colors

|                                                                 |
|:----------------------------------------------------------------|
| ![red color](./assets/photo/colors/Red.png)                     |
| ![green color](./assets/photo/colors/Green.png)                 |
| ![yellow color](./assets/photo/colors/Yellow.png)               |
| ![blue color](./assets/photo/colors/Blue.png)                   |
| ![magenta color](./assets/photo/colors/Magenta.png)             |
| ![cyan color](./assets/photo/colors/Cyan.png)                   |
| ![white color](./assets/photo/colors/White.png)                 | 
| ![black color](./assets/photo/colors/Black.png)                 |
| ![brightBlack color](./assets/photo/colors/BrightBlack.png)     |
| ![brightRed color](./assets/photo/colors/BrightRed.png)         |
| ![brightGreen color](./assets/photo/colors/BrightGreen.png)     |
| ![brightYellow color](./assets/photo/colors/BrightYellow.png)   |
| ![brightBlue color](./assets/photo/colors/BrightBlue.png)       |
| ![brightMagenta color](./assets/photo/colors/BrightMagenta.png) |
| ![brightCyan color](./assets/photo/colors/BrightCyan.png)       |
| ![brightWhite color](./assets/photo/colors/BrightWhite.png)     |
>You can also use the following color codes :
>
>![color code](./assets/photo/colors/code.png)

### Supported Background Colors
|                                                                      |
|:---------------------------------------------------------------------| 
| ![black background](./assets/photo/backgrounds/Black.png)                 |
| ![red background](./assets/photo/backgrounds/Red.png)                     |
| ![green background](./assets/photo/backgrounds/Green.png)                 |
| ![yellow background](./assets/photo/backgrounds/Yellow.png)               |
| ![blue background](./assets/photo/backgrounds/Blue.png)                   |
| ![magenta background](./assets/photo/backgrounds/Magenta.png)             |
| ![cyan background](./assets/photo/backgrounds/Cyan.png)                   |
| ![white background](./assets/photo/backgrounds/White.png)                 |
| ![brightBlack background](./assets/photo/backgrounds/BrightBlack.png)     |
| ![brightRed background](./assets/photo/backgrounds/BrightRed.png)         |
| ![brightGreen background](./assets/photo/backgrounds/BrightGreen.png)     |
| ![brightYellow background](./assets/photo/backgrounds/BrightYellow.png)   |
| ![brightBlue background](./assets/photo/backgrounds/BrightBlue.png)       |
| ![brightMagenta background](./assets/photo/backgrounds/BrightMagenta.png) |
| ![brightCyan background](./assets/photo/backgrounds/BrightCyan.png)       |
| ![brightWhite background](./assets/photo/backgrounds/BrightWhite.png)     |
>You can also use the following background codes :
>
>![color code](./assets/photo/backgrounds/code.png)

## Quick Start

The following example demonstrates how to print text in desired format and color:
```go
package main

import (
	"github.com/ErfanMomeniii/colorful"
)

func main() {
	colorful.Println(colorful.RedColor, colorful.BlueBackground, 
		"RedColorBlueBackground", colorful.ResetColor)
}

```
:warning: **If you dont use ResetColor**  ,Next time it will be printed with the same color as the one you had been set 

The following example demonstrates how to print text using color and background code:
```go
package main

import (
	"github.com/ErfanMomeniii/colorful"
)

func main() {
	colorful.Println(colorful.GetBackgroundColorFromCode(1), colorful.GetBackgroundColorFromCode(4), 
		"RedColorBlueBackground", colorful.ResetColor)
}

```
## Contributing
Pull requests are welcome. For changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
