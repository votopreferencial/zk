package main

import (
	"fmt"
	"sort"
)

// Estrutura para armazenar informações sobre uma candidata
type Candidata struct {
	Nome  string
	Votos int
}

// Função para calcular os totais de votos ponderados
func calcularTotais(votos [][]string, candidatasValidas []string, pesos []int) []Candidata {
	resultados := make(map[string]int)

	// Inicializando o mapa de resultados com as candidatas válidas
	for _, candidata := range candidatasValidas {
		resultados[candidata] = 0
	}

	// Aplicando os pesos para cada voto
	for _, voto := range votos {
		for i, candidata := range voto {
			if i < len(pesos) && contains(candidatasValidas, candidata) {
				resultados[candidata] += pesos[i]
			}
		}
	}

	// Convertendo o mapa de resultados para uma lista de Candidatas
	var listaResultados []Candidata
	for nome, total := range resultados {
		listaResultados = append(listaResultados, Candidata{Nome: nome, Votos: total})
	}

	// Ordenando as candidatas por número de votos
	sort.Slice(listaResultados, func(i, j int) bool {
		return listaResultados[i].Votos > listaResultados[j].Votos
	})

	return listaResultados
}

// Função para verificar se a candidata está na lista de candidatas válidas
func contains(candidatasValidas []string, candidata string) bool {
	for _, c := range candidatasValidas {
		if c == candidata {
			return true
		}
	}
	return false
}

// Função principal que simula a apuração
func main() {
	// Lista de candidatas válidas
	candidatasValidas := []string{
		"Tainá de Paula (PT)", "Rosa Fernandes (PSD)", "Joyce Trindade (PSD)", "Helena Vieira (PSD)",
		"Vera Lins (PP)", "Monica Benicio (PSOL)", "Tânia Bastos (REPUBLICANOS)", "Talita Galhardo (PSDB)",
		"Thais Ferreira (PSOL)", "Tatiana Roque (PSB)", "Maíra do MST (PT)", "Gigi Castilho (REPUBLICANOS)",
	}

	// Simulando uma tabela de votos onde cada linha representa um voto com as preferências
	votos := [][]string{
		{"Tainá de Paula (PT)", "Monica Benicio (PSOL)", "Thais Ferreira (PSOL)"},
		{"Tatiana Roque (PSB)", "Joyce Trindade (PSD)", "Maíra do MST (PT)"},
		{"Rosa Fernandes (PSD)", "Vera Lins (PP)", "Tânia Bastos (REPUBLICANOS)"},
		{"Monica Benicio (PSOL)", "Thais Ferreira (PSOL)", "Tainá de Paula (PT)"},
	}

	// Definindo os pesos para o sistema de apuração
	pesos := []int{3, 2, 1}

	// Chamando a função para calcular os totais
	resultados := calcularTotais(votos, candidatasValidas, pesos)

	// Exibindo os resultados
	fmt.Println("Resultados da apuração STV:")
	for _, candidata := range resultados {
		fmt.Printf("%s: %d votos ponderados\n", candidata.Nome, candidata.Votos)
	}
}