package infrastructure

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ProcessPDFBytesToExcel(pdfBytes []byte) (*excelize.File, error) {
	// Configura o comando Camelot para converter PDF para CSV em memória
	cmd := exec.Command("camelot", "-f", "csv", "-p", "1-end", "stream", "-")
	cmd.Stdin = bytes.NewReader(pdfBytes) // Passa o PDF como entrada para o Camelot
	var camelotOutput bytes.Buffer
	cmd.Stdout = &camelotOutput
	cmd.Stderr = &camelotOutput

	// Executa o comando Camelot
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("falha ao executar Camelot: %v : %s", err, camelotOutput.String())
	}

	// Inicializa o leitor CSV com a saída do Camelot
	r := csv.NewReader(bufio.NewReader(&camelotOutput))
	totalPages := 0
	var outputs [][12]string

	// Detecta o número total de páginas e processa os dados
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("erro ao ler CSV: %v", err)
		}

		// Determina o total de páginas
		if len(record) >= 5 && strings.Contains(record[4], "1 of") {
			totalPages, _ = strconv.Atoi(record[4][5:])
		}
	}

	if totalPages == 0 {
		return nil, fmt.Errorf("não foi possível obter o número total de páginas")
	}

	xlsx := excelize.NewFile()

	// Processa as páginas e extrai os dados do CSV
	for i := 1; i <= totalPages; i++ {
		var lines [12]string
		counter := -1
		isFirst := true

		r := csv.NewReader(bufio.NewReader(bytes.NewReader(camelotOutput.Bytes())))

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, fmt.Errorf("erro ao ler CSV: %v", err)
			}

			if counter == -1 {
				if strings.Contains(record[1], "Transaction ID") {
					counter = 0
				}
			} else {
				match, _ := regexp.MatchString("^\\d\\d\\s...\\s\\d\\d\\d\\d", record[0])
				if match && !isFirst {
					counter = 0
					lines[3] = strings.Replace(lines[3], ",", "", 10)
					outputs = append(outputs, lines)
				}

				if counter >= 3 {
					lines[10] = record[1]
					lines[11] = record[2]
				} else {
					lines[counter*4] = record[0]
					lines[counter*4+1] = record[1]
					lines[counter*4+2] = record[2]
					if len(record) >= 5 {
						lines[counter*4+3] = record[4]
					} else {
						lines[counter*4+3] = record[3]
					}
				}

				counter++
				if isFirst {
					isFirst = false
				}
			}
		}

		lines[3] = strings.Replace(lines[3], ",", "", 10)
		outputs = append(outputs, lines)
	}

	// Cria o arquivo Excel com os dados processados
	for i, data := range outputs {
		sheetName := fmt.Sprintf("Page %d", i+1)
		xlsx.NewSheet(sheetName)
		for colIndex, value := range data[:11] {
			cell := fmt.Sprintf("%s%d", string(rune(65+colIndex)), i+1)
			xlsx.SetCellValue(sheetName, cell, value)
		}
	}

	return xlsx, nil
}
