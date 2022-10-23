echo "~~~~~ Starting database ~~~~~~~~~"
make docker-up
make test-renew-seed
echo "~~~~~ Run test ~~~~~~~~~"
TEST_RESULT_DIR="./test/test-result"
# go test ./...
go test -v -coverprofile ${TEST_RESULT_DIR}/cover.out ./...
go tool cover -html ${TEST_RESULT_DIR}/cover.out -o ${TEST_RESULT_DIR}/test-cover.html
open ${TEST_RESULT_DIR}/test-cover.html