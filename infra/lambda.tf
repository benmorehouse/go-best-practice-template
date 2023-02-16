data "archive_file" "dummy" {
  output_path = "dist.zip"
  type        = "zip"
  source {
    content  = "dummy dummy"
    filename = "dummy.txt"
  }
}

resource "aws_lambda_function" "budget_service" {
  # If the file is not in the current working directory you will need to include a
  # path.module in the filename.
  function_name = "budget-service-go"
  publish       = true
  role          = aws_iam_role.iam_for_lambda.arn

  filename      = data.archive_file.dummy.output_path
  
  # runtime 
  runtime       = "go1.x"
  handler       = "bin/cmd" # must match with what the gihtub actions packages

  # LAMBDA CI is done through codebuild/codepipeline
  lifecycle {
    ignore_changes = [s3_key, s3_bucket, layers, filename]
  }
}

# Required by lambda provisioned concurrency
resource "aws_lambda_alias" "alias" {
  name             = "latest"
  description      = "alias pointing to the latest published version of the lambda"
  function_name    = aws_lambda_function.budget_service.function_name
  function_version = aws_lambda_function.budget_service.version

  lifecycle {
    ignore_changes = [
      description,
      routing_config
    ]
  }
}

resource "aws_cloudwatch_log_group" "lambda_group" {
  name = "BudgetServiceLogGroup"

  tags = {
    Environment = "production"
    Application = "BudgetService"
  }
}

data "aws_route53_zone" "zone" {
  name = "budgetballerz.a.extra.app"
}

module "acm" {
  source  = "terraform-aws-modules/acm/aws"
  version = "~> 4.0"

  domain_name  = "budget-service-go.budgetballerz.a.extra.app"
  zone_id      = data.aws_route53_zone.zone.zone_id

  wait_for_validation = true
}

resource "aws_route53_record" "api" {
  zone_id = data.aws_route53_zone.zone.zone_id
  name    = "budget-service-go"
  type    = "A"

  alias {
    name                   = module.api_gateway.apigatewayv2_domain_name_configuration[0].target_domain_name
    zone_id                = module.api_gateway.apigatewayv2_domain_name_configuration[0].hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "budget-service-go-lambda-permission"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.budget_service.function_name
  principal     = "apigateway.amazonaws.com"

  # The /*/*/* part allows invocation from any stage, method and resource path
  # within API Gateway REST API.
  # source_arn = "${module.api_gateway.apigatewayv2_api_execution_arn}/*/*/*"
  # look at using source account option instead. This will then mean that 
  # any api gateway can reach this lambda from within this account.
}

module "api_gateway" {
  source = "terraform-aws-modules/apigateway-v2/aws"

  name          = "budget-service-go"
  description   = "Budget Service written in go"
  protocol_type = "HTTP"

  domain_name                 = "budget-service-go.budgetballerz.a.extra.app"
  domain_name_certificate_arn = module.acm.acm_certificate_arn

  create_api_domain_name = true

  cors_configuration = {
    allow_headers = ["content-type", "x-amz-date", "authorization", "x-api-key", "x-amz-security-token", "x-amz-user-agent"]
    allow_methods = ["*"]
    allow_origins = ["*"]
  }

  # set up to have separate environments for staging vs production
  default_stage_access_log_destination_arn = aws_cloudwatch_log_group.lambda_group.arn
  default_stage_access_log_format          = "$context.identity.sourceIp - - [$context.requestTime] \"$context.httpMethod $context.routeKey $context.protocol\" $context.status $context.responseLength $context.requestId $context.integrationErrorMessage"

  # Routes and integrations
  integrations = {
    "$default" = {
      lambda_arn = aws_lambda_function.budget_service.arn
    }
  }
}
