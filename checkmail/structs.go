package main

import "fmt"

type EmailVerInfo struct {
	Email      string
	Username   string
	Domain     string
	SyntaxOK   bool
	Valid      bool
	DomainInfo *DomainInfo
	Error      error
}

func (e *EmailVerInfo) String() string {
	fmtstr := "%-15v:%v\n"
	res := fmt.Sprintf(fmtstr, "Email", e.Email)
	if e.SyntaxOK {
		res += fmt.Sprintf(fmtstr, "Syntax ok", "😀")
	} else {
		res += fmt.Sprintf(fmtstr, "Syntax nicht ok", "\u26D4")
	}
	if e.Valid {
		res += fmt.Sprintf(fmtstr, "gültig", "😀")
	} else {
		res += fmt.Sprintf(fmtstr, "ungültig", "\u26D4")
	}
	res += fmt.Sprintf(fmtstr, "Benutzer", e.Username)
	res += fmt.Sprintf(fmtstr, "Domäne", e.Domain)
	res += fmt.Sprintf(fmtstr, "Mx", e.DomainInfo)
	if e.Error != nil {
		res += fmt.Sprintf(fmtstr, "Fehler", e.Error)
	}
	return res
}
