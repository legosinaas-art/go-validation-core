package service

func (s *MigrantService) CheckGos(passportID string) bool {
	if passportID == "0000" {
		return false
	}
	return true
}
