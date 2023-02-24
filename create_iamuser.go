	package main

	import (
		"fmt"

		"github.com/aws/aws-sdk-go/aws"
		"github.com/aws/aws-sdk-go/aws/session"
		"github.com/aws/aws-sdk-go/service/iam"
	)


	func main() {
		// Create a session object that the SDK uses to communicate with AWS
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		// Create a new IAM service client
		svc := iam.New(sess)

		// Set the username of the new IAM user
		userName := "my-new-iam-user"

		// Create the new IAM user
		_, err := svc.CreateUser(&iam.CreateUserInput{
			UserName: aws.String(userName),
		})

		// Check for errors
		if err != nil {
			fmt.Println("Error creating IAM user:", err)
			return
		}

		// Print a success message
		fmt.Printf("Created IAM user: %s\n", userName)
	}
