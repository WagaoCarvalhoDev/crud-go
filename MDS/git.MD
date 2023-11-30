git init
git add .
git commit -m "create: model/user_domain.go"
git config --global --list --show-origin
git branch -M main
git remote add origin https://github.com/WagaoCarvalhoDev/crud-go.git
git push -u origin main
git log --oneline
