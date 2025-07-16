package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"golang.org/x/sync/errgroup"
)

const Reset = "\033[0m"
const Yellow = "\033[33m"
const Red = "\033[31m"

func main() {
	dir, err := os.MkdirTemp("", "executables*")
	if err != nil {
		log.Fatal("Failed to make tmp dir", err)
	}
	defer os.RemoveAll(dir)

	files := [25]string{}

	errGrp := new(errgroup.Group)
	errGrp.SetLimit(6)

	for p := range 25 {
		exe := fmt.Sprintf("%s%s%d.exe", dir, string(os.PathSeparator), p)
		files[p] = exe

		errGrp.Go(func() error {
			prob := exec.Command("go", "build", "-o", exe, fmt.Sprintf("%02d/%02d.go", p+1, p+1))

			if _, err := prob.Output(); true {
				if err != nil {
					if err, ok := err.(*exec.ExitError); ok {
						log.Println("Output:", string(err.Stderr))
					}
					log.Println("Failed to compile problem", p+1, err)
					return err
				}
			}

			return nil
		})
	}

	err = errGrp.Wait()
	if err != nil {
		log.Fatal("Failed to compile...", err)
	}

	totalTime := 0

	for p := range 25 {
		prob := exec.Command(fmt.Sprintf("%s", files[p]))
		prob.Dir = fmt.Sprintf("%02d/", p+1)

		t1 := time.Now()

		if out, err := prob.Output(); true {
			if err != nil {
				if err, ok := err.(*exec.ExitError); ok {
					log.Println("Output:", string(err.Stderr))
				}
				log.Fatal("Failed to run problem", p+1, err)
			}

			dur := time.Since(t1)
			totalTime += int(dur)

			if dur > 150*1000000 {
				fmt.Print(Red)
			} else if dur > 90*1000000 {
				fmt.Print(Yellow)
			}

			fmt.Println("Problem", p+1, "took", dur, "Answer", string(out))
			fmt.Print(Reset)
		}
	}

	fmt.Println("Total execution time:", totalTime/int(time.Millisecond), "ms")

}
