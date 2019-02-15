This project is a very simple set of helpers designed to make it easier to implement "design by contract" in Go. There is no meta-programming magic or code generation, only a small utility file. The tradeoff is that this project only intends to make it easier to implement contracts, but it will not stop you from implementing bad contracts, nor will it make sure you don't forget to close off all points of entry with proper invariant and condition checking.

This project can be considered beta or proof of concept, there are really no guarantees about the API's consistency.  Luckily it's extremely small, small enough to just peruse and take code as needed without vendoring, even.

You can view the documention at: https://godoc.org/github.com/samuelhorwitz/contract