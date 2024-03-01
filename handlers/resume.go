package handlers

import (
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetResume(w http.ResponseWriter, r *http.Request) {
	awsRegion, ok := r.Context().Value("AWSRegion").(string)
	if !ok {
		log.Error("AWS region not found in context")
		InternalError500(w, "AWSRegion", ErrAWSRegionNotFound)
		return
	}

	bucket, ok := r.Context().Value("Bucket").(string)
	if !ok {
		log.Error("Bucket not found in context")
		InternalError500(w, "Bucket", ErrBucketNotFound)
		return
	}

	bucketKey, ok := r.Context().Value("BucketKey").(string)
	if !ok {
		log.Error("Bucket key not found in context")
		InternalError500(w, "BucketKey", ErrBucketKeyNotFound)
		return
	}

	fields := log.Fields{"AWSRegion": awsRegion, "Bucket": bucket, "BucketKey": bucketKey}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		log.WithFields(fields).Errorf("%+v", err)
		InternalError500(w, "AWS Session", ErrNoResponseFromAWS)
		return
	}

	svc := s3.New(sess)

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(bucketKey),
	})
	if err != nil {
		log.WithFields(fields).Errorf("%+v", err)
		InternalError500(w, "AWS Response", ErrNoResponseFromAWS)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=Eli_Fuchsman_resume.pdf")
	w.Header().Set("Content-Type", aws.StringValue(resp.ContentType))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
