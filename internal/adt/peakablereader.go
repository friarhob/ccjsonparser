package adt

import (
	"bufio"
)

type PeakableReader struct {
	reader *bufio.Reader
	buffer *Queue
}

func NewPeakableReader(reader *bufio.Reader) *PeakableReader {
	return &PeakableReader{reader: reader, buffer: NewQueue()}
}

func (pb *PeakableReader) PopRune() (rune, error) {
	if pb.buffer.IsEmpty() {
		nextRune, _, err := pb.reader.ReadRune()
		if err != nil {
			return nextRune, err
		}

		pb.buffer.Enqueue(nextRune)
	}

	nextRune, err := pb.buffer.Dequeue()
	if err != nil {
		return rune(0), err
	}
	return nextRune.(rune), nil
}

func (pb *PeakableReader) PeekRune() (rune, error) {
	if pb.buffer.IsEmpty() {
		nextRune, _, err := pb.reader.ReadRune()
		if err != nil {
			return nextRune, err
		}

		pb.buffer.Enqueue(nextRune)
	}

	nextRune, err := pb.buffer.Peek()
	if err != nil {
		return rune(0), err
	}
	return nextRune.(rune), nil
}
