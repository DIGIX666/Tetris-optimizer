package main

/*
----------------------------------- RECURSION ------------------------------------

	The "recursion" function recursively places tetrominos on the board, searching
	for a valid position for each piece. If it finds a solution, it returns true; otherwise,
	it continues the search after removing the tetromino.

\----------------------------------------------------------------------------------
*/
func recursion(b board, array []tetrimino, idx int) bool {
	if idx == len(array) {
		return true
	}
	for i := range b {
		for j := range b[i] {
			if b.check(i, j, array[idx]) {
				b.put(i, j, idx, array[idx])
				if recursion(b, array, idx+1) {
					return true
				}
				b.remove(i, j, array[idx])
			}
		}
	}
	return false
}

/*
-------------------------------------- SOLVE ------------------------------------

	The “solve” function takes an array of tetriminos and returns a board with the solution.
	It begins with a board of minimum size to fit all tetriminos, then uses the “recursion” function to place them.
	If successful, it returns the board; if not, it enlarges the board and retries until a solution is found.

\--------------------------------------------------------------------------------
*/
func Solve(array []tetrimino) board {
	var res board
	square := 2
	for square*square < len(array)*4 {
		square += 1
	}
	for done := false; !done; {
		res = makeBoard(square)
		done = recursion(res, array, 0)
		square += 1
	}
	return res
}
