// package block
package main 


// import block
import (
	"fmt"
	"goskills/hillel/router"
)



// main function
func main () {

	/* fmt.Println("Hello, playground") */
    /* var router *gin.Engine равно router := value */


    r := router.New()

    
    // запускать router
    r.Run()
}


/*
Lunch code

/var/www/goskills/hillel$ go build
/var/www/goskills/hillel$ ./hillel

*/
/* 
# https://newsapi.org
# https://github.com/gin-gonic/gin
Language core

- functions
- structs
- pointers
- methods
- interfaces
- errors
- don't panic!
- goroutines
- channels
- simple web server
*/

/*
Compile binary "hillel бинарный"
$ go build ( will generate file hillel in our case )
можно под любой платформе компилировать программу
это делается через $ go build и указать какие-то параметры
*/