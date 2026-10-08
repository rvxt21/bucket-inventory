#!/bin/sh
awslocal s3api head-bucket --bucket "${S3_BUCKET_NAME}" 2>/dev/null || awslocal s3 mb "s3://${S3_BUCKET_NAME}"
