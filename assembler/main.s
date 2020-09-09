"".add STEXT nosplit size=60 args=0x18 locals=0x10
	0x0000 00000 (main.go:3)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $16-24
	0x0000 00000 (main.go:3)	SUBQ	$16, SP
	0x0004 00004 (main.go:3)	MOVQ	BP, 8(SP)
	0x0009 00009 (main.go:3)	LEAQ	8(SP), BP
	0x000e 00014 (main.go:3)	PCDATA	$0, $-2
	0x000e 00014 (main.go:3)	PCDATA	$1, $-2
	0x000e 00014 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:3)	PCDATA	$0, $0
	0x000e 00014 (main.go:3)	PCDATA	$1, $0
	0x000e 00014 (main.go:3)	MOVQ	$0, "".~r2+40(SP)
	0x0017 00023 (main.go:4)	MOVQ	$0, "".sum(SP)
	0x001f 00031 (main.go:5)	MOVQ	"".a+24(SP), AX
	0x0024 00036 (main.go:5)	ADDQ	"".b+32(SP), AX
	0x0029 00041 (main.go:5)	MOVQ	AX, "".sum(SP)
	0x002d 00045 (main.go:6)	MOVQ	AX, "".~r2+40(SP)
	0x0032 00050 (main.go:6)	MOVQ	8(SP), BP
	0x0037 00055 (main.go:6)	ADDQ	$16, SP
	0x003b 00059 (main.go:6)	RET
	0x0000 48 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 48 c7  H...H.l$.H.l$.H.
	0x0010 44 24 28 00 00 00 00 48 c7 04 24 00 00 00 00 48  D$(....H..$....H
	0x0020 8b 44 24 18 48 03 44 24 20 48 89 04 24 48 89 44  .D$.H.D$ H..$H.D
	0x0030 24 28 48 8b 6c 24 08 48 83 c4 10 c3              $(H.l$.H....
"".main STEXT size=114 args=0x0 locals=0x28
	0x0000 00000 (main.go:10)	TEXT	"".main(SB), ABIInternal, $40-0
	0x0000 00000 (main.go:10)	MOVQ	TLS, CX
	0x0009 00009 (main.go:10)	PCDATA	$0, $-2
	0x0009 00009 (main.go:10)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (main.go:10)	PCDATA	$0, $-1
	0x0010 00016 (main.go:10)	CMPQ	SP, 16(CX)
	0x0014 00020 (main.go:10)	PCDATA	$0, $-2
	0x0014 00020 (main.go:10)	JLS	107
	0x0016 00022 (main.go:10)	PCDATA	$0, $-1
	0x0016 00022 (main.go:10)	SUBQ	$40, SP
	0x001a 00026 (main.go:10)	MOVQ	BP, 32(SP)
	0x001f 00031 (main.go:10)	LEAQ	32(SP), BP
	0x0024 00036 (main.go:10)	PCDATA	$0, $-2
	0x0024 00036 (main.go:10)	PCDATA	$1, $-2
	0x0024 00036 (main.go:10)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (main.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (main.go:10)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (main.go:11)	PCDATA	$0, $0
	0x0024 00036 (main.go:11)	PCDATA	$1, $0
	0x0024 00036 (main.go:11)	MOVQ	$1, (SP)
	0x002c 00044 (main.go:11)	MOVQ	$2, 8(SP)
	0x0035 00053 (main.go:11)	CALL	"".add(SB)
	0x003a 00058 (main.go:11)	MOVQ	16(SP), AX
	0x003f 00063 (main.go:11)	MOVQ	AX, ""..autotmp_0+24(SP)
	0x0044 00068 (main.go:11)	CALL	runtime.printlock(SB)
	0x0049 00073 (main.go:11)	MOVQ	""..autotmp_0+24(SP), AX
	0x004e 00078 (main.go:11)	MOVQ	AX, (SP)
	0x0052 00082 (main.go:11)	CALL	runtime.printint(SB)
	0x0057 00087 (main.go:11)	CALL	runtime.printnl(SB)
	0x005c 00092 (main.go:11)	CALL	runtime.printunlock(SB)
	0x0061 00097 (main.go:12)	MOVQ	32(SP), BP
	0x0066 00102 (main.go:12)	ADDQ	$40, SP
	0x006a 00106 (main.go:12)	RET
	0x006b 00107 (main.go:12)	NOP
	0x006b 00107 (main.go:10)	PCDATA	$1, $-1
	0x006b 00107 (main.go:10)	PCDATA	$0, $-2
	0x006b 00107 (main.go:10)	CALL	runtime.morestack_noctxt(SB)
	0x0070 00112 (main.go:10)	PCDATA	$0, $-1
	0x0070 00112 (main.go:10)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 76 55 48 83 ec 28 48 89 6c 24 20 48  H;a.vUH..(H.l$ H
	0x0020 8d 6c 24 20 48 c7 04 24 01 00 00 00 48 c7 44 24  .l$ H..$....H.D$
	0x0030 08 02 00 00 00 e8 00 00 00 00 48 8b 44 24 10 48  ..........H.D$.H
	0x0040 89 44 24 18 e8 00 00 00 00 48 8b 44 24 18 48 89  .D$......H.D$.H.
	0x0050 04 24 e8 00 00 00 00 e8 00 00 00 00 e8 00 00 00  .$..............
	0x0060 00 48 8b 6c 24 20 48 83 c4 28 c3 e8 00 00 00 00  .H.l$ H..(......
	0x0070 eb 8e                                            ..
	rel 12+4 t=17 TLS+0
	rel 54+4 t=8 "".add+0
	rel 69+4 t=8 runtime.printlock+0
	rel 83+4 t=8 runtime.printint+0
	rel 88+4 t=8 runtime.printnl+0
	rel 93+4 t=8 runtime.printunlock+0
	rel 108+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".add SDWARFLOC size=0
go.info."".add SDWARFINFO size=82
	0x0000 03 22 22 2e 61 64 64 00 00 00 00 00 00 00 00 00  ."".add.........
	0x0010 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0a  ................
	0x0020 73 75 6d 00 04 00 00 00 00 02 91 68 0f 61 00 00  sum........h.a..
	0x0030 03 00 00 00 00 01 9c 0f 62 00 00 03 00 00 00 00  ........b.......
	0x0040 02 91 08 0f 7e 72 32 00 01 03 00 00 00 00 02 91  ....~r2.........
	0x0050 10 00                                            ..
	rel 0+0 t=24 type.int+0
	rel 8+8 t=1 "".add+0
	rel 16+8 t=1 "".add+60
	rel 26+4 t=30 gofile..C:\Users\12948\go\src\go-codes\assembler\main.go+0
	rel 37+4 t=29 go.info.int+0
	rel 49+4 t=29 go.info.int+0
	rel 60+4 t=29 go.info.int+0
	rel 74+4 t=29 go.info.int+0
go.range."".add SDWARFRANGE size=0
go.debuglines."".add SDWARFMISC size=18
	0x0000 04 02 0a 11 f6 60 06 41 06 6a 06 41 04 01 03 7b  .....`.A.j.A...{
	0x0010 06 01                                            ..
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 0+0 t=24 type.int+0
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+114
	rel 27+4 t=30 gofile..C:\Users\12948\go\src\go-codes\assembler\main.go+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=20
	0x0000 04 02 03 04 14 0a eb 9c 06 5f 06 02 1e f6 71 04  ........._....q.
	0x0010 01 03 77 01                                      ..w.
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
