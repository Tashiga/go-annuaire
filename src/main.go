package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type contact struct {
	Nom string `json:"nom"`
	Tel string `json:"tel"`
}

var annuaire = make(map[string]contact)

const fichier = "annuaire.json"

func importAnnuaire() {
	file, err := os.ReadFile(fichier)
	if err != nil {
		return // Fichier non trouvé, on garde une map vide
	}
	json.Unmarshal(file, &annuaire)
}

func saveAnnuaire() {
	data, err := json.MarshalIndent(annuaire, "", "  ")
	if err != nil {
		fmt.Println("Erreur de sauvegarde :", err)
		return
	}
	os.WriteFile(fichier, data, 0644)
}

func add(nom, tel string) string {
	if _, exists := annuaire[nom]; exists {
		return "Ce contact existe déjà."
	}
	annuaire[nom] = contact{Nom: nom, Tel: tel}
	saveAnnuaire()
	return "Contact ajouté."
}

func search(nom string) string {
	if c, ok := annuaire[nom]; ok {
		return fmt.Sprintf("Contact trouvé : %s - %s", c.Nom, c.Tel)
	}
	return "Contact introuvable."
}

func lister() []contact {
	var liste []contact
	for _, c := range annuaire {
		liste = append(liste, c)
	}
	return liste
}

func supprimer(nom string) string {
	if _, ok := annuaire[nom]; ok {
		delete(annuaire, nom)
		saveAnnuaire()
		return "Contact supprimé."
	}
	return "Contact introuvable."
}

func update(nom, tel string) string {
	if _, ok := annuaire[nom]; ok {
		annuaire[nom] = contact{Nom: nom, Tel: tel}
		saveAnnuaire()
		return "Contact modifié."
	}
	return "Contact introuvable."
}

func main() {
	//charger l'annuaire
	importAnnuaire()

	//les actions
	var action, nom, tel string
	flag.StringVar(&action, "action", "", "Action à effectuer: ajouter, rechercher, lister, supprimer, modifier")
	flag.StringVar(&nom, "nom", "", "Nom du contact")
	flag.StringVar(&tel, "tel", "", "Numéro de téléphone")
	flag.Parse()

	switch action {
	case "ajouter":
		fmt.Println(add(nom, tel))
	case "rechercher":
		fmt.Println(search(nom))
	case "lister":
		for _, c := range lister() {
			fmt.Printf("%s : %s\n", c.Nom, c.Tel)
		}
	case "supprimer":
		fmt.Println(supprimer(nom))
	case "modifier":
		fmt.Println(update(nom, tel))
	default:
		fmt.Println("Action inconnue")
	}
}
