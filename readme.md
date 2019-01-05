# chingsrv 💰

Simple server that chings when you call `/ching`.

## Usage on a Raspberry Pi

Since `github.com/hajimehoshi/oto` relies on Cgo, cross compiling is not possible. Also, the sample rate is adjusted to work well on the Pi's internal audio DAC, which is not very great quality wise, but easy to use.

- Clone the repository
- `go install`
- Adjust volume `amixer cset numid=1 -- number` where number is between -10200 and +400 (100 is ok)
- Run `chingsrv <WAV File>`
