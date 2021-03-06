package main

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
	}
}

//
//func main(){
//	ch := make(chan int) // Create a new channel.
//	go generate(ch)      // Start generate() as a subprocess.
//	for {
//		prime := <-ch
//		fmt.Print(prime, "\n")
//		ch1 := make(chan int)
//		go filter(ch, ch1, prime)
//		ch = ch1
//	}
//
//}

type Result1 struct {
}

type Result2 struct {
}

// fds反倒是

type Result3 struct {
}
