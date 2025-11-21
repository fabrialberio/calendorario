package views

const (
	destTerm           = "/term"
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
	DestAdminTerm           = destAdminBase + destTerm
	DestAdminCalendar       = destAdminBase + destCalendar
	DestAdminTimetableClass = destAdminBase + destTimetableClass

	DestSecretary = destSecretaryBase + "/"

	DestTeacher = destTeacherBase + "/"
)
