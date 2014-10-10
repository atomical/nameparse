package nameparse

import (
  "testing"
)

func TestParseName1(t *testing.T) {

  n := &NameParse{}

  n.Parse("Mr. Adam T Hallett")

  if n.Suffix != "" {
    t.Errorf("Could not find the correct suffix: %v", n.Suffix)
  }

  if n.Salutation != "Mr." {
    t.Errorf("Could not find the correct saluation: %v", n.Salutation)
  }

  if n.FirstName != "Adam" {
    t.Errorf("Could not find the correct first name: %v", n.FirstName)
  }

  if n.LastName != "Hallett" {
    t.Errorf("Could not find the correct last name: %v", n.LastName)
  }
  
}

func TestParseName2(t *testing.T) {

  n := &NameParse{}

  n.Parse("Mike Jones")

  if n.Suffix != "" {
    t.Errorf("Could not find the correct suffix: %v", n.Suffix)
  }

  if n.Salutation != "" {
    t.Errorf("Could not find the correct saluation: %v", n.Salutation)
  }

  if n.FirstName != "Mike" {
    t.Errorf("Could not find the correct first name: %v", n.FirstName)
  }

  if n.LastName != "Jones" {
    t.Errorf("Could not find the correct last name: %v", n.LastName)
  }
  
}