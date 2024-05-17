package main

const blockSize = 20
const dot = '.'
const hashTag = '#'
const newLine = '\n'

/*
-------------------------------------- IS VALID BLOCK ---------------------------

	The "isValidBlock" function checks a byte array "data" for validity, using a loop
	for each byte and a "switch" statement to classify characters as dots ('.'),
	hashtags ('#'), or line breaks ('\n'). It counts dots and hashtags,
	verifies the regularity of line breaks, and returns true if the conditions of
	12 dots and 4 hashtags are met, otherwise false.

--------------------------------------------------------------------------------
*/

func isValidBlock(data []byte) bool {
	numDots, numHashtags := 0, 0
	for i := range data {
		switch data[i] {
		case dot:
			numDots++
		case hashTag:
			numHashtags++
		case newLine:
			if i%5 != 4 {
				return false
			}
		default:
			return false
		}
	}
	return numDots == 12 && numHashtags == 4
}

/*
-------------------------------------- VALIDATION FILE ---------------------------

	The “ValidationFile” function checks a byte array “data” for valid tetriminos using a loop
	to process each 20-byte block with the “isValidBlock” function. It converts valid blocks to tetriminos,
	checks them with “isValid”, and adds to the “res” array if successful.
	The function returns a boolean “ok” indicating overall success and the “res” array,
	breaking early if any block fails validation.

--------------------------------------------------------------------------------
*/

func ValidationFile(data []byte) (ok bool, res []tetrimino) {
	for i := 0; i < len(data); {
		if i+blockSize > len(data) {
			break
		}
		if !isValidBlock(data[i : i+blockSize]) {
			break
		}
		t := blocktoTetri(data[i : i+blockSize])
		if !t.isValid() {
			break
		}
		res = append(res, t)
		i += blockSize
		if i == len(data) {
			ok = true
			break
		}
		if data[i] != '\n' {
			break
		}
		i += 1
	}
	return ok, res
}
