package main

import (
	"context"
	"fmt"
	"time"
)

// Full example, putting all together
// We have a function that calls three web services. We send data to two
// of thse services, take the results and send them to the third, returning
// the result. This process must take less than 50 ms, or an error is returned

func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
	// context that times out in 50 ms. When there is a ctx variable, use
	// its timer support instead of calling time.After
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	// Reaching the timeout cancels the context
	defer cancel() // use it to avoid resources leaking
	// processor instance with a series of channels that we'll use to communicate
	// with our goroutines
	p := processor{
		// all are buffered, so that the goroutines that write to them
		// can exit after writing without waiting for a read.
		outA: make(chan Aout, 1),
		outB: make(chan Bout, 1),
		inC: make(chan Cin, 1),
		outC: make(chan COut, 1),
		errs: make(chan error, 2), // it could have 2 errors written on it
	}
	p.launch(ctx, data)
	inputC, err := p.waitForAB(ctx)
	if err != nil {
		return Cout{}, err
	}
	p.inC <- inputC
	out, err := p.waitForC(ctx)
	return out, err
}

type processor struct {
	outA chan Aout 
	outB chan Bout 
	inC  chan Cin
	outC chan COut 
	errs chan error
}

// launch method on processor to start three goroutines
func (p *processor) launch(ctx context.Context, data Input) {
	go func() {
		// the goroutines are very similar for A and B, they call their 
		// method and if an error is present, they write it in the  p.err 
		// channel. If not, they write over their channel
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err 
			return
		}
		p.outA <- aOut
	}()

	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- bOut
	}()

	go func() {
		select {
		case <-ctx.Done(): // is triggered if the context is cancelled
			return
		case inputC := <-p.inC: // if data to call the func is available
			cOut, err := getResultC(ctx, inputC)
			if err != nil {
				p.err != nil {
					p.errs <- err
					return
				}
				p.outC <- cOut
			}
		}()
	}
}

// After the goroutines are launched, we call the waitForAB method
func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
	var inputC cIn
	count := 0
	// for select loop to populate inputC, an instance of CIn, the
	// input parameter for getResultC
	for count < 2 {
		select {
		// the first two read from channels and populate the fields in inputC
		case a := <-p.outA:
			inputC.A = a 
			count++
		case b := <-p.outB:
			inputC.b = b 
			count++
		// Thse two handle error conditions
		case err := <-p.errs:
			// written error 
			return CIn{}, err
		case <-ctx.Done():
			// if context was canceled
			return CIn{}, ctx.Err()
		}
	}
	return inputC, nil
}

// Back in GatherAndProcess() we perform a standard nil check on the error
// if all well, we write the inputC value to the p.inC channel and then call 
// the waitForC method on processor
func (p *processor) waitForC(ctx context.Context) (COut, error) {
	// If getResultC completed successfully, we read its output from the p.out 
	// channel and return it.
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.errs:
		// If getResultC returns an error, we read the error
		// from the p.errs channel and return it
		return cOut{}, err 
	case <-ctx.Done():
		// return an error if the context is cancelled
		return COut{}, ctx.Err()
	}
}

// By structuring our code with goroutines, channels and select statements,
// we separate the individual steps, allow independent parts to run and 
// complete in any order, and cleanly exchange data between the dependent parts.
// We make sure that no part of the program hangs and we can properly handle 
// timeouts from within functions


func main() {

}