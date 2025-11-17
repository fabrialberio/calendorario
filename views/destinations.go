package views

const (
	destCalendar       = "/calendar"
	destTimetableClass = "/timetableclass"
	destAdminBase      = "/admin"
	destSecretaryBase  = "/secretary"
	destTeacherBase    = "/teacher"
)

const (
	DestLogin  = "/login"
	DestLogout = "/logout"

	DestAdmin               = destAdminBase + "/"
	DestAdminCalendar       = destAdminBase + destCalendar
	DestAdminTimetableClass = destAdminBase + destTimetableClass

	DestSecretary = destSecretaryBase + "/"

	DestTeacher = destTeacherBase + "/"
)
