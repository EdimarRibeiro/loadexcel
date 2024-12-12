package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/EdimarRibeiro/loadexcel/internal/infrastructure"
)

type FileController struct{}

func CreateFileController() *FileController {
	return &FileController{}
}

func (c *FileController) CreateFileHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		fmt.Printf("Erro ao ler o corpo da requisição: %v\n", err)
		return
	}
	fmt.Println("Tamanho do corpo recebido:", len(body))

	if len(body) < 4 || string(body[:4]) != "%PDF" {
		http.Error(w, "O arquivo enviado não é um PDF válido", http.StatusBadRequest)
		return
	}

	excelFile, err := infrastructure.ProcessPDFBytesToExcel(body)
	if err != nil {
		http.Error(w, "Erro ao processar o PDF: "+err.Error(), http.StatusInternalServerError)
		fmt.Printf("Erro ao processar o PDF: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", `attachment; filename="output.xlsx"`)
	w.WriteHeader(http.StatusOK)

	if err := excelFile.Write(w); err != nil {
		http.Error(w, "Erro ao escrever o arquivo Excel: "+err.Error(), http.StatusInternalServerError)
		fmt.Printf("Erro ao escrever o arquivo Excel: %v\n", err)
		return
	}

	fmt.Println("Arquivo Excel gerado e enviado com sucesso.")
}
