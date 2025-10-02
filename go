function findLongest(array) {
    let result = 0;
    let max = 0;

    for (let i = 0; i < array.length; i++) {
        let prev = max;
        max = Math.max(array[i].toString().length, max);

        if (max !== prev) {
            result = array[i];
        }
    };

    return result;
}

function getRank(number) {
    if (number == 0) {
        return 1;
    }
    return Math.floor(Math.log10(Math.abs(number))) + 1;
}

function findLongest2(array) {
    if (array.length === 0) {
        return null; 
    }

    let longestNumber = array[0];
    let maxRank = getRank(array[0]);

    for (let i = 1; i < array.length; i++) {
        const currentRank = getRank(array[i]);
        
        if (currentRank > maxRank) {
            maxRank = currentRank;
            longestNumber = array[i];
        }
    }

    return longestNumber;
}

const arr = Array.from({ length: 3_000_001 }, (_, i) => i);

console.time("a");
console.log(findLongest2(arr));
console.timeEnd("a");
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите что-то: ")
	reader := bufio.NewScanner(os.Stdin)

	ready := reader.Scan()

	fmt.Println(ready)
	fmt.Println(reader.Text())
}
