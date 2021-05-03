make build
make install
rm -rf ./example/.terraform.lock.hcl terraform.*
cd example
terraform init
terraform apply -auto-approve
cd ../