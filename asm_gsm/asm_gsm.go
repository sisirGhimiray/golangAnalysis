package asmgsm

import (
	"example.com/golang/csvReading"
)


var(
	// Reading file contains Additional Surveillance Measure, Graded Surveillance Measure stocks name provided by nse
	AsmStock=csvreading.ExecuteReading("./asm_gsm/asm-latest.csv")
	GsmStock=csvreading.ExecuteReading("./asm_gsm/gsm-latest.csv")
)


