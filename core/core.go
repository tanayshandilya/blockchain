package core

const (
	txn_min_fee = 0.0000005
	txn_max_fee = 0.0005

	ChainVersion = "1.0.0"
	EventVersion = "1.0.0"

	event_GenesisData      = "genesis"
	event_GenesisTimestamp = "1996-05-31 00:00:00.0000000 +0000 UTC"

	event_CreateQuestion  = "create_question"
	event_SubmitQuestion  = "submit_question"
	event_ReviewQuestion  = "review_question"
	event_ConfirmQuestion = "confirm_question"

	event_RequestLicense = "request_license"
	event_CreateLicense  = "create_license"
	event_ConfirmLicense = "confirm_license"

	event_CreateExamAppointment  = "create_exam_appointment"
	event_WitnessExamAppointment = "witness_exam_appointment"
	event_ConfirmExamAppointment = "confirm_exam_appointment"
	event_AssignExamAppointment  = "assign_exam_appointment"

	event_RequestExamAppointment = "request_exam_appointment"
	event_ApproveExamAppointment = "approve_exam_appointment"
	event_RejectExamAppointment  = "reject_exam_appointment"
	event_AcceptExamAppointment  = "accept_exam_appointment"
)
