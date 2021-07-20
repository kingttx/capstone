----
main_test.go
----
package main

 


func TestHelloServerLogic(t *testing.T) {
    assert.Equal(t, "Hello, Tom!", HelloServerLogic("Tom"))
}
