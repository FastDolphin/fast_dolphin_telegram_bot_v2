package utils

var AdminCommands = map[string]string{
	"/start":                         "Начать",
	"/fetch_all_requests":            "Получить все запросы",
	"/read_logs":                     "Читать логи",
	"/desc_groups":                   "Описание групп",
	"/get_training_plan":             "Получить тренировку",
	"/get_training_for_current_week": "Получить тренировку на эту неделю",
	"/get_calender_week":             "Какая сейчас календарная неделя?",
	"/get_achievements":              "Посмотреть свои достижения",
	"/get_parent_training":           "Получить тренировку для родителей",
}

var UserCommands = map[string]string{
	"/start":               "Начать",
	"/contact":             "Как с нами связаться?",
	"/new_request":         "Оставить заявку",
	"/get_achievements":    "Посмотреть свои достижения",
	"/get_parent_training": "Получить тренировку для родителей",
}

var CoachCommands = map[string]string{
	"/start":                         "Начать",
	"/desc":                          "Объяснения",
	"/get_training_plan":             "Получить тренировку",
	"/get_training_for_current_week": "Получить тренировку на эту неделю",
	"/desc_groups":                   "Описание групп",
	"/get_calender_week":             "Какая сейчас календарная неделя?",
	"/get_achievements":              "Посмотреть свои достижения",
	"/get_parent_training":           "Получить тренировку для родителей",
}
