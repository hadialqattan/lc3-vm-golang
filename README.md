# LC-3 VM

![LC-3 VM by Hadi Alqattan](docs/carbon-lc3-vm.png)

[LC-3](https://en.wikipedia.org/wiki/LC-3) VM written in Go. LC-3 or Little Computer 3 is
a fictional computer system that is designed to teach students how to code in assembly language.
This project is based on an [article](https://justinmeiners.github.io/lc3-vm/) written by Justin Meiners and Ryan Pendleton.

---
## Installation

```bash
$ go get github.com/hadialqattan/lc3-vm-golang
```

## Usage

Run a compiled program:
```bash
$ lc3-vm-golang programs/bin/rogue.obj
```

Compile & Run an LC-3 assembly program (Docker is required):
```
$ ./compiler/runner.sh hello_world # `hello_world.asm` should be placed in `programs/source/`
```

---
## Useful Resources

* [LC-3 Online Simulator](https://wchargin.github.io/lc3web/)
* [LC-3 Assembly Manual And Examples](http://people.cs.georgetown.edu/~squier/Teaching/HardwareFundamentals/LC3-trunk/docs/LC3-AssemblyManualAndExamples.pdf)
* [Ryan Pendleton's 2048 game implementation](https://github.com/rpendleton/lc3-2048)
* [Justin Meiners's Rogue game implementation](https://github.com/justinmeiners/lc3-rogue)

---
## Copyright ©

👤 **Hadi Alqattan**

* Github: [@hadialqattan](https://github.com/hadialqattan)
* Email: [alqattanhadizaki@gmail.com](<mailto:alqattanhadizaki@gmail.com>)

📝 **License**

Copyright © 2020 [Hadi Alqattan](https://github.com/hadialqattan).<br />
This project is [MIT](LICENSE) licensed.

---
Give a ⭐️ if this project helped you!
