package main

import "fmt"

func main() {
	nombreHeures := 730.0
	
	prixAzureParHeure := 0.0960 //D2s_v3
	prixAWSParHeure := 0.0416 //t3 medium
	prixGCPParHeure := 0.0670 //e2-standard-2

	coutTotalAzure := nombreHeures * prixAzureParHeure
	coutTotalAWS := nombreHeures * prixAWSParHeure
	coutTotalGCP := nombreHeures * prixGCPParHeure

	fmt.Println("===========================================")
	fmt.Println("   COMPARATEUR DE COÛT CLOUD (MENSUEL)    ")
	fmt.Println("===========================================")
	
	fmt.Printf("Estimation Coût Azure : %.2f USD\n", coutTotalAzure)
	fmt.Printf("Estimation Coût AWS   : %.2f USD\n", coutTotalAWS)
	fmt.Printf("Estimation Coût GCP   : %.2f USD\n", coutTotalGCP) // Corrigé ici
	fmt.Println("-------------------------------------------")

	// Logique simplifiée : On cherche le plus petit des trois
	if coutTotalGCP < coutTotalAWS && coutTotalGCP < coutTotalAzure {
		fmt.Println("VERDICT : GCP est le plus économique.")
	} else if coutTotalAWS < coutTotalAzure {
		fmt.Println("VERDICT : AWS est le plus économique.")
	} else {
		fmt.Println("VERDICT : Azure est le plus économique.")
	}
	fmt.Println("===========================================")
}