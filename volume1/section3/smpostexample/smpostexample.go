// An example of creating a new instance of a Social Media Post type.
package main

import (
	"fmt"

	"github.com/yuanyewa/gofullstack/volume1/section3/smpostexample/mypackage"

	"github.com/yuanyewa/gofullstack/volume1/section3/socialmedia"
)

func main() {

	myPost := socialmedia.NewPost("EngineerKamesh", socialmedia.Moods["thrilled"], "Go is awesome!", "Check out the Go web site!", "https://golang.org", "", "", []string{"go", "golang", "programming language"})
	fmt.Printf("myPost: %+v\n", myPost)
	fmt.Println(mypackage.XX)

	var m mypackage.MyS
	m = *m.Update("Eva", 16, []string{"SD", "CA", "US"}) // is there a better way without need to assign back to m?
	fmt.Println(m)
}
