package nameparse

import (
  "strings"
)

type NameParse struct {
  Salutation string
  Suffix string
  FirstName string
  LastName string
  MiddleInitials string
  Initials string
}

func (n * NameParse ) Parse( fullName string ) {
  nameParts   := strings.Split(fullName, " ")
  numWords    := len(nameParts)
  
  var start int
  if salutation, found := parseSalutation(nameParts[0]); found != false {
    n.Salutation = salutation
    start = 1
  } else {
    start = 0
  }
  
  if suffix, found := parseSuffix(nameParts[numWords - 1]); found != false {
    n.Suffix = suffix
  }

  word := nameParts[start]

  if isInitial(word) {
    if isInitial( nameParts[start + 1] ) {
      n.FirstName = word
    } else {
      n.Initials = word
    }
  } else {
    n.FirstName = nameParts[start + 1]
  }

  // currently we don't support compound names
  n.LastName = nameParts[numWords -1]
}

func parseSalutation( namePart string ) ( string, bool ) {
  word := normalizeWord(namePart)

  if word == "mr" || word == "master" || word == "mister" {
    return "Mr.",true
  } else if word == "mrs" {
    return "Mrs.", true
  } else if word == "miss" || word == "ms" {
    return "Ms.", true
  } else if word == "dr" {
    return "Dr.", true
  } else if word == "rev" {
    return "Rev.", true
  } else if word == "fr" { 
    return "Fr.", true
  } else {
    return "", false
  }

}

func parseSuffix( namePart string ) ( string, bool ) {
  commonSuffixes := []string{"I","II","III","IV","V","Senior","Junior","Jr","Sr","PhD","APR","RPh","PE","MD","MA","DMD","CME",
      "BVM","CFRE","CLU","CPA","CSC","CSJ","DC","DD","DDS","DO","DVM","EdD","Esq",
      "JD","LLD","OD","OSB","PC","Ret","RGS","RN","RNC","SHCJ","SJ","SNJM","SSMO",
      "USA","USAF","USAFR","USAR","USCG","USMC","USMCR","USN","USNR"}

  for i := range(commonSuffixes) { 
    if strings.ToLower(commonSuffixes[i]) == namePart {
      return commonSuffixes[i], true
    }
  }

  return "", false
}

func isInitial( word string ) bool {
  if len(normalizeWord(word)) > 0 {
    return true
  } else {
    return false
  }
}

func normalizeWord( word string ) string {
  return strings.ToLower(strings.Replace(word, ".", "", -1))
}
