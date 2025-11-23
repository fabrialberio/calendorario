package views

const (
	destTerm           = "/term"
	destVacation       = "/vacation"
	destCalendar       = "/calendar"
	destTimetableClass = "/timetable_class"

	destAdminBase     = "/admin"
	destSecretaryBase = "/secretary"
	destTeacherBase   = "/teacher"
)

const (
	DestLogin  = "/login"
	DestLogout = "/logout"

	DestAdmin               = destAdminBase + "/"
	DestAdminLoadTerm       = destAdminBase + "/load_term"
	DestAdminTerm           = destAdminBase + destTerm
	DestAdminVacation       = destAdminBase + destVacation
	DestAdminCalendar       = destAdminBase + destCalendar
	DestAdminTimetableClass = destAdminBase + destTimetableClass

	DestSecretary = destSecretaryBase + "/"

	DestTeacher = destTeacherBase + "/"
)
