package builtincmd

import (
	"fmt"
	"os"

	"github.com/jhaksh-24/shell/utils"
)

func Type (parts []string) {
	if len(parts) <= 1 {
		return
	}
	
	builtins := []string{"exit", "echo", "type", "cd", "history"}

	for _, val := range builtins {
		if val == parts[1] {
			fmt.Println(val, "is a shell builtin")
			return
		}
	}
	
	if file, exists := utils.SearchBinInPath(parts[1]); exists {
		fmt.Fprintf(os.Stdout, "%s is %s\n", parts[1], file)
		return
	}

	fmt.Printf("%s: not found\n", parts[1])
}
