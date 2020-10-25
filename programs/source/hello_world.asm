;--------------------------------------------------------------------------
; A "Hello world" program written for the LC3-VM.
; Copyright (c) 2020 Hadi Alqattan | MIT License.
;--------------------------------------------------------------------------

.ORIG x3000

LEA R0, hello    ; R0 = &hello
TRAP x22         ; PUTS (print char array at addr in R0)
hello .STRINGZ "\n-- HELLO WORLD! --\n\n"

LEA	R0,	anykey   ; R0 = &anykey
TRAP x22         ; PUTS (print char array at addr in R0)
TRAP x20         ; GETC (wait for any key to be pressed)
anykey .STRINGZ	"press any key to exit..."

HALT
.END
