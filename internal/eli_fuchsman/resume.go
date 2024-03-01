package elifuchsman

// func getHomeDirectory() (string, error) {
// 	if homeDir, err := os.UserHomeDir(); err == nil {
// 		return homeDir, nil
// 	}

// 	currentUser, err := user.Current()
// 	if err != nil {
// 		return "", err
// 	}
// 	return currentUser.HomeDir, nil
// }

// func (c *EliFuchsmanClient) RetrieveResume(region string, bucketName string, objectKey string) error {
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String(region),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	svc := s3.New(sess)

// 	homeDir, err := getHomeDirectory()
// 	if err != nil {
// 		return err
// 	}

// 	downloadPath := filepath.Join(homeDir, "Downloads", "Eli Fuchsman resume.pdf")

// 	file, err := os.Create(downloadPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	resp, err := svc.GetObject(&s3.GetObjectInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(objectKey),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	_, err = io.Copy(file, resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("Resume downloaded successfully to %s\n", downloadPath)

// 	return nil
// }
