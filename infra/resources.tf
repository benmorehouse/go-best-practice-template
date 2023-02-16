# Route53 DNS record which points to API Gateway
# Lambda 
# DynamoDB
# API Gateway endpoint

# -------------------------------------------------- # 
# Route53 # 
# -------------------------------------------------- # 

/*
data "aws_route53_zone" "selected" {
  name         = "budgetballerz.a.extra.app"
  private_zone = false
}

resource "aws_route53_record" "budget_service" {
  zone_id = data.aws_route53_zone.selected.zone_id
  name    = "budget-service-go.budgetballerz.a.extra.app"
  type    = "A"
  ttl     = 300
  records = ["budget-service-go.budgetballerz.a.extra.app"]
}
*/
# -------------------------------------------------- # 
# Data stores # 
# -------------------------------------------------- # 
resource "aws_dynamodb_table" "example" {
  name             = "budget-service-go"
  hash_key         = "UserID"
  range_key        = "CreatedAt"
  billing_mode     = "PAY_PER_REQUEST"
  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "UserID"
    type = "S"
  }

  attribute {
    name = "CreatedAt"
    type = "S"
  }
}
