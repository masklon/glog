package handlers

import (
	"errors"
	"fmt"
	"github.com/ml444/glog/filters"
	"github.com/ml444/glog/formatters"
	"github.com/ml444/glog/message"
	"io"
	"os"
)

const terminator = '\n'

type StreamHandler struct {
	//BaseHandler
	stream io.Writer

	formatter formatters.IFormatter
	filter    filters.IFilter
}

func NewStreamHandler(formatter formatters.IFormatter, filter filters.IFilter) (*StreamHandler, error) {
	return &StreamHandler{
		filter:    filter,
		formatter: formatter,
		stream:    os.Stdout,
	}, nil
}

func (h *StreamHandler) Init(dir, name string) error {
	return nil
}

func (h *StreamHandler) format(record *message.Entry) ([]byte, error) {
	if h.formatter != nil {
		return h.formatter.Format(record)
	}
	return nil, nil
}

func (h *StreamHandler) emit(msg []byte) error {
	msg = append(msg, terminator)
	_, err := h.stream.Write(msg)
	if err != nil {
		return err
	}
	h.Flush()
	return nil
}

func (h *StreamHandler) Emit(record *message.Entry) error {
	if h.filter != nil {
		if ok := h.filter.Filter(record); !ok {
			return errors.New(fmt.Sprintf("Filter out this msg: %v", record))
		}
	}

	msgByte, err := h.format(record)
	if err != nil {
		return err
	}

	//h.Acquire()
	err = h.emit(msgByte)
	//h.Release()
	return err
}

// Flushes the stream.
func (h *StreamHandler) Flush() {
	/*
	   	self.acquire()
	      try:
	   	   if self.stream and hasattr(self.stream, "flush"):
	   		   self.stream.flush()
	      finally:
	   	   self.release()
	*/
}
func (h *StreamHandler) Sync() error {
	return nil
}

func (h StreamHandler) SetStream(stream io.Writer) {
	h.stream = stream
}
