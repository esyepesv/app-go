COVERAGE_PERCENTAGE=$(go tool cover -func=coverage.out | grep total: | awk '{print $3}' | sed 's/\%//g')
 
if [ "$COVERAGE_PERCENTAGE" -lt "70" ]; then
    echo "Code coverage ($COVERAGE_PERCENTAGE%) is below the required threshold (70%)"
    exit 1  # Exit with a non-zero status to indicate an error
else
    echo "Code coverage ($COVERAGE_PERCENTAGE%) meets the required threshold"
fi