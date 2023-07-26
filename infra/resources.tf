# -------------------------------------------------- # 
# Data stores # 
# -------------------------------------------------- # 
resource "aws_dynamodb_table" "example" {
  name             = "example-service-go"
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
