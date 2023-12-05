git init
git add .
git commit -m "create: user_controller.go and convert_domain_to_response.go "
git config --global --list --show-origin
git branch -M main
git remote add origin https://github.com/WagaoCarvalhoDev/crud-go.git
git push -u origin main
git log --oneline
