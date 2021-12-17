package shared

import "math"

func BinaryToInteger(binaryString *string) (value int) {
	n := 0
	for i := len(*binaryString) - 1; i >= 0; i-- {
		if (*binaryString)[i] == '1' {
			value += (int(math.Pow(float64(2), float64(n))))
		}
		n++
	}
	return
}

func HexToBinary(hexString *string) (binary string) {
	hexMap := make(map[string]string)

	hexMap["0"] = "0000"
	hexMap["1"] = "0001"
	hexMap["2"] = "0010"
	hexMap["3"] = "0011"
	hexMap["4"] = "0100"
	hexMap["5"] = "0101"
	hexMap["6"] = "0110"
	hexMap["7"] = "0111"
	hexMap["8"] = "1000"
	hexMap["9"] = "1001"
	hexMap["A"] = "1010"
	hexMap["B"] = "1011"
	hexMap["C"] = "1100"
	hexMap["D"] = "1101"
	hexMap["E"] = "1110"
	hexMap["F"] = "1111"

	for _, hex := range *hexString {
		binary = binary + hexMap[string(hex)]
	}

	return

}
