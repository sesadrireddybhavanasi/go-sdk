package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
    // create a new AWS session
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    })
    if err != nil {
        fmt.Println("Error creating session", err)
        os.Exit(1)
    }

    // create a new EC2 service client
    svc := ec2.New(sess)

    // create a keypair
    keyPairName := "doc"
    keyPairInput := &ec2.CreateKeyPairInput{
        KeyName: aws.String(keyPairName),
    }
    keyPairOutput, err := svc.CreateKeyPair(keyPairInput)
    if err != nil {
        fmt.Println("Error creating key pair", err)
        os.Exit(1)
    }

    // download the keypair
    err = ioutil.WriteFile(keyPairName+".pem", []byte(*keyPairOutput.KeyMaterial), 0600)
    if err != nil {
        fmt.Println("Error downloading key pair", err)
        os.Exit(1)
    }

    // create an EC2 instance
    instanceInput := &ec2.RunInstancesInput{
        ImageId:      aws.String("ami-0dfcb1ef8550277af"),
        InstanceType: aws.String("t2.micro"),
        MinCount:     aws.Int64(1),
        MaxCount:     aws.Int64(1),
        KeyName:      aws.String(keyPairName),
    }
    instanceOutput, err := svc.RunInstances(instanceInput)
    if err != nil {
        fmt.Println("Error creating instance", err)
        os.Exit(1)
    }

    // print the instance ID and the private key file path
    fmt.Println("Instance created:", *instanceOutput.Instances[0].InstanceId)
    fmt.Println("Key pair downloaded to:", keyPairName+".pem")
}
