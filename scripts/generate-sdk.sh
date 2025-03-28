#!/bin/bash
set -e

# This script fetches the latest Mistral AI OpenAPI spec and generates a Go SDK

# Check if jq is installed
if ! command -v jq &>/dev/null; then
    echo "Error: jq is required but not installed. Please install it first."
    exit 1
fi

# Download Swagger Codegen CLI v3 if not present
SWAGGER_CODEGEN_JAR="swagger-codegen-cli.jar"
if [ ! -f "$SWAGGER_CODEGEN_JAR" ]; then
    echo "Downloading Swagger Codegen CLI v3..."
    wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.68/swagger-codegen-cli-3.0.68.jar -O $SWAGGER_CODEGEN_JAR
fi

# Fetch latest OpenAPI spec
echo "Fetching the latest OpenAPI spec from Mistral AI..."
RAW_URL="https://raw.githubusercontent.com/mistralai/platform-docs-public/main/openapi.yaml"
curl -L $RAW_URL -o openapi.yaml

# Get the latest commit info
echo "Getting commit information..."
REPO_URL="https://api.github.com/repos/mistralai/platform-docs-public/commits?path=openapi.yaml&per_page=1"
COMMIT_INFO=$(curl -s $REPO_URL)

# Extract commit hash and message
COMMIT_HASH=$(echo $COMMIT_INFO | jq -r '.[0].sha')
COMMIT_MESSAGE=$(echo $COMMIT_INFO | jq -r '.[0].commit.message')

# Try to extract version from commit message
VERSION_PATTERN="v[0-9]+\.[0-9]+\.[0-9]+"
if [[ $COMMIT_MESSAGE =~ $VERSION_PATTERN ]]; then
    VERSION=${BASH_REMATCH[0]}
else
    # Fallback: use date and short commit hash
    VERSION="v$(date +%Y%m%d)-${COMMIT_HASH:0:7}"
fi

echo "Using version: $VERSION"

# Generate Go client
echo "Generating Go SDK..."
java -jar $SWAGGER_CODEGEN_JAR generate \
    -i openapi.yaml \
    -l go \
    -o . \
    --additional-properties packageName=mistral

go mod tidy
# Git operations
echo "Committing changes..."
git add openapi.yaml .
git commit -m "Update Mistral SDK to $VERSION" || echo "No changes to commit"

echo "Tagging version $VERSION..."
git tag -a "$VERSION" -m "Mistral SDK $VERSION" || echo "Tag $VERSION already exists"

echo "Done! SDK updated to version $VERSION"
echo "You can now push the changes with: git push origin main && git push origin $VERSION"
