Data Types

bool : 0 or 1 
string : 32bit or ""

int = {
	int: [32bit]
	int8: -128
	int16: -32.768
	int32: -2.147.483.648
	int64: -9.224.372.036.854.775.808
	uint: [32bit or 64bit]
	uint8: 0 ~ 255
	uint16: 0 ~ 65.535
	uint32: 0 ~ 4.294.967.295
	uint64: 0 ~ 18.446.744.073.709.551.615
	uintptr: Grande o suficiente para guardar o padrão de bits de 	qualquer ponteiro
}

byte: uint8
rune: int32 (for integers)

float = {
	float32: 754 de 32 bits
	float64: 754 de 64 bits
}

complex = {
	complex64: float32 e parte imaginária
	complex128: float64 e parte imaginária
}

