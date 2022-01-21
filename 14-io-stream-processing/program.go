package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"sync"
)

type chInt chan int

type opLock struct {
	count int
	sync.Mutex
}

func (opLock *opLock) incr() {
	opLock.Lock()
	opLock.count++
	opLock.Unlock()
}

func (opLock *opLock) decr() int {
	opLock.Lock()
	if opLock.count > 0 {
		opLock.count--
	}
	tmp := opLock.count
	opLock.Unlock()
	return tmp
}

func main() {

	splitter := make(chInt)
	sourceIntraMutex := &opLock{}
	sourceIntraMutex.incr()
	sourceIntraMutex.incr()

	wg := &sync.WaitGroup{}
	// read from file one and push to splitter
	wg.Add(1)
	go source("data1.dat", splitter, wg, sourceIntraMutex)
	// read from file two and push to splitter
	wg.Add(1)
	go source("data2.dat", splitter, wg, sourceIntraMutex)

	eventCh := make(chInt)
	oddCh := make(chInt)
	wg.Add(1)
	go split(splitter, eventCh, oddCh, wg)

	outFileName := "result.dat"
	wg.Add(1)
	go merge(eventCh, oddCh, outFileName, wg)

	wg.Wait()
}

func source(fileName string, splitter chInt, wg *sync.WaitGroup, srcMutex *opLock) {
	fileReader, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		fileReader.Close()
		if srcMutex.decr() == 0 {
			// both files read - now close splitter
			close(splitter)
		}
		wg.Done()
	}()
	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		if no, err := strconv.Atoi(scanner.Text()); err == nil {
			splitter <- no
		}
	}
}

func split(inChannel, eventCh, oddCh chInt, wg *sync.WaitGroup) {
	defer func() {
		close(oddCh)
		close(eventCh)
		wg.Done()
	}()
	for val := range inChannel {
		if val%2 == 0 {
			eventCh <- val
		} else {
			oddCh <- val
		}
	}
}

func naiveAdd(ch chInt, result *int, lwg *sync.WaitGroup) {
	defer lwg.Done()

	sum := 0
	for val := range ch {
		sum += val
	}
	*result = sum
	return
}
func writeOut(f *os.File, output string, sum int) {
	text := output + strconv.Itoa(sum) + "\n"
	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
}
func merge(eventCh, oddCh chInt, outputFile string, wg *sync.WaitGroup) {
	defer wg.Done()

	localWg := &sync.WaitGroup{}
	var oddSum int
	var evenSum int

	localWg.Add(1)
	go naiveAdd(oddCh, &oddSum, localWg)

	localWg.Add(1)
	go naiveAdd(eventCh, &evenSum, localWg)

	localWg.Wait()

	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writeOut(f, "Odd Sum is ", oddSum)
	writeOut(f, "Even Sum is ", evenSum)
}
