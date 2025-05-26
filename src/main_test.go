package main

import "testing"

// add et search
func Test1(t *testing.T) {
	annuaire = make(map[string]contact)
	add("Alice Martin", "0622335544")
	result := search("Alice Martin")
	expected := "Contact trouvé : Alice Martin - 0622335544"
	if result != expected {
		t.Errorf("rechercher = %v; want %v", result, expected)
	}
}

// add already exist
func Test2(t *testing.T) {
	annuaire = make(map[string]contact)
	add("Charlie Dupont", "0788112233")
	result := add("Charlie Dupont", "0999888777")
	expected := "Ce contact existe déjà."
	if result != expected {
		t.Errorf("ajouter doublon = %v; want %v", result, expected)
	}
}

// delete
func Test3(t *testing.T) {
	annuaire = make(map[string]contact)
	add("Lucie Bernard", "0644556677")
	result := supprimer("Lucie Bernard")
	expected := "Contact supprimé."
	if result != expected {
		t.Errorf("supprimer = %v; want %v", result, expected)
	}
}

// update
func Test4(t *testing.T) {
	annuaire = make(map[string]contact)
	add("Thomas Durand", "0677889900")
	result := add("David", "2222")
	expected := "Contact modifié."
	if result != expected {
		t.Errorf("modifier = %v; want %v", result, expected)
	}
}
