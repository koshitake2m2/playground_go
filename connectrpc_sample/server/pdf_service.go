package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"example.com/aaa/gen/proto/pdf"
	"example.com/aaa/gen/proto/pdf/pdfconnect"

	"connectrpc.com/connect"
)

type PdfServer struct{}

func NewPdfService() pdfconnect.PdfServiceHandler {
	return &PdfServer{}
}

func (s *PdfServer) GetPdf(
	ctx context.Context,
	req *connect.Request[pdf.PdfRequest],
	stream *connect.ServerStream[pdf.PdfChunk],
) error {
	filePath := "../static/" + req.Msg.Filename
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open PDF file: %v", err)
		return connect.NewError(connect.CodeInternal, err)
	}
	defer file.Close()

	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Failed to read file: %v", err)
			return connect.NewError(connect.CodeInternal, err)
		}

		err = stream.Send(&pdf.PdfChunk{Content: buffer[:n]})
		if err != nil {
			log.Printf("Failed to send chunk: %v", err)
			return err
		}

		time.Sleep(50 * time.Millisecond)
	}

	return nil
}
