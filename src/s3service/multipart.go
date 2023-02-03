package s3service

import (
	"bytes"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/forfam/authentication-service/src/utils/logger"
)

const (
	maxPartSize = int64(5 * 1024 * 1024)
	maxRetries  = 3
)

func MultipartUpload(
	bucket string,
	key string,
	file *multipart.FileHeader,
	log *logger.Logger,
) (string, error) {
	fileBuffer, err := file.Open()
	if err != nil {
		fmt.Println("Error while file open", err)
		return "", err
	}
	defer fileBuffer.Close()

	fileSize := file.Size
	buffer := make([]byte, fileSize)
	fileBuffer.Read(buffer)

	input := &s3.CreateMultipartUploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String(file.Header["Content-Type"][0]),
	}

	upload, err := s3Client.CreateMultipartUpload(input)
	if err != nil {
		log.Error("Something went wrong while creating multipart upload: " + err.Error())
		return "", err
	}

	log.Info("Multipart upload created for: " + key)

	var curr, partLength int64
	var remaining = fileSize
	var completedParts []*s3.CompletedPart
	partNumber := 1
	for curr = 0; remaining != 0; curr += partLength {
		if remaining < maxPartSize {
			partLength = remaining
		} else {
			partLength = maxPartSize
		}
		completedPart, err := uploadPart(upload, buffer[curr:curr+partLength], partNumber, log)
		if err != nil {
			fmt.Println(err.Error())
			err := abortMultipartUpload(upload)
			if err != nil {
				fmt.Println(err.Error())
			}
			return "", err
		}
		remaining -= partLength
		partNumber++
		completedParts = append(completedParts, completedPart)
	}

	completeResponse, err := completeMultipartUpload(upload, completedParts)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	fmt.Printf("Successfully uploaded file: %s\n", completeResponse.String())
	return key, nil
}

func completeMultipartUpload(
	resp *s3.CreateMultipartUploadOutput,
	completedParts []*s3.CompletedPart,
) (*s3.CompleteMultipartUploadOutput, error) {
	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	}
	return s3Client.CompleteMultipartUpload(completeInput)
}

func uploadPart(
	resp *s3.CreateMultipartUploadOutput,
	fileBytes []byte,
	partNumber int,
	log *logger.Logger,
) (*s3.CompletedPart, error) {
	tryNum := 1
	partInput := &s3.UploadPartInput{
		Body:          bytes.NewReader(fileBytes),
		Bucket:        resp.Bucket,
		Key:           resp.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      resp.UploadId,
		ContentLength: aws.Int64(int64(len(fileBytes))),
	}

	for tryNum <= maxRetries {
		uploadResult, err := s3Client.UploadPart(partInput)
		if err != nil {
			if tryNum == maxRetries {
				if aerr, ok := err.(awserr.Error); ok {
					return nil, aerr
				}
				return nil, err
			}
			log.Warning("Retrying to upload part")
			tryNum++
		} else {
			log.Info("Uploaded part")
			return &s3.CompletedPart{
				ETag:       uploadResult.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			}, nil
		}
	}
	return nil, nil
}

func abortMultipartUpload(resp *s3.CreateMultipartUploadOutput) error {
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
	}
	_, err := s3Client.AbortMultipartUpload(abortInput)
	return err
}
