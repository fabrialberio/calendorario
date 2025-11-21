package views

const (
	destTerms    = "/terms"
	destCalendar = "/calendar"
	destClasses  = "/classes"

	destAdminBase     = "/admin"
	destSecretaryBase = "/secretary"
	destTeacherBase   = "/teacher"
)

const (
	DestLogin  = "/login"
	DestLogout = "/logout"

	DestAdmin         = destAdminBase + "/"
	DestAdminTerms    = destAdminBase + destTerms
	DestAdminCalendar = destAdminBase + destCalendar
	DestAdminClasses  = destAdminBase + destClasses

	DestSecretary = destSecretaryBase + "/"

	DestTeacher = destTeacherBase + "/"
)
