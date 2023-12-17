package handlers

import (
	"fmt"
	"log"

	"anilkhadka.com.np/task-management/internal/models"
	"anilkhadka.com.np/task-management/internal/services"
	"anilkhadka.com.np/task-management/utils"
)

func SendDailyMail() {
	// Get assignees from the database with tasks not marked as DONE
	taskService := services.NewTaskService()
	assignees, err := taskService.AssigneeWithUndoneTasks()
	if err != nil {
		fmt.Println("Error getting assignees:", err)
		return
	}

	// Iterate over assignees and send reminders
	for _, assignee := range assignees {
		sendReminderEmail(assignee)
	}
}

func sendReminderEmail(assignee models.AssigneeWithTask) {
	// Implement your email sending logic here
	// Use a library like gomail or net/smtp
	fmt.Printf("Sending email to assignee %d with tasks %+v\n", assignee.AssigneeID, assignee.TaskDueDate)

	emailService := services.NewEmailService()

	templateFile := "internal/templates/assignee_email.html"
	emailTemplate, err := utils.ReadFile(templateFile)
	if err != nil {
		log.Fatal(err)
	}
	emailBody := utils.ParseTemplate(emailTemplate, assignee)

	err = emailService.SendEmail(assignee.AssigneeEmail, emailBody)

	if err != nil {
		log.Println(err)
	}
}
