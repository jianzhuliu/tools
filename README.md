"# tools" 

...create a new repository on the command line
echo "# tools" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M master
git remote add github https://github.com/jianzhuliu/tools.git
git push -u github master

...or push an existing repository from the command line
git remote add github https://github.com/jianzhuliu/tools.git
git branch -M master
git push -u github master