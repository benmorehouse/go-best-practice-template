# -------------------------------------------------- # 
# Lambda # 
# -------------------------------------------------- # 
#
# added AWSLambdaBasicExecutionRole manually as a policy -- add this to tf later
#
# Also have to make sure that API gateway has permissions: 
#
# 2023-02-15T15:08:37.713-08:00
# Copy
# 64.136.145.109 - - [15/Feb/2023:23:08:37 +0000] "GET $default HTTP/1.1" 500 35 AZzecgvEPHcEPew= The IAM role configured on the integration or API Gateway doesn't have permissions to call the integration. Check the permissions and try again.
resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "dynamodb_access" {
  name = "dynamodb_access_budget_service_go"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "dynamodb:*"
        ]
        Resource = [
          aws_dynamodb_table.example.arn,
          "*"
        ]
      }
    ]
  })

  role = aws_iam_role.iam_for_lambda.name
}

data "aws_iam_policy" "aws_lambda_basic_execution" {
  name = "AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "attach_lambda_exec" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = data.aws_iam_policy.aws_lambda_basic_execution.arn
}

