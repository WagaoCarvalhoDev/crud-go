git init
git add .
git commit -m "create: repository "
git config --global --list --show-origin
git branch -M main
git remote add origin https://github.com/WagaoCarvalhoDev/crud-go.git
git push -u origin main
git log --oneline
