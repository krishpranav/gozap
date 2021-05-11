package version

import (
	"fmt"
	"os"
)

// Banner is show banner
func Banner() {
	fmt.Fprintln(os.Stderr, "          ,/")
	fmt.Fprintln(os.Stderr, "        ,'/")
	fmt.Fprintln(os.Stderr, "      ,' /")
	fmt.Fprintln(os.Stderr, "    ,'  /_____,")
	fmt.Fprintln(os.Stderr, "  .'____    ,'                     MZAP")
	fmt.Fprintln(os.Stderr, "        /  ,'     [ Multiple target ZAP scanner ]")
	fmt.Fprintln(os.Stderr, "       / ,'       [ "+VERSION+" ] [ by @krishpranav ]")
	fmt.Fprintln(os.Stderr, "      /,'")
	fmt.Fprintln(os.Stderr, "     /'")
	fmt.Fprintln(os.Stderr, "")
}
